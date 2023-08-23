import api from '../../api'

const submission = {
	state: {
		submission: {},
	},
	mutations: {
		setSubmissions(state, submissions) {
			state.submissions = submissions;
		},
	},
	actions: {
		async fetchSubmissions({ commit }) {
			const response = await api.getSubmissions();
			commit("setSubmissions", response.data);
		},
	},
	getters: {
		getSubmissions: (state) => {
			return state.submissions;
		},
		getSubmissionById: (state) => (id) => {
			return state.submissions.find((submission) => submission.id == id);
		},
	}
}

export default submission
