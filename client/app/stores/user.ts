export const useUserStore = defineStore(
  "user",
  () => {
    const { user, clear } = useUserSession();
    const intolerances = ref<string[]>([]);
    const savedRecipes = ref<any[]>([]);
    const dislikedIngredients = ref<string[]>([]);

    async function logOut() {
      intolerances.value = [];
      savedRecipes.value = [];
      dislikedIngredients.value = [];
      await clear();
      navigateTo("/");
    }

    return { user, intolerances, dislikedIngredients, savedRecipes, logOut };
  },
  { persist: true },
);
