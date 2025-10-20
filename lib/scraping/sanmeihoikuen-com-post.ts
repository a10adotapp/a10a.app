import { load } from "cheerio";
import type { ScrapingClient } from "./client";

export async function sanmeihoikuenComPost(this: ScrapingClient): Promise<{
  url: string;
  title: string;
  date: string;
  categories: string[];
  imageUris: string[];
}[]> {
  const result: {
    url: string;
    title: string;
    date: string;
    categories: string[];
    imageUris: string[];
  }[] = [];

  try {
    const response = await fetch("https://sanmeihoikuen.com/post");

    if (response.status !== 200) {
      throw new Error(`fetch failed (${JSON.stringify({
        url: "https://sanmeihoikuen.com/post",
        status: response.status,
        body: await response.text(),
      })})`);
    }

    const responseData = await response.text();

    const posts$ = load(responseData);

    const links = posts$(".newsPost>.articleBox>a");

    for (const link of links) {
      const url = posts$(link).prop("href");

      if (!url) {
        continue;
      }

      const response = await fetch(url);

      if (response.status !== 200) {
        throw new Error(`fetch failed (${JSON.stringify({
          url,
          status: response.status,
          body: await response.text(),
        })})`);
      }

      const responseData = await response.text();

      const post$ = load(responseData);

      const title = post$(".newsType").text().trim();

      const date = post$(".entryInfo time").text().trim();

      const categories = Array.from(post$(".entryInfo .post-categories li")).map((elem) => {
        return post$(elem).text().trim();
      });

      let imageUris = Array.from(post$(".entryContents img")).map((elem) => {
        return post$(elem).prop("src");
      }).filter((imageUri) => {
        return imageUri !== undefined;
      }).filter((imageUri) => {
        if (imageUri.endsWith("favicon.png")) {
          return false;
        }

        return true;
      });

      imageUris = await Promise.all(imageUris.map(async (imageUri) => {
        if (this.fileDownloader) {
          return await this.fileDownloader.download(imageUri);
        }

        return imageUri;
      }));

      result.push({
        url,
        title,
        date,
        categories,
        imageUris,
      });
    }
  } catch (err) {
    console.error({
      action: "ScrapingClient.sanmeihoikuenComPost",
      error: err,
    });
  }

  return result;
}
