import { getApiUrl } from "~~/server/utils/bff";

export default defineEventHandler(async (event) => {
  return proxyRequest(event, getApiUrl(event));
});
