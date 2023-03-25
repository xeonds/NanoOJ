const state = {
  token: null,
  isLoggedIn: false,
};

const mutations = {
  SET_TOKEN(state, token) {
    state.token = token;
  },
  SET_LOGGED_IN_STATUS(state, isLoggedIn) {
    state.isLoggedIn = isLoggedIn;
  },
};

const actions = {
  async login({ commit }, token) {
    commit("SET_TOKEN", token);
  },
  async setLoggedInStatus({ commit }, isLoggedIn) {
    commit("SET_LOGGED_IN_STATUS", isLoggedIn);
  },
};

const getters = {
  getToken: state => state.token,
  getLoggedInStatus: state => state.isLoggedIn,
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
