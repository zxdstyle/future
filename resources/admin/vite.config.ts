import { URL, fileURLToPath } from "node:url"
import { defineConfig } from "vite"
import { setupVitePlugins } from "./build"

export default defineConfig({
    plugins: setupVitePlugins(),
    resolve: {
        alias: {
            "~": fileURLToPath(new URL("./", import.meta.url)),
            "@": fileURLToPath(new URL("./src", import.meta.url)),
        },
    },
    server: {
        proxy: {
            "^/api": {
                target: "http://127.0.0.1:80",
                changeOrigin: true,
            },
        },
    },
})
