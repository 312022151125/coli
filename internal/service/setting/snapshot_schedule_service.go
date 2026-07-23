// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package service

import (
	"context"
	"errors"

	"github.com/312022151125/coli/internal/event"
	eventbus "github.com/312022151125/coli/internal/event/bus"
	commonModel "github.com/312022151125/coli/internal/model/common"
	model "github.com/312022151125/coli/internal/model/setting"
	coreSetting "github.com/312022151125/coli/internal/setting"
	fmtUtil "github.com/312022151125/coli/internal/util/format"
	"github.com/312022151125/coli/pkg/viewer"
)

// GetSnapshotScheduleSetting 获取定时快照计划。缺省值由 setting 引擎处理。
func (settingService *SettingService) GetSnapshotScheduleSetting(
	setting *model.SnapshotSchedule,
) error {
	v, err := coreSetting.Get(context.Background(), settingService.durableKV, coreSetting.Snapshot)
	if err != nil {
		return err
	}
	*setting = v
	return nil
}

// UpdateSnapshotScheduleSetting 更新定时快照计划
func (settingService *SettingService) UpdateSnapshotScheduleSetting(
	ctx context.Context,
	newSetting *model.SnapshotScheduleDto,
) error {
	// 鉴权
	userid := viewer.MustFromContext(ctx).UserID()
	user, err := settingService.commonService.CommonGetUserByUserId(ctx, userid)
	if err != nil {
		return err
	}
	if !user.IsAdmin {
		return errors.New(commonModel.NO_PERMISSION_DENIED)
	}

	updated := model.SnapshotSchedule{
		Enable:         newSetting.Enable,
		CronExpression: newSetting.CronExpression,
	}

	// 验证 Cron 表达式是否合法
	if err := fmtUtil.ValidateCrontabExpression(updated.CronExpression); err != nil {
		return errors.New(commonModel.INVALID_CRON_EXPRESSION)
	}

	if err := coreSetting.Set(ctx, settingService.durableKV, coreSetting.Snapshot, updated); err != nil {
		return err
	}

	// 写入成功后再发布事件，避免失败时出现幽灵事件。
	eventbus.Notify(context.Background(), settingService.bus, event.UpdateSnapshotSchedule{Schedule: updated})
	return nil
}
