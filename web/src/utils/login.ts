import { jwtDecode } from "jwt-decode";

const keys = ["token", "role", "username", "userid"];

const getter = (key: string) => () => {
    const value = window.sessionStorage.getItem(key);
    return value ? value : "";
};

const setter = (key: string) => (value: string) => {
    window.sessionStorage.setItem(key, value);
};

const remove = (key: string) => () => {
    window.sessionStorage.removeItem(key);
};

export const isAdmin = () => {
    const role = window.sessionStorage.getItem("role");
    return role ? role == "0" || role == "1" : false;
};

export const isLogin = () => {
    const token = window.sessionStorage.getItem("token");
    if (typeof token === "string") {
        const { exp } = jwtDecode(token) as { exp: number };
        if (Date.now() < exp * 1000) {
            return true;
        }
    }
    return false;
};

export const logout = () => {
    keys.forEach((key) => remove(key)());
    window.location.href = "/login";
};

export const getToken = getter("token");
export const setToken = setter("token");
export const getUsername = getter("username");
export const setUsername = setter("username");
export const getUserid = getter("userid");
export const setUserid = setter("userid");
export const getRole = getter("role");
export const setRole = setter("role");
