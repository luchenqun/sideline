import { createRouter, createWebHistory } from "vue-router";

import Employers from "../views/Employers.vue";
import Employer from "../views/Employer.vue";
import Portfolio from "../views/Portfolio.vue";

const routerHistory = createWebHistory();
const routes = [
  { path: "/", component: Portfolio },
  { path: "/portfolio", name: "portfolio", component: Portfolio },
  { path: "/employers", name: "employers", component: Employers },
  { path: "/employer/:address", name: "employer", component: Employer },
];

const router = createRouter({
  history: routerHistory,
  routes,
});

export default router;
