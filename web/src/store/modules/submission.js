// store/modules/submission.js



export default {
  namespaced: true,
  state:{
    submissions: [],
  },
  mutations:{
    setSubmissions(state, submissions) {
      state.submissions = submissions;
    },
  },
  actions: {
    async fetchSubmissions({ commit }) {
      const response = await this.axios.get("submissions");
      commit("setSubmissions", response.data);
    },
  },
  getters:{
    getSubmissionById: (state) => (id) => {
      return state.submissions.find((submission) => submission.id === id);
    },
  },
};
