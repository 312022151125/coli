// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import type { RouteLocationNormalized } from 'vue-router'

/** Public origin for canonical, Open Graph, and sitemap (override via `VITE_HUB_SITE_ORIGIN`). */
export const HUB_SITE_ORIGIN = (
  import.meta.env.VITE_HUB_SITE_ORIGIN ?? 'https://coli.dev/hub'
).replace(/\/+$/, '')

const DEFAULT_DESCRIPTION =
  'Browse public posts from independent instances in one shared stream.'

const KEYWORDS =
  'microblog, timeline, self-hosted, open source, RSS, feed, aggregate, decentralized'

const ROUTES = {
  home: {
    title: 'Public voices, gathered in one place.',
    description: DEFAULT_DESCRIPTION,
    path: '/',
    jsonLdType: 'WebSite',
  },
  explore: {
    title: 'Explore the public timeline — coli.dev Hub',
    description:
      'Browse public posts from independent instances in one shared stream.',
    path: '/explore',
    jsonLdType: 'CollectionPage',
  },
} as const

const OG_IMAGE = `${HUB_SITE_ORIGIN}/android-chrome-512x512.png`

function ensureMeta(attr: 'name' | 'property', key: string, content: string) {
  const selector = attr === 'name' ? `meta[name="${key}"]` : `meta[property="${key}"]`
  let el = document.head.querySelector<HTMLMetaElement>(selector)
  if (!el) {
    el = document.createElement('meta')
    el.setAttribute(attr, key)
    document.head.appendChild(el)
  }
  el.setAttribute('content', content)
}

function ensureLinkCanonical(href: string) {
  let el = document.head.querySelector<HTMLLinkElement>('link[rel="canonical"]')
  if (!el) {
    el = document.createElement('link')
    el.setAttribute('rel', 'canonical')
    document.head.appendChild(el)
  }
  el.setAttribute('href', href)
}

function ensureJsonLd(content: string) {
  let el = document.head.querySelector<HTMLScriptElement>('script[data-hub-jsonld="route"]')
  if (!el) {
    el = document.createElement('script')
    el.type = 'application/ld+json'
    el.dataset.hubJsonld = 'route'
    document.head.appendChild(el)
  }
  el.textContent = content
}

/**
 * Updates document title and social / SEO tags after client navigation (SPA).
 * Initial values also exist in `index.html` for crawlers and first paint.
 */
export function applyHubRouteMeta(to: RouteLocationNormalized): void {
  const isExplore = to.name === 'explore' || to.path.startsWith('/explore')
  const cfg = isExplore ? ROUTES.explore : ROUTES.home
  const url = `${HUB_SITE_ORIGIN}${to.path === '/' ? '/' : to.path}`

  document.title = cfg.title

  ensureMeta('name', 'description', cfg.description)
  ensureMeta('name', 'keywords', KEYWORDS)

  ensureMeta('property', 'og:type', 'website')
  ensureMeta('property', 'og:site_name', 'coli.dev Hub')
  ensureMeta('property', 'og:title', cfg.title)
  ensureMeta('property', 'og:description', cfg.description)
  ensureMeta('property', 'og:url', url)
  ensureMeta('property', 'og:locale', 'en_US')
  ensureMeta('property', 'og:image', OG_IMAGE)

  ensureMeta('name', 'twitter:card', 'summary_large_image')
  ensureMeta('name', 'twitter:title', cfg.title)
  ensureMeta('name', 'twitter:description', cfg.description)
  ensureMeta('name', 'twitter:image', OG_IMAGE)

  ensureLinkCanonical(url)
  ensureJsonLd(
    JSON.stringify({
      '@context': 'https://schema.org',
      '@type': cfg.jsonLdType,
      name: cfg.title,
      description: cfg.description,
      url,
      isPartOf: {
        '@type': 'WebSite',
        name: 'coli.dev Hub',
        url: `${HUB_SITE_ORIGIN}/`,
      },
      image: OG_IMAGE,
      inLanguage: 'en-US',
    }),
  )
}
