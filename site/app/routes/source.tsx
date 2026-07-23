// SPDX-License-Identifier: AGPL-3.0-or-later
// Copyright (C) 2025-2026 lin-snow

import type { Route } from "./+types/source";
import { Link } from "react-router";
import { siteUrl, absoluteUrl } from "../site";

export function meta(_: Route.MetaArgs) {
  const title = "Source code — coli.dev";
  const description =
    "Corresponding source code and license information for coli.dev under AGPL-3.0-or-later.";
  const canonical = absoluteUrl("/source");

  return [
    { title },
    { name: "description", content: description },
    { property: "og:title", content: title },
    { property: "og:description", content: description },
    { property: "og:url", content: canonical },
    { property: "og:type", content: "website" },
    { name: "twitter:card", content: "summary" },
    { name: "twitter:title", content: title },
    { name: "twitter:description", content: description },
  ];
}

export default function SourcePage() {
  return (
    <div className="mx-auto max-w-[42rem] px-4 py-12 text-sand-12 sm:px-6">
      <header className="mb-8 border-b border-sand-6 pb-6">
        <Link
          to="/"
          className="text-sm font-medium text-sand-11 hover:text-sand-12"
        >
          ← Back to home
        </Link>
        <h1 className="mt-4 font-sans text-2xl font-semibold leading-tight text-sand-12">
          Source code — coli.dev
        </h1>
        <p className="mt-2 text-sm text-sand-11">
          Complete corresponding source code availability and licensing.
        </p>
      </header>

      <main className="space-y-8 text-sm leading-relaxed text-sand-11">
        <section className="space-y-3">
          <h2 id="source" className="font-sans text-base font-semibold text-sand-12">
            Source archive
          </h2>
          <p>
            In accordance with the GNU Affero General Public License (AGPL-3.0-or-later), complete corresponding source code for this service release is available for download:
          </p>
          <div className="mt-3 rounded-lg border border-sand-6 bg-sand-2/50 p-4">
            <a
              href="https://coli.dev/source/source.tar.gz"
              className="inline-flex items-center gap-2 font-medium text-sand-12 underline decoration-sand-6 underline-offset-4 hover:decoration-sand-12"
            >
              📦 Download source.tar.gz
            </a>
            <p className="mt-2 text-xs text-sand-11">
              Version-matched source code tarball containing backend Go modules, web frontend SPA, Hub aggregator, and site documentation.
            </p>
          </div>
        </section>

        <section className="space-y-3">
          <h2 id="license" className="font-sans text-base font-semibold text-sand-12">
            License & Rights
          </h2>
          <p>
            This project is licensed under the{" "}
            <strong className="text-sand-12 font-medium">
              GNU Affero General Public License v3.0 or later (AGPL-3.0-or-later)
            </strong>
            .
          </p>
          <p>
            You are free to run, study, modify, and redistribute this software or host your own independent instances, provided that any modified versions hosted as a network service also make their corresponding source code available under AGPL-3.0.
          </p>
        </section>
      </main>

      <footer className="mt-12 border-t border-sand-6 pt-6 text-xs text-sand-11">
        <p>© 2025-2026 coli.dev. Distributed under AGPL-3.0-or-later.</p>
      </footer>
    </div>
  );
}
