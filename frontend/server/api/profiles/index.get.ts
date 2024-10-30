export default eventHandler(async (event) => {
  const profiles = await useDB().query.auth.findFirst({
    columns: {
      email: true,
    },
  });
  return profiles;
});
