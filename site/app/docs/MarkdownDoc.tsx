import { Link } from "react-router";
import type { ComponentProps, ComponentPropsWithoutRef } from "react";
import { useMemo, useRef } from "react";
import "highlight.js/styles/github.css";
import ReactMarkdown from "react-markdown";
import rehypeHighlight from "rehype-highlight";
import remarkGfm from "remark-gfm";
import type { TocItem } from "./toc";

function DocLink({
  href,
  children,
}: {
  href?: string;
  children?: React.ReactNode;
}) {
  if (!href) return <span>{children}</span>;
  if (/^https?:\/\//i.test(href)) {
    return (
      <a
        href={href}
        target="_blank"
        rel="noreferrer noopener"
        className="text-sand-12 underline decoration-sand-11/45 underline-offset-2"
      >
        {children}
      </a>
    );
  }
  if (href.startsWith("/docs")) {
    return (
      <Link
        to={href}
        className="text-sand-12 underline decoration-sand-11/45 underline-offset-2"
      >
        {children}
      </Link>
    );
  }
  if (href.endsWith(".md")) {
    let path = href.replace(/\.md$/, "").replace(/^\.\//, "");
    while (path.startsWith("../")) path = path.slice(3);
    const to = path === "" || path === "README" ? "/docs" : `/docs/${path}`;
    return (
      <Link
        to={to}
        className="text-sand-12 underline decoration-sand-11/45 underline-offset-2"
      >
        {children}
      </Link>
    );
  }
  return (
    <a
      href={href}
      className="text-sand-12 underline decoration-sand-11/45 underline-offset-2"
    >
      {children}
    </a>
  );
}

/** Smaller sans body; headings in Space Grotesk; code blocks via rehype-highlight + github.css */
const markdownClass =
  "docs-markdown max-w-none font-sans text-[0.875rem] leading-relaxed text-sand-12 [&_h1]:font-sans [&_h1]:text-xl [&_h1]:font-semibold [&_h1]:leading-snug [&_h1]:text-sand-12 [&_h1]:mt-8 [&_h1]:mb-3 [&_h1]:first:mt-0 [&_h2]:font-sans [&_h2]:text-base [&_h2]:font-semibold [&_h2]:leading-snug [&_h2]:text-sand-12 [&_h2]:mt-6 [&_h2]:mb-2 [&_h3]:font-sans [&_h3]:text-[0.9375rem] [&_h3]:font-semibold [&_h3]:text-sand-12 [&_h3]:mt-5 [&_h3]:mb-1.5 [&_p]:my-2.5 [&_ul]:my-2.5 [&_ul]:list-disc [&_ul]:pl-5 [&_ol]:my-2.5 [&_ol]:list-decimal [&_ol]:pl-5 [&_li]:my-1 [&_blockquote]:my-3 [&_blockquote]:border-l-2 [&_blockquote]:border-sand-6 [&_blockquote]:pl-3.5 [&_blockquote]:italic [&_blockquote]:text-sand-11 [&_a]:font-medium [&_a]:text-sand-12 [&_a]:underline [&_a]:decoration-sand-6 [&_a]:underline-offset-4 [&_a]:transition-colors [&_a]:hover:decoration-sand-12 [&_hr]:my-8 [&_hr]:border-dashed [&_hr]:border-sand-6 [&_img]:my-4 [&_img]:rounded-md [&_img]:border [&_img]:border-sand-6 [&_table]:my-4 [&_table]:w-full [&_table]:text-left [&_th]:border-b [&_th]:border-sand-6 [&_th]:pb-2 [&_th]:font-semibold [&_td]:border-b [&_td]:border-sand-6 [&_td]:py-2";

export function MarkdownDoc({
  content,
  tocItems,
}: {
  content: string;
  tocItems?: TocItem[];
}) {
  const tocPtr = useRef(0);
  tocPtr.current = 0;
  const rehypePlugins = useMemo(() => [rehypeHighlight], []);

  const components = useMemo(() => {
    const toc = tocItems ?? [];
    const base: Record<string, React.ComponentType<unknown>> = {
      a: DocLink as React.ComponentType<unknown>,
    };

    if (toc.length === 0) {
      return base;
    }

    const H2 = ({ children, ...props }: ComponentPropsWithoutRef<"h2">) => {
      const item = toc[tocPtr.current];
      if (item?.depth === 2) {
        tocPtr.current++;
        return (
          <h2 id={item.id} {...props}>
            {children}
          </h2>
        );
      }
      return <h2 {...props}>{children}</h2>;
    };

    const H3 = ({ children, ...props }: ComponentPropsWithoutRef<"h3">) => {
      const item = toc[tocPtr.current];
      if (item?.depth === 3) {
        tocPtr.current++;
        return (
          <h3 id={item.id} {...props}>
            {children}
          </h3>
        );
      }
      return <h3 {...props}>{children}</h3>;
    };

    return {
      ...base,
      h2: H2,
      h3: H3,
    };
  }, [tocItems]);

  return (
    <div className={markdownClass}>
      <ReactMarkdown
        remarkPlugins={[remarkGfm]}
        rehypePlugins={
          rehypePlugins as ComponentProps<typeof ReactMarkdown>["rehypePlugins"]
        }
        components={components}
      >
        {content}
      </ReactMarkdown>
    </div>
  );
}
