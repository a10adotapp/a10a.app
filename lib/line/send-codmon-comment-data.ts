import type z from "zod";
import type { commentResponseSchema } from "../codmon/schema";
import type { LineClient } from "./client";

export async function sendCodmonCommentData(this: LineClient, data: z.output<typeof commentResponseSchema>["data"][number]): Promise<void> {
  await this.sendMessage([
    {
      type: "flex",
      altText: "Codmon Comment",
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
                      text: "保護者",
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
                          text: "夜のごきげん",
                        },
                        {
                          type: "text",
                          text: data.content.mood_evening,
                        },
                      ],
                    },
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "食事（夜）",
                        },
                        {
                          type: "text",
                          text: data.content.meal_evening,
                          wrap: true,
                        },
                      ],
                    },
                    ...((data.content.evacuation_evening && data.content.evacuation_evening_times) ? [
                      {
                        type: "box",
                        layout: "horizontal",
                        contents: [
                          {
                            type: "text",
                            text: "排便（夜）",
                          },
                          {
                            type: "text",
                            text: `${data.content.evacuation_evening}: ${data.content.evacuation_evening_times}`,
                          },
                        ],
                      },
                    ] : []),
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "入眠",
                        },
                        {
                          type: "text",
                          text: data.content.sleep,
                        },
                      ],
                    },
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "起床",
                        },
                        {
                          type: "text",
                          text: data.content.wake,
                        },
                      ],
                    },
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "朝のごきげん",
                        },
                        {
                          type: "text",
                          text: data.content.mood_morning,
                        },
                      ],
                    },
                    ...((data.content.evacuation_morning && data.content.evacuation_morning_times) ? [
                      {
                        type: "box",
                        layout: "horizontal",
                        contents: [
                          {
                            type: "text",
                            text: "排便（朝）",
                          },
                          {
                            type: "text",
                            text: `${data.content.evacuation_morning}: ${data.content.evacuation_morning_times}`,
                          },
                        ],
                      },
                    ] : []),
                    {
                      type: "box",
                      layout: "horizontal",
                      contents: [
                        {
                          type: "text",
                          text: "食事（朝）",
                        },
                        {
                          type: "text",
                          text: data.content.meal_morning,
                          wrap: true,
                        },
                      ],
                    },
                    ...((data.content.temprature && data.content.temprature_time) ? [
                      {
                        type: "box",
                        layout: "horizontal",
                        contents: [
                          {
                            type: "text",
                            text: "体温（朝）",
                          },
                          {
                            type: "text",
                            text: `${data.content.temprature}度（${data.content.temprature_time}）`,
                          },
                        ],
                      },
                    ] : []),
                  ],
                },
                {
                  type: "box",
                  layout: "vertical",
                  margin: "lg",
                  contents: [
                    {
                      type: "text",
                      text: data.content.memo,
                      wrap: true,
                    },
                  ],
                },
              ],
            },
          },
        ],
      },
    },
  ]);
}
