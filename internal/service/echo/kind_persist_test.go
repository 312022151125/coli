// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package service_test

import (
	"context"
	"testing"

	echoModel "github.com/312022151125/coli/internal/model/echo"
	echoService "github.com/312022151125/coli/internal/service/echo"
	"github.com/312022151125/coli/internal/test/helpers"
	commonmock "github.com/312022151125/coli/internal/test/mocks/commonmock"
	echomock "github.com/312022151125/coli/internal/test/mocks/echomock"
	filemock "github.com/312022151125/coli/internal/test/mocks/filemock"
	txmock "github.com/312022151125/coli/internal/test/mocks/txmock"
	"github.com/312022151125/coli/pkg/busen"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// TestPostEcho_KindNormalizedAndPersisted 确认 PostEcho 把 kind 归一化后写入 repository，
// 且非法 kind 在事务前直接拒绝，不触达 CreateEcho。
func TestPostEcho_KindNormalizedAndPersisted(t *testing.T) {
	t.Run("uppercase/whitespace kind is normalized before persist", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()
		tx.EXPECT().Run(mock.Anything, mock.Anything).RunAndReturn(runTx).Once()
		repo.EXPECT().GetTagsByNames(mock.Anything, mock.Anything).Return([]*echoModel.Tag{}, nil).Once()

		var created echoModel.Echo
		repo.EXPECT().
			CreateEcho(mock.Anything, mock.Anything).
			Run(func(_ context.Context, e *echoModel.Echo) { created = *e }).
			Return(nil).
			Once()
		repo.EXPECT().InvalidateEchoCaches().Once()
		saved := helpers.NewEcho()
		repo.EXPECT().GetEchosById(mock.Anything, mock.Anything).Return(&saved, nil).Once()
		file.EXPECT().ConfirmTempFiles(mock.Anything, mock.Anything).Return(nil).Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{Content: "hello", Kind: "  PROJECT  "}
		require.NoError(t, svc.PostEcho(helpers.CtxAsUser(adminID), in))

		assert.Equal(t, echoModel.KindProject, created.Kind)
	})

	t.Run("empty kind defaults to note", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()
		tx.EXPECT().Run(mock.Anything, mock.Anything).RunAndReturn(runTx).Once()
		repo.EXPECT().GetTagsByNames(mock.Anything, mock.Anything).Return([]*echoModel.Tag{}, nil).Once()

		var created echoModel.Echo
		repo.EXPECT().
			CreateEcho(mock.Anything, mock.Anything).
			Run(func(_ context.Context, e *echoModel.Echo) { created = *e }).
			Return(nil).
			Once()
		repo.EXPECT().InvalidateEchoCaches().Once()
		saved := helpers.NewEcho()
		repo.EXPECT().GetEchosById(mock.Anything, mock.Anything).Return(&saved, nil).Once()
		file.EXPECT().ConfirmTempFiles(mock.Anything, mock.Anything).Return(nil).Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{Content: "hello"}
		require.NoError(t, svc.PostEcho(helpers.CtxAsUser(adminID), in))

		assert.Equal(t, echoModel.KindNote, created.Kind)
	})

	t.Run("invalid kind rejected before transaction", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{Content: "hello", Kind: "bogus"}
		err := svc.PostEcho(helpers.CtxAsUser(adminID), in)
		require.ErrorContains(t, err, "unsupported echo kind")
		// tx/repo 均不应被触达
	})
}

// TestUpdateEcho_KindPersistedWhenProvidedEmptyLeftUntouched 确认 UpdateEcho 提供 kind 时归一化写入，
// 未提供（空串）时不覆盖 repository 侧既有 kind（避免 MCP/局部更新意外把已有 project 重置为 note）。
func TestUpdateEcho_KindPersistedWhenProvidedEmptyLeftUntouched(t *testing.T) {
	t.Run("provided kind is normalized and persisted", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()
		tx.EXPECT().Run(mock.Anything, mock.Anything).RunAndReturn(runTx).Once()
		repo.EXPECT().GetTagsByNames(mock.Anything, mock.Anything).Return([]*echoModel.Tag{}, nil).Once()

		var updated echoModel.Echo
		repo.EXPECT().
			UpdateEcho(mock.Anything, mock.Anything).
			Run(func(_ context.Context, e *echoModel.Echo) { updated = *e }).
			Return(nil).
			Once()
		repo.EXPECT().InvalidateEchoCaches(echoID).Once()
		file.EXPECT().ConfirmTempFiles(mock.Anything, mock.Anything).Return(nil).Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{ID: echoID, Content: "updated", Kind: "Startup"}
		require.NoError(t, svc.UpdateEcho(helpers.CtxAsUser(adminID), in))

		assert.Equal(t, echoModel.KindStartup, updated.Kind)
	})

	t.Run("empty kind left empty so repository skips the column", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()
		tx.EXPECT().Run(mock.Anything, mock.Anything).RunAndReturn(runTx).Once()
		repo.EXPECT().GetTagsByNames(mock.Anything, mock.Anything).Return([]*echoModel.Tag{}, nil).Once()

		var updated echoModel.Echo
		repo.EXPECT().
			UpdateEcho(mock.Anything, mock.Anything).
			Run(func(_ context.Context, e *echoModel.Echo) { updated = *e }).
			Return(nil).
			Once()
		repo.EXPECT().InvalidateEchoCaches(echoID).Once()
		file.EXPECT().ConfirmTempFiles(mock.Anything, mock.Anything).Return(nil).Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{ID: echoID, Content: "updated"}
		require.NoError(t, svc.UpdateEcho(helpers.CtxAsUser(adminID), in))

		assert.Equal(t, "", updated.Kind, "empty kind must stay empty so repository.UpdateEcho skips overwriting it")
	})

	t.Run("invalid kind rejected before transaction", func(t *testing.T) {
		repo := echomock.NewMockRepository(t)
		common := commonmock.NewMockService(t)
		file := filemock.NewMockService(t)
		tx := txmock.NewMockTransactor(t)
		bus := helpers.NewTestBus(t)

		common.EXPECT().
			CommonGetUserByUserId(mock.Anything, adminID).
			Return(helpers.NewUser(helpers.AsAdmin), nil).
			Once()

		svc := echoService.NewEchoService(tx, common, file, repo, func() *busen.Bus { return bus })
		in := &echoModel.Echo{ID: echoID, Content: "updated", Kind: "bogus"}
		err := svc.UpdateEcho(helpers.CtxAsUser(adminID), in)
		require.ErrorContains(t, err, "unsupported echo kind")
	})
}
