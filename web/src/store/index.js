import { createStore } from "vuex";
import info from "./modules/info.js";
import user from "./modules/user.js";
import contest from "./modules/contest.js";
import notification from "./modules/notification.js";
import problem from "./modules/problem.js";
import submission from "./modules/submission.js";
import createPersistedState from "vuex-persistedstate";
import jwtDecode from "jwt-decode";

export default createStore({
	actions: {
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
		async fetchRanks({ commit }) {
			// TODO:use info from store.users, order users by rank
		},
	},
	modules: {
		problem, 
		user,
		info,
		contest,
		notification,
		submission
	});
