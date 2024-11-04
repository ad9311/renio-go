import resolve from "@rollup/plugin-node-resolve";
import commonjs from "@rollup/plugin-commonjs";

export default {
  input: "web/static/js/index.js",
  output: {
    file: "web/static/js/build/bundle.js",
    format: "iife",
    name: "renio",
  },
  plugins: [resolve(), commonjs()],
};
