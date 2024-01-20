import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {vitePluginForArco} from '@arco-plugins/vite-vue'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        vitePluginForArco({
            style: 'css'
        })
    ],
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
