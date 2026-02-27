import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Link from "next/link";
import { Analytics } from "@vercel/analytics/next";
import { GitHubLink } from "./github-link";
import "./globals.css";

const inter = Inter({ subsets: ["latin"] });

const siteUrl = "https://golos.sh";

export const metadata: Metadata = {
  metadataBase: new URL(siteUrl),
  title: {
    default: "Golos - Voice to Text, Instantly",
    template: "%s | Golos",
  },
  description:
    "A free, open-source alternative to WisprFlow. Hold a hotkey, speak, release — your words get pasted. macOS CLI powered by Deepgram.",
  keywords: [
    "speech to text",
    "voice to text",
    "macOS",
    "CLI",
    "push to talk",
    "dictation",
    "Deepgram",
    "WisprFlow alternative",
    "open source",
  ],
  authors: [{ name: "Basil Yusuf", url: "https://github.com/basilysf1709" }],
  creator: "Basil Yusuf",
  icons: {
    icon: "/mascot.png",
    apple: "/mascot.png",
  },
  openGraph: {
    type: "website",
    siteName: "Golos",
    title: "Golos - Voice to Text, Instantly",
    description:
      "A free, open-source alternative to WisprFlow. Hold a hotkey, speak, release — your words get pasted.",
    url: siteUrl,
    images: [{ url: "/og.png", width: 1200, height: 630, alt: "Golos - Voice to Text, Instantly" }],
  },
  twitter: {
    card: "summary_large_image",
    title: "Golos - Voice to Text, Instantly",
    description:
      "A free, open-source alternative to WisprFlow. Hold a hotkey, speak, release — your words get pasted.",
    images: ["/og.png"],
  },
  alternates: {
    canonical: siteUrl,
  },
};

const jsonLd = {
  "@context": "https://schema.org",
  "@type": "SoftwareApplication",
  name: "Golos",
  applicationCategory: "UtilitiesApplication",
  operatingSystem: "macOS",
  offers: {
    "@type": "Offer",
    price: "0",
    priceCurrency: "USD",
  },
  description:
    "A free, open-source alternative to WisprFlow. Hold a hotkey, speak, release — your words get pasted. macOS CLI powered by Deepgram.",
  url: "https://golos.sh",
  downloadUrl: "https://github.com/basilysf1709/golos",
  author: {
    "@type": "Person",
    name: "Basil Yusuf",
    url: "https://github.com/basilysf1709",
  },
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en" className="dark">
      <head>
        <script
          type="application/ld+json"
          dangerouslySetInnerHTML={{ __html: JSON.stringify(jsonLd) }}
        />
      </head>
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
              <Link href="/blog" className="hover:text-neutral-100 transition-colors">
                Blog
              </Link>
              <GitHubLink />
            </div>
          </div>
        </nav>
        <main className="mx-auto max-w-3xl px-6 py-16">{children}</main>
        <Analytics />
      </body>
    </html>
  );
}
