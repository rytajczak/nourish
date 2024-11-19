const apiUrl = useRuntimeConfig().public.apiUrl;

interface GetMeResponse {
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
  const { secure } = await requireUserSession(event);

  try {
    const response = await $fetch<GetMeResponse>(`${apiUrl}/users/me`, {
      headers: { Authorization: `Bearer ${secure?.idToken}` },
    });

    await setUserSession(event, {
      secure: {
        spoonName: response.spoonCredential.username,
        spoonHash: response.spoonCredential.hash,
      },
    });

    return {
      profile: response.profile,
      intolerances: response.intolerances,
      dislikedIngredients: response.dislikedIngredients,
      savedRecipes: response.savedRecipes,
    };
  } catch (error) {
    return null;
  }
});
