const config = useRuntimeConfig();

const desiredNutrients = ["Calories", "Protein", "Carbohydrates", "Fat"];

interface RecipeResponse {
  results: Recipe[];
  offset: number;
  number: number;
  totalResults: number;
}

export interface Recipe {
  vegetarian: boolean;
  vegan: boolean;
  glutenFree: boolean;
  dairyFree: boolean;
  veryHealthy: boolean;
  cheap: boolean;
  veryPopular: boolean;
  sustainable: boolean;
  lowFodmap: boolean;
  weightWatcherSmartPoints: number;
  gaps: string;
  preparationMinutes: number | null;
  cookingMinutes: number | null;
  aggregateLikes: number;
  healthScore: number;
  creditsText: string;
  license: string;
  sourceName: string;
  pricePerServing: number;
  id: number;
  title: string;
  readyInMinutes: number;
  servings: number;
  sourceUrl: string;
  image: string;
  imageType: string;
  nutrition: Nutrition;
  summary: string;
  cuisines: string[];
  dishTypes: string[];
  diets: string[];
  occasions: string[];
  analyzedInstructions: AnalyzedInstruction[];
  spoonacularScore: number;
  spoonacularSourceUrl: string;
}

interface Nutrition {
  nutrients: Nutrient[];
  properties: Property[];
  flavonoids: Flavonoid[];
  ingredients: Ingredient[];
  caloricBreakdown: CaloricBreakdown;
  weightPerServing: WeightPerServing;
}

interface Nutrient {
  name: string;
  amount: number;
  unit: string;
  percentOfDailyNeeds: number;
}

interface Property {
  name: string;
  amount: number;
  unit: string;
}

interface Flavonoid {
  name: string;
  amount: number;
  unit: string;
}

interface Ingredient {
  id: number;
  name: string;
  amount: number;
  unit: string;
  nutrients: Nutrient[];
}

interface CaloricBreakdown {
  percentProtein: number;
  percentFat: number;
  percentCarbs: number;
}

interface WeightPerServing {
  amount: number;
  unit: string;
}

interface AnalyzedInstruction {
  name: string;
  steps: Step[];
}

interface Step {
  number: number;
  step: string;
  ingredients: IngredientReference[];
  equipment: Equipment[];
}

interface IngredientReference {
  id: number;
  name: string;
  localizedName: string;
  image: string;
}

interface Equipment {
  id: number;
  name: string;
  localizedName: string;
  image: string;
}

export default defineCachedEventHandler(async (event) => {
  const query = getQuery(event);
  const resp = await $fetch<RecipeResponse>(
    `${config.public.apiUrl}/recipes/search`,
    {
      query: {
        query: query.query,
      },
    },
  );
  return resp;
});
