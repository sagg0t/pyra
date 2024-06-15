import path from 'path';
import fs from 'fs';

const config = {
  sourcemap: "external",
  entrypoints: [
      "view/assets/js/main.js",
  ],
  outdir: path.join(process.cwd(), "public/assets"),
};

const build = async (config) => {
  const result = await Bun.build(config);

  if (!result.success) {
    if (process.argv.includes("--watch")) {
      console.error("Build failed");
      for (const message of result.logs) {
        console.error(message);
      }
      return;
    } else {
      throw new AggregateError(result.logs, "Build failed");
    }
  }
};

(async () => {
  await build(config);

  if (process.argv.includes("--watch")) {
    fs.watch(path.join(process.cwd(), "view/assets/js"), { recursive: true }, (_eventType, filename) => {
      console.log(`File changed: ${filename}. Rebuilding...`);
      build(config);
    });
  } else {
    process.exit(0);
  }
})();
