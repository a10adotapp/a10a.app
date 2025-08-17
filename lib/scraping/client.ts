import { sanmeihoikuenComPost } from "./sanmeihoikuen-com-post";

export class ScrapingClient {
  static init(): ScrapingClient {
    return new ScrapingClient();
  }

  async sanmeihoikuenComPost(): Promise<{
    url: string;
    title: string;
    date: string;
    categories: string[];
    imageUris: string[];
  }[]> {
    return sanmeihoikuenComPost.bind(this)();
  }
}
