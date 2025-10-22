import { FileDownloader } from "@/lib/file-downloader/client";
import { RouteProps } from "@/lib/props/route-props";
import z from "zod";

const paramsSchema = z.object({
  slug: z.array(z.string()),
});

export async function GET(_: Request, props: RouteProps): Promise<Response> {
  const params = paramsSchema.parse(await props.params);

  const fileDownloader = FileDownloader.init();

  const fileData = await fileDownloader.readDownloadFile(params.slug);

  if (!fileData) {
    return Response.json({
      message: "not found",
    }, {
      status: 404,
    });
  }

  return new Response(fileData);
}
