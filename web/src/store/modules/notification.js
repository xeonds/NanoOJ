import api from '../../api';

const notification = {
	state: {
		notification: {
			message: '',
			type: '',
			show: false
		}
	},
	mutations: {
		addNotification(state, notification) {
			state.notifications.push(notification);
		},
		updateNotification(state, updatedNotification) {
			const index = state.notifications.findIndex(
				(notification) => notification.id === updatedNotification.id
			);
			if (index !== -1) {
				state.notifications.splice(index, 1, updatedNotification);
			}
		},
		deleteNotification(state, notificationId) {
			state.notifications = state.notifications.filter(
				(notification) => notification.id !== notificationId
			);
		},
		setNotifications(state, notifications) {
			state.notifications = notifications;
		},
	},
	actions: {
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
		async fetchNotifications({ commit }) {
			const response = await api.getNotifications();
			commit("setNotifications", response.data);
		},
	},
	getters: {
		getNotificationById: (state) => (id) => {
			return state.notifications.find((notification) => notification.id == id);
		},
		getNotifications: (state) => {
			return state.notifications;
		},
	},
}

export default notification;
