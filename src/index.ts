import { ScrapingClient } from "@/lib/scraping/client";
import { serve } from "@hono/node-server";
import { Hono } from "hono";
import TurndownService from "turndown";
import type { Variables } from "./context";
import { sanmeihoikuen } from "./sanmeihoikuen";

const app = new Hono<{
  Variables: Variables;
}>();

app.use(async (c, next) => {
  c.set("turndownService", new TurndownService());

  c.set("scrapingClient", ScrapingClient.init());

  await next();
});

app.route("/sanmeihoikuen", sanmeihoikuen);

serve({
  fetch: app.fetch,
  port: 3000
}, (info) => {
  console.info("Server is running", JSON.stringify({ info }));
});
