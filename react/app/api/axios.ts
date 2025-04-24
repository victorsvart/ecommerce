import axios from "axios";
import * as fs from "fs";

const api = axios.create({
  baseURL: "http://localhost:8080/v1/api",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
});

export default api;
