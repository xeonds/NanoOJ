import { createApp } from 'vue'
import router from './utils/router'
// support for responsive ui class
import "element-plus/theme-chalk/display.css";
import App from './App.vue'

createApp(App).use(router).mount('#app')
