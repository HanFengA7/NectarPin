import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'

// import 'mdui/mdui.css';
// import 'mdui';

import ArcoVueIcon from '@arco-design/web-vue/es/icon';

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(ArcoVueIcon);
app.use(createPinia())
app.use(router)

app.mount('#app')