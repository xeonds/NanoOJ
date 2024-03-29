import { jwtDecode } from "jwt-decode";

const getter = (key: string) => () => {
    const value = window.sessionStorage.getItem(key);
    return value ? value : "";
};
const remove = (key: string) => () => {
    window.sessionStorage.removeItem(key);
};
const keys = ["token", "name", "id", "permission", "expire"];

export const isAdmin = () => {
    const role = getPermission();
    return role ? role == "0" || role == "1" : false;
};
export const isLogin = () => (Date.now() < new Date(getExpire()).getTime());
export const logout = () => {
    keys.forEach((key) => remove(key)());
    window.location.href = "/";
};
export const getToken = getter("token");
export const getUsername = getter("name");
export const getUserID = getter("id");
export const getPermission = getter("permission");
export const getExpire = getter("expire");
export const setToken  =  (value: string) => {
    window.sessionStorage.setItem("token", value);
    const { permission, name, id, expire } = jwtDecode(value) as { permission: number, name: string, id: number, expire: Date}
    window.sessionStorage.setItem("name", name);
    window.sessionStorage.setItem("id", id.toString());
    window.sessionStorage.setItem("permission", permission.toString());
    window.sessionStorage.setItem("expire", expire.toString());
};