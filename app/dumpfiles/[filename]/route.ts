import { Filedump } from "@/lib/filedump/client";
import { extname } from "node:path";
import z from "zod";

const paramsSchema = z.object({
  filename: z.string(),
});

export async function GET(
  _: Request,
  props: {
    params: Promise<Record<string, string[] | string | undefined>>;
  },
) {
  const params = paramsSchema.parse(await props.params);

  const filedump = Filedump.init();

  const ext = extname(params.filename);

  const content = filedump.readDumpfile(params.filename);

  if (ext.endsWith("json")) {
    return Response.json(JSON.parse(content));
  }

  return new Response(content);
}
