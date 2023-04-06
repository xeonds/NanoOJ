import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  server: {
    proxy: {
      "/api/v1": {
        target: "http://www.jiujiuer.xyz:808",
        changeOrigin: true,
        pathRewrite: {
          "^/api/v1": "",
        },
      },
    },
  },
});
