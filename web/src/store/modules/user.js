// store/modules/user.js

const state = {
  users: [],
};

const mutations = {
  setUsers(state, users) {
    state.users = users;
  },
};

const actions = {
  async fetchUsers({ commit }) {
    const response = await this.axios.get("users");
    commit("setUsers", response.data);
  },
};

const getters = {
  getUserById: (state) => (id) => {
    return state.users.find((user) => user.id === id);
  },
};

export default {
  state,
  mutations,
  actions,
  getters,
};

