import './globals.css';
import { Header } from './components/Header';
import type { Metadata } from "next";

export const metadata: Metadata = {
  title: 'Navin\'s Stock Trading App',
}

export default function RootLayout({ children, }: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" data-theme="autumn">
      <body>
        <Header />
        {children}
      </body>
    </html>
  );
}
