import type { Metadata } from "next";
import Link from "next/link";

export const metadata: Metadata = {
  title: "Blog",
  description:
    "Updates, guides, and tips from the Golos team — the free, open-source voice-to-text CLI for macOS.",
  alternates: { canonical: "https://golos.sh/blog" },
  openGraph: {
    title: "Blog | Golos",
    description:
      "Updates, guides, and tips from the Golos team.",
    url: "https://golos.sh/blog",
  },
};

const posts = [
  {
    href: "/vs/wisprflow",
    title: "Golos vs WisprFlow — Free Open-Source Alternative",
    date: "Feb 26, 2026",
    summary:
      "WisprFlow charges $8-24/mo for voice-to-text. Golos does the same thing for free. Here's how they compare.",
  },
  {
    href: "/vs/superwhisper",
    title: "Golos vs SuperWhisper — Free Open-Source Alternative",
    date: "Feb 26, 2026",
    summary:
      "SuperWhisper uses local Whisper models for dictation. Golos takes a different approach with cloud-based Deepgram. Here's the full comparison.",
  },
  {
    href: "/vs/dragon",
    title: "Golos vs Dragon Naturally Speaking — Free Alternative",
    date: "Feb 26, 2026",
    summary:
      "Dragon has been the dictation standard for decades, but it's expensive and discontinued on Mac. Golos is a free, modern alternative.",
  },
];

export default function Blog() {
  return (
    <div className="space-y-12">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Blog</h1>
        <p className="mt-2 text-neutral-400">
          Updates, guides, and tips.
        </p>
      </div>

      <div className="space-y-6">
        {posts.map((post) => (
          <Link
            key={post.href}
            href={post.href}
            className="block space-y-1 rounded-lg border border-neutral-800 p-5 hover:border-neutral-600 transition-colors"
          >
            <time className="text-xs text-neutral-500">{post.date}</time>
            <h2 className="text-lg font-semibold">{post.title}</h2>
            <p className="text-sm text-neutral-400">{post.summary}</p>
          </Link>
        ))}
      </div>
    </div>
  );
}
