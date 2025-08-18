import { Filedump } from "@/lib/filedump/client";

export async function GET(): Promise<Response> {
  const filedump = Filedump.init();

  const dumpfiles = filedump.listDumpfile();

  return Response.json(dumpfiles);
}
