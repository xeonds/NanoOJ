import { http } from '@/utils/http'

export default {
    login: (data: any = {}) => http.post('/user/login', data),
    register: (data: any = {}) => http.post('/user/register', data),

    getUsers: () => http.get("/users"),
    getProblems: () => http.get("/problems"),
    getContests: () => http.get("/contests"),
    getSubmissions: () => http.get("/submissions"),
    getNotifications: () => http.get("/notifications"),
    getUserInfo: (id: number) => http.get(`/users/${id}`),
    getProblemInfo: (id: number) => http.get(`/problems/${id}`),
    getContestInfo: (id: number) => http.get(`/contests/${id}`),
    getSubmissionInfo: (id: number) => http.get(`/submissions/${id}`),
    getNotificationInfo: (id: number) => http.get(`/notifications/${id}`),

    // need token
    addProblems: (data: any) => http.post("/problems", data),
    addContests: (data: any) => http.post("/contests", data),
    addSubmissions: (data: any) => http.post("/actions/submit", data),
    addNotifications: (data: any) => http.post("/notifications", data),
    updateUser: (data: any, id: number) => http.put(`/users`, id, data),
    updateProblem: (data: any, id: number) => http.put(`/problems`, id, data),
    updateContests: (data: any, id: number) => http.put(`/contests`, id, data),
    updateNotification: (data: any, id: number) => http.put(`/notifications`, id, data),
    deleteUser: (id: number) => http.del('/users', id),
    deleteProblem: (id: number) => http.del('/problems', id),
    deleteContest: (id: number) => http.del('/contests', id),
    deleteNotification: (id: number) => http.del('/notifications', id),
}
