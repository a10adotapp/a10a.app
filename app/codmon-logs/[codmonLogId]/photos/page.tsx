import { photoSchema } from "@/lib/entity/codmon-log/photos";
import prisma from "@/prisma";
import Image from "next/image";
import z from "zod";

export default async function Page(props: {
  params: Promise<{
    codmonLogId: string;
  }>;
}) {
  const params = await props.params;

  const codmonLog = await prisma.codmonLog.findFirstOrThrow({
    where: {
      deletedAt: null,
      id: params.codmonLogId,
    },
  });

  const photos = z.array(photoSchema).parse(codmonLog.photos);

  return photos.map((photo, i) => (
    <div key={`${i}:${photo.uri}`}>
      {photo.uri ? (
        <Image
          key={photo.id}
          alt={photo.uri}
          height={512}
          width={512}
          src={photo.uri} />
      ) : (
        <div>
          {photo.id}
        </div>
      )}
    </div>
  ));
}
