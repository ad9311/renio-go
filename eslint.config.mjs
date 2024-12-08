import globals from 'globals';
import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import eslintPluginPrettier from 'eslint-plugin-prettier/recommended';
import eslintConfigPrettier from 'eslint-config-prettier';

export default tseslint.config(
  eslint.configs.recommended,
  tseslint.configs.strict,
  tseslint.configs.stylistic,
  eslintPluginPrettier,
  eslintConfigPrettier,
  {
    files: ['**/*.{js,mjs,cjs,ts}'],
  },
  {
    ignores: ['web/static/js/build/bundle.js'],
  },
  {
    languageOptions: {
      globals: globals.browser,
    },
  },
  {
    rules: {
      'prettier/prettier': 'error',
      'no-unused-vars': 'off',
      'no-else-return': 'error',
      eqeqeq: 'error',
      'no-console': 'warn',
      curly: 'error',
      semi: 'error',
    },
  }
);
