import axios from "axios";

export default {
  namespaced: true,
  state:{
    ranks: [],
  },
  mutations:{
    setRanks(state, problems) {
      state.problems = problems;
    },
  },
  actions:{
    async fetchRanks({ commit }) {
      const response = await axios.get("ranks");
      commit("setRanks", response.data);
    },
  },
  getters:{
    getRankById: (state) => (id) => {
      return state.ranks.find((rank) => rank.id === id);
    },
    getRanks: (state) => (id) => {
      return state.ranks;
    },
  },
};
