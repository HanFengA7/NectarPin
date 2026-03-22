import { createApp } from 'vue'
import { createPinia } from 'pinia'

import '@/assets/style.css'

import App from './App.vue'
import { applyDocumentTitleFromRoute } from './lib/page-title'
import router from './router'
import { useSiteStore } from './stores/site'

const app = createApp(App)
const pinia = createPinia()

app.use(pinia)

const siteStore = useSiteStore()
siteStore.$subscribe(() => {
  applyDocumentTitleFromRoute(router.currentRoute.value)
})

app.use(router)

app.mount('#app')
