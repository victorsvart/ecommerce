import {
  type RouteConfig,
  index,
  layout,
  route,
} from "@react-router/dev/routes";

export default [
  index("routes/home.tsx"),
  route("login", "./routes/login/Login.tsx"),
  layout("./routes/auth/auth.tsx", [
    route("userSettings", "./routes/auth/usersettings/usersettings.tsx"),
    route("products", "./routes/auth/products/products.tsx"),
  ]),
] satisfies RouteConfig;
