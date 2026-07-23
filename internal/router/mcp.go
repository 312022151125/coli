// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package router

import (
	"github.com/312022151125/coli/internal/handler"
	"github.com/312022151125/coli/internal/middleware"
	authModel "github.com/312022151125/coli/internal/model/auth"
)

func setupMCPRoutes(groups *AppRouterGroup, h *handler.Bundle) {
	g := groups.MCPRouterGroup
	g.Use(
		middleware.RateLimit(20, 40),
		middleware.OriginGuard(nil),
		middleware.RequireAudience(authModel.AudienceMCPRemote),
	)
	g.POST("", h.MCPHandler.ServeEndpoint())
	g.GET("", h.MCPHandler.ServeEndpoint())
}
