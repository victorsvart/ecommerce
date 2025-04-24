import axios, { type AxiosRequestConfig } from "axios";

// Base configuration
const baseConfig = {
  baseURL: "http://localhost:8080/v1/api",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
};

// Create a client-side instance (used in browser)
export const baseApi = axios.create(baseConfig);

/**
 * Creates an API instance with the proper cookie handling for both client and server environments
 * @param requestHeaders Optional headers from the server request (for SSR)
 */
export default function createApi(requestHeaders?: Headers | null) {
  if (typeof window !== "undefined" || !requestHeaders) {
    return baseApi;
  }

  const cookie = requestHeaders.get("cookie");
  const serverConfig: AxiosRequestConfig = {
    ...baseConfig,
    headers: {
      ...baseConfig.headers,
      ...(cookie ? { cookie } : {}),
    },
  };

  return axios.create(serverConfig);
}
