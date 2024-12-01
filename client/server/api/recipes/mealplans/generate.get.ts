import { getSpoonUrl, spoonHeaders } from "~~/server/utils/bff";

export default defineEventHandler(async (event) => {
  const url = getSpoonUrl(event);
  return proxyRequest(event, url, {
    fetchOptions: { headers: spoonHeaders },
  });
});
