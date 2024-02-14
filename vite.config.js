import { defineConfig, loadEnv  } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig(({ mode }) => {
    const env = loadEnv(mode, process.cwd());

    const API_URL = `${env.VITE_API_URL ?? 'http://localhost:8080'}`;
    const PORT = `${env.VITE_PORT ?? '3000'}`;

    return {
        plugins: [react()],
        server: {
            proxy: {
                '/api': {
                    target: API_URL,
                    rewriteHost: true,
                }
            },
            port: PORT,
        }
    }
});