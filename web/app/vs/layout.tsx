import Link from "next/link";

export default function VsLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="space-y-16">
      {children}
      <section className="space-y-4 border-t border-neutral-800 pt-12">
        <h2 className="text-xl font-semibold">Other Comparisons</h2>
        <div className="flex flex-wrap gap-3 text-sm">
          <Link
            href="/vs/wisprflow"
            className="rounded-full border border-neutral-800 px-4 py-1.5 text-neutral-400 hover:border-neutral-600 hover:text-neutral-100 transition-colors"
          >
            Golos vs WisprFlow
          </Link>
          <Link
            href="/vs/superwhisper"
            className="rounded-full border border-neutral-800 px-4 py-1.5 text-neutral-400 hover:border-neutral-600 hover:text-neutral-100 transition-colors"
          >
            Golos vs SuperWhisper
          </Link>
          <Link
            href="/vs/dragon"
            className="rounded-full border border-neutral-800 px-4 py-1.5 text-neutral-400 hover:border-neutral-600 hover:text-neutral-100 transition-colors"
          >
            Golos vs Dragon
          </Link>
        </div>
      </section>
    </div>
  );
}
