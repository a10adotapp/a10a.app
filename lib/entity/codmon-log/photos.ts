import z from "zod";

export const photoSchema = z.object({
  uri: z.string(),
});

export type Photo = z.output<typeof photoSchema>;
