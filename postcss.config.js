import postCSSImport from "postcss-import";
import tailwindCSS from "tailwindcss";
import autoprefixer from "autoprefixer";
import postCSSNested from "postcss-nested";
import postCSSFlexbugsFixes from "postcss-flexbugs-fixes";

/** @type {import('postcss-load-config').Config} */
export default {
    plugins: [
        postCSSImport,
        tailwindCSS,
        autoprefixer,
        postCSSNested,
        postCSSFlexbugsFixes,
    ]
}
