// store/modules/problem.js
import axios from "axios";

export default {
  namespaced: true,
  state:{
    problems: [],
  },
  mutations:{
    setProblems(state, problems) {
      state.problems = problems;
    },
  },
  actions:{
    async fetchProblems({ commit }) {
      const response = await axios.get("problems");
      commit("setProblems", response.data);
    },
  },
  getters:{
    getProblemById: (state) => (id) => {
      return state.problems.find((problem) => problem.id === id);
    },
    getProblems: (state) => (id) => {
      return state.problems;
    },
  },
};
