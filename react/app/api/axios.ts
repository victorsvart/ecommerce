import axios, {
  AxiosError,
  type AxiosRequestConfig,
  type AxiosResponse,
} from "axios";

export class ApiError extends Error {
  constructor(message: string, public status: number) {
    super(message);
    this.name = "ApiError";
    this.cause = status;
  }
}

// Base configuration
const baseConfig = {
  baseURL: process.env.PUBLIC_API_URL || "http://localhost:8080/v1/api",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
  },
};

// Create a client-side instance (used in browser)
export const baseApi = axios.create(baseConfig);

// Error interceptor
baseApi.interceptors.response.use(
  (response: Promise<AxiosResponse>) => response,
  (error: AxiosError) => {
    // Handle timeout errors
    if (error.code === "ECONNABORTED") {
      throw new ApiError("Request timed out. Please try again.", 408);
    }

    // Handle network errors
    if (!error.response) {
      throw new ApiError("Network error. Please check your connection.", 500);
    }

    // Extract error details from response
    const status = error.response.status;
    let message: string;

    const data = error.response.data as any;
    if (typeof data === "string") {
      message = data;
    } else if (data?.message) {
      message = data.message;
    } else if (data?.error) {
      message = data.error;
    } else {
      message = `Request failed with status ${status}`;
    }

    throw new ApiError(message, status);
  }
);

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
