export const useUserStore = defineStore(
  "user",
  () => {
    const profile = ref<Profile | null>(null);
    const intolerances = ref<string[]>([]);
    const savedRecipes = ref<any[]>([]);

    const setUser = (user: {
      profile: Profile;
      intolerances: string[];
      savedRecipes: any[];
    }) => {
      profile.value = user.profile;
      intolerances.value = user.intolerances ?? [];
      savedRecipes.value = user.savedRecipes ?? [];
    };

    const logOut = () => {
      profile.value = null;
      intolerances.value = [];
      savedRecipes.value = [];
    };

    return { profile, intolerances, savedRecipes, setUser, logOut };
  },
  { persist: true },
);
