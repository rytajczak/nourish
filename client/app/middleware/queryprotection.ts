export default defineNuxtRouteMiddleware((route) => {
  const query = route.query.query;
  if (query === undefined || query === "") {
    return navigateTo("/planner");
  }
});
