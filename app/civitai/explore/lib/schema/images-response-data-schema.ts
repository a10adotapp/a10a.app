import { z } from "zod";

export const imagesResponseDataSchema = z.object({
  items: z.array(z.object({
    id: z.number(),
    createdAt: z.string().pipe(z.coerce.date()),
    url: z.string(),
    baseModel: z.string().nullable(),
    stats: z.object({
      cryCount: z.number(),
      laughCount: z.number(),
      likeCount: z.number(),
      dislikeCount: z.number(),
      heartCount: z.number(),
      commentCount: z.number(),
    }),
    meta: z.object({
      prompt: z.string().optional(),
      negativePrompt: z.string().optional(),
    }).nullable(),
  })),
  metadata: z.object({
    nextCursor: z.string(),
  }),
});
