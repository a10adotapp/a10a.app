import { format } from "@formkit/tempo";
import type z from "zod";
import type { handoutResponseSchema } from "../codmon/schema";
import type { LineClient } from "./client";

export async function sendCodmonHandoutData(this: LineClient, data: z.output<typeof handoutResponseSchema>): Promise<void> {
  await this.sendMessage([
    {
      type: "flex",
      altText: "Codmon Handout",
      contents: {
        type: "carousel",
        contents: [
          {
            type: "bubble",
            body: {
              type: "box",
              layout: "vertical",
              contents: [
                {
                  type: "box",
                  layout: "horizontal",
                  contents: [
                    {
                      type: "text",
                      text: format({
                        date: data.publishFromDateTime,
                        format: "YYYY-MM-DD",
                        tz: "Asia/Tokyo"
                      }),
                      color: "#999999",
                      flex: 1,
                    },
                    {
                      type: "text",
                      text: "お知らせ",
                      color: "#999999",
                      flex: 0,
                    },
                  ],
                  alignItems: "center",
                },
                {
                  type: "box",
                  layout: "vertical",
                  margin: "lg",
                  contents: [
                    {
                      type: "text",
                      text: data.title,
                      wrap: true,
                    },
                  ],
                },
                {
                  type: "box",
                  layout: "vertical",
                  margin: "lg",
                  contents: [
                    {
                      type: "text",
                      text: data.description?.substring(0, 2000) || "未入力",
                      wrap: true,
                    },
                  ],
                },
              ],
            },
          },
          ...data.attachments.map((attachment) => ({
            type: "bubble",
            body: {
              type: "box",
              layout: "vertical",
              paddingAll: "none",
              contents: [
                {
                  type: "image",
                  url: attachment.url,
                  size: "full",
                  aspectMode: "cover",
                  flex: 1,
                  action: {
                    "type": "uri",
                    "label": "show original image",
                    "uri": attachment.url,
                  },
                },
              ],
            },
          })),
        ],
      },
    },
  ]);
}
