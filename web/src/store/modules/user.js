const user={
	state:{
		userInfo: {
			// id: 0,
			// username: "",
			// nickname: "",
			// email: "",
			// avatar: "",
			// role: "",
		},
	},
	mutations:{
		setUsers(state, users) {
			state.users = users;
		},
		setUserInfo(state, userInfo) {
			state.userInfo = userInfo;
		},
		deleteUser(state, userID) {
			state.users = state.users.filter((user) => user.id !== userID);
		},
		setUserInfo({commit}, userInfo){
			commit("setUserInfo", userInfo);
		},
	},
	actions:{
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
	},
	getters:{
		getUsers: (state) => {
			return state.users;
		},
		getUserInfo: (state) => {
			return state.userInfo;
		},
		getUserById: (state) => (id) => {
			return state.users.find((user) => user.id == id);
		},
		getUserInfo(state){
			return state.userInfo;
		},
	},
}

export default user;
