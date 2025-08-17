import type z from "zod";
import type { timelineCommentDataSchema } from "../codmon/schema";
import type { LineClient } from "./client";

export async function sendCodmonTimelineData(this: LineClient, data: z.output<typeof timelineCommentDataSchema>): Promise<void> {
  await this.sendMessage([
    {
      type: "flex",
      altText: "Codmon Timeline",
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
                      text: data.display_date,
                      color: "#999999",
                      flex: 1,
                    },
                    {
                      type: "text",
                      text: data.insert_administrator_name,
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
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "午前の機嫌",
                        },
                        {
                          type: "text",
                          text: data.content.mood_morning,
                        },
                      ],
                    },
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "午後の機嫌",
                        },
                        {
                          type: "text",
                          text: data.content.mood_afternoon,
                        },
                      ],
                    },
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "午睡",
                        },
                        {
                          type: "text",
                          text: data.content.sleepings,
                        },
                      ],
                    },
                    ...(data.content.evacuations
                      ? data.content.evacuations.map((evacuation) => ({
                        type: "box",
                        layout: "horizontal",
                        contents: [
                          {
                            type: "text",
                            text: evacuation.evacuation,
                          },
                          {
                            type: "text",
                            text: evacuation.evacuation_time,
                          },
                        ],
                      }))
                      : []),
                  ],
                },
                {
                  type: "box",
                  layout: "vertical",
                  margin: "lg",
                  contents: [
                    {
                      type: "text",
                      text: this.turndownService?.turndown(data.content.memo.replace(/\n/g, "  \n")) ?? "",
                      wrap: true,
                    },
                  ],
                },
              ],
            },
          },
          ...(data.thumbnail_url ? [
            {
              type: "bubble",
              body: {
                type: "box",
                layout: "vertical",
                paddingAll: "none",
                contents: [
                  {
                    type: "image",
                    url: data.thumbnail_url,
                    size: "full",
                    aspectMode: "cover",
                    flex: 1,
                  },
                ],
              },
            },
          ] : []),
          ...(data.thumbnail_url2 ? [
            {
              type: "bubble",
              body: {
                type: "box",
                layout: "vertical",
                paddingAll: "none",
                contents: [
                  {
                    type: "image",
                    url: data.thumbnail_url2,
                    size: "full",
                    aspectMode: "cover",
                    flex: 1,
                  },
                ],
              },
            },
          ] : []),
          ...(data.thumbnail_url3 ? [
            {
              type: "bubble",
              body: {
                type: "box",
                layout: "vertical",
                paddingAll: "none",
                contents: [
                  {
                    type: "image",
                    url: data.thumbnail_url3,
                    size: "full",
                    aspectMode: "cover",
                    flex: 1,
                  },
                ],
              },
            },
          ] : []),
        ],
      },
    },
  ]);
}
