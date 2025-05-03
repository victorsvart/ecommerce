import createApi, { ApiError } from "~/api/axios";
import type { Route } from "./+types/usersetting";
import Profile from "~/components/settings/Profile";
import { data, redirect, useActionData } from "react-router";
import { ValidateForm } from "~/errors/form-validator";
import { z } from "zod";
import { useEffect, useState } from "react";
import { AlertError } from "~/components/AlertError";

export interface LoaderData {
  data: UserSettings | null;
  status: number;
  statusText: string | null;
}

export interface UserSettings {
  fullName: string;
  name: string;
  surname: string;
  email: string;
  contact: string;
}

const UserEdit = z.object({
  name: z.string(),
  surname: z.string(),
  email: z.string().email("Email inválido"),
  contact: z.string().regex(/^\(\d{2}\)\s9\d{4}-\d{4}$/, {
    message: "Número de telefone inválido",
  }),
});

export async function loader({ request }: Route.LoaderArgs) {
  const api = createApi(request.headers);
  return await api
    .get("/users", {
      headers: { cookie: request.headers.get("cookie") },
    })
    .then((res) => {
      return data({
        data: res.data as UserSettings,
        status: 200,
        statusText: null,
      } as LoaderData);
    })
    .catch((error) => {
      if (error instanceof ApiError) {
        return data({
          data: null,
          status: error.cause,
          statusText: error.message,
        } as LoaderData);
      }

      console.error("error:", error);
      return data({
        data: null,
        status: error.cause,
        statusText: error.message,
      } as LoaderData);
    });
}

export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  let result = await ValidateForm(UserEdit, formData);
  const input = result.data;

  const api = createApi();
  return await api
    .put(
      "/users",
      {
        name: input.name,
        surname: input.surname,
        email: input.email,
        contact: input.contact,
      },
      {
        headers: { cookie: request.headers.get("cookie") },
      }
    )
    .then(() => redirect("/userSettings"))
    .catch((error) => {
      if (error instanceof ApiError) {
        return data({
          data: null as UserSettings | null,
          status: error.cause,
          statusText: error.message,
        } as LoaderData);
      }

      console.error("error:", error);
      return data({
        data: null as UserSettings | null,
        status: error.cause,
        statusText: error.message,
      } as LoaderData);
    });
}

export default function UserSettings({ loaderData }: Route.ComponentProps) {
  const actionData = useActionData<LoaderData>();
  const [showError, setShowError] = useState(true);
  useEffect(() => {
    setShowError(true);
  }, [actionData]);

  return (
    <div className="min-h-screen bg-gray-900 text-white p-10">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-2xl font-semibold mb-6">Perfil</h1>
        <Profile user={loaderData.data} />
      </div>
      <div className="m-4">
        {actionData?.statusText && showError && (
          <AlertError
            message={actionData.statusText}
            onClose={() => setShowError(false)}
          />
        )}
      </div>
    </div>
  );
}
