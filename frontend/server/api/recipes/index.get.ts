const config = useRuntimeConfig();

export default defineCachedEventHandler(async (event) => {
  const query = getQuery(event);
  const resp = await $fetch(`${config.public.API_URL}/recipes/search`, {
    query: {
      query: query.query,
    },
  });

  return resp;
});
