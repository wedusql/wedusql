import { createApp } from "vue";
import App from "./App.vue";
import { router } from "./router";
import primevue from "./plugins/primevue";

createApp(App).use(router).use(primevue).mount("#app");
