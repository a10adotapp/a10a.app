import type TurndownService from "turndown";
import type z from "zod";
import type { commentResponseSchema, handoutResponseSchema, timelineCommentDataSchema } from "../codmon/schema";
import { sendCodmonCommentData } from "./send-codmon-comment-data";
import { sendCodmonHandoutData } from "./send-codmon-handout-data";
import { sendCodmonTimelineData } from "./send-codmon-timeline-data";
import { sendSanmeihoikuenComPost } from "./send-sanmeihoikuen-com-post";

export class LineClient {
  channelId: string | undefined;
  channelSecret: string | undefined;
  channelAccessToken: string | undefined;
  destination: string | undefined;
  turndownService: TurndownService | undefined;

  static init(deps: {
    turndownService: TurndownService;
  }): LineClient {
    const self = new LineClient();

    self.channelId = process.env.LINE_MESSAGING_API_CHANNEL_ID;
    self.channelSecret = process.env.LINE_MESSAGING_API_CHANNEL_SECRET;
    self.channelAccessToken = process.env.LINE_MESSAGING_API_CHANNEL_ACCESS_TOKEN;
    self.destination = process.env.LINE_MESSAGING_API_DESTINATION;
    self.turndownService = deps.turndownService;

    return self;
  }

  async sendMessage(messages: readonly Record<string, unknown>[]): Promise<void> {
    if (!this.channelAccessToken) {
      throw new Error("no channel access token");
    }

    try {
      const response = await fetch(
        "https://api.line.me/v2/bot/message/push",
        {
          method: "POST",
          headers: {
            "content-type": "application/json",
            "authorization": `Bearer ${this.channelAccessToken}`,
          },
          body: JSON.stringify({
            to: this.destination,
            messages,
          }),
        },
      );

      if (response.status !== 200) {
        throw new Error(`send line message failed (${JSON.stringify({
          status: response.status,
          body: await response.text(),
        })})`);
      }
    } catch (err) {
      console.error({
        action: "LineClient.sendMessage",
        error: err,
      });
    }
  }

  async sendCodmonTimelineData(data: z.output<typeof timelineCommentDataSchema>): Promise<void> {
    await sendCodmonTimelineData.bind(this)(data);
  }

  async sendCodmonCommentData(data: z.output<typeof commentResponseSchema>["data"][number]): Promise<void> {
    await sendCodmonCommentData.bind(this)(data);
  }

  async sendCodmonHandoutData(data: z.output<typeof handoutResponseSchema>): Promise<void> {
    await sendCodmonHandoutData.bind(this)(data);
  }

  async sendSanmeihoikuenComPost(data: {
    url: string;
    title: string;
    date: string;
    categories: string[];
    imageUris: string[];
  }): Promise<void> {
    await sendSanmeihoikuenComPost.bind(this)(data);
  }
}
