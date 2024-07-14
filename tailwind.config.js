import forms from "@tailwindcss/forms";

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./**/*.{templ,css,js,html}",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        forms,
    ],
}

