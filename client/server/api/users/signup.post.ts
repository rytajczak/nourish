const apiUrl = useRuntimeConfig().public.apiUrl;

export default defineEventHandler(async (event) => {
  const { user, secure } = await requireUserSession(event);
  const rawBody = await readBody(event);

  const body = {
    ...user,
    diet: rawBody.diet,
    calories: parseInt(rawBody.calories),
    protein: parseInt(rawBody.protein),
    carbs: parseInt(rawBody.carbs),
    fat: parseInt(rawBody.fat),
    intolerances: rawBody.intolerances,
  };

  return await $fetch(`${apiUrl}/users/signup`, {
    method: "POST",
    headers: { Authorization: `Bearer ${secure?.idToken}` },
    body,
  });
});
