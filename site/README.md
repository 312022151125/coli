# coli.dev — source offer & privacy page

Minimal client-side SPA (React Router 7, React 19, Vite, Tailwind CSS v4) serving the AGPL-3.0 corresponding-source offer required by `SourceURL` (`internal/version`) and the privacy policy. No SSR.

## What lives here

| Path                     | Purpose                                                             |
| ------------------------ | -------------------------------------------------------------------- |
| `app/routes/source.tsx`  | `/` and `/source` — AGPL §13 corresponding-source offer            |
| `app/routes/privacy.tsx` | `/privacy` — privacy policy (Markdown-backed)                      |
| `app/docs/`              | Markdown rendering helpers (`MarkdownDoc.tsx`, `registry.ts`)       |
| `content/privacy.md`     | Privacy policy source                                              |
| `public/`                | Static assets (`logo.svg`, `_redirects` for SPA fallback on static hosts) |

## Prerequisites

- **Node.js** (LTS recommended)
- **pnpm** (workspace uses `pnpm` at the repo root)

## Commands

```bash
pnpm install
pnpm dev          # dev server (Vite; default http://localhost:5173)
pnpm build        # output: build/client/ (static assets)
pnpm start        # serve build/client locally (serve -s)
pnpm typecheck    # react-router typegen + tsc
pnpm lint
pnpm format       # or pnpm format:check
```

## Environment

| Variable        | Purpose                                                                       |
| --------------- | ------------------------------------------------------------------------------ |
| `VITE_SITE_URL` | Canonical site origin (no trailing slash) for OG URLs. Default: `https://coli.dev`. |

When deploying to a custom domain, align `VITE_SITE_URL` and update `public/sitemap.xml` / `public/robots.txt` if needed.

## Production build

- Run `pnpm build`.
- Deploy the contents of `build/client/` to any static host.
- Configure **SPA fallback** to `index.html` for client-side routes (this repo includes `public/_redirects` for Netlify-style hosts).

## License

Content and code follow the same terms as the parent repository unless noted otherwise.
