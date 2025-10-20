/**
 * For a detailed explanation regarding each configuration property, visit:
 * https://jestjs.io/docs/configuration
 */

import type { Config } from "jest";
import nextJest from "next/jest.js";

const createJestConfig = nextJest({
  dir: "./",
});

const config: Config = {
  coverageProvider: "v8",
  testEnvironment: "jsdom",
};

export default (async () => ({
  ...(await createJestConfig(config)),
  transformIgnorePatterns: [
    "node_modules/(?!mime)",
  ],
  transform: {
    "\\.[jt]sx?$": "ts-jest",
  },
}))();
