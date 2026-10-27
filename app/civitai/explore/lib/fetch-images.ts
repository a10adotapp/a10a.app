import z from "zod";
import { imagesResponseDataSchema } from "./schema/images-response-data-schema";

export async function fetchImages(options?: {
  nextCursor?: string | null | undefined;
} | undefined): Promise<z.output<typeof imagesResponseDataSchema>> {
  const url = [
    "https://civitai.red/api/v1/images",
    `?${new URLSearchParams({
      sort: "Newest",
      limit: "200",
      withMeta: "true",
      nsfw: "X",
      baseModels: "Illustrious",
      ...(options?.nextCursor ? {
        cursor: options.nextCursor,
      } : {}),
    })}`,
  ].join("");

  const response = await fetch(url);

  const rawResponseData = await response.text();

  const responseData = imagesResponseDataSchema.parse(JSON.parse(rawResponseData));

  return responseData;
}
