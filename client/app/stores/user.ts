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

    async function loadUser(data: any) {
      profile.value = data.profile;
      intolerances.value = data.intolerances;
      dislikedIngredients.value = data.dislikedIngredients;
      savedRecipes.value = data.savedRecipes;
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
      loadUser,
      logOut,
    };
  },
  { persist: true },
);
