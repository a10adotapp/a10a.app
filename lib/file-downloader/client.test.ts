import { FileDownloader } from "./client";

describe("FileDownloader", () => {
  it("downloads file from url", async () => {
    const client = FileDownloader.init();

    const result = await client.download("https://picsum.photos/200");

    console.log({ result });
  });
});
