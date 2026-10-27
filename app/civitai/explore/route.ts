import { withCronAuth } from "@/lib/api-wrapper/cron-auth-wrapper";
import prisma from "@/prisma";
import { CivitaiTag } from "@/prisma/generated/client";
import { extractTagsFromPrompt } from "./lib/extract-tags-from-prompt";
import { fetchImages } from "./lib/fetch-images";
import { positiveScoreThreshold, safeLoopCount } from "./lib/settings";
import { sleep } from "./lib/sleep";

export const POST = withCronAuth(async () => {
  const now = Date.now();

  let nextCursor: string | null = null;

  for (let loopCount = 0; ; loopCount++) {
    console.log(`loopCount: ${loopCount}`);

    const imageResponseData = await fetchImages({
      nextCursor,
    });

    for (const item of imageResponseData.items) {
      const positiveScore = [
        item.stats.heartCount,
        item.stats.laughCount,
        item.stats.likeCount,
      ].reduce((sum, score) => (sum + score), 0);

      if (positiveScore < positiveScoreThreshold) {
        continue;
      }

      let civitaiItem = await prisma.civitaiItem.findFirst({
        where: {
          id: item.id,
        },
      });

      if (civitaiItem) {
        continue;
      }

      civitaiItem = await prisma.civitaiItem.create({
        data: {
          id: item.id,
          createdAt: item.createdAt,
          url: item.url,
          baseModel: item.baseModel || "",
          prompt: item.meta?.prompt || "",
          negativePrompt: item.meta?.negativePrompt || "",
        },
      });

      const tags = [
        ...(extractTagsFromPrompt(item.meta?.prompt || "")),
        ...(extractTagsFromPrompt(item.meta?.negativePrompt || "")),
      ];

      const civitaiTags: CivitaiTag[] = [];

      for (const tag of tags) {
        let civitaiTag = await prisma.civitaiTag.findFirst({
          where: {
            value: tag,
          },
        });

        if (!civitaiTag) {
          civitaiTag = await prisma.civitaiTag.create({
            data: {
              value: tag,
            },
          });
        }

        civitaiTags.push(civitaiTag);
      }

      for (const civitaiTag of civitaiTags) {
        await prisma.civitaiItemTag.create({
          data: {
            civitaiItemId: civitaiItem.id,
            civitaiTagId: civitaiTag.id,
          },
        });
      }
    }

    const lastItem = imageResponseData.items.at(imageResponseData.items.length - 1);

    if (!lastItem || ((now - lastItem.createdAt.getTime()) > (24 * 60 * 60 * 1000))) {
      console.log(`loop finished: ${JSON.stringify({
        now,
        "lastItem?.createdAt.getDate()": lastItem?.createdAt.getDate(),
      })}`);

      break;
    }

    if (loopCount > safeLoopCount) {
      console.log(`loop aborted. (${JSON.stringify({
        safeLoopCount,
      })})`);

      break;
    }

    nextCursor = imageResponseData.metadata.nextCursor;

    await sleep(1);
  }

  return Response.json({
    message: "ok",
  });
});
