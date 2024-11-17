const apiUrl = useRuntimeConfig().public.apiUrl;

export default defineEventHandler(async (event) => {
  const { secure } = await requireUserSession(event);

  return await $fetch(`${apiUrl}/users/me`, {
    headers: { Authorization: `Bearer ${secure?.idToken}` },
  });
});
