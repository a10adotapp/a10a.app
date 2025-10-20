import { FileDownloader } from "@/lib/file-downloader/client";
import { sanmeihoikuenComPost } from "./sanmeihoikuen-com-post";

export class ScrapingClient {
  fileDownloader: FileDownloader | undefined;

  static init(): ScrapingClient {
    const self = new ScrapingClient();

    self.fileDownloader = FileDownloader.init();

    return self;
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
