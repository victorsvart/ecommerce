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
  return (
    <div>
      <div>
        <p>{loaderData.data.fullName}</p>
        <p>{loaderData.data.name}</p>
        <p>{loaderData.data.surname}</p>
        <p>{loaderData.data.email}</p>
      </div>
    </div>
  );
}
