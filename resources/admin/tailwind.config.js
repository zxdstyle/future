import { colors } from "./src/config/theme/dark"

/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: "class",
    content: ["./index.html", "./src/**/*.{js,ts,jsx,tsx}"],
    theme: {
        container: {
            center: true,
        },
        extend: {
            boxShadow: {
                primary: "0 25px 50px -12px hsla(0,0%,100%,.05)",
                secondary: "14px 17px 40px 4px rgb(112 144 176 / 18%)",
                normal: "0 2px 4px 0 rgba(0,0,0,.05)",
                strong: "0 2px 8px 0 rgba(0, 0, 0, 0.1)",
            },
            colors,
            height: {
                128: "32rem",
                160: "40rem",
            },
            minWidth: {
                128: "32rem",
                160: "40rem",
            },
            blur: {
                "4xl": "84px",
            },
        },
    },
    plugins: [],
    corePlugins: {
        preflight: false,
    },
}
