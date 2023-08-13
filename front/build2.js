import * as esbuild from 'esbuild'
import sveltePlugin from "esbuild-svelte";

let ctx = await esbuild.context({
  entryPoints: ["src/main.js"],
  mainFields: ["svelte", "browser", "module", "main"],
  conditions: ["svelte", "browser"],
  minify: false,
  sourcemap: true,
  bundle: true,
  outdir: "dist",
  plugins: [sveltePlugin()],
  logLevel: "info",
  alias: {
    'components': './src/components',
    'views': './src/views',
    'assets': './src/assets',
  }
})

await ctx.watch()

let { host, port } = await ctx.serve({
  servedir: 'dist',
})
