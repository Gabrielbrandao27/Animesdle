import axios from "axios";

const axiosInstance = axios.create({
  baseURL: "http://localhost:8080", // seu backend em Go
  headers: {
    "Content-Type": "application/json",
  },
});

export default axiosInstance;
