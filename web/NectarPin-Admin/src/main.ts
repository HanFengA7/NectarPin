import './assets/main.css'

import {createApp} from 'vue'
import {createPinia} from 'pinia'

// import 'mdui/mdui.css';
// import 'mdui';
import ECharts from 'vue-echarts' // 引入ECharts
import 'echarts';

import ArcoVueIcon from '@arco-design/web-vue/es/icon';

import App from './App.vue'
import router from './router'

const app = createApp(App)
app.use(ArcoVueIcon);
app.use(createPinia())
app.use(router)
app.component('ECharts', ECharts)
app.mount('#app')
