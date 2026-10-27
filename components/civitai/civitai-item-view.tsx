"use client";

import { bookmark } from "@/actions/civitai-item/bookmark";
import { ThumbsUp } from "lucide-react";
import { useRouter } from "next/navigation";


export function CivitaiItemView(props: {
  civitaiItem: {
    id: number;
    createdAt: Date;
    url: string;
    baseModel: string;
    prompt: string;
    negativePrompt: string;
    isBookmarked: boolean;
  };
}) {
  const router = useRouter();

  return (
    <button
      onClick={async () => {
        console.log(props.civitaiItem);

        await bookmark(props.civitaiItem.id);

        router.refresh();
      }}
      className="relative cursor-pointer">
      {/* eslint-disable-next-line @next/next/no-img-element */}
      <img
        key={props.civitaiItem.id.toFixed(0)}
        src={props.civitaiItem.url}
        alt={props.civitaiItem.id.toFixed(0)}
        loading="lazy"
        className="w-auto h-auto aspect-square object-cover" />

      {props.civitaiItem.isBookmarked && (
        <div className="absolute top-1 left-1 w-10 aspect-square flex justify-center items-center rounded-full border-2 bg-white">
          <ThumbsUp />
        </div>
      )}
    </button>
  );
}
