import { createStore } from "vuex";

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
  actions: {},
  modules: {},
});
