import { useState } from "react";
import api from "../../api/axios";
import type { AxiosError } from "axios";
import { LoginInput } from "~/components/LoginInput";
import { AlertError } from "~/components/AlertError";
import { useForm, type SubmitHandler } from "react-hook-form";

type Input = {
  email: string;
  password: string;
};

export default function Login() {
  const [error, setError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<Input>();

  const onSubmit: SubmitHandler<Input> = async (data) => {
    try {
      await api.post("/auth/login", data);
    } catch (error) {
      const err = error as AxiosError<{ data: string }>;

      if (err.response) {
        setError(err.response.data.data);
      } else {
        setError("Um erro interno ocorreu. Tente novamente.");
      }
    }
  };

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

              <form
                onSubmit={handleSubmit(onSubmit)}
                className="shadow-md rounded-lg p-6"
              >
                <div className="text-center mb-4">
                  <h4 className="text-lg font-semibold text-gray-400">
                    Iniciar Sessão
                  </h4>
                </div>
                <LoginInput register={register} errors={errors} />
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

                {error && (
                  <div className="mt-4">
                    <AlertError
                      message={error}
                      onClose={() => setError(null)}
                    />
                  </div>
                )}
              </form>
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
