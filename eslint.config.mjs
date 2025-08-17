import eslint from "@eslint/js"
import tseslint from "typescript-eslint";
import honoConfig from "@hono/eslint-config"

const config = tseslint.config(
  eslint.configs.recommended,
  tseslint.configs.recommendedTypeChecked,
  {
    ignores: [
      "eslint.config.mjs",
    ],
  },
  ...honoConfig,
  {
    languageOptions: {
      parserOptions: {
        ecmaVersion: "latest",
        projectService: true,
        tsconfigRootDir: import.meta.dirname,
      },
    },
  },
  {
    rules: {
      semi: "error",
      quotes: ["error", "double"],

      "import-x/order": "off",

      "@typescript-eslint/consistent-type-definitions": ["error", "type"],
      "@typescript-eslint/require-await": "warn",
      "@typescript-eslint/no-extraneous-class": "warn",
      "@typescript-eslint/no-unsafe-assignment": "warn",
    },
  },
);

export default config;
