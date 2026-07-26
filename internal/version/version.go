// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

// Package version is the single source of truth for build / release metadata
// of the coli.dev binary: semantic version, git commit, build time, and the
// AGPL corresponding-source offer URL.
//
// Const values are source-controlled (bump in code on release).
// Var values are injected at build time via -ldflags "-X .../version.Commit=...".
// See Makefile / docker/build.Dockerfile / .github/workflows/release.yml.
package version

import "fmt"

const (
	// Version is the current semantic version. Bump on release.
	Version = "5.4.6"

	// SourceURL is the first-party corresponding source page URL, required by
	// AGPL-3.0 §13 for network-deployed use.
	SourceURL = "https://coli.dev/source"
)

// Commit is the short git commit hash, injected at build time.
// Defaults to "unknown" so `go run` / unbranded local builds still compile.
var Commit = "unknown"

// BuildTime is the RFC3339 build timestamp, injected at build time.
// Empty when not injected.
var BuildTime = ""

// Info bundles all build / release metadata. Useful for handlers that want
// to return a single struct instead of hand-spelling every field.
type Info struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildTime string `json:"build_time"`
	SourceURL string `json:"source_url"`
}

// Get returns a snapshot of the current build / release metadata.
func Get() Info {
	return Info{
		Version:   Version,
		Commit:    Commit,
		BuildTime: BuildTime,
		SourceURL: SourceURL,
	}
}

// String returns a human-readable one-line summary, e.g.
// "coli.dev v5.4.6 (commit abc123, built 2026-01-01T00:00:00Z)".
func (i Info) String() string {
	return fmt.Sprintf("coli.dev v%s (commit %s, built %s)", i.Version, i.Commit, i.BuildTime)
}
