import z from "zod";

export const loginResponseSchema = z.object({
  success: z.literal(true),
  data: z.object({
    session_id: z.string(),
  }),
});

export const timelineTopicDataSchema = z.object({
  id: z.string(),
  kind: z.literal("1"),
  display_date: z.string(),
  insert_administrator_name: z.string(),
  title: z.string(),
  content: z.string(),
});

export const timelineCommentDataSchema = z.object({
  id: z.string(),
  kind: z.literal("4"),
  display_date: z.string(),
  insert_administrator_name: z.string(),
  content: z.string().transform((v) => {
    return z.object({
      memo: z.string(),
      mood_morning: z.string(),
      mood_afternoon: z.string(),
      sleepings: z.string(),
      evacuations: z.array(
        z.object({
          evacuation: z.string(),
          evacuation_time: z.string(),
        }),
      ).optional(),
    }).parse(JSON.parse(v));
  }),
  thumbnail_url: z.string(),
  thumbnail_url2: z.string(),
  thumbnail_url3: z.string(),
});

export const timelineResponseDataSchema = z.object({
  id: z.string(),
  kind: z.literal("7"),
});

export const timelineResponseSchema = z.object({
  success: z.literal(true),
  data: z.array(
    z.discriminatedUnion("kind", [
      timelineTopicDataSchema,
      timelineCommentDataSchema,
      timelineResponseDataSchema,
    ]),
  ),
});

export const commentResponseSchema = z.object({
  success: z.literal(true),
  data: z.array(
    z.object({
      id: z.string(),
      display_date: z.string(),
      content: z.string().transform((v) => {
        return z.object({
          mood_evening: z.string(),
          mood_morning: z.string(),
          evacuation_evening: z.string().optional(),
          evacuation_evening_times: z.string().optional(),
          evacuation_morning: z.string().optional(),
          evacuation_morning_times: z.string().optional(),
          meal_evening: z.string(),
          meal_morning: z.string(),
          sleep: z.string(),
          wake: z.string(),
          temprature: z.string().optional(),
          temprature_time: z.string().optional(),
          memo: z.string(),
        }).parse(JSON.parse(v));
      }),
    }),
  ),
});

export const handoutsResponseSchema = z.object({
  handouts: z.array(
    z.object({
      handoutId: z.string(),
      title: z.string(),
      description: z.string().nullable(),
      publishFromDateTime: z.string().transform((v) => (new Date(v))),
      categories: z.array(
        z.object({
          categoryName: z.string(),
        }),
      ),
    }),
  ),
});

export const handoutResponseSchema = z.object({
  handoutId: z.string(),
  title: z.string(),
  description: z.string().nullable(),
  publishFromDateTime: z.string().transform((v) => (new Date(v))),
  categories: z.array(
    z.object({
      categoryName: z.string(),
    }),
  ),
  attachments: z.array(
    z.object({
      fileName: z.string(),
      url: z.string(),
    }),
  ),
});
