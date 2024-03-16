import { createRouter, createWebHashHistory } from "vue-router";

const HomeView = () => import("@/views/HomeView.vue");

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
    children: [
      {
        path: "/",
        name: "homeview",
        component: () => import("@/views/home/HomeView.vue"),
      },
      {
        path: "/problem",
        name: "problemset",
        component: () => import("@/views/home/ProblemsetView.vue"),
      },
      {
        path: "/contest",
        name: "contest",
        component: () => import("@/views/home/ContestsetView.vue"),
      },
      {
        path: "/contest/:id",
        name: "contest_item",
        component: () => import("@/views/home/ContestView.vue"),
      },
      {
        path: "/status",
        name: "status",
        component: () => import("@/views/home/StatusView.vue"),
      },
      {
        path: "/ranklist",
        name: "ranklist",
        component: () => import("@/views/home/RanklistView.vue"),
      },
      {
        path: "/about",
        name: "about",
        component: () => import("@/views/home/AboutView.vue"),
      },
      {
        path: "/profile",
        name: "profile",
        component: () => import("@/views/home/ProfileView.vue"),
      },
      {
        path: "/problem/:id",
        name: "problem",
        component: () => import("@/views/home/ProblemView.vue"),
        props: true,
      },
    ],
  },
  {
    path: "/admin",
    name: "admin",
    component: () => import("@/views/AdminView.vue"),
  },
  {
    path: "/register",
    name: "register",
    component: () => import("@/views/LoginView.vue"),
  },
  {
    path: "/login",
    name: "login",
    component: () => import("@/views/LoginView.vue"),
  },
  {
    path: "/editor",
    name: "editor",
    component: () => import("@/views/EditorView.vue"),
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
