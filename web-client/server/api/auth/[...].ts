import type { AuthConfig } from "@auth/core/types";
import { NuxtAuthHandler } from "#auth";
import Credentials from "@auth/core/providers/credentials";

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

        return { email: "ryantratajczak@gmail.com" };
      },
    }),
  ],
};

export default NuxtAuthHandler(authOptions, runtimeConfig);
