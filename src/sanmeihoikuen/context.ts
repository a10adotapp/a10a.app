import type { CodmonClient } from "@/lib/codmon/client";
import type { LineClient } from "@/lib/line/client";
import type { Variables as ParentVariables } from "../context";

export type Variables = ParentVariables & {
  lineClient: LineClient;
  codmonClient: CodmonClient;
};
