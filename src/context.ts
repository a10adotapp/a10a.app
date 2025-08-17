import type { ScrapingClient } from "@/lib/scraping/clinet";
import type TurndownService from "turndown";

export type Variables = {
  turndownService: TurndownService;
  scrapingClient: ScrapingClient;
};
