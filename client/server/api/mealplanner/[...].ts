import { getSpoonUrl } from "~~/server/utils/bff";

export default defineEventHandler(async (event) => {
  const { secure } = await requireUserSession(event);
  const adjustedUrl = getSpoonUrl(event).replace(
    `/me/`,
    `/${secure?.spoonName}/`,
  );

  const url = new URL(adjustedUrl);
  url.searchParams.set("hash", secure?.spoonHash ?? "");

  return proxyRequest(event, url.toString(), {
    fetchOptions: { headers: spoonHeaders },
  });
});
