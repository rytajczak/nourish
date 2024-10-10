import type { AuthConfig } from "@auth/core/types";
import { NuxtAuthHandler } from "#auth";
import Credentials from "@auth/core/providers/credentials";
import Google from "@auth/core/providers/google";

const runtimeConfig = useRuntimeConfig();

export const authOptions: AuthConfig = {
  secret: runtimeConfig.authJs.secret,
  basePath: "/api/auth",
  providers: [
    Credentials({
      credentials: {
        email: {},
        password: {},
      },
      authorize: async (Credentials) => {
        console.log(Credentials);
        // TODO: think harder on what this should return
        return { email: "ryantratajczak@gmail.com" };
      },
    }),
    Google({
      clientId: runtimeConfig.google.clientId,
      clientSecret: runtimeConfig.google.clientSecret,
    }),
  ],
};

export default NuxtAuthHandler(authOptions, runtimeConfig);
