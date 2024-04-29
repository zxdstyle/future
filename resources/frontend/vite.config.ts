import { defineConfig } from 'vite';
import solid from 'vite-plugin-solid';
import { URL, fileURLToPath } from 'node:url';

export default defineConfig({
    plugins: [solid()],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url)),
        },
    },
});
