import { parse } from "cookie";
export async function getSession(request: Request): Promise<string | null> {
  const cookieHeader = request.headers.get("Cookie");
  const cookies = parse(cookieHeader || "");
  return cookies.auth_token;
}

export async function setSession(request: Request): Promise<void> {
  request.headers.get("Cookie");
}
