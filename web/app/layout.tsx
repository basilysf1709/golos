import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Link from "next/link";
import { GitHubLink } from "./github-link";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Golos - Voice to Text, Instantly",
  description:
    "Hold a hotkey, speak, release. Your words get pasted. A macOS CLI for push-to-talk speech recognition.",
  icons: {
    icon: "/mascot.png",
    apple: "/mascot.png",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="dark">
      <body
        className={`${inter.className} bg-neutral-950 text-neutral-100 antialiased`}
      >
        <nav className="sticky top-0 z-50 border-b border-neutral-800 bg-neutral-950/80 backdrop-blur-sm">
          <div className="mx-auto flex max-w-3xl items-center justify-between px-6 py-4">
            <Link href="/" className="flex items-center gap-2 text-lg font-bold tracking-tight">
              <img src="/mascot.png" alt="Golos" width={20} height={20} />
              Golos
            </Link>
            <div className="flex items-center gap-6 text-sm text-neutral-400">
              <Link href="/docs" className="hover:text-neutral-100 transition-colors">
                Docs
              </Link>
              <GitHubLink />
            </div>
          </div>
        </nav>
        <main className="mx-auto max-w-3xl px-6 py-16">{children}</main>
      </body>
    </html>
  );
}
