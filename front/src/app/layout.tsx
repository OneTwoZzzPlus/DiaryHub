import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "DiaryHub",
  description: "Учебный проект",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ru">
      <body>
        {children}
      </body>
    </html>
  );
}
