const config = useRuntimeConfig();

export default defineEventHandler(async (event) => {
  const idToken = getCookie(event, "idToken");
  const rawBody = await readBody(event);
  const body = {
    diet: rawBody.diet,
    calories: parseInt(rawBody.calories),
    protein: parseInt(rawBody.protein),
    carbs: parseInt(rawBody.carbs),
    fat: parseInt(rawBody.fat),
    intolerances: rawBody.intolerances,
  };

  const res = await $fetch(`${config.public.apiUrl}/users/me/preferences`, {
    method: "PUT",
    headers: {
      Authorization: `Bearer ${idToken}`,
    },
    body,
  });
  console.log(res);
});
