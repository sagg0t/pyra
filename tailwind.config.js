import forms from "@tailwindcss/forms";

/** @type {import('tailwindcss').Config} */
export default {
    content: [
        "./view/**/*.{html,css,js}",
    ],
    theme: {
        extend: {},
    },
    plugins: [
        forms,
    ],
}

