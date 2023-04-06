export default {
  //auth
  setToken(state, payload) {
    state.token = payload;
    state.isLoggedIn = true;
  },
  setLoggedInStatus(state, isLoggedIn) {
    state.isLoggedIn = isLoggedIn;
  },
  setNotifications(state, notifications) {
    state.notifications = notifications;
  },
  setProblems(state, problems) {
    state.problems = problems;
  },
  setRanks(state, ranks) {
    state.ranks = ranks;
  },
  setSubmissions(state, submissions) {
    state.submissions = submissions;
  },
  setUsers(state, users) {
    state.users = users;
  },
  setUserInfo(state, userInfo) {
    state.userInfo = userInfo;
  },
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
};
