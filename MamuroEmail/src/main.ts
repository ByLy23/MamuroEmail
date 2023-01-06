import { createApp } from "vue";
import { createStore } from "vuex";
import App from "./App.vue";
import router from "./router";
import type { EmailInterface } from "./interface/contentInterface";

import "./assets/main.css";
const store = createStore({
  state() {
    return {
      query: "",
      contentSize: 0,
      totalPages: 0,
      Response: [] as EmailInterface[],
      fetching: false,
      isError: false,
      test: 0,
    };
  },
  mutations: {
    addTest(state) {
      state.test += 1;
    },
  },
});

const app = createApp(App);
app.use(store);
app.use(router);

app.mount("#app");
