import forms from "@tailwindcss/forms";

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./view/**/*.{templ,css,js,html,go}",
        "!./view/**/*_templ.go"
    ],
    theme: {
        extend: {},
    },
    plugins: [
        forms,
    ],
}

