import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"
import { Header, Footer } from '@/components/ui'

const inter = Inter({ subsets: ["latin"] })

export const metadata: Metadata = {
  title: "Catalog Samael",
  description: "made by Samael",
}

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode
}>) {
  return (
    <html lang="ru">
      <body className={inter.className}>
        <Header />
        <main className="min-h-screen container mx-auto">
          {children}
        </main>
        <Footer />
      </body>
    </html>
  )
}
