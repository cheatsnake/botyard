import { defineConfig } from "vite";
import { VitePWA } from "vite-plugin-pwa";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        react(),
        VitePWA({
            registerType: "autoUpdate",
            workbox: { cleanupOutdatedCaches: true },
            manifest: {
                lang: "en",
                name: "Botyard",
                short_name: "Botyard",
                icons: [
                    {
                        src: "icon-48x48.png",
                        sizes: "48x48",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-72x72.png",
                        sizes: "72x72",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-96x96.png",
                        sizes: "96x96",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-128x128.png",
                        sizes: "128x128",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-144x144.png",
                        sizes: "144x144",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-152x152.png",
                        sizes: "152x152",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-192x192.png",
                        sizes: "192x192",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-384x384.png",
                        sizes: "384x384",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                    {
                        src: "icon-512x512.png",
                        sizes: "512x512",
                        type: "image/png",
                        purpose: "maskable any",
                    },
                ],
            },
        }),
    ],
    server: {
        proxy: {
            "/v1": { target: "http://localhost:7007", changeOrigin: true },
            "/static": { target: "http://localhost:7007", changeOrigin: true },
        },
    },
    build: {
        chunkSizeWarningLimit: 700,
    },
});
