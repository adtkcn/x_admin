import App from "./App";
import { createSSRApp } from "vue";

import { installMethods } from "@/methods/index";
import * as Pinia from "pinia";


export function createApp() {
  const app = createSSRApp(App);

  app.use(Pinia.createPinia());

  installMethods(app);

  return {
    app,
    Pinia, // 此处必须将 Pinia 返回
  };
}
