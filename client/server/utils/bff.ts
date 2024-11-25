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
  "x-rapidapi-key": key,
  "x-rapidapi-host": host,
};

export interface Nutrient {
  name: string;
  amount: number;
  unit: string;
  percentOfDailyNeeds: number;
}

export interface NutritionSummary {
  nutrients: Nutrient[];
}

export interface RecipeValue {
  readyInMinutes: number;
  sourceUrl: string;
  servings: number;
  id: number;
  title: string;
  imageType: string;
}

export interface RecipeItem {
  id: number;
  slot: number;
  position: number;
  type: string;
  value: RecipeValue;
}

export interface Day {
  date: number;
  day: string;
  items: RecipeItem[];
  nutritionSummary: NutritionSummary;
  nutritionSummaryBreakfast: NutritionSummary;
  nutritionSummaryLunch: NutritionSummary;
  nutritionSummaryDinner: NutritionSummary;
}

export interface WeeklyPlan {
  days: Day[];
}

export function getApiUrl(event: H3Event) {
  const proxyUrl = useRuntimeConfig().public.apiUrl;
  return joinURL(proxyUrl, event.path.replace(/^\/api\//, ""));
}

export function getSpoonUrl(event: H3Event) {
  const proxyUrl = `https://${host}`;
  return joinURL(proxyUrl, event.path.replace(/^\/api\//, ""));
}
