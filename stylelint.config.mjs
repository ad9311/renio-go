export default {
  extends: [
    'stylelint-config-standard-scss',
    'stylelint-config-tailwindcss/scss',
    'stylelint-config-prettier-scss',
  ],
  plugins: [],
  rules: {},
  ignoreFiles: ['web/static/css/build/**/*.css'],
};
