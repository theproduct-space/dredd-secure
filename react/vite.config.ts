import { defineConfig } from "vite";
import { nodeResolve } from "@rollup/plugin-node-resolve";
import tsconfigPaths from "vite-tsconfig-paths";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [nodeResolve(), react(), tsconfigPaths()],
  optimizeDeps: {
    esbuildOptions: {
      define: {
        global: "globalThis",
      },
    },
  },
});
