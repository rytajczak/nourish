import type { AuthConfig } from "@auth/core/types";
import { NuxtAuthHandler } from "#auth";

const runtimeConfig = useRuntimeConfig();

export const authOptions: AuthConfig = {
  secret: runtimeConfig.authJs.secret,
  basePath: "/api/auth",
  providers: [],
};

export default NuxtAuthHandler(authOptions, runtimeConfig);
