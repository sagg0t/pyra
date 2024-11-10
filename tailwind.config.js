import forms from "@tailwindcss/forms";

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./{view,pkg}/**/*.{templ,css,js,html}",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        forms,
    ],
}

