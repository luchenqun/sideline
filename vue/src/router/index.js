import { createRouter, createWebHistory } from "vue-router";

import Developers from "../views/Developers.vue";
import Developer from "../views/Developer.vue";
import Employers from "../views/Employers.vue";
import Employer from "../views/Employer.vue";
import Tasks from "../views/Tasks.vue";
import Task from "../views/Task.vue";
import Portfolio from "../views/Portfolio.vue";

const routerHistory = createWebHistory();
const routes = [
  { path: "/", component: Portfolio },
  { path: "/portfolio", name: "portfolio", component: Portfolio },
  { path: "/employers", name: "employers", component: Employers },
  { path: "/employer/:address", name: "employer", component: Employer },
  { path: "/developers", name: "developers", component: Developers },
  { path: "/developer/:address", name: "developer", component: Developer },
  { path: "/tasks", name: "tasks", component: Tasks },
  { path: "/task/:id", name: "task", component: Task },
];

const router = createRouter({
  history: routerHistory,
  routes,
});

export default router;
