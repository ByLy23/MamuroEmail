import { createApp } from "vue";
import App from "./App.vue";
import type { EmailInterface } from "./interface/contentInterface";

import "./assets/main.css";

const app = createApp(App);

app.mount("#app");
