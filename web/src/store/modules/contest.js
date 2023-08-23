import api from "../../api";

const contest = {
	state: {
		contests: {},
	},
	mutations: {
		setContests(state, contests) {
			state.contests = contests;
		},
		addContest(state, contest) {
			state.contests.push(contest);
		},
		updateContest(state, updatedContest) {
			const index = state.contests.findIndex(
				(contest) => contest.id == updatedContest.id
			);
			if (index != -1) {
				state.contests.splice(index, 1, updatedContest);
			}
		},
		deleteContest(state, contestId) {
			state.contests = state.contests.filter(
				(contest) => contest.id != contestId
			);
		},
	},
	actions: {
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
		async fetchContests({ commit }) {
			const response = await api.getContests();
			commit("setContests", response.data);
		},
		async fetchContest({ commit }, contestId) {
			const response = await api.getContest(contestId);
			commit("setContest", response.data);
		},
	},
	getters: {
		getContests: (state) => {
			return state.contests;
		},
		getContestById: (state) => (id) => {
			return state.contests.find((contest) => contest.id == id);
		},
	}
};

export default contest;
