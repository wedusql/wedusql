import { App } from "vue";

import "primevue/resources/primevue.min.css";
import "primeflex/primeflex.css";
import "primevue/resources/themes/lara-light-blue/theme.css";

import PrimeVue from "primevue/config";
import ToastService from 'primevue/toastservice';


export default {
  install: (app: App) => {
    app.use(PrimeVue);
    app.use(ToastService);
  },
};
