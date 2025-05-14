import postCSSFlexbugsFixes from "postcss-flexbugs-fixes";

/** @type {import('postcss-load-config').Config} */
export default {
    plugins: {
        "@tailwindcss/postcss": {},
        "@tailwindcss/nesting": {},
        "postcss-flexbugs-fixes": {},
    }
}
