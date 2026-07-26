// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import { type RouteConfig, index, route } from "@react-router/dev/routes";

export default [
  index("routes/source.tsx", { id: "source-index" }),
  route("privacy", "routes/privacy.tsx"),
  route("source", "routes/source.tsx"),
] satisfies RouteConfig;
