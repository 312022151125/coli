// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package service

import "testing"

// TestNormalizeEchoKind 覆盖 kind 归一化的全部分支：trim/lowercase、空串归 note、
// 四个合法值原样通过、非法值直接拒绝（不静默落回 note）。
func TestNormalizeEchoKind(t *testing.T) {
	cases := []struct {
		name    string
		in      string
		want    string
		wantErr bool
	}{
		{"empty defaults to note", "", "note", false},
		{"whitespace only defaults to note", "   ", "note", false},
		{"uppercase normalized", "PROJECT", "project", false},
		{"mixed case with whitespace", "  Startup  ", "startup", false},
		{"note passthrough", "note", "note", false},
		{"idea passthrough", "idea", "idea", false},
		{"todo passthrough", "todo", "todo", false},
		{"invalid rejected", "meme", "", true},
		{"github historical value rejected", "githubproj", "", true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := normalizeEchoKind(tc.in)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("normalizeEchoKind(%q) expected error, got nil", tc.in)
				}
				return
			}
			if err != nil {
				t.Fatalf("normalizeEchoKind(%q) unexpected error: %v", tc.in, err)
			}
			if got != tc.want {
				t.Fatalf("normalizeEchoKind(%q) = %q, want %q", tc.in, got, tc.want)
			}
		})
	}
}
