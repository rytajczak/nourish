export default defineNuxtRouteMiddleware(() => {
  const { loggedIn, clear } = useUserSession();
  const { logOut } = useUserStore();

  if (!loggedIn.value) {
    logOut();
    clear();
    return navigateTo("/");
  }
});
