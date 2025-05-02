import { Outlet, redirect } from "react-router";
import type { Route } from "../../+types/root";
import createApi from "~/api/axios";
import { Navbar } from "~/components/Nav";

export async function loader({ request }: Route.LoaderArgs) {
  const api = createApi(request.headers);
  return api
    .get("/auth/me", {
      headers: { cookie: request.headers.get("cookie") },
    })
    .then(() => null)
    .catch((error: any) => {
      const params = new URLSearchParams({
        errors: "Session expired. Please login again.",
      });
      return redirect(`/login?${params.toString()}`);
    });
}

export default function Authenticated({ loaderData }: Route.ComponentProps) {
  return (
    <div className="min-h-screen w-full bg-gray-900 text-white">
      <Navbar />
      <main className="w-full">
        <Outlet />
      </main>
    </div>
  );
}
