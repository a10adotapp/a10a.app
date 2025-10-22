import { createHash } from "crypto";
import { mkdirSync, readFileSync, statSync, writeFileSync } from "fs";
import mime from "mime";
import path from "path";

export class FileDownloader {
  publicFileDirname: string | undefined;

  appUrl: string | undefined;

  static init(): FileDownloader {
    const self = new FileDownloader();

    self.publicFileDirname = process.env.PUBLIC_FILE_DIRNAME;
    self.appUrl = process.env.APP_URL;

    if (self.publicFileDirname) {
      mkdirSync(self.publicFileDirname, {
        recursive: true,
      });
    }

    return self;
  }

  async download(uri: string): Promise<string> {
    if (!this.publicFileDirname) {
      throw new Error("no dirname");
    }

    const response = await fetch(uri);

    if (!response.ok) {
      throw new Error(`response not ok (status: ${response.status}, body: ${await response.text()})`);
    }

    const responseData = await response.blob();

    const fileData = Buffer.from(await responseData.arrayBuffer());

    const fileName = [
      createHash("md5")
        .update(fileData)
        .digest("hex"),
      mime.getExtension(responseData.type),
    ].join(".");

    writeFileSync(
      path.join(this.publicFileDirname, "download", fileName),
      fileData,
    );

    const url = URL.parse(
      path.join("download-files", fileName),
      this.appUrl,
    )?.toString();

    if (!url) {
      throw new Error("no url generated");
    }

    return url;
  }

  async readDownloadFile(slug: string[]): Promise<Buffer<ArrayBuffer> | null> {
    if (!this.publicFileDirname) {
      throw new Error("no dirname");
    }

    const filePath = path.join(this.publicFileDirname, ...slug);

    const stat = statSync(filePath, {
      throwIfNoEntry: false,
    });

    if (!stat) {
      return null;
    }

    return readFileSync(filePath);
  }
}
