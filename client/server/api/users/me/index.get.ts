import { User } from "#auth-utils";
const config = useRuntimeConfig();

export default defineEventHandler(async (event) => {
  const idToken = getCookie(event, "idToken");
  const user = await $fetch<User>(`${config.public.apiUrl}/users/me`, {
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
  })
  return user;
});
