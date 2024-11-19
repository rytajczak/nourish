export interface Profile {
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

    const profile = ref<Profile>({} as Profile);
    const intolerances = ref<string[]>([]);
    const dislikedIngredients = ref<string[]>([]);
    const savedRecipes = ref<any[]>([]);

    async function getUser() {
      const response = await $fetch("/api/users/me");
      console.log(response);
    }

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
