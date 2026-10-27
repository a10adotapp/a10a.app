import { ReactNode } from "react";

import "./global.css";

export default async function Layout(props: {
  children: ReactNode;
}) {
  return (
    <html>
      <body>
        {props.children}
      </body>
    </html>
  );
}
