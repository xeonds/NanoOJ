// store/modules/submission.js

const state = {
  submissions: [],
};

const mutations = {
  setSubmissions(state, submissions) {
    state.submissions = submissions;
  },
};

const actions = {
  async fetchSubmissions({ commit }) {
    const response = await axios.get("submissions");
    commit("setSubmissions", response.data);
  },
};

const getters = {
  getSubmissionById: (state) => (id) => {
    return state.submissions.find((submission) => submission.id === id);
  },
};

export default {
  state,
  mutations,
  actions,
  getters,
};

