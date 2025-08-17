import type { ScrapingClient } from "@/lib/scraping/client";
import type TurndownService from "turndown";

export type Variables = {
  turndownService: TurndownService;
  scrapingClient: ScrapingClient;
};
