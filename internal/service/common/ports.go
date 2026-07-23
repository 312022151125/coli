// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package service

import (
	"context"

	"github.com/gin-gonic/gin"
	commonModel "github.com/312022151125/coli/internal/model/common"
	echoModel "github.com/312022151125/coli/internal/model/echo"
	userModel "github.com/312022151125/coli/internal/model/user"
)

type Service interface {
	CommonGetUserByUserId(ctx context.Context, userId string) (userModel.User, error)
	GetOwner() (userModel.User, error)
	GetHeatMap(timezone string) ([]commonModel.Heatmap, error)
	GenerateRSS(ctx *gin.Context) (string, error)
	GetWebsiteTitle(websiteURL string) (string, error)
}

type CommonRepository interface {
	GetUserByUserId(ctx context.Context, id string) (userModel.User, error)
	GetOwner(ctx context.Context) (userModel.User, error)
	GetAllEchos(ctx context.Context, showPrivate bool) ([]echoModel.Echo, error)
	GetHeatMap(ctx context.Context, startTime, endTime int64) ([]int64, error)
	TrackRSSCacheKey(cacheKey string)
}
