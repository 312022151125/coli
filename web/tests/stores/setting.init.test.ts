// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import { beforeEach, describe, expect, it, vi } from 'vitest'
import { createPinia, setActivePinia } from 'pinia'
import { useSettingStore } from '@/stores/setting'
import { useUserStore } from '@/stores/user'

const apiMocks = vi.hoisted(() => ({
  fetchGetSettings: vi.fn(),
  fetchGetS3Settings: vi.fn(),
  fetchGetOAuth2Settings: vi.fn(),
  fetchGetAllWebhooks: vi.fn(),
  fetchListAccessTokens: vi.fn(),
  fetchGetBackupScheduleSetting: vi.fn(),
  fetchGetAgentSettings: vi.fn(),
  fetchGetAgentInfo: vi.fn(),
  fetchHelloEch0: vi.fn(),
}))

vi.mock('@/service/api', () => apiMocks)

describe('settingStore.init', () => {
  beforeEach(() => {
    setActivePinia(createPinia())
    vi.clearAllMocks()

    apiMocks.fetchGetSettings.mockResolvedValue({
      code: 1,
      data: {
        site_title: 'coli.dev',
        server_logo: '/timeline-mark.svg',
        server_name: 'coli.dev',
        server_url: '',
        allow_register: true,
        default_locale: 'vi-VN',
        ICP_number: '',
        footer_content: '',
        footer_link: '',
        meting_api: '',
        custom_css: '',
        custom_js: '',
      },
    })
    apiMocks.fetchGetS3Settings.mockResolvedValue({ code: 1, data: { enable: false } })
    apiMocks.fetchGetAgentInfo.mockResolvedValue({
      code: 1,
      data: { enable: false, model: '', protocol: 'openai' },
    })
    apiMocks.fetchHelloEch0.mockResolvedValue({ code: 1, data: {} })
  })

  it('匿名初始化时不请求 S3 设置', async () => {
    const userStore = useUserStore()
    userStore.user = null

    const settingStore = useSettingStore()
    await settingStore.init()

    expect(apiMocks.fetchGetS3Settings).not.toHaveBeenCalled()
  })

  it('管理员初始化时请求 S3 设置', async () => {
    const userStore = useUserStore()
    userStore.user = {
      id: 'admin-1',
      username: 'admin',
      is_admin: true,
      locale: 'vi-VN',
    }

    const settingStore = useSettingStore()
    await settingStore.init()

    expect(apiMocks.fetchGetS3Settings).toHaveBeenCalledTimes(1)
  })
})
