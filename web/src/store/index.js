import { createStore } from "vuex";
import problem from "./modules/problem";
import user from "./modules/user";
import submission from "./modules/submission";
import notification from "./modules/notification";

export default createStore({
  state: {
    problemData: [{ id: 0, title: "Hello", note: "Hello, world!" }],
    token: "",
    isLoggedIn: false,
  },
  getters: {},
  mutations: {
    setProblemData(state, payload) {
      state.problemData = payload;
    },
    setToken(state, payload) {
      state.token = payload;
      state.isLoggedIn = true;
    },
  },
  modules: {
    problem,
    user,
    submission,
    notification,
  },
  actions: {
    async init({dispatch}){
      await dispatch("problem/fetchProblems");
      await dispatch("user/fetchUsers");
      await dispatch("submission/fetchSubmissions");
      await dispatch("notification/fetchNotifications");
    }
  },
});
