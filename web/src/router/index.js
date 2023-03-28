import { createRouter, createWebHashHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/problemset",
    name: "problemset",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/ProblemsetView.vue"),
  },
  {
    path: "/contest",
    name: "contest",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/ContestView.vue"),
  },
  {
    path: "/status",
    name: "status",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/StatusView.vue"),
  },
  {
    path: "/ranklist",
    name: "ranklist",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/RanklistView.vue"),
  },
  {
    path: "/about",
    name: "about",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/AboutView.vue"),
  },
  {
    path: "/register",
    name: "register",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/LoginView.vue"),
  },
  {
    path: "/login",
    name: "login",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/LoginView.vue"),
  },
  {
    path: "/problem/:id",
    name: "problem",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/ProblemView.vue"),
  },
  {
    path: "/admin",
    name: "admin",
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/AdminView.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
