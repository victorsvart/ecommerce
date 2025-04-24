import { redirect, Form, data } from "react-router";
import api from "../../api/axios";
import { LoginInput } from "~/components/LoginInput";
import type { Route } from "./+types/Login";
import { AlertError } from "~/components/AlertError";
import { useEffect, useState } from "react";

export async function loader() {
  try {
    await api.get("/auth/me");
    return redirect("/home");
  } catch {
    return null;
  }
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  const email = formData.get("email");
  const password = formData.get("password");

  return api
    .post("/auth/login", { email, password })
    .then(() => redirect("/home"))
    .catch((error: any) => {
      if (error.response) {
        return data({
          error: {
            status: error.response.status,
            message: error.response.data,
          },
        });
      }

      return data({
        error: {
          status: 500,
          message: "Unknown error",
        },
      });
    });
}

export default function Login({ actionData }: Route.ComponentProps) {
  const [showError, setShowError] = useState(true);

  useEffect(() => {
    setShowError(true);
  }, [actionData]);
  return (
    <div className="w-full">
      <div className="flex flex-col md:flex-row min-h-screen">
        <div className="w-full md:w-1/2 lg:w-1/4 dark:bg-gray-900">
          <div className="flex justify-center items-center h-screen">
            <div className="w-full px-4">
              <div className="flex justify-center items-center mb-6">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="40"
                  height="40"
                  fill="currentColor"
                  className="text-gray-400 me-2"
                  viewBox="0 0 16 16"
                >
                  <path d="m.334 0 4.358 4.359h7.15v7.15l4.358 4.358V0zM.2 9.72l4.487-4.488v6.281h6.28L6.48 16H.2z" />
                </svg>
                <h4 className="text-xl font-semibold text-gray-400 mb-0 ml-2">
                  ECommerce
                </h4>
              </div>

              <Form method="post" className="shadow-md rounded-lg p-6">
                <div className="text-center mb-4">
                  <h4 className="text-lg font-semibold text-gray-400">
                    Iniciar Sessão
                  </h4>
                </div>

                <LoginInput />

                <div className="mt-4">
                  <button
                    type="submit"
                    className="w-full dark:bg-blue-600 text-white py-2 px-4 rounded hover:bg-blue-700 transition"
                  >
                    Entrar
                  </button>
                </div>

                <hr className="my-6 border-gray-300" />

                <p className="text-center text-sm text-gray-300 mb-3">
                  É um recrutador e gostaria de testar a aplicação?
                </p>
                <button
                  type="button"
                  className="w-full border border-blue-600 text-blue-600 py-2 px-4 rounded hover:bg-blue-500 transition"
                >
                  <span className="hover:text-white-600">
                    Logar como recrutador
                  </span>
                </button>
                <div className="m-4">
                  {actionData?.error?.message && showError && (
                    <AlertError
                      message={actionData.error.message}
                      onClose={() => setShowError(false)}
                    />
                  )}
                </div>
              </Form>
            </div>
          </div>
        </div>

        <div className="w-full md:w-3/4 stylish-gradient">
          <div className="w-full h-full"></div>
        </div>
      </div>
    </div>
  );
}
