import { parse, stringify } from "zipson/lib";

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
    const profile = ref<Profile>({
      diet: "none",
      calories: 0,
      protein: 0,
      carbs: 0,
      fat: 0,
    });

    /**
     * The user's intolerances
     */
    const intolerances = ref<string[]>([]);

    /**
     * The user's saved recipes
     */
    const savedRecipes = ref<any[]>([]);

    /**
     *
     * @param data the data for creating a user
     */
    async function createUser(data: any) {
      const response = await $fetch("/api/users", {
        method: "POST",
        body: data,
      });
      if (!response) return false;

      profile.value = response.profile;
      intolerances.value = response.intolerances;
      savedRecipes.value = response.savedRecipes;
      return true;
    }

    /**
     * Get the user's profile, intolerances, disliked ingredients, and saved recipes
     * @returns Whether the user was successfully fetched
     */
    async function getUser() {
      const response = await $fetch("/api/users/me");
      if (!response) return false;

      profile.value = response.profile;
      intolerances.value = response.intolerances;
      savedRecipes.value = response.savedRecipes;
      return true;
    }

    async function updateProfile() {
      await $fetch("/api/users/me/profile");
    }

    const saveRecipe = debounce(() => {
      console.log("hello");
    }, 300);

    /**
     * Log out the user and clear the user session
     */
    async function signOut() {
      profile.value = {} as Profile;
      intolerances.value = [];
      savedRecipes.value = [];
      await clear();
      navigateTo("/");
    }

    return {
      profile,
      intolerances,
      savedRecipes,
      createUser,
      getUser,
      saveRecipe,
      updateProfile,
      signOut,
    };
  },
  {
    persist: {
      storage: sessionStorage,
      serializer: {
        serialize: (value) => stringify(value),
        deserialize: (value) => parse(value),
      },
    },
  },
);
