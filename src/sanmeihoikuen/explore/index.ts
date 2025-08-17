import prisma from "@/prisma";
import { CodmonLogKind } from "@/prisma/generated/client";
import { addDay, format } from "@formkit/tempo";
import { Hono } from "hono";
import type { Variables } from "./context";

export const explore = new Hono<{
  Variables: Variables;
}>();

explore.post("/", async (c) => {
  const codmonClient = c.get("codmonClient");

  const lineClient = c.get("lineClient");

  const scrapingClient = c.get("scrapingClient");

  const now = new Date();

  const nowDateString = format({
    date: now,
    format: "YYYY-MM-DD",
    tz: "Asia/Tokyo",
  });

  const yesterdayDateString = format({
    date: addDay(now, -1),
    // date: addDay(now, -30),
    format: "YYYY-MM-DD",
    tz: "Asia/Tokyo",
  });

  const timeline = await codmonClient.timeline({
    startDate: yesterdayDateString,
    endDate: nowDateString,
  });

  for (const data of timeline.data) {
    if (data.kind === "4") {
      const codmonLog = await prisma.codmonLog.findFirst({
        where: {
          deletedAt: null,
          kind: CodmonLogKind.TimelineData,
          codmonId: data.id,
        },
      });

      if (!codmonLog) {
        await lineClient.sendCodmonTimelineData(data);

        await prisma.codmonLog.create({
          data: {
            kind: CodmonLogKind.TimelineData,
            codmonId: data.id,
          },
        });
      }
    }
  }

  const comments = await codmonClient.comments({
    startDate: yesterdayDateString,
    endDate: nowDateString,
  });

  for (const data of comments.data) {
    const codmonLog = await prisma.codmonLog.findFirst({
      where: {
        deletedAt: null,
        kind: CodmonLogKind.CommentData,
        codmonId: data.id,
      },
    });

    if (!codmonLog) {
      await lineClient.sendCodmonCommentData(data);

      await prisma.codmonLog.create({
        data: {
          kind: CodmonLogKind.CommentData,
          codmonId: data.id,
        },
      });
    }
  }

  const handouts = await codmonClient.handouts();

  for (const data of handouts.handouts) {
    const codmonLog = await prisma.codmonLog.findFirst({
      where: {
        deletedAt: null,
        kind: CodmonLogKind.HandoutData,
        codmonId: data.handoutId,
      },
    });

    if (!codmonLog) {
      const handout = await codmonClient.handout(data.handoutId);

      console.log(JSON.stringify(handout, null, 2));

      await lineClient.sendCodmonHandoutData(handout);

      await prisma.codmonLog.create({
        data: {
          kind: CodmonLogKind.HandoutData,
          codmonId: data.handoutId,
        },
      });

      break;
    }
  }

  const posts = await scrapingClient.sanmeihoikuenComPost();

  for (const post of posts) {
    const scrapingUrl = await prisma.scrapingUrl.findFirst({
      where: {
        deletedAt: null,
        url: post.url,
      },
    });

    if (!scrapingUrl) {
      await lineClient.sendSanmeihoikuenComPost(post);

      await prisma.scrapingUrl.create({
        data: {
          url: post.url,
        },
      });
    }
  }

  return c.json({
    message: "ok",
  });
});
