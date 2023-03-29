import v_axios from "../axios/index";
import jwtDecode from "jwt-decode";

export default {
  async init({ dispatch }) {
    const token = localStorage.getItem("token");
    if (token) {
      const data = jwtDecode(token);
      if (data.exp >= Date.now() / 1000) {
        dispatch("logout");
      }
      dispatch("login", token);
      dispatch("fetchUserInfo", data.user_id);
    }
    await dispatch("fetchProblems");
    await dispatch("fetchUsers");
    await dispatch("fetchSubmissions");
    await dispatch("fetchNotifications");
  },
  async login({ commit, dispatch }, token) {
    commit("setToken", token);
    commit("setLoggedInStatus", true);
    v_axios.defaults.headers.common["Authorization"] = token;
    localStorage.setItem("token", token);
    const decode = jwtDecode(token);
    await dispatch("fetchUserInfo", decode.user_id);
  },
  logout({ commit }) {
    commit("setToken", "");
    commit("setLoggedInStatus", false);
    delete v_axios.defaults.headers.common["Authorization"];
    localStorage.removeItem("token");
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
  async fetchUserInfo({ commit }, userID) {
    const response = await v_axios.get("/users/" + userID);
    commit("setUserInfo", response.data);
  },
  async createProblem({ commit }, problem) {
    const response = await v_axios.post("/problems", problem);
    commit("addProblem", response.data);
  },
  async updateProblem({ commit }, problem) {
    const response = await v_axios.put(
      `/problems/${problem.ProblemID}`,
      problem
    );
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
