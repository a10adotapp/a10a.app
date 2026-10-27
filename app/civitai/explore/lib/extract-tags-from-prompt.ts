export function extractTagsFromPrompt(prompt: string): string[] {
  return prompt
    .replace(/\s+/g, " ")
    .replace(/[\<\>\(\)]/g, "")
    .replace(/:\d+/g, "")
    .split(",")
    .map((tag) => (tag.trim()))
    .filter((tag) => (tag !== ""));
}
