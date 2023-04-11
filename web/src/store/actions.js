import api from "../api";
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
    await dispatch("fetchUsers");
    await dispatch("fetchContests");
    await dispatch("fetchProblems");
    await dispatch("fetchSubmissions");
    await dispatch("fetchNotifications");
  },
  async login({ commit, dispatch }, token) {
    commit("setToken", token);
    commit("setLoggedInStatus", true);
    localStorage.setItem("token", token);
    const decode = jwtDecode(token);
    await dispatch("fetchUserInfo", decode.user_id);
  },
  logout({ commit }) {
    commit("setToken", "");
    commit("setLoggedInStatus", false);
    localStorage.removeItem("token");
  },
  async fetchNotifications({ commit }) {
    const response = await api.getNotifications();
    commit("setNotifications", response.data);
  },
  async fetchProblems({ commit }) {
    const response = await api.getProblems();
    commit("setProblems", response.data);
  },
  async fetchContests({ commit }) {
    const response = await api.getContests();
    commit("setContests", response.data);
  },
  async fetchRanks({ commit }) {
    // TODO:use info from store.users, order users by rank
  },
  async fetchSubmissions({ commit }) {
    const response = await api.getSubmissions();
    commit("setSubmissions", response.data);
  },
  async fetchUsers({ commit }) {
    const response = await api.getUsers();
    commit("setUsers", response.data);
  },
  async fetchUserInfo({ commit }, userID) {
    const response = await api.getUserInfo(userID);
    commit("setUserInfo", response.data);
  },
  async deleteUser({ commit }, userID) {
    await api.deleteUser(userID);
    commit("deleteUser", userID);
  },
  async createProblem({ commit }, problem) {
    const response = await api.addProblems(problem);
    commit("addProblem", response.data);
  },
  async updateProblem({ commit }, problem) {
    const response = await api.updateProblem(problem, problem.id);
    commit("updateProblem", response.data);
  },
  async deleteProblem({ commit }, problemId) {
    await api.deleteProblem(problemId);
    commit("deleteProblem", problemId);
  },
  async createNotification({ commit }, notification) {
    const response = await api.addNotifications(notification);
    commit("addNotification", response.data);
  },
  async updateNotification({ commit }, notification) {
    const response = await api.updateNotification(
      notification,
      notification.id
    );
    commit("updateNotification", response.data);
  },
  async deleteNotification({ commit }, notificationId) {
    await api.deleteNotification(notificationId);
    commit("deleteNotification", notificationId);
  },
  async createContest({ commit }, contest) {
    const response = await api.addContests(contest);
    commit("addContest", response.data);
  },
  async updateContest({ commit }, contest) {
    const response = await api.updateContest(contest, contest.id);
    commit("updateContest", response.data);
  },
  async deleteContest({ commit }, contestId) {
    await api.deleteContest(contestId);
    commit("deleteContest", contestId);
  },
};
