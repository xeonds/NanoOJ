// store/modules/notification.js

const state = {
  notifications: [],
};

const mutations = {
  setNotifications(state, notifications) {
    state.notifications = notifications;
  },
};

const actions = {
  async fetchNotifications({ commit }) {
    const response = await this.axios.get("notifications");
    commit("setNotifications", response.data);
  },
};

const getters = {
  getNotificationById: (state) => (id) => {
    return state.notifications.find((notification) => notification.id === id);
  },
  getNotifications: (state) => (id) => {
    return state.notifications;
  },
};

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters,
};
