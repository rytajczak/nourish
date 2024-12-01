import { getApiUrl } from "~~/server/utils/bff";

export default defineEventHandler((event) => {
  console.log("proxied request:", event.method, getApiUrl(event));
});
