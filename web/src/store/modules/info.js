import api from '../../api'

const info = {
	state: {
		login: {
			token: "",
			isLoggedIn: false,
		},
	},
	mutations: {
		setToken(state, payload) {
			state.token = payload;
			state.isLoggedIn = true;
		},
		setLoggedInStatus(state, isLoggedIn) {
			state.isLoggedIn = isLoggedIn;
		},
	},
	actions: {
	},
	getters: {
		getToken: (state) => state.token,
		getLoggedInStatus: (state) => state.isLoggedIn,
	}
}

export default info
