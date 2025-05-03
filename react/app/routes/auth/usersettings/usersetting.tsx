import createApi, { ApiError } from "~/api/axios";
import type { Route } from "./+types/usersetting";
import Profile from "~/components/settings/Profile";
import { data } from "react-router";

interface LoaderData {
  data: UserSettings;
  status: number;
  statusText: string;
}

export interface UserSettings {
  fullName: string;
  name: string;
  surname: string;
  email: string;
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
  return (
    <div className="min-h-screen bg-gray-900 text-white p-10">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-2xl font-semibold mb-6">Perfil</h1>
        <Profile user={loaderData.data} />
      </div>
    </div>
  );
}
