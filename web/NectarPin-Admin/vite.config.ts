import {fileURLToPath, URL} from 'node:url'

import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {vitePluginForArco} from '@arco-plugins/vite-vue'
import fs from 'fs'

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        vitePluginForArco({
            style: 'css'
        })
    ],
    server: {

        https: {
            key: fs.readFileSync('../../certs/localhost+2-key.pem'),
            cert: fs.readFileSync('../../certs/localhost+2.pem')
        }
    },
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./src', import.meta.url))
        }
    }
})
