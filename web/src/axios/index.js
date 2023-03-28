import axios from "axios";

const v_axios = axios.create({
  baseURL: "/api/v1/",
});

export default v_axios;
