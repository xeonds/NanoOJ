import axios from "axios";
import jwtDecode from "jwt-decode";

const v_axios = axios.create({
  baseURL: "/api/v1/",
});
v_axios.interceptors.request.use(
  (req) => {
    console.log("configuring");
    if (!req.needToken) return req;

    const token = localStorage.getItem("token");
    if (!token) return Promise.reject("No token");
    const data = jwtDecode(token);
    if (data.exp <= Date.now() / 1000) return Promise.reject("Token expired");
    req.headers.Authorization = token;

    return req;
  },
  (err) => {
    console.log("error");
    return Promise.reject(err);
  }
);
v_axios.interceptors.response.use(
  (req) => {
    return req;
  },
  (err) => {
    console.log("error");
    return Promise.reject(err);
  }
);

export default v_axios;
