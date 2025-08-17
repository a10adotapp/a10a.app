import { CodmonClient } from "@/lib/codmon/client";
import { LineClient } from "@/lib/line/client";
import { Hono } from "hono";
import type { Variables } from "./context";
import { explore } from "./explore";

export const sanmeihoikuen = new Hono<{
  Variables: Variables;
}>();

sanmeihoikuen.use(async (c, next) => {
  const turndownService = c.get("turndownService");

  c.set("lineClient", LineClient.init({
    turndownService,
  }));

  c.set("codmonClient", await CodmonClient.init());

  await next();
});

sanmeihoikuen.route("/explore", explore);
