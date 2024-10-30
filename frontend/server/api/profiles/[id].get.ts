export default eventHandler(async (event) => {
  const { user } = await getUserSession(event);
  const profiles = await useDB().query.auth.findFirst({});
  return profiles;
});
