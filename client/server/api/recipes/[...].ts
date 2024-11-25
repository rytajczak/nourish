export default defineEventHandler(async (event) => {
  return proxyRequest(event, getApiUrl(event));
});
