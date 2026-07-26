// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import { describe, expect, it } from 'vitest'
import router from '@/router/index'

describe('register route', () => {
  it('resolves /register to the register route name using AuthPage', () => {
    const resolved = router.resolve('/register')

    expect(resolved.name).toBe('register')
    expect(resolved.meta.title).toBe('Register')
    expect(resolved.meta.noindex).toBe(true)
  })

  it('registers /register as a distinct public route rendering AuthView', async () => {
    const auth = router.resolve('/auth')
    const register = router.resolve('/register')

    expect(register.path).not.toBe(auth.path)

    const authLoader = auth.matched[0]?.components?.default as () => Promise<{ default: unknown }>
    const registerLoader = register.matched[0]?.components?.default as () => Promise<{
      default: unknown
    }>
    const [authModule, registerModule] = await Promise.all([authLoader(), registerLoader()])

    expect(registerModule.default).toBe(authModule.default)
  })
})
