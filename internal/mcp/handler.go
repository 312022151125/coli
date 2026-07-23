// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package mcp

import (
	"github.com/gin-gonic/gin"
	commentService "github.com/312022151125/coli/internal/service/comment"
	commonService "github.com/312022151125/coli/internal/service/common"
	connectService "github.com/312022151125/coli/internal/service/connect"
	copilotService "github.com/312022151125/coli/internal/service/copilot"
	dashboardService "github.com/312022151125/coli/internal/service/dashboard"
	echoService "github.com/312022151125/coli/internal/service/echo"
	fileService "github.com/312022151125/coli/internal/service/file"
	settingService "github.com/312022151125/coli/internal/service/setting"
	userService "github.com/312022151125/coli/internal/service/user"
)

type Handler struct {
	server *Server
}

func NewHandler(
	echoSvc echoService.Service,
	userSvc userService.Service,
	commentSvc commentService.Service,
	fileSvc fileService.Service,
	commonSvc commonService.Service,
	connectSvc connectService.Service,
	agentSvc copilotService.SummaryService,
	settingSvc settingService.Service,
	dashboardSvc dashboardService.Service,
) *Handler {
	registry := NewRegistry()
	adapter := NewAdapter(echoSvc, userSvc, commentSvc, fileSvc, commonSvc, connectSvc, agentSvc, settingSvc, dashboardSvc)
	adapter.RegisterAll(registry)
	return &Handler{server: NewServer(registry)}
}

func (h *Handler) ServeEndpoint() gin.HandlerFunc {
	return func(c *gin.Context) {
		h.server.ServeHTTP(c.Writer, c.Request)
	}
}
