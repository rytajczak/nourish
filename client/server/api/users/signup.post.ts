const apiUrl = useRuntimeConfig().public.apiUrl;

interface SignupResponse {
  profile: {
    diet: string;
    calories: number;
    protein: number;
    carbs: number;
    fat: number;
  };
  intolerances: string[];
  dislikedIngredients: string[];
  savedRecipes: any[];
  spoonCredential: {
    username: string;
    hash: string;
  };
}

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

  console.log(body);

  try {
    const response = await $fetch<SignupResponse>(`${apiUrl}/users/signup`, {
      method: "POST",
      headers: { Authorization: `Bearer ${secure?.idToken}` },
      body,
    });

    await setUserSession(event, {
      secure: {
        spoonName: response.spoonCredential.username,
        spoonHash: response.spoonCredential.hash,
      },
    });
    return true;
  } catch (error) {
    return false;
  }
});
