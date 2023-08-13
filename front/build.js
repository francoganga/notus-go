import esbuild from "esbuild";
import sveltePlugin from "esbuild-svelte";

let result = await esbuild
  .build({
    entryPoints: ["src/main.js"],
    mainFields: ["svelte", "browser", "module", "main"],
    conditions: ["svelte", "browser"],
    minify: true,
    bundle: true,
    outdir: "../public",
    plugins: [sveltePlugin()],
    logLevel: "info",
    alias: {
      'components': './src/components',
      'views': './src/views',
      'assets': './src/assets',
    }
  })

console.log(result)
