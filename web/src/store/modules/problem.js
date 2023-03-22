// store/modules/problem.js

const state = {
  problems: [],
};

const mutations = {
  setProblems(state, problems) {
    state.problems = problems;
  },
};

const actions = {
  async fetchProblems({ commit }) {
    const response = await this.axios.get("problems");
    commit("setProblems", response.data);
  },
};

const getters = {
  getProblemById: (state) => (id) => {
    return state.problems.find((problem) => problem.id === id);
  },
};

export default {
  state,
  mutations,
  actions,
  getters,
};
