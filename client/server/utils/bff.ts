import { H3Event } from "h3";
import { joinURL } from "ufo";

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

export function getApiUrl(event: H3Event) {
  const proxyUrl = useRuntimeConfig().public.apiUrl;
  return joinURL(proxyUrl, event.path.replace(/^\/api\//, ""));
}
