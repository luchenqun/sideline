import starportLibrary from "@starport/vue";
import { createApp } from "vue";

import App from "./App.vue";
import router from "./router";
import store from "./store";

import ElementPlus from "element-plus";
import "element-plus/dist/index.css";
import "./style/global.css";

import VMdEditor from "@kangc/v-md-editor";
import "@kangc/v-md-editor/lib/style/base-editor.css";
import githubTheme from "@kangc/v-md-editor/lib/theme/github.js";
import "@kangc/v-md-editor/lib/theme/style/github.css";

// highlightjs
import hljs from "highlight.js";

VMdEditor.use(githubTheme, {
  Hljs: hljs,
});

const app = createApp(App);
app.use(store).use(router).use(starportLibrary).use(ElementPlus).use(VMdEditor).mount("#app");
