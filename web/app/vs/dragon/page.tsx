import type { Metadata } from "next";
import Link from "next/link";
import { CopyBlock } from "../../copy-block";

export const metadata: Metadata = {
  title: "Golos vs Dragon Naturally Speaking — Free Alternative",
  description:
    "Looking for a free Dragon Naturally Speaking alternative? Golos is an open-source voice-to-text CLI for macOS with push-to-talk dictation powered by Deepgram.",
  keywords: [
    "dragon naturally speaking alternative",
    "dragon naturally speaking alternatives",
    "alternative to dragon naturally speaking",
    "alternatives to dragon naturally speaking",
    "alternatives to dragon dictate",
    "dragon professional dictation software",
    "nuance dragon alternative",
    "free dictation software",
    "dictation software mac",
    "medical dictation software mac",
    "best dictation software",
    "voice recognition software",
    "speech to text mac",
  ],
  alternates: { canonical: "https://golos.sh/vs/dragon" },
  openGraph: {
    title: "Golos vs Dragon Naturally Speaking — Free Alternative",
    description:
      "Golos is a free, open-source alternative to Dragon Naturally Speaking for macOS.",
    url: "https://golos.sh/vs/dragon",
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

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "FAQPage",
  mainEntity: [
    {
      "@type": "Question",
      name: "Is Golos a free alternative to Dragon Naturally Speaking?",
      acceptedAnswer: {
        "@type": "Answer",
        text: "Yes. Golos is completely free and open source. Dragon Professional costs $200-500. Golos works on macOS with push-to-talk dictation — no voice training required.",
      },
    },
    {
      "@type": "Question",
      name: "Does Dragon Naturally Speaking work on Mac?",
      acceptedAnswer: {
        "@type": "Answer",
        text: "Nuance discontinued Dragon Dictate for Mac. Golos is a free, open-source alternative that is built for macOS from the ground up. Install with one command and start dictating in under a minute.",
      },
    },
  ],
};

export default function VsDragon() {
  return (
    <>
      <script
        type="application/ld+json"
        dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
      />
      <div>
        <p className="text-sm font-medium text-green-400">Comparison</p>
        <h1 className="mt-2 text-3xl font-bold tracking-tight sm:text-4xl">
          Golos vs Dragon Naturally Speaking
        </h1>
        <p className="mt-4 text-neutral-400">
          Dragon Naturally Speaking (now Nuance Dragon) has been the gold
          standard in dictation software for decades. But it&apos;s expensive,
          Windows-focused, and overkill for many users. Golos is a free,
          open-source alternative for macOS that gets the job done.
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
                <th className="px-4 py-3">Dragon</th>
              </tr>
            </thead>
            <tbody className="text-neutral-300">
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Price</td>
                <td className="px-4 py-3"><Free /></td>
                <td className="px-4 py-3 text-neutral-400">$200-500 (one-time)</td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Open source</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">macOS support</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3 text-neutral-400">Dragon Dictate (discontinued)</td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Windows support</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
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
                <td className="px-4 py-3 text-neutral-400">Custom dictionary</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Setup time</td>
                <td className="px-4 py-3 text-neutral-400">~1 minute</td>
                <td className="px-4 py-3 text-neutral-400">30+ minutes</td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Voice training required</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">CLI interface</td>
                <td className="px-4 py-3"><Check /></td>
                <td className="px-4 py-3"><Cross /></td>
              </tr>
              <tr className="border-b border-neutral-800">
                <td className="px-4 py-3 text-neutral-400">Voice commands / macros</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
              <tr>
                <td className="px-4 py-3 text-neutral-400">Medical vocabulary</td>
                <td className="px-4 py-3"><Cross /></td>
                <td className="px-4 py-3"><Check /></td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>

      {/* Why Golos */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Why choose Golos over Dragon?</h2>
        <div className="space-y-6">
          <div>
            <h3 className="font-medium text-neutral-200">Actually works on macOS</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Nuance discontinued Dragon Dictate for Mac years ago. If you&apos;re
              on macOS, Dragon isn&apos;t even an option anymore. Golos is built
              for macOS from the ground up.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Free and open source</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Dragon Professional costs $500. Golos is free. You bring your own
              Deepgram API key (free tier available), and the code is fully open
              on GitHub.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">No voice training</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Dragon requires you to read passages aloud to train its model.
              Golos works immediately with Deepgram&apos;s Nova-3 engine — no
              training, no calibration.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Install in 60 seconds</h3>
            <p className="mt-1 text-sm text-neutral-400">
              One curl command and you&apos;re ready. No installers, no license
              keys, no account creation, no registration.
            </p>
          </div>
          <div>
            <h3 className="font-medium text-neutral-200">Built for developers</h3>
            <p className="mt-1 text-sm text-neutral-400">
              Golos is a CLI tool. Pipe output to stdout, run it as a background
              daemon, configure via TOML, and integrate it into your terminal
              workflow with Cursor, Claude Code, or any editor.
            </p>
          </div>
        </div>
      </section>

      {/* When Dragon might be better */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">When Dragon might be better</h2>
        <p className="text-sm text-neutral-400">
          Dragon is still the best choice for specialized professional use cases
          — particularly medical and legal dictation with industry-specific
          vocabulary. It also has deep voice command and macro support that lets
          you control your computer by voice. If you&apos;re on Windows and need
          enterprise-grade dictation with compliance certifications, Dragon is
          hard to beat.
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
