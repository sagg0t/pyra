import postCSSImport from "postcss-import";
import tailwindCSS from "tailwindcss";
import autoprefixer from "autoprefixer";
import postCSSFlexbugsFixes from "postcss-flexbugs-fixes";

/** @type {import('postcss-load-config').Config} */
export default {
    plugins: {
        "postcss-import": {},
        "tailwindcss/nesting": {},
        tailwindcss: {},
        autoprefixer: {},
        "postcss-flexbugs-fixes": {},
    }
}
