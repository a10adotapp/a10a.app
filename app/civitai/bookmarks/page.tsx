import { CivitaiItemView } from "@/components/civitai/civitai-item-view";
import { Header } from "@/components/civitai/header";
import prisma from "@/prisma";
import { redirect } from "next/navigation";
import z from "zod";
import { NextButton } from "./_components/next-button";

const paramsSchema = z.object({
  page: z.coerce.number().optional(),
});

export default async function Page(props: {
  searchParams: Promise<Record<string, unknown>>;
}) {
  const params = paramsSchema.parse(await props.searchParams);

  if (!params.page) {
    return redirect([
      "/civitai/bookmarks",
      `?${new URLSearchParams({
        page: "1",
      })}`,
    ].join(""));
  }

  const pageSize = 100;

  const civitaiItems = await prisma.civitaiItem.findMany({
    where: {
      isBookmarked: true,
    },
    orderBy: [
      {
        createdAt: "desc",
      },
      {
        id: "desc",
      },
    ],
    take: pageSize,
    skip: pageSize * (params.page - 1),
  });

  const firstCivitaiItem = civitaiItems.at(0);
  const lastCivitaiItem = civitaiItems.at(civitaiItems.length - 1);

  return (
    <div className="flex flex-col gap-4">
      <div />

      <Header />

      <div className="flex items-center gap-4">
        <div className="flex-1 flex justify-center items-center gap-4">
          {firstCivitaiItem && (
            <div className="break-keep">
              {firstCivitaiItem.createdAt.toISOString()}
            </div>
          )}

          <div>➡️</div>

          {lastCivitaiItem && (
            <div className="break-keep">
              {lastCivitaiItem.createdAt.toISOString()}
            </div>
          )}
        </div>
      </div>

      <div className="grid grid-cols-5 gap-1">
        {civitaiItems.map((civitaiItem) => (
          <CivitaiItemView
            key={civitaiItem.id.toFixed(0)}
            civitaiItem={{
              id: civitaiItem.id,
              createdAt: civitaiItem.createdAt,
              url: civitaiItem.url,
              baseModel: civitaiItem.baseModel,
              prompt: civitaiItem.prompt,
              negativePrompt: civitaiItem.negativePrompt,
              isBookmarked: civitaiItem.isBookmarked,
            }} />
        ))}
      </div>

      <div className="flex justify-center items-center gap-4">
        <button
          className="py-2 px-4 border rounded-lg cursor-pointer">
          &laquo;
          prev
        </button>

        <NextButton />
      </div>

      <div className="h-8" />
    </div>
  );
}
