import { data } from "react-router";
import type { Route } from "./+types/usersetting";
import createApi, { ApiError } from "~/api/axios";

interface UserSettings {
  fullName: string;
  name: string;
  surname: string;
  email: string;
}

interface LoaderData {
  data: UserSettings;
  status: number;
  statusText: string;
}

export async function loader({ request }: Route.LoaderArgs) {
  const api = createApi(request.headers);
  return await api
    .get("/users", {
      headers: { cookie: request.headers.get("cookie") },
    })
    .then((res) => {
      return data({ data: res.data as UserSettings }, {
        status: 200,
        statusText: "",
      } as LoaderData);
    })
    .catch((error) => {
      if (error instanceof ApiError) {
        return data({ data: {} as UserSettings }, {
          status: error.cause as number,
          statusText: error.message,
        } as LoaderData);
      }

      console.error("error:", error);
      return data({ data: {} as UserSettings }, {
        status: error.cause as number,
        statusText: error.message,
      } as LoaderData);
    });
}

export default function UserSettings({ loaderData }: Route.ComponentProps) {
  const user = loaderData.data;

  return (
    <div className="min-h-screen bg-gray-900 text-white p-10">
      <h1 className="text-2xl font-semibold mb-6">Profile</h1>

      <div className="flex flex-col md:flex-row bg-gray-800 p-6 rounded-lg shadow-lg gap-8">
        <div className="flex flex-col items-center">
          <div className="w-24 h-24 rounded-full bg-gray-700 flex items-center justify-center text-2xl font-bold">
            {user.name?.charAt(0).toUpperCase()}
          </div>
          <button className="mt-4 text-sm text-red-400 hover:text-red-500">
            Remove photo
          </button>
        </div>

        <form className="flex-1 grid grid-cols-1 md:grid-cols-2 gap-4">
          <div className="md:col-span-2">
            <label className="block text-sm text-gray-400">
              Nome de usuário
            </label>
            <input
              disabled
              value={user.fullName}
              className="w-full brightness-50 p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
            />
            <p className="text-xs text-gray-500 mt-1">
              * Você não pode mudar seu nome de usuário
            </p>
          </div>

          <div className="w-full flex gap-4">
            <div className="flex-1 w-full">
              <label className="block text-sm text-gray-400">
                Primeiro nome
              </label>
              <input
                value={user.name}
                className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
            </div>
            <div className="flex-1 w-full">
              <label className="block text-sm text-gray-400">Sobrenome</label>
              <input
                value={user.surname}
                className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white"
              />
            </div>
          </div>
          <div className="md:col-span-2">
            <label className="block text-sm text-gray-400">Bio</label>
            <textarea
              rows={3}
              className="w-full p-2 bg-gray-700 border border-gray-600 rounded-md text-white resize-none"
              placeholder="Fale sobre você..."
            />
            <p className="text-right text-xs text-gray-400 mt-1">
              111 characters remaining
            </p>
          </div>

          <div className="md:col-span-2 flex justify-between items-center">
            <a
              href="/profile/public"
              className="text-sm text-blue-400 hover:underline"
            >
              <p>Exibir perfil público</p>
            </a>
            <button
              type="submit"
              className="bg-blue-600 hover:bg-blue-700 text-white font-medium py-2 px-4 rounded-md"
            >
              Salvar
            </button>
          </div>
        </form>
      </div>
    </div>
  );
}
