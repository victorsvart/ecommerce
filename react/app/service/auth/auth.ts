import { z } from "zod";
import { ApiError, baseApi } from "~/api/axios";
import { ValidateForm } from "~/errors/form-validator";

const LoginInput = z.object({
  email: z.string().email("Endereço de email inválido"),
  password: z.string(),
});

export class AuthService {
  async login(formData: FormData) {
    const result = await ValidateForm(LoginInput, formData);
    const input = result.data;
    const res = await baseApi.post("/auth/login", {
      email: input.email,
      password: input.password,
    });

    const cookies = res.headers["set-cookie"];
    if (!cookies) {
      throw new ApiError(
        "Authentication successful but no session cookie received?",
        500
      );
    }

    return {
      cookies: Array.isArray(cookies) ? cookies.join(", ") : cookies,
    };
  }
}

export const authService = new AuthService();
