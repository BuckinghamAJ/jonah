/// <reference types="vitest" />
import { defineConfig } from "vite";
import solidPlugin from "vite-plugin-solid";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [tailwindcss(), solidPlugin()],
  build: {
    target: "esnext",
  },
  resolve: {
    conditions: ["development", "browser"],
  },
  test: {
    environment: "jsdom",
    globals: true,
    setupFiles: ["./src/setupTests.ts"],
    deps: {
      optimizer: {
        web: {
          include: [],
        },
      },
    },
  },
});
