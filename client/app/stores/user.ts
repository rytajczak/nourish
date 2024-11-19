interface Profile {
  diet: string;
  calories: number;
  protein: number;
  carbs: number;
  fat: number;
}

export const useUserStore = defineStore(
  "user",
  () => {
    const { clear } = useUserSession();

    /**
     * The user's profile
     */
    const profile = ref<Profile>({} as Profile);

    /**
     * The user's intolerances
     */
    const intolerances = ref<string[]>([]);

    /**
     * The user's disliked ingredients
     */
    const dislikedIngredients = ref<string[]>([]);

    /**
     * The user's saved recipes
     */
    const savedRecipes = ref<any[]>([]);

    /**
     * Get the user's profile, intolerances, disliked ingredients, and saved recipes
     * @returns Whether the user was successfully fetched
     */
    async function getUser(): Promise<boolean> {
      const response = await $fetch("/api/users/me");
      if (!response) return false;

      profile.value = response.profile;
      intolerances.value = response.intolerances;
      dislikedIngredients.value = response.dislikedIngredients;
      savedRecipes.value = response.savedRecipes;
      return true;
    }

    /**
     * Log out the user and clear the user session
     */
    async function logOut() {
      profile.value = {} as Profile;
      intolerances.value = [];
      savedRecipes.value = [];
      dislikedIngredients.value = [];
      await clear();
      navigateTo("/");
    }

    return {
      profile,
      intolerances,
      dislikedIngredients,
      savedRecipes,
      getUser,
      logOut,
    };
  },
  { persist: true },
);
