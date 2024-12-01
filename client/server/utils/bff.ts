import { H3Event } from "h3";
import { joinURL } from "ufo";

const host = useRuntimeConfig().spoonacularApiHost;
const key = useRuntimeConfig().spoonacularApiKey;

export interface UserResourceResponse {
  spoonCredential: {
    username: string;
    hash: string;
  };
  profile: {
    diet: string;
    calories: number;
    protein: number;
    carbs: number;
    fat: number;
  };
  intolerances: string[];
  savedRecipes: any[];
}

export const spoonHeaders = {
  "X-Rapidapi-Key": key,
  "X-Rapidapi-Host": host,
};

export function getApiUrl(event: H3Event) {
  const proxyUrl = useRuntimeConfig().public.apiUrl;
  return joinURL(proxyUrl, event.path.replace(/^\/api\//, ""));
}

export function getSpoonUrl(event: H3Event) {
  const proxyUrl = `https://${host}`;
  return joinURL(proxyUrl, event.path.replace(/^\/api\//, ""));
}
