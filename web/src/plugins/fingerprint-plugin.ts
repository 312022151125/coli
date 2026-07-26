// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import type { Plugin } from 'vite'

export function fingerprintPlugin(): Plugin {
  const banner = `/*! coli.dev | AGPL-3.0-or-later */`

  return {
    name: 'coli-fingerprint',
    apply: 'build',
    generateBundle(_, bundle) {
      for (const chunk of Object.values(bundle)) {
        if (chunk.type === 'chunk' && chunk.isEntry) {
          chunk.code = banner + '\n' + chunk.code
        }
      }
    },
  }
}
