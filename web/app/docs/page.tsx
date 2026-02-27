import type { Metadata } from "next";
import { CopyBlock } from "../copy-block";

export const metadata: Metadata = {
  title: "Docs",
  description:
    "Complete CLI reference for Golos — commands, flags, dictionary, configuration, hotkeys, and environment variables.",
  alternates: {
    canonical: "https://golos.sh/docs",
  },
  openGraph: {
    title: "Docs | Golos",
    description:
      "Complete CLI reference for Golos — commands, flags, dictionary, configuration, hotkeys, and environment variables.",
    url: "https://golos.sh/docs",
  },
};

function Code({ children }: { children: React.ReactNode }) {
  return (
    <code className="rounded bg-neutral-800 px-1.5 py-0.5 text-sm text-neutral-200">
      {children}
    </code>
  );
}

function Section({
  id,
  title,
  children,
}: {
  id: string;
  title: string;
  children: React.ReactNode;
}) {
  return (
    <section id={id} className="scroll-mt-24 space-y-4">
      <h2 className="text-xl font-semibold">{title}</h2>
      {children}
    </section>
  );
}

function CommandRow({
  command,
  description,
}: {
  command: string;
  description: string;
}) {
  return (
    <div className="flex items-start justify-between gap-4 border-b border-neutral-800 py-3">
      <code className="shrink-0 text-sm text-neutral-200">{command}</code>
      <span className="text-right text-sm text-neutral-500">{description}</span>
    </div>
  );
}

export default function Docs() {
  return (
    <div className="space-y-16">
      <div>
        <h1 className="text-3xl font-bold tracking-tight">Documentation</h1>
        <p className="mt-2 text-neutral-400">
          Everything you need to use Golos.
        </p>
      </div>

      {/* TOC */}
      <nav className="flex flex-wrap gap-3 text-sm">
        {[
          ["#getting-started", "Getting Started"],
          ["#commands", "Commands"],
          ["#flags", "Flags"],
          ["#dictionary", "Dictionary"],
          ["#config", "Configuration"],
          ["#hotkeys", "Hotkeys"],
          ["#env", "Environment Variables"],
        ].map(([href, label]) => (
          <a
            key={href}
            href={href}
            className="rounded-full border border-neutral-800 px-3 py-1 text-neutral-400 hover:border-neutral-600 hover:text-neutral-100 transition-colors"
          >
            {label}
          </a>
        ))}
      </nav>

      {/* Getting Started */}
      <Section id="getting-started" title="Getting Started">
        <p className="text-sm text-neutral-400">
          Install Golos with a single command:
        </p>
        <CopyBlock text="curl -fsSL https://raw.githubusercontent.com/basilysf1709/golos/main/install.sh | bash" />
        <p className="text-sm text-neutral-400">
          Then configure your Deepgram API key:
        </p>
        <CopyBlock text="golos setup" />
        <p className="text-sm text-neutral-400">
          Grant Accessibility permission to your terminal in{" "}
          <strong className="text-neutral-200">
            System Settings &rarr; Privacy &amp; Security &rarr; Accessibility
          </strong>
          , then start Golos:
        </p>
        <CopyBlock text="golos" />
        <p className="text-sm text-neutral-400">
          Hold <strong className="text-neutral-200">Right Option</strong> to
          record. Release to transcribe and paste.
        </p>
      </Section>

      {/* Commands */}
      <Section id="commands" title="Commands">
        <div className="rounded-lg border border-neutral-800 px-4">
          <CommandRow command="golos" description="Run in foreground" />
          <CommandRow command="golos -d" description="Run in background (detached)" />
          <CommandRow command="golos stop" description="Stop the background process" />
          <CommandRow command="golos setup" description="Configure your Deepgram API key" />
          <CommandRow command="golos version" description="Print version, commit, and build date" />
          <CommandRow command='golos add "phrase" "replacement"' description="Add a dictionary replacement" />
          <CommandRow command='golos delete "phrase"' description="Remove a dictionary entry" />
          <CommandRow command="golos list" description="List all dictionary entries" />
          <CommandRow
            command="golos import file.toml"
            description="Import dictionary entries from a TOML file"
          />
        </div>
      </Section>

      {/* Flags */}
      <Section id="flags" title="Flags">
        <p className="text-sm text-neutral-400">
          Flags are passed when running Golos (not with subcommands like{" "}
          <Code>add</Code> or <Code>stop</Code>).
        </p>
        <div className="rounded-lg border border-neutral-800 px-4">
          <CommandRow
            command="--output <mode>"
            description='Output mode: "clipboard" (default) or "stdout"'
          />
          <CommandRow
            command="--hotkey <key>"
            description='Push-to-talk key (default: "right_option")'
          />
          <CommandRow
            command="-d, --detach"
            description="Run as a background daemon"
          />
        </div>

        <h3 className="pt-2 text-sm font-medium text-neutral-300">Examples</h3>
        <CopyBlock text="golos --output stdout" />
        <p className="text-sm text-neutral-500">
          Print transcription to the terminal instead of pasting to clipboard.
        </p>
        <CopyBlock text="golos --hotkey right_command" />
        <p className="text-sm text-neutral-500">
          Use Right Command as the push-to-talk key.
        </p>
        <CopyBlock text="golos -d --output stdout" />
        <p className="text-sm text-neutral-500">
          Run in background with stdout output.
        </p>
      </Section>

      {/* Dictionary */}
      <Section id="dictionary" title="Dictionary">
        <p className="text-sm text-neutral-400">
          The dictionary replaces spoken words with custom text. Say
          &ldquo;period&rdquo; and get <Code>.</Code> in your output. The
          dictionary file lives at{" "}
          <Code>~/.config/golos/dictionary.toml</Code>.
        </p>

        <h3 className="pt-2 text-sm font-medium text-neutral-300">
          Add entries
        </h3>
        <CopyBlock text='golos add "period" "."' />
        <CopyBlock text='golos add "new line" "\n"' />

        <h3 className="pt-2 text-sm font-medium text-neutral-300">
          Remove an entry
        </h3>
        <CopyBlock text='golos delete "period"' />

        <h3 className="pt-2 text-sm font-medium text-neutral-300">
          List all entries
        </h3>
        <CopyBlock text="golos list" />

        <h3 className="pt-2 text-sm font-medium text-neutral-300">
          Import from a file
        </h3>
        <CopyBlock text="golos import dictionary.example.toml" />
        <p className="text-sm text-neutral-400">
          The TOML file should have a <Code>[words]</Code> table:
        </p>
        <pre className="overflow-x-auto rounded-lg border border-neutral-800 bg-neutral-900 px-4 py-3 text-sm">
          <code>{`[words]
"period" = "."
"comma" = ","
"question mark" = "?"
"new line" = "\\n"
"new paragraph" = "\\n\\n"
"open paren" = "("
"close paren" = ")"
"arrow" = "->"
"fat arrow" = "=>"`}</code>
        </pre>
      </Section>

      {/* Config */}
      <Section id="config" title="Configuration">
        <p className="text-sm text-neutral-400">
          Config lives at <Code>~/.config/golos/config.toml</Code>. Run{" "}
          <Code>golos setup</Code> to create it, or edit directly:
        </p>
        <pre className="overflow-x-auto rounded-lg border border-neutral-800 bg-neutral-900 px-4 py-3 text-sm">
          <code>{`deepgram_api_key = "your-key"
hotkey = "right_option"
output_mode = "clipboard"
sample_rate = 16000
language = "en-US"
overlay = true`}</code>
        </pre>

        <div className="rounded-lg border border-neutral-800 px-4">
          <CommandRow command="deepgram_api_key" description="Your Deepgram API key (required)" />
          <CommandRow command="hotkey" description='Push-to-talk key (default: "right_option")' />
          <CommandRow command="output_mode" description='"clipboard" or "stdout" (default: "clipboard")' />
          <CommandRow command="sample_rate" description="Audio sample rate in Hz (default: 16000)" />
          <CommandRow command="language" description='Deepgram language code (default: "en-US")' />
          <CommandRow command="overlay" description="Show visual overlay while recording (default: true)" />
        </div>

        <h3 className="pt-2 text-sm font-medium text-neutral-300">
          Resolution order
        </h3>
        <p className="text-sm text-neutral-400">
          Settings are resolved in this order. Later sources override earlier
          ones:
        </p>
        <ol className="list-inside list-decimal space-y-1 text-sm text-neutral-400">
          <li>Built-in defaults</li>
          <li>
            <Code>.env</Code> file in the current directory
          </li>
          <li>
            <Code>~/.config/golos/config.toml</Code>
          </li>
          <li>Environment variables</li>
          <li>CLI flags (<Code>--output</Code>, <Code>--hotkey</Code>)</li>
        </ol>
      </Section>

      {/* Hotkeys */}
      <Section id="hotkeys" title="Hotkeys">
        <p className="text-sm text-neutral-400">
          Valid values for the <Code>hotkey</Code> config or{" "}
          <Code>--hotkey</Code> flag:
        </p>
        <div className="rounded-lg border border-neutral-800 px-4">
          <div className="flex items-center justify-between border-b border-neutral-800 py-3">
            <code className="text-sm text-neutral-200">right_option</code>
            <span className="text-sm text-neutral-500">
              Right Option / Alt key <span className="text-neutral-600">(alias: right_alt)</span>
            </span>
          </div>
          <div className="flex items-center justify-between border-b border-neutral-800 py-3">
            <code className="text-sm text-neutral-200">right_command</code>
            <span className="text-sm text-neutral-500">
              Right Command key <span className="text-neutral-600">(alias: right_cmd)</span>
            </span>
          </div>
          <div className="flex items-center justify-between border-b border-neutral-800 py-3">
            <code className="text-sm text-neutral-200">fn</code>
            <span className="text-sm text-neutral-500">Fn key</span>
          </div>
          <div className="flex items-center justify-between border-b border-neutral-800 py-3">
            <code className="text-sm text-neutral-200">f18</code>
            <span className="text-sm text-neutral-500">F18 key</span>
          </div>
          <div className="flex items-center justify-between py-3">
            <code className="text-sm text-neutral-200">f19</code>
            <span className="text-sm text-neutral-500">F19 key</span>
          </div>
        </div>
        <CopyBlock text="golos --hotkey fn" />
      </Section>

      {/* Env vars */}
      <Section id="env" title="Environment Variables">
        <p className="text-sm text-neutral-400">
          These override the matching config.toml values:
        </p>
        <div className="rounded-lg border border-neutral-800 px-4">
          <CommandRow command="DEEPGRAM_API_KEY" description="Overrides deepgram_api_key" />
          <CommandRow command="GOLOS_OUTPUT" description="Overrides output_mode" />
          <CommandRow command="GOLOS_HOTKEY" description="Overrides hotkey" />
        </div>
        <CopyBlock text='DEEPGRAM_API_KEY="your-key" golos' />
        <p className="text-sm text-neutral-500">
          You can also put variables in a <Code>.env</Code> file in your
          working directory.
        </p>
      </Section>
    </div>
  );
}
