import type { RecipePreview } from "~~/server/api/recipes/search.get";

export const useUserStore = defineStore(
  "user",
  () => {
    const { user, clear } = useUserSession();
    const intolerances = ref<string[]>([]);
    const savedRecipes = ref<RecipePreview[]>([]);
    const dislikedIngredients = ref<string[]>([]);

    async function logOut() {
      await clear();
      navigateTo("/");
    }

    async function fetchSavedRecipes() {}

    return { user, intolerances, dislikedIngredients, savedRecipes, logOut };
  },
  { persist: true },
);
