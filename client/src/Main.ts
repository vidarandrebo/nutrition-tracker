import { createApp } from "vue";
import App from "./App.vue";
import router from "./Router.ts";
import { pinia } from "./Pinia.ts";

const app = createApp(App);
app.use(router);
app.use(pinia);
app.mount("#app");
