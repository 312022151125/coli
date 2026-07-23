// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

// Package handler 暴露仪表盘相关的 HTTP 接口。
package handler

import (
	"context"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	commonModel "github.com/312022151125/coli/internal/model/common"
	service "github.com/312022151125/coli/internal/service/dashboard"
	jwtUtil "github.com/312022151125/coli/internal/util/jwt"
	"github.com/312022151125/coli/internal/visitor"
	logUtil "github.com/312022151125/coli/pkg/log"
)

type DashboardHandler struct {
	dashboardService service.Service
}

func NewDashboardHandler(dashboardService service.Service) *DashboardHandler {
	return &DashboardHandler{
		dashboardService: dashboardService,
	}
}

type (
	GetSystemLogsInput struct {
		Tail    string `query:"tail" doc:"返回最近多少条（默认 200）"`
		Level   string `query:"level" doc:"日志级别过滤"`
		Keyword string `query:"keyword" doc:"关键词过滤"`
	}
	GetVisitorStatsInput struct{}

)

type (
	LogsOutput         = commonModel.Result[[]logUtil.LogEntry]
	VisitorStatsOutput = commonModel.Result[[]visitor.DayStat]
)


// GetSystemLogs 获取系统历史日志（admin:settings）。成功响应预设显式 message_key（localizeResult 不覆盖）。
func (dashboardHandler *DashboardHandler) GetSystemLogs(ctx context.Context, in *GetSystemLogsInput) (LogsOutput, error) {
	tail := 200
	if rawTail := strings.TrimSpace(in.Tail); rawTail != "" {
		n, err := strconv.Atoi(rawTail)
		if err != nil || n <= 0 {
			return LogsOutput{}, commonModel.NewBizErrorWithMessageKey(
				commonModel.ErrCodeInvalidQuery, commonModel.INVALID_QUERY_PARAMS, commonModel.MsgKeyDashboardTailBad, nil,
			)
		}
		tail = n
	}

	logs, err := dashboardHandler.dashboardService.GetSystemLogs(service.SystemLogQuery{
		Tail:    tail,
		Level:   in.Level,
		Keyword: in.Keyword,
	})
	if err != nil {
		return LogsOutput{}, err
	}
	result := commonModel.OK(logs, "获取系统日志成功")
	result.MessageKey = commonModel.MsgKeyDashboardLogsOk
	return result, nil
}

// GetVisitorStats 获取近七天访客统计（admin:settings）。service 无 error 返回，故补 nil。
func (dashboardHandler *DashboardHandler) GetVisitorStats(ctx context.Context, _ *GetVisitorStatsInput) (VisitorStatsOutput, error) {
	return commonModel.OK(dashboardHandler.dashboardService.GetVisitorStats()), nil
}

func (dashboardHandler *DashboardHandler) WSSubscribeSystemLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "missing token"})
			return
		}

		token = strings.Trim(token, `"`)
		if _, err := jwtUtil.ParseToken(token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			return
		}

		err := dashboardHandler.dashboardService.WSSubscribeSystemLogs(
			ctx.Writer,
			ctx.Request,
			service.SystemLogStreamFilter{
				Level:   ctx.Query("level"),
				Keyword: ctx.Query("keyword"),
			},
		)
		if err != nil {
			logUtil.GetLogger().Error("WebSocket Subscribe System Logs Failed", logUtil.Err(err))
		}
	}
}

func (dashboardHandler *DashboardHandler) SSESubscribeSystemLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Query("token")
		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "missing token"})
			return
		}

		token = strings.Trim(token, `"`)
		if _, err := jwtUtil.ParseToken(token); err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			return
		}

		err := dashboardHandler.dashboardService.SSESubscribeSystemLogs(
			ctx.Writer,
			ctx.Request,
			service.SystemLogStreamFilter{
				Level:   ctx.Query("level"),
				Keyword: ctx.Query("keyword"),
			},
		)
		if err != nil {
			logUtil.GetLogger().Error("SSE Subscribe System Logs Failed", logUtil.Err(err))
		}
	}
}
