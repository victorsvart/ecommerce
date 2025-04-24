import createApi from "~/api/axios";
import type { Route } from "./+types/home";
import { redirect } from "react-router";

export async function loader({ request }: Route.LoaderArgs) {
  const api = createApi(request.headers);
  return api
    .get("/auth/me", {
      headers: { cookie: request.headers.get("cookie") },
    })
    .then(() => {
      return null;
    })
    .catch((error: any) => {
      return redirect("/login");
    });
}

export default function Home() {
  return <div>Hello world</div>;
}
