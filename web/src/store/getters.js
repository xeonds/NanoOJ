export default {
  getToken: (state) => state.token,
  getLoggedInStatus: (state) => state.isLoggedIn,
  getNotificationById: (state) => (id) => {
    return state.notifications.find((notification) => notification.id == id);
  },
  getNotifications: (state) => {
    return state.notifications;
  },
  getProblemById: (state) => (id) => {
    return state.problems.find((problem) => problem.ProblemID == id);
  },
  getProblems: (state) => {
    return state.problems;
  },
  getRankById: (state) => (id) => {
    return state.ranks.find((rank) => rank.id == id);
  },
  getRanks: (state) => {
    return state.ranks;
  },
  getSubmissions: (state) => {
    return state.submissions;
  },
  getSubmissionById: (state) => (id) => {
    return state.submissions.find((submission) => submission.id == id);
  },
  getUserInfo: (state) => {
    return state.userInfo;
  },
  getUserById: (state) => (id) => {
    return state.users.find((user) => user.id == id);
  },
  getContests: (state) => {
    return state.contests;
  },
};
