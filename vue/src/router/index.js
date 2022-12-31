import { createRouter, createWebHistory } from "vue-router";

import Employers from "../views/Employers.vue";
import Portfolio from "../views/Portfolio.vue";

const routerHistory = createWebHistory();
const routes = [
  { path: "/", component: Portfolio },
  { path: "/portfolio", component: Portfolio },
  { path: "/employers", component: Employers },
];

const router = createRouter({
  history: routerHistory,
  routes,
});

export default router;
