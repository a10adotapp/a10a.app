import { CodmonLog } from "@/prisma/generated/client";
import type z from "zod";
import { timelineAbsenseDataSchema } from "../codmon/schema";
import type { LineClient } from "./client";

export async function sendCodmonTimelineAbsenseData(
  this: LineClient,
  data: z.output<typeof timelineAbsenseDataSchema>,
  codmonLog: CodmonLog,
): Promise<void> {
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
                      text: data.insert_administrator_name || "",
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
              ],
            },
          },
          {
            type: "bubble",
            body: {
              type: "box",
              layout: "vertical",
              contents: [
                {
                  type: "button",
                  action: {
                    type: "uri",
                    label: "写真を見る",
                    uri: `${process.env.APP_URL}/codmon-logs/${codmonLog.id}/photos`,
                  },
                },
              ],
            },
          },
        ],
      },
    },
  ]);
}
