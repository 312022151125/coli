// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package migration

import (
	"fmt"

	commonModel "github.com/312022151125/coli/internal/model/common"
	"gorm.io/gorm"
)

type echoKindBackfillMigrator struct{}

// NewEchoKindBackfillMigrator 回填历史 echos 行的空 kind 值为 note：
// AutoMigrate 的列默认值只作用于新增列的既有行（依赖方言支持），此处作为兜底,
// 确保 NULL/空白 kind 一律落回 note，不让旧 post 因新字段而丢失或渲染空徽章。
func NewEchoKindBackfillMigrator() Migrator {
	return &echoKindBackfillMigrator{}
}

func (m *echoKindBackfillMigrator) Name() string {
	return "echo_kind_backfill_migrator"
}

func (m *echoKindBackfillMigrator) Key() string {
	return commonModel.EchoKindBackfilledKey
}

func (m *echoKindBackfillMigrator) CanRerun() bool {
	return false
}

func (m *echoKindBackfillMigrator) Migrate(db *gorm.DB) error {
	if db == nil {
		return fmt.Errorf("database not initialized")
	}
	return db.Exec(
		`UPDATE echos SET kind = 'note' WHERE kind IS NULL OR TRIM(kind) = ''`,
	).Error
}
