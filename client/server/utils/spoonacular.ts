const host = useRuntimeConfig().spoonacularApiHost;
const key = useRuntimeConfig().spoonacularApiKey;

export const spoonUrl = `https://${host}`;

export const spoonHeaders = {
  "x-rapidapi-key": key,
  "x-rapidapi-host": host,
};
