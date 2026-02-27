import Image from "next/image";
import Link from "next/link";
import { CopyBlock } from "./copy-block";

function Step({
  number,
  title,
  children,
}: {
  number: number;
  title: string;
  children: React.ReactNode;
}) {
  return (
    <div className="flex gap-4">
      <div className="flex h-8 w-8 shrink-0 items-center justify-center rounded-full border border-neutral-700 text-sm font-medium text-neutral-300">
        {number}
      </div>
      <div>
        <h3 className="font-medium">{title}</h3>
        <div className="mt-1 text-sm text-neutral-400">{children}</div>
      </div>
    </div>
  );
}

export default function Home() {
  return (
    <div className="space-y-16">
      {/* Hero */}
      <section className="flex flex-col items-center text-center">
        <Image
          src="/mascot.png"
          alt="Golos mascot"
          width={120}
          height={120}
          className="mb-6"
          priority
        />
        <h1 className="text-4xl font-bold tracking-tight sm:text-5xl">
          Golos
        </h1>
        <p className="mt-4 max-w-md text-lg text-neutral-400">
          Hold a hotkey, speak, release. Your words get pasted.
        </p>
        <p className="mt-2 text-sm text-neutral-500">
          A free, open-source alternative to WisprFlow. macOS CLI powered by
          Deepgram.
        </p>
      </section>

      {/* Install */}
      <section className="space-y-4">
        <h2 className="text-xl font-semibold">Install</h2>
        <CopyBlock text="curl -fsSL https://raw.githubusercontent.com/basilysf1709/golos/main/install.sh | bash" />
        <p className="text-sm text-neutral-500">
          Requires macOS and Homebrew. The installer handles portaudio and places
          the binary at <code className="text-neutral-300">/usr/local/bin/golos</code>.
        </p>
      </section>

      {/* Quick Start */}
      <section className="space-y-6">
        <h2 className="text-xl font-semibold">Quick Start</h2>
        <div className="space-y-6">
          <Step number={1} title="Install Golos">
            Run the curl command above. It installs portaudio and the golos
            binary.
          </Step>
          <Step number={2} title="Grant Accessibility Permission">
            <span>
              Go to{" "}
              <strong className="text-neutral-200">
                System Settings &rarr; Privacy &amp; Security &rarr;
                Accessibility
              </strong>{" "}
              and enable your terminal app (Terminal, iTerm2, Alacritty, etc.).
            </span>
          </Step>
          <Step number={3} title="Configure your API key">
            <span>
              Run{" "}
              <code className="rounded bg-neutral-800 px-1.5 py-0.5 text-neutral-200">
                golos setup
              </code>{" "}
              to enter your{" "}
              <a
                href="https://console.deepgram.com"
                target="_blank"
                rel="noopener noreferrer"
                className="underline hover:text-neutral-200"
              >
                Deepgram API key
              </a>
              .
            </span>
          </Step>
        </div>
        <div className="mt-6">
          <CopyBlock text="golos" />
          <p className="mt-2 text-sm text-neutral-500">
            Hold <strong className="text-neutral-300">Right Option</strong> to
            record, release to transcribe and paste.
          </p>
        </div>
      </section>

      {/* Links */}
      <section className="flex gap-4 text-sm">
        <Link
          href="/docs"
          className="flex items-center gap-1.5 rounded-lg border border-neutral-700 px-4 py-2 font-medium hover:border-neutral-500 transition-colors"
        >
          Read the Docs
          <svg width="14" height="14" viewBox="0 0 16 16" fill="none" stroke="currentColor" strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M4.5 11.5L11.5 4.5M11.5 4.5H5.5M11.5 4.5V10.5" />
          </svg>
        </Link>
        <a
          href="https://github.com/basilysf1709/golos"
          target="_blank"
          rel="noopener noreferrer"
          className="flex items-center gap-1.5 rounded-lg border border-neutral-700 px-4 py-2 font-medium hover:border-neutral-500 transition-colors"
        >
          <svg width="16" height="16" viewBox="0 0 16 16" fill="currentColor">
            <path d="M8 0C3.58 0 0 3.58 0 8c0 3.54 2.29 6.53 5.47 7.59.4.07.55-.17.55-.38 0-.19-.01-.82-.01-1.49-2.01.37-2.53-.49-2.69-.94-.09-.23-.48-.94-.82-1.13-.28-.15-.68-.52-.01-.53.63-.01 1.08.58 1.23.82.72 1.21 1.87.87 2.33.66.07-.52.28-.87.51-1.07-1.78-.2-3.64-.89-3.64-3.95 0-.87.31-1.59.82-2.15-.08-.2-.36-1.02.08-2.12 0 0 .67-.21 2.2.82.64-.18 1.32-.27 2-.27s1.36.09 2 .27c1.53-1.04 2.2-.82 2.2-.82.44 1.1.16 1.92.08 2.12.51.56.82 1.27.82 2.15 0 3.07-1.87 3.75-3.65 3.95.29.25.54.73.54 1.48 0 1.07-.01 1.93-.01 2.2 0 .21.15.46.55.38A8.01 8.01 0 0016 8c0-4.42-3.58-8-8-8z" />
          </svg>
          GitHub
        </a>
      </section>
    </div>
  );
}
