import { User } from "#auth-utils";

export default defineEventHandler(async (event) => {
  const idToken = getCookie(event, "idToken");
  const user = await $fetch<User>(`http://localhost:8080/users/me`, {
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
  });
  return user;
});
