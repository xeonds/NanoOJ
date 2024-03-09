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
    addProblems: (data: any) => http.post("/problems", data, { needToken: true } as RequestConfig),
    addContests: (data: any) => http.post("/contests", data, { needToken: true } as RequestConfig),
    addSubmissions: (data: any) => http.post("/submissions", data, { needToken: true } as RequestConfig),
    addNotifications: (data: any) => http.post("/notifications", data, { needToken: true } as RequestConfig),
    updateUser: (data: any, id: number) => http.put(`/users/${id}`, data, { needToken: true } as RequestConfig),
    updateProblem: (data: any, id: number) => http.put(`/problems/${id}`, data, { needToken: true } as RequestConfig),
    updateContests: (data: any, id: number) => http.put(`/contests/${id}`, data, { needToken: true } as RequestConfig),
    updateNotification: (data: any, id: number) => http.put(`/notifications/${id}`, data, { needToken: true } as RequestConfig),
    deleteUser: (id: number) => http.delete(`/users/${id}`, { needToken: true } as RequestConfig),
    deleteProblem: (id: number) => http.delete(`/problems/${id}`, { needToken: true } as RequestConfig),
    deleteContest: (id: number) => http.delete(`/contests/${id}`, { needToken: true } as RequestConfig),
    deleteNotification: (id: number) => http.delete(`/notifications/${id}`, { needToken: true } as RequestConfig),
}

interface RequestConfig extends AxiosRequestConfig {
    needToken?: boolean
}