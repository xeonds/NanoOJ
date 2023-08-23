import api from '../../api';

const problem = {
	state: {
	},
	mutations: {
		addProblem(state, problem) {
			state.problems.push(problem);
		},
		updateProblem(state, updatedProblem) {
			const index = state.problems.findIndex(
				(problem) => problem.id == updatedProblem.id
			);
			if (index !== -1) {
				state.problems.splice(index, 1, updatedProblem);
			}
		},
		deleteProblem(state, problemId) {
			state.problems = state.problems.filter(
				(problem) => problem.ProblemID != problemId
			);
		},
		setProblems(state, problems) {
			state.problems = problems;
		},
	},
	actions: {
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
		async fetchProblems({ commit }) {
			const response = await api.getProblems();
			commit("setProblems", response.data);
		},
	},
	getters: {
		getProblemById: (state) => (id) => {
			return state.problems.find((problem) => problem.id == id);
		},
		getProblems: (state) => {
			return state.problems;
		},
	}
};

export default problem;
