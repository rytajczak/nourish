export const useUserStore = defineStore(
  "user",
  () => {
    const { user, clear } = useUserSession();

    const logOut = async () => {
      await clear();
      return navigateTo("/");
    };

    return { user, logOut };
  },
  { persist: true },
);
