import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [react()],
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
