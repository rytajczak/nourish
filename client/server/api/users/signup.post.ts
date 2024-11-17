const apiUrl = useRuntimeConfig().public.apiUrl;

export default defineEventHandler(async (event) => {
  const { user, secure } = await requireUserSession(event);

  const body = {
    ...user,
  };

  return await $fetch(`${apiUrl}/users/signup`, {
    method: "POST",
    headers: { Authorization: `Bearer ${secure?.idToken}` },
    body,
  });
});
