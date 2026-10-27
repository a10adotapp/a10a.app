"use client";

import { useRouter } from "next/navigation";

export function CivitaiTagView(props: {
  civitaiTag: {
    id: string;
    value: string;
    itemCount: number;
  };
}) {
  const router = useRouter();

  return (
    <button
      className="py-2 px-4 border rounded-lg flex items-center gap-2 cursor-pointer"
      onClick={() => {
        router.push([
          "/civitai",
          `/tags/${props.civitaiTag.id}`,
        ].join(""));
      }}>
      <div>
        {props.civitaiTag.value}
      </div>

      <div className="min-w-8 rounded-full bg-gray-200 text-sm">
        {Intl.NumberFormat().format(props.civitaiTag.itemCount)}
      </div>
    </button>
  );
}
