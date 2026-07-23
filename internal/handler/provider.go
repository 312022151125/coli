// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package handler

import (
	"github.com/google/wire"
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

var (
	WebSet       = wire.NewSet(webHandler.NewWebHandler)
	UserSet      = wire.NewSet(userHandler.NewUserHandler)
	AuthSet      = wire.NewSet(authHandler.NewAuthHandler)
	EchoSet      = wire.NewSet(echoHandler.NewEchoHandler)
	FileSet      = wire.NewSet(fileHandler.NewFileHandler)
	CommentSet   = wire.NewSet(commentHandler.NewCommentHandler)
	InitSet      = wire.NewSet(initHandler.NewInitHandler)
	CommonSet    = wire.NewSet(commonHandler.NewCommonHandler)
	SettingSet   = wire.NewSet(settingHandler.NewSettingHandler)
	ConnectSet   = wire.NewSet(connectHandler.NewConnectHandler)
	DashboardSet = wire.NewSet(dashboardHandler.NewDashboardHandler)
	CopilotSet   = wire.NewSet(copilotHandler.NewCopilotHandler)
	EmbeddingSet = wire.NewSet(embeddingHandler.NewEmbeddingHandler)
	MigrationSet = wire.NewSet(migratorHandler.NewMigrationHandler)
	MCPSet       = wire.NewSet(mcp.NewHandler)
)
