import type { LineClient } from "./client";

export async function sendSanmeihoikuenComPost(this: LineClient, data: {
  url: string;
  title: string;
  date: string;
  categories: string[];
  imageUris: string[];
}): Promise<void> {
  await this.sendMessage([
    {
      type: "flex",
      altText: "Sanmeihoikuen Blog",
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
                      text: data.date,
                      color: "#999999",
                      flex: 1,
                    },
                    {
                      type: "text",
                      text: "ブログ",
                      color: "#999999",
                      flex: 0,
                    },
                  ],
                  alignItems: "center",
                },
                {
                  type: "box",
                  layout: "horizontal",
                  alignItems: "center",
                  spacing: "md",
                  margin: "lg",
                  contents: data.categories.map((category) => ({
                    type: "box",
                    layout: "vertical",
                    flex: 0,
                    justifyContent: "center",
                    paddingStart: "md",
                    paddingEnd: "md",
                    backgroundColor: "#999999",
                    cornerRadius: "md",
                    contents: [
                      {
                        type: "text",
                        text: category,
                        color: "#ffffff",
                        size: "sm",
                      },
                    ],
                  })),
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
            footer: {
              type: "box",
              layout: "vertical",
              contents: [
                {
                  type: "button",
                  action: {
                    type: "uri",
                    label: "記事を開く",
                    uri: data.url,
                  },
                },
              ],
            },
          },
          ...data.imageUris.slice(0, 9).map((imageUri) => ({
            type: "bubble",
            body: {
              type: "box",
              layout: "vertical",
              paddingAll: "none",
              contents: [
                {
                  type: "image",
                  url: imageUri,
                  size: "full",
                  aspectMode: "cover",
                  flex: 1,
                  action: {
                    "type": "uri",
                    "label": "show original image",
                    "uri": imageUri,
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
