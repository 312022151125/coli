// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import { describe, expect, it, vi } from 'vitest'
import { mount } from '@vue/test-utils'
import { createPinia, setActivePinia } from 'pinia'
import { createRouter, createWebHistory } from 'vue-router'
import type * as VueI18n from 'vue-i18n'
import AuthPage from '@/views/auth/modules/AuthPage.vue'
vi.mock('@/service/api', () => ({
  fetchGetOAuth2Status: vi.fn().mockResolvedValue({ code: 1, data: { oauth_ready: false, provider: '' } }),
  fetchGetPasskeyStatus: vi.fn().mockResolvedValue({ code: 1, data: { passkey_ready: false } }),
  fetchPasskeyLoginBegin: vi.fn(),
  fetchPasskeyLoginFinish: vi.fn(),
}))

vi.mock('vue-i18n', async (importOriginal) => {
  const actual = await importOriginal<typeof VueI18n>()
  return {
    ...actual,
    useI18n: () => ({ t: (key: string) => key }),
  }
})

function buildTestRouter() {
  return createRouter({
    history: createWebHistory(),
    routes: [
      { path: '/auth', name: 'auth', component: AuthPage },
      { path: '/register', name: 'register', component: AuthPage },
    ],
  })
}

describe('AuthPage register mode', () => {
  it('starts in register mode when mounted on the /register route', async () => {
    setActivePinia(createPinia())
    const router = buildTestRouter()
    await router.push('/register')
    await router.isReady()

    const wrapper = mount(AuthPage, {
      global: { plugins: [router] },
    })
    await wrapper.vm.$nextTick()

    expect(wrapper.find('h2').text()).toBe('authPage.register')
  })

  it('switches mode reactively when navigating /auth -> /register without remount', async () => {
    setActivePinia(createPinia())
    const router = buildTestRouter()
    await router.push('/auth')
    await router.isReady()

    const wrapper = mount(AuthPage, {
      global: { plugins: [router] },
    })
    await wrapper.vm.$nextTick()
    expect(wrapper.find('h2').text()).toBe('authPage.login')

    await router.push('/register')
    await wrapper.vm.$nextTick()

    expect(wrapper.find('h2').text()).toBe('authPage.register')
  })
})
