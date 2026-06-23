import z from "zod";

export const photoSchema = z.object({
  id: z.string(),
  uri: z.string().optional(),
});

export type Photo = z.output<typeof photoSchema>;
