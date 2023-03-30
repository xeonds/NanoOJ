import { AxiosRequestConfig } from 'axios'
import v_axios from '../axios/index.js'

export default {
    login: (data: any = {}) => v_axios.post('/user/login', data),
    register: (data: any = {}) => v_axios.post('/user/register', data),

    getUsers: () => v_axios.get("/users"),
    getProblems: () => v_axios.get("/problems"),
    getContests: () => v_axios.get("/contests"),
    getSubmissions: () => v_axios.get("/submissions"),
    getNotifications: () => v_axios.get("/notifications"),
    getUserInfo: (id: number) => v_axios.get(`/users/${id}`),
    getProblemInfo: (id: number) => v_axios.get(`/problems/${id}`),
    getContestInfo: (id: number) => v_axios.get(`/contests/${id}`),
    getSubmissionInfo: (id: number) => v_axios.get(`/submissions/${id}`),
    getNotificationInfo: (id: number) => v_axios.get(`/notifications/${id}`),

    // need token
    addProblems: (data: any) => v_axios.post("/problems", data, { needToken: true } as RequestConfig),
    addContests: (data: any) => v_axios.post("/contests", data, { needToken: true } as RequestConfig),
    addSubmissions: (data: any) => v_axios.post("/submissions", data, { needToken: true } as RequestConfig),
    addNotifications: (data: any) => v_axios.post("/notifications", data, { needToken: true } as RequestConfig),
    updateUser: (data: any, id: number) => v_axios.put(`/users/${id}`, data, { needToken: true } as RequestConfig),
    updateProblem: (data: any, id: number) => v_axios.put(`/problems/${id}`, data, { needToken: true } as RequestConfig),
    updateContests: (data: any, id: number) => v_axios.put(`/contests/${id}`, data, { needToken: true } as RequestConfig),
    updateNotification: (data: any, id: number) => v_axios.put(`/notifications/${id}`, data, { needToken: true } as RequestConfig),
    deleteUser: (id: number) => v_axios.delete(`/users/${id}`, { needToken: true } as RequestConfig),
    deleteProblem: (id: number) => v_axios.delete(`/problems/${id}`, { needToken: true } as RequestConfig),
    deleteContest: (id: number) => v_axios.delete(`/contests/${id}`, { needToken: true } as RequestConfig),
    deleteNotification: (id: number) => v_axios.delete(`/notifications/${id}`, { needToken: true } as RequestConfig),
}

interface RequestConfig extends AxiosRequestConfig {
    needToken?: boolean
}