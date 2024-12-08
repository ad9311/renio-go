import globals from "globals";
import pluginJs from "@eslint/js";
import tseslint from "typescript-eslint";

export default [
  {
    files: ["**/*.{js,mjs,cjs,ts}"],
  },
  {
    languageOptions: { globals: globals.browser },
  },
  {
    ignores: ["web/static/js/build/bundle.js"],
  },
  {
    rules: {
      "prettier/prettier": "error",
      "react-hooks/exhaustive-deps": "off",
      "no-unused-vars": "off",
      "no-else-return": "error",
      "eqeqeq": "error",
      "no-console": "warn",
      "@typescript-eslint/no-unused-vars": "error",
      curly: "error",
      semi: "error",
    },
  },
  pluginJs.configs.recommended,
  ...tseslint.configs.recommended,
];
