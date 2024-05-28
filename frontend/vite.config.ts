import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import tsconfigPaths from 'vite-tsconfig-paths';

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    host: true,
    port: 3000,
    // proxy: {
    //   "/hello": "http://localhost:9000",
    //   "/api": "http://localhost:9000",
    // },
  },
  plugins: [react(), tsconfigPaths()],
});
