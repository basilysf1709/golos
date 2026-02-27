import type { Metadata } from "next";
import Link from "next/link";
import { CopyBlock } from "../../copy-block";

export const metadata: Metadata = {
  title: "Golos vs WisprFlow — Free Open-Source Alternative",
  description:
    "Compare Golos and WisprFlow for voice-to-text dictation on macOS. Golos is a free, open-source WisprFlow alternative with push-to-talk speech recognition.",
  keywords: [
    "wispr flow alternative",
    "wispr flow alternatives",
    "wisprflow alternative",
    "wispr flow free",
    "wispr flow cost",
    "wispr flow pricing",
    "wispr flow reviews",
    "dictation software mac",
    "voice to text mac",
    "speech to text mac",
    "free dictation software",
  ],
  alternates: { canonical: "https://golos.sh/vs/wisprflow" },
  openGraph: {
    title: "Golos vs WisprFlow — Free Open-Source Alternative",
    description:
      "Compare Golos and WisprFlow for voice-to-text dictation. Golos is free and open-source.",
    url: "https://golos.sh/vs/wisprflow",
  },
};

function Check() {
  return <span className="text-green-400">Yes</span>;
}
function Cross() {
  return <span className="text-neutral-500">No</span>;
}
function Free() {
  return <span className="text-green-400">Free</span>;
}

export default function VsWisprFlow() {
  return (
    <>
      <div>
        <p className="text-sm font-medium text-green-400">Comparison</p>
        <h1 className="mt-2 text-3xl font-bold tracking-tight sm:text-4xl">
          Golos vs WisprFlow
        </h1>
        <p className="mt-4 text-neutral-400">
          WisprFlow is a popular voice-to-text tool for macOS. Golos is a free,
          open-source alternative that does the same thing — hold a hotkey,
          speak, and your words get pasted — without the subscription.
        </p>
      </div>

      {/* Comparison table */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Feature Comparison</h2>
        <div className="overflow-x-auto rounded-lg border border-neutral-800">
          <table className="w-full text-left text-sm">
            <thead>
              <tr className="border-b border-neutral-700 bg-neutral-900 text-xs uppercase tracking-wider text-neutral-500">
                <th className="px-4 py-3">Feature</th>
                <th className="px-4 py-3">Golos</th>
                <th className="px-4 py-3">WisprFlow</th>
              </tr>
            </thead>
            <tbody className="text-neutral-300">
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Price</td>
                <td className="px-4 py-3"><Free /></td>
                <td className="px-4 py-3 text-neutral-400">$8-24/mo</td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Open source</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Push-to-talk</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Paste to any app</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">macOS support</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Custom dictionary</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Background mode</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Configurable hotkey</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">CLI interface</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">stdout output</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">GUI app</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">AI rewriting / commands</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr>
                <td className="px-4 py-3 text-neutral-400">Windows / Linux</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      {/* Why Golos */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Why choose Golos over WisprFlow?</h2>
        <div className="space-y-6">
          <div>
            <h3 className="font-medium text-neutral-200">Completely free</h3>
            <p className="mt-1 text-sm text-neutral-400">
              WisprFlow charges $8-24/month. Golos is free forever. You only pay
              for your Deepgram API usage, which has a generous free tier for
              personal use.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Open source</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos is fully open source on GitHub. You can audit the code, contribute
              features, or fork it. Your voice data flows directly to Deepgram — no
              middleman server.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Built for developers</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos is a CLI tool that fits into developer workflows. Pipe output to
              stdout, run it in the background, configure it via TOML files, and
              manage it from the terminal. Use it with Claude Code, Cursor, or any
              editor.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Lightweight and fast</h3>
            <p className="mt-1 text-sm text-neutral-400">
              A single binary with no Electron, no GUI overhead. Golos uses minimal
              resources and starts instantly.
            </p>
          </div>
        </div>
      </section>

      {/* When WisprFlow might be better */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">When WisprFlow might be better</h2>
        <p className="text-sm text-neutral-400">
          WisprFlow has a polished GUI, AI-powered rewriting, and a command mode
          that lets you give instructions to modify text after dictation. If you
          want a point-and-click experience with AI features built in,
          WisprFlow is a solid product. Golos is for people who want a simple,
          free, no-frills CLI tool that does one thing well.
        </p>
      </section>

      {/* Get started */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Switch to Golos</h2>
        <p className="text-sm text-neutral-400">
          Install in one command and start dictating in under a minute:
        </p>
        <CopyBlock text="curl -fsSL https://raw.githubusercontent.com/basilysf1709/golos/main/install.sh | bash" />
        <div className="flex gap-4 text-sm">
          <Link
            href="/docs"
            className="rounded-lg border border-neutral-700 px-4 py-2 font-medium hover:border-neutral-500 transition-colors"
          >
            Read the Docs
          </Link>
          <a
            href="https://github.com/basilysf1709/golos"
            target="_blank"
            rel="noopener noreferrer"
            className="rounded-lg border border-neutral-700 px-4 py-2 font-medium hover:border-neutral-500 transition-colors"
          >
            View on GitHub
          </a>
        </div>
      </section>
    </>
  );
}
