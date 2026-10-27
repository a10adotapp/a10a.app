import { Header } from "@/components/civitai/header";
import prisma from "@/prisma";
import { CivitaiTagView } from "./_components/civitai-tag-view";

export default async function Page() {
  const civitaiTags = await prisma.civitaiTag.findMany({
    orderBy: [
      {
        itemCount: "desc",
      },
      {
        value: "asc",
      },
      {
        id: "asc",
      },
    ],
    take: 100,
  });

  return (
    <div className="flex flex-col gap-4">
      <div />

      <Header />

      <div className="flex flex-wrap gap-2">
        {civitaiTags.map((civitaiTag) => (
          <CivitaiTagView
            key={civitaiTag.id}
            civitaiTag={{
              id: civitaiTag.id,
              value: civitaiTag.value,
              itemCount: civitaiTag.itemCount,
            }} />
        ))}
      </div>

      <div className="h-8" />
    </div>
  );
}
