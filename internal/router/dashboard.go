// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package router

import (
	"net/http"

	"github.com/danielgtaylor/huma/v2"
	"github.com/312022151125/coli/internal/handler"
	"github.com/312022151125/coli/internal/middleware"
	authModel "github.com/312022151125/coli/internal/model/auth"
	authService "github.com/312022151125/coli/internal/service/auth"
)

// setupDashboardRoutes 仅保留实时日志订阅走裸 gin：SSE 流 + WebSocket。
func setupDashboardRoutes(appRouterGroup *AppRouterGroup, h *handler.Bundle) {
	appRouterGroup.AuthRouterGroup.GET(
		"/system/logs/stream",
		middleware.RequireScopes(authModel.ScopeAdminSettings),
		h.DashboardHandler.SSESubscribeSystemLogs(),
	)
	appRouterGroup.WSRouterGroup.GET("/system/logs", h.DashboardHandler.WSSubscribeSystemLogs())
}

// registerDashboard 注册仪表盘的 JSON 端点（admin:settings）。
func registerDashboard(api huma.API, h *handler.Bundle, revoker authService.TokenRevoker) {

	route(api, secured(revoker, authModel.ScopeAdminSettings), huma.Operation{
		OperationID: "dashboard-system-logs",
		Method:      http.MethodGet,
		Path:        "/system/logs",
		Summary:     "获取系统历史日志",
		Tags:        []string{"Dashboard"},
	}, h.DashboardHandler.GetSystemLogs)

	route(api, secured(revoker, authModel.ScopeAdminSettings), huma.Operation{
		OperationID: "dashboard-visitor-stats",
		Method:      http.MethodGet,
		Path:        "/system/visitor-stats",
		Summary:     "获取近七天访客统计",
		Tags:        []string{"Dashboard"},
	}, h.DashboardHandler.GetVisitorStats)
}
