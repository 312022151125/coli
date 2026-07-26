// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package migration_test

import (
	"fmt"
	"testing"

	"github.com/312022151125/coli/internal/database"
	dbMigration "github.com/312022151125/coli/internal/database/migration"
	commonModel "github.com/312022151125/coli/internal/model/common"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestEchoKindBackfillMigrator_BackfillsEmptyKindAndMarksDone(t *testing.T) {
	dsn := fmt.Sprintf("file:%s?mode=memory&cache=shared", t.Name())
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		t.Fatalf("open sqlite failed: %v", err)
	}
	database.SetDB(db)
	if err := database.MigrateDB(); err != nil {
		t.Fatalf("migrate db failed: %v", err)
	}

	// 绕过 GORM column default，模拟 schema 变更前遗留的空 kind 行
	// （schema 已有 NOT NULL 约束，NULL 值在当前 schema 下已不可能插入；
	// 空字符串 '' 是唯一现实的历史遗留场景，对应旧版本代码未写入 kind 时的默认值）。
	for _, stmt := range []string{
		`INSERT INTO echos (id, content, kind, user_id, private, created_at) VALUES ('e-empty', 'legacy empty', '', 'u1', false, 100)`,
		`INSERT INTO echos (id, content, kind, user_id, private, created_at) VALUES ('e-kept', 'already project', 'project', 'u1', false, 102)`,
	} {
		if err := db.Exec(stmt).Error; err != nil {
			t.Fatalf("insert echo failed: %v", err)
		}
	}

	dbMigration.Migrate(
		db,
		dbMigration.WithStopOnError(),
		dbMigration.WithMigrators(dbMigration.NewEchoKindBackfillMigrator()),
	)

	rows := map[string]string{}
	var results []struct {
		ID   string
		Kind string
	}
	if err := db.Raw("SELECT id, kind FROM echos ORDER BY id").Scan(&results).Error; err != nil {
		t.Fatalf("query echos failed: %v", err)
	}
	for _, r := range results {
		rows[r.ID] = r.Kind
	}

	if rows["e-empty"] != "note" {
		t.Fatalf("expected e-empty backfilled to note, got %q", rows["e-empty"])
	}
	if rows["e-kept"] != "project" {
		t.Fatalf("expected e-kept to stay project, got %q", rows["e-kept"])
	}

	var marker commonModel.KeyValue
	if err := db.Where("key = ?", commonModel.EchoKindBackfilledKey).First(&marker).Error; err != nil {
		t.Fatalf("expected migrator marker, got err: %v", err)
	}
}
