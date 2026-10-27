export function Header() {
  return (
    <div className="flex justify-end items-center gap-4">
      <div>
        <a
          href={[
            "/civitai",
          ].join("")}
          className="py-2 px-4 border rounded-lg cursor-pointer">
          images
        </a>
      </div>

      <div>
        <a
          href={[
            "/civitai/bookmarks",
          ].join("")}
          className="py-2 px-4 border rounded-lg cursor-pointer">
          bookmarks
        </a>
      </div>

      <div>
        <a
          href={[
            "/civitai/tags",
          ].join("")}
          className="py-2 px-4 border rounded-lg cursor-pointer">
          tags
        </a>
      </div>
    </div>
  );
}
