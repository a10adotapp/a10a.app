import { mkdirSync, readdirSync, readFileSync, writeFileSync } from "fs";
import { join } from "path";

export class Filedump {
  dirname: string | undefined;

  static init(): Filedump {
    const self = new Filedump();

    self.dirname = process.env.FILEDUMP_DIRNAME;

    if (self.dirname) {
      mkdirSync(self.dirname, {
        recursive: true,
      });
    }

    return self;
  }

  dumpTextFile(filename: string, content: string): void {
    if (!this.dirname) {
      throw new Error("no dirname");
    }

    writeFileSync(join(this.dirname, filename), content, {
      encoding: "utf8",
    });
  }

  listDumpfile(): string[] {
    if (!this.dirname) {
      throw new Error("no dirname");
    }

    const dirents = readdirSync(this.dirname, {
      encoding: "utf8",
      withFileTypes: true,
      recursive: true,
    });

    return dirents.map((dirent) => {
      if (dirent.isDirectory()) {
        return null;
      }

      return dirent.name;
    }).filter((file) => {
      return file !== null;
    });
  }

  readDumpfile(filename: string): string {
    if (!this.dirname) {
      throw new Error("no dirname");
    }

    return readFileSync(join(this.dirname, filename), {
      encoding: "utf8",
    });
  }
}
