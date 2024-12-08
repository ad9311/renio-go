import resolve from '@rollup/plugin-node-resolve';
import commonjs from '@rollup/plugin-commonjs';
import typescript from '@rollup/plugin-typescript';

export default {
  input: 'web/static/js/index.ts',
  output: {
    file: 'web/static/js/build/bundle.js',
    format: 'iife',
    name: 'renio',
  },
  plugins: [resolve(), commonjs(), typescript()],
};
