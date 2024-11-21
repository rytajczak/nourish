const host = useRuntimeConfig().spoonacularApiHost;
const key = useRuntimeConfig().spoonacularApiKey;

export const spoonUrl = `https://${host}`;

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
