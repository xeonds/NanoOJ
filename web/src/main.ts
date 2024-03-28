import { createApp } from 'vue'
import router from './utils/router'
import i18n from './utils/i18n'
// support for responsive ui class
import "element-plus/theme-chalk/display.css";
import App from './App.vue'

createApp(App).use(router).use(i18n).mount('#app')
