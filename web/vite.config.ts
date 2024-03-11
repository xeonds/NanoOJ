import path from 'path'
import { defineConfig } from 'vite'
import Vue from '@vitejs/plugin-vue'
import AutoImport from 'unplugin-auto-import/vite'
import Components from 'unplugin-vue-components/vite'
import { ElementPlusResolver } from 'unplugin-vue-components/resolvers'
import Inspect from 'vite-plugin-inspect'

const pathSrc = path.resolve(__dirname, 'src')
export default defineConfig({
    plugins: [
        Vue(),
        AutoImport({
            imports: ['vue', 'vue-router', '@vueuse/core'],
            resolvers: [ElementPlusResolver()],
            vueTemplate: true,
            dts: path.resolve(pathSrc, 'auto-imports.d.ts')
        }),
        Components({
            resolvers: [ElementPlusResolver()],
            dts: path.resolve(pathSrc, 'typings', 'components.d.ts'),
        }),
        Inspect(),
    ],
    resolve: {
        alias: {
            '@': pathSrc,
        },
    }
})
