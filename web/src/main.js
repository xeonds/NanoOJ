import { createApp } from "vue";
import v_axios from "./axios/index.js";
import VueAxios from "vue-axios";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "element-plus/theme-chalk/display.css";
import "element-plus/theme-chalk/dark/css-vars.css";

const app = createApp(App);

app.use(store);
app.use(router);
app.use(ElementPlus);
app.use(VueAxios, v_axios);
app.mount("#app");
