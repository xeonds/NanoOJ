import v_axios from "../axios/index";

export default {
  async init({ dispatch }) {
    await dispatch("fetchProblems");
    await dispatch("fetchUsers");
    await dispatch("fetchSubmissions");
    await dispatch("fetchNotifications");
  },
  //auth
  async login({ commit }, token) {
    commit("setToken", token);
    v_axios.defaults.headers.common["Authorization"] = token;
  },
  async setLoggedInStatus({ commit }, isLoggedIn) {
    commit("setLoggedInStatus", isLoggedIn);
  },
  async fetchNotifications({ commit }) {
    const response = await v_axios.get("/notifications");
    commit("setNotifications", response.data);
  },
  async fetchProblems({ commit }) {
    const response = await v_axios.get("/problems");
    commit("setProblems", response.data);
  },
  async fetchRanks({ commit }) {
    const response = await v_axios.get("/ranks");
    commit("setRanks", response.data);
  },
  async fetchSubmissions({ commit }) {
    const response = await v_axios.get("/submissions");
    commit("setSubmissions", response.data);
  },
  async fetchUsers({ commit }) {
    const response = await v_axios.get("/users");
    commit("setUsers", response.data);
  },
  async createProblem({ commit }, problem) {
    const response = await v_axios.post("/problems", problem);
    commit("addProblem", response.data);
  },
  async updateProblem({ commit }, problem) {
    const response = await v_axios.put(`/problems/${problem.id}`, problem);
    commit("updateProblem", response.data);
  },
  async deleteProblem({ commit }, problemId) {
    await v_axios.delete(`/problems/${problemId}`);
    commit("deleteProblem", problemId);
  },
  async createNotification({ commit }, notification) {
    const response = await v_axios.post("/notifications", notification);
    commit("addNotification", response.data);
  },
  async updateNotification({ commit }, notification) {
    const response = await v_axios.put(
      `/notifications/${notification.id}`,
      notification
    );
    commit("updateNotification", response.data);
  },
  async deleteNotification({ commit }, notificationId) {
    await v_axios.delete(`/notifications/${notificationId}`);
    commit("deleteNotification", notificationId);
  },
};
