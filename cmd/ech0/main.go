// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

package main

import (
	_ "time/tzdata"

	"github.com/312022151125/coli/cmd"
	"github.com/312022151125/coli/internal/bootstrap"
)

func main() {
	bootstrap.Bootstrap()
	cmd.Execute()
}
