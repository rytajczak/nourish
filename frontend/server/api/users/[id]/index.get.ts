export default defineEventHandler(async (event) => {
  const id = getRouterParam(event, "id");
  console.log(id);
  const user = await $fetch(`http://localhost:8080/users/${id}`);
  return user;
});
