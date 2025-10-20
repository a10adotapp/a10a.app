import { Filedump } from "@/lib/filedump/client";
import { format } from "@formkit/tempo";
import type z from "zod";
import { FileDownloader } from "../file-downloader/client";
import { commentResponseSchema, handoutResponseSchema, handoutsResponseSchema, loginResponseSchema, timelineResponseSchema } from "./schema";

export class CodmonClient {
  sessionId: string | undefined;

  filedump: Filedump | undefined;

  fileDownloader: FileDownloader | undefined;

  static async init(): Promise<CodmonClient> {
    const self = new CodmonClient();

    const response = await fetch("https://ps-api.codmon.com/api/v2/parent/login?auto=true&__env__=myapp", {
      method: "POST",
      headers: {
        "content-type": "application/json; charset=utf-8",
      },
      body: JSON.stringify({
        login_id: process.env.CODMON_LOGIN_ID,
        login_password: process.env.CODMON_LOGIN_PASSWORD,
        use_db_replica: 1,
        osCode: "1",
        deviceToken: "eHjhRBhUx0XVnU-VfG0ed7:APA91bEdudujeSU4qt9nawbCkgjNo6MyxEXS9MiYHllvoxhadeWvqDnwDEeH2LU8obc99yDmB78tbZOEtISO9GCiwuRkTcFqO-FUSY6REhGx7_zdenR_mXMP_G732RpTf31EUlHqQ9Bp",
        apnsTokenToDelete: "865c538a82e23ac7306dbd10de66183744168e5e786547cba7a5adb437e23c14"
      }),
    });

    const responseData = loginResponseSchema.parse(await response.json());

    self.sessionId = responseData.data.session_id;

    self.filedump = Filedump.init();

    self.fileDownloader = FileDownloader.init();

    return self;
  }

  authHeaders(): Record<string, string> {
    if (!this.sessionId) {
      throw new Error("no session id");
    }

    return {
      "cookie": [
        `CODMONSESSID=${this.sessionId}`,
        `PHPSESSID=${this.sessionId}`,
      ].join("; "),
      "authorization": this.sessionId,
    };
  }

  async timeline({
    startDate,
    endDate,
  }: {
    startDate: string;
    endDate: string;
  }): Promise<z.output<typeof timelineResponseSchema>> {
    const response = await fetch(
      `https://ps-api.codmon.com/api/v2/parent/timeline/?${new URLSearchParams({
        "__env__": "myapp",
        "service_id": "37102",
        "listpage": "1",
        "search_type[]": "new_all",
        "start_date": startDate,
        "end_date": endDate,
        "use_image_edge": "true",
      })}`,
      {
        method: "GET",
        headers: {
          ...this.authHeaders(),
        },
      });

    const responseData = await response.json();

    this.filedump?.dumpTextFile(
      [
        "codmon",
        "timeline",
        format({
          date: new Date(),
          format: "YYYYMMDDHHmmss",
          tz: "Asia/Tokyo",
        }),
        "json",
      ].join("."),
      JSON.stringify(responseData, null, 2),
    );

    const timelineResponse = timelineResponseSchema.parse(responseData);

    if (this.fileDownloader) {
      for (const data of timelineResponse.data) {
        if (data.kind === "4") {
          if (data.thumbnail_url) {
            data.thumbnail_url = await this.fileDownloader.download(data.thumbnail_url);
          }

          if (data.thumbnail_url) {
            data.thumbnail_url = await this.fileDownloader.download(data.thumbnail_url);
          }

          if (data.thumbnail_url) {
            data.thumbnail_url = await this.fileDownloader.download(data.thumbnail_url);
          }
        }
      }
    }

    return timelineResponse;
  }

  async comments({
    startDate,
    endDate,
  }: {
    startDate: string;
    endDate: string;
  }): Promise<z.output<typeof commentResponseSchema>> {
    const response = await fetch(
      `https://ps-api.codmon.com/api/v2/parent/comments/?${new URLSearchParams({
        "__env__": "myapp",
        "relation_id": "3874066",
        "search_kind": "2",
        "search_start_display_date": startDate,
        "search_end_display_date": endDate,
      })}`,
      {
        method: "GET",
        headers: {
          ...this.authHeaders(),
        },
      });

    const responseData = await response.json();

    this.filedump?.dumpTextFile(
      [
        "codmon",
        "comments",
        format({
          date: new Date(),
          format: "YYYYMMDDHHmmss",
          tz: "Asia/Tokyo",
        }),
        "json",
      ].join("."),
      JSON.stringify(responseData, null, 2),
    );

    return commentResponseSchema.parse(responseData);
  }

  async handouts(): Promise<z.output<typeof handoutsResponseSchema>> {
    const response = await fetch(
      `https://api-reference-room.codmon.com/v1/handouts/forParents?${new URLSearchParams({
        "page": "1",
      })}`,
      {
        method: "GET",
        headers: {
          ...this.authHeaders(),
        },
      });

    const responseData = await response.json();

    this.filedump?.dumpTextFile(
      [
        "codmon",
        "handouts",
        format({
          date: new Date(),
          format: "YYYYMMDDHHmmss",
          tz: "Asia/Tokyo",
        }),
        "json",
      ].join("."),
      JSON.stringify(responseData, null, 2),
    );

    return handoutsResponseSchema.parse(responseData);
  }

  async handout(id: string): Promise<z.output<typeof handoutResponseSchema>> {
    const response = await fetch(
      `https://api-reference-room.codmon.com/v1/handouts/${id}/forParents`,
      {
        method: "GET",
        headers: {
          ...this.authHeaders(),
        },
      });

    const responseData = await response.json();

    this.filedump?.dumpTextFile(
      [
        "codmon",
        "handout",
        format({
          date: new Date(),
          format: "YYYYMMDDHHmmss",
          tz: "Asia/Tokyo",
        }),
        "json",
      ].join("."),
      JSON.stringify(responseData, null, 2),
    );

    const handoutResponse = handoutResponseSchema.parse(responseData);

    if (this.fileDownloader) {
      for (const attachment of handoutResponse.attachments) {
        attachment.url = await this.fileDownloader.download(attachment.url);
      }
    }

    return handoutResponse;
  }
}
