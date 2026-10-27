"use client";

import { useRouter, useSearchParams } from "next/navigation";

export function NextButton() {
  const searchParams = useSearchParams();

  const router = useRouter();

  const page = Number(searchParams.get("page") || "1");

  return (
    <button
      className="py-2 px-4 border rounded-lg cursor-pointer"
      onClick={() => {
        router.push([
          "/civitai",
          `?${new URLSearchParams({
            page: (page + 1).toFixed(0),
          })}`,
        ].join(""));
      }}>
      next
      &raquo;
    </button>
  );
}
