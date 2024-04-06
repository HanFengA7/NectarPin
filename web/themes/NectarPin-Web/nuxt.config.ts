// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {enabled: true},
    css: [
        '~/assets/css/main.css',
    ],
    modules: [
        'arco-design-nuxt-module',
        '@nuxtjs/tailwindcss',
    ]
})
