export interface Recipe {
  aggregateLikes: number;
  analyzedInstructions: {
    name: string;
    steps: {
      equipment: {
        id: number;
        image: string;
        localizedName: string;
        name: string;
      }[];
      ingredients: {
        id: number;
        image: string;
        localizedName: string;
        name: string;
      }[];
      length?: {
        number: number;
        unit: string;
      };
      number: number;
      step: string;
    }[];
  }[];
  cheap: boolean;
  cookingMinutes: number | null;
  creditsText: string;
  cuisines: string[];
  dairyFree: boolean;
  diets: string[];
  dishTypes: string[];
  extendedIngredients: {
    aisle: string;
    amount: number;
    consistency: string;
    id: number;
    image: string;
    measures: {
      metric: {
        amount: number;
        unitLong: string;
        unitShort: string;
      };
      us: {
        amount: number;
        unitLong: string;
        unitShort: string;
      };
    };
    meta: string[];
    name: string;
    nameClean: string;
    original: string;
    originalName: string;
    unit: string;
  }[];
  gaps: string;
  glutenFree: boolean;
  healthScore: number;
  id: number;
  image: string;
  imageType: string;
  instructions: string;
  lowFodmap: boolean;
  occasions: string[];
  originalId: number | null;
  preparationMinutes: number | null;
  pricePerServing: number;
  readyInMinutes: number;
  servings: number;
  sourceName: string;
  sourceUrl: string;
  spoonacularScore: number;
  spoonacularSourceUrl: string;
  summary: string;
  sustainable: boolean;
  title: string;
  vegan: boolean;
  vegetarian: boolean;
  veryHealthy: boolean;
  veryPopular: boolean;
  weightWatcherSmartPoints: number;
}
