import { Html, Head, Main, NextScript } from "next/document";

export default function Document() {
  return (
    <Html className="light">
      <Head />
      <body className="bg-light-theme dark:bg-gray-900">
        <Main />
        <NextScript />
      </body>
    </Html>
  );
}
