"use server";

import prisma from "@/prisma";

export async function bookmark(id: number) {
  const civitaiItem = await prisma.civitaiItem.findFirstOrThrow({
    where: {
      id,
    },
  });

  await prisma.civitaiItem.update({
    where: {
      id: civitaiItem.id,
    },
    data: {
      isBookmarked: !civitaiItem.isBookmarked,
    },
  });
}
