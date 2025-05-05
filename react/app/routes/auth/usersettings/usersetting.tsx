import makeApi, { ApiError } from "~/api/axios";
import type { Route } from "./+types/usersetting";
import Profile, { SkeletonProfile } from "~/components/settings/Profile";
import {
  Await,
  data,
  redirect,
  useActionData,
  useRouteError,
} from "react-router";
import { ValidateForm, ValidationError } from "~/errors/form-validator";
import { z } from "zod";
import { Suspense, useEffect, useState } from "react";
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
  name: z.string().min(1, "Nome é obrigatório!"),
  surname: z.string().min(1, "Sobrenome é obrigatório!"),
  email: z.string().email("Email inválido"),
  contact: z.string().regex(/^\(\d{2}\)\s9\d{4}-\d{4}$/, {
    message: "Número de telefone inválido",
  }),
});

export async function loader({ request }: Route.LoaderArgs) {
  const api = makeApi(request.headers);
  return {
    userPromise: api
      .get("/users")
      .then((response) => response.data as UserSettings),
  };
}
export async function action({ request }: Route.ActionArgs) {
  const formData = await request.formData();
  try {
    const result = await ValidateForm(UserEdit, formData);
    const { name, surname, email, contact } = result.data;
    const api = makeApi(request.headers);

    await api.put("/users", { name, surname, email, contact });
    return redirect("/userSettings");
  } catch (error) {
    if (error instanceof ApiError || error instanceof ValidationError) {
      return data({
        data: null,
        status: error.cause,
        statusText: error.message,
      });
    }

    console.error("Unexpected error:", error);
    return data({
      data: null,
      status: 500,
      statusText: "Unexpected error",
    });
  }
}

export default function UserSettings({ loaderData }: Route.ComponentProps) {
  const actionData = useActionData();
  console.log(actionData);
  const [showError, setShowError] = useState(true);

  useEffect(() => {
    setShowError(true);
  }, [actionData]);

  return (
    <div className="min-h-screen bg-gray-900 text-white p-10">
      <div className="max-w-3xl mx-auto">
        <h1 className="text-2xl font-semibold mb-6">Perfil</h1>

        <Suspense fallback={<SkeletonProfile />}>
          <Await resolve={loaderData.userPromise}>
            {(res) => <Profile user={res} />}
          </Await>
        </Suspense>
      </div>

      {actionData?.statusText && showError && (
        <div className="w-94 mx-auto mt-4">
          <AlertError
            message={actionData.statusText}
            onClose={() => setShowError(false)}
          />
        </div>
      )}
    </div>
  );
}
