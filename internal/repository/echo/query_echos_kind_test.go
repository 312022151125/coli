// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package repository

import (
	"testing"

	commonModel "github.com/312022151125/coli/internal/model/common"
	echoModel "github.com/312022151125/coli/internal/model/echo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

// seedEchoWithKind 插入一条带指定 kind 的 echo 行；kind 传空串模拟绕过 GORM 默认值的历史遗留行
// （现网场景：schema 变更前创建、从未被 backfill migrator 处理过的行）。
func seedEchoWithKind(t *testing.T, db *gorm.DB, id, kind string) {
	t.Helper()
	require.NoError(t, db.Exec(
		`INSERT INTO echos (id, content, kind, user_id, private, created_at) VALUES (?, ?, ?, 'u1', false, ?)`,
		id, "content-"+id, kind, 100,
	).Error)
}

func TestEchoRepository_QueryEchos_KindFilter(t *testing.T) {
	repo, db := newEchoRepo(t)
	seedEchoWithKind(t, db, "e-note", "note")
	seedEchoWithKind(t, db, "e-project", "project")
	seedEchoWithKind(t, db, "e-startup", "startup")
	seedEchoWithKind(t, db, "e-idea", "idea")

	t.Run("single kind filter returns only that kind", func(t *testing.T) {
		echos, total, err := repo.QueryEchos(
			commonModel.EchoQueryDto{Page: 1, PageSize: 10, Kinds: []string{"project"}},
			true,
		)
		require.NoError(t, err)
		assert.Equal(t, int64(1), total)
		require.Len(t, echos, 1)
		assert.Equal(t, "e-project", echos[0].ID)
	})

	t.Run("multi kind filter uses IN semantics", func(t *testing.T) {
		echos, total, err := repo.QueryEchos(
			commonModel.EchoQueryDto{Page: 1, PageSize: 10, Kinds: []string{"startup", "idea"}},
			true,
		)
		require.NoError(t, err)
		assert.Equal(t, int64(2), total)
		require.Len(t, echos, 2)
		ids := echoIDs(echos)
		assert.Contains(t, ids, "e-startup")
		assert.Contains(t, ids, "e-idea")
	})

	t.Run("empty kinds means no filter applied", func(t *testing.T) {
		echos, total, err := repo.QueryEchos(
			commonModel.EchoQueryDto{Page: 1, PageSize: 10},
			true,
		)
		require.NoError(t, err)
		assert.Equal(t, int64(4), total)
		assert.Len(t, echos, 4)
	})
}

// TestEchoRepository_QueryEchos_LegacyEmptyKindTreatedAsNote 覆盖历史遗留空 kind 行：
// SQL 层 COALESCE(NULLIF(...),'note') 让它命中 kinds=[note] 过滤，
// Go 侧读出后也被 normalize 成 "note" 而不是裸露空字符串（不能让旧帖子在前端渲染空徽章）。
func TestEchoRepository_QueryEchos_LegacyEmptyKindTreatedAsNote(t *testing.T) {
	repo, db := newEchoRepo(t)
	seedEchoWithKind(t, db, "e-legacy", "")

	echos, total, err := repo.QueryEchos(
		commonModel.EchoQueryDto{Page: 1, PageSize: 10, Kinds: []string{echoModel.KindNote}},
		true,
	)
	require.NoError(t, err)
	assert.Equal(t, int64(1), total, "legacy empty kind should match the note filter via COALESCE/NULLIF")
	require.Len(t, echos, 1)
	assert.Equal(t, echoModel.KindNote, echos[0].Kind, "empty kind must be normalized to note in the Go response")
}
