// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package handler

import (
	authHandler "github.com/312022151125/coli/internal/handler/auth"
	commentHandler "github.com/312022151125/coli/internal/handler/comment"
	commonHandler "github.com/312022151125/coli/internal/handler/common"
	connectHandler "github.com/312022151125/coli/internal/handler/connect"
	copilotHandler "github.com/312022151125/coli/internal/handler/copilot"
	dashboardHandler "github.com/312022151125/coli/internal/handler/dashboard"
	echoHandler "github.com/312022151125/coli/internal/handler/echo"
	embeddingHandler "github.com/312022151125/coli/internal/handler/embedding"
	fileHandler "github.com/312022151125/coli/internal/handler/file"
	initHandler "github.com/312022151125/coli/internal/handler/init"
	migratorHandler "github.com/312022151125/coli/internal/handler/migrator"
	settingHandler "github.com/312022151125/coli/internal/handler/setting"
	userHandler "github.com/312022151125/coli/internal/handler/user"
	webHandler "github.com/312022151125/coli/internal/handler/web"
	"github.com/312022151125/coli/internal/mcp"
)

type Bundle struct {
	WebHandler       *webHandler.WebHandler
	UserHandler      *userHandler.UserHandler
	AuthHandler      *authHandler.AuthHandler
	EchoHandler      *echoHandler.EchoHandler
	FileHandler      *fileHandler.FileHandler
	CommentHandler   *commentHandler.CommentHandler
	InitHandler      *initHandler.InitHandler
	CommonHandler    *commonHandler.CommonHandler
	SettingHandler   *settingHandler.SettingHandler
	ConnectHandler   *connectHandler.ConnectHandler
	MigrationHandler *migratorHandler.MigrationHandler
	DashboardHandler *dashboardHandler.DashboardHandler
	CopilotHandler   *copilotHandler.CopilotHandler
	EmbeddingHandler *embeddingHandler.EmbeddingHandler
	MCPHandler       *mcp.Handler
}

func NewBundle(
	webHandler *webHandler.WebHandler,
	userHandler *userHandler.UserHandler,
	authHandler *authHandler.AuthHandler,
	echoHandler *echoHandler.EchoHandler,
	fileHandler *fileHandler.FileHandler,
	commentHandler *commentHandler.CommentHandler,
	initHandler *initHandler.InitHandler,
	commonHandler *commonHandler.CommonHandler,
	settingHandler *settingHandler.SettingHandler,
	connectHandler *connectHandler.ConnectHandler,
	migratorHandler *migratorHandler.MigrationHandler,
	dashboardHandler *dashboardHandler.DashboardHandler,
	copilotHandler *copilotHandler.CopilotHandler,
	embeddingHandler *embeddingHandler.EmbeddingHandler,
	mcpHandler *mcp.Handler,
) *Bundle {
	return &Bundle{
		WebHandler:       webHandler,
		UserHandler:      userHandler,
		AuthHandler:      authHandler,
		EchoHandler:      echoHandler,
		FileHandler:      fileHandler,
		CommentHandler:   commentHandler,
		InitHandler:      initHandler,
		CommonHandler:    commonHandler,
		SettingHandler:   settingHandler,
		ConnectHandler:   connectHandler,
		MigrationHandler: migratorHandler,
		DashboardHandler: dashboardHandler,
		CopilotHandler:   copilotHandler,
		EmbeddingHandler: embeddingHandler,
		MCPHandler:       mcpHandler,
	}
}
