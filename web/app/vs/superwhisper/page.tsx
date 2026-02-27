import type { Metadata } from "next";
import Link from "next/link";
import { CopyBlock } from "../../copy-block";

export const metadata: Metadata = {
  title: "Golos vs SuperWhisper — Free Open-Source Alternative",
  description:
    "Compare Golos and SuperWhisper for speech-to-text on macOS. Golos is a free, open-source SuperWhisper alternative with push-to-talk dictation.",
  keywords: [
    "superwhisper alternative",
    "superwhisper alternatives",
    "superwhisper pricing",
    "superwhisper free",
    "superwhisper vs macwhisper",
    "macwhisper vs superwhisper",
    "superwhisper for windows",
    "betterdictation",
    "speech to text mac",
    "dictation software mac",
    "whisper app replacement",
    "free dictation software",
  ],
  alternates: { canonical: "https://golos.sh/vs/superwhisper" },
  openGraph: {
    title: "Golos vs SuperWhisper — Free Open-Source Alternative",
    description:
      "Compare Golos and SuperWhisper for speech-to-text. Golos is free and open-source.",
    url: "https://golos.sh/vs/superwhisper",
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

export default function VsSuperWhisper() {
  return (
    <>
      <div>
        <p className="text-sm font-medium text-green-400">Comparison</p>
        <h1 className="mt-2 text-3xl font-bold tracking-tight sm:text-4xl">
          Golos vs SuperWhisper
        </h1>
        <p className="mt-4 text-neutral-400">
          SuperWhisper is a macOS dictation app powered by OpenAI Whisper. Golos
          is a free, open-source alternative that gives you push-to-talk
          speech-to-text from the terminal — no subscription required.
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
                <th className="px-4 py-3">SuperWhisper</th>
              </tr>
            </thead>
            <tbody className="text-neutral-300">
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Price</td>
                <td className="px-4 py-3"><Free /></td>
                <td className="px-4 py-3 text-neutral-400">$8-16/mo</td>
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
                <td className="px-4 py-3 text-neutral-400">STT engine</td>
                <td className="px-4 py-3 text-neutral-400">Deepgram Nova-3</td>
                <td className="px-4 py-3 text-neutral-400">OpenAI Whisper</td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Processing</td>
                <td className="px-4 py-3 text-neutral-400">Cloud (Deepgram)</td>
                <td className="px-4 py-3 text-neutral-400">Local + Cloud</td>
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
                <td className="px-4 py-3 text-neutral-400">Local-only mode</td>
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
        <h2 className="text-xl font-semibold">Why choose Golos over SuperWhisper?</h2>
        <div className="space-y-6">
          <div>
            <h3 className="font-medium text-neutral-200">Completely free</h3>
            <p className="mt-1 text-sm text-neutral-400">
              SuperWhisper starts at $8/month. Golos is free and open source. You
              bring your own Deepgram API key, which has a generous free tier
              that covers most personal usage.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Developer-first workflow</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos is a CLI tool. Run it in the background, pipe output to
              stdout, configure via TOML files, and integrate it into scripts.
              Perfect for developers using Cursor, Claude Code, or terminal-based
              editors.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Faster transcription</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos uses Deepgram Nova-3 for cloud-based transcription, which
              returns results near-instantly. SuperWhisper&apos;s local mode is
              private but can be slower, especially on older hardware.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Lightweight</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos is a single binary — no large AI models downloaded to your
              machine. SuperWhisper&apos;s local Whisper models can take several
              gigabytes of disk space.
            </p>
          </div>
        </div>
      </section>

      {/* When SuperWhisper might be better */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">When SuperWhisper might be better</h2>
        <p className="text-sm text-neutral-400">
          SuperWhisper shines if you need fully offline, local-only transcription
          with no data leaving your machine. It also has a polished GUI with
          visual feedback and doesn&apos;t require any API keys. If privacy is your
          top priority and you don&apos;t mind the subscription, SuperWhisper is a
          solid choice.
        </p>
      </section>

      {/* Get started */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Try Golos for free</h2>
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
