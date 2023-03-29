import { createRouter, createWebHashHistory } from "vue-router";

const HomeView = () => import("../views/HomeView.vue");
const ProblemsetView = () => import("../views/ProblemsetView.vue");
const ContestView = () => import("../views/ContestView.vue");
const StatusView = () => import("../views/StatusView.vue");
const RanklistView = () => import("../views/RanklistView.vue");
const AboutView = () => import("../views/AboutView.vue");
const LoginView = () => import("../views/LoginView.vue");
const ProblemView = () => import("../views/ProblemView.vue");
const AdminView = () => import("../views/AdminView.vue");
const ProfileView = () => import("../views/ProfileView.vue");

const routes = [
  {
    path: "/",
    name: "home",
    component: HomeView,
  },
  {
    path: "/problemset",
    name: "problemset",
    component: ProblemsetView,
  },
  {
    path: "/contest",
    name: "contest",
    component: ContestView,
  },
  {
    path: "/status",
    name: "status",
    component: StatusView,
  },
  {
    path: "/ranklist",
    name: "ranklist",
    component: RanklistView,
  },
  {
    path: "/about",
    name: "about",
    component: AboutView,
  },
  {
    path: "/register",
    name: "register",
    component: LoginView,
  },
  {
    path: "/login",
    name: "login",
    component: LoginView,
  },
  {
    path: "/profile",
    name: "profile",
    component: ProfileView,
  },
  {
    path: "/problem/:id",
    name: "problem",
    component: ProblemView,
  },
  {
    path: "/admin",
    name: "admin",
    component: AdminView,
  },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
