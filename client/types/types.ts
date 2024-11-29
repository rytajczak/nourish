// Measure Interfaces
export interface Measure {
  amount: number;
  unitLong: string;
  unitShort: string;
}

export interface Measures {
  metric: Measure;
  us: Measure;
}

// Ingredient Interface
export interface Ingredient {
  aisle: string;
  amount: number;
  consistency: string;
  id: number;
  image: string;
  measures: Measures;
  meta: string[];
  name: string;
  nameClean: string;
  original: string;
  originalName: string;
  unit: string;
}

// Analyzed Instruction Interfaces
export interface AnalyzedInstructionStep {
  number: number;
  step: string;
  ingredients: Array<{
    id: number;
    image: string;
    localizedName: string;
    name: string;
  }>;
  equipment: Array<{
    id: number;
    image: string;
    localizedName: string;
    name: string;
  }>;
  length?: {
    number: number;
    unit: string;
  };
}

export interface AnalyzedInstruction {
  name: string;
  steps: AnalyzedInstructionStep[];
}

// Recipe Interfaces
export interface WinePairing {
  pairedWines: string[];
  pairingText: string;
  productMatches: any[];
}

export interface Recipe {
  aggregateLikes: number;
  analyzedInstructions: AnalyzedInstruction[];
  cheap: boolean;
  cookingMinutes: number;
  creditsText: string;
  cuisines: string[];
  dairyFree: boolean;
  diets: string[];
  dishTypes: string[];
  extendedIngredients: Ingredient[];
  gaps: string;
  glutenFree: boolean;
  healthScore: number;
  id: number;
  image: string;
  imageType: string;
  instructions: string;
  lowFodmap: boolean;
  occasions: string[];
  originalId: null | number;
  preparationMinutes: number;
  pricePerServing: number;
  readyInMinutes: number;
  servings: number;
  sourceName: string;
  sourceUrl: string;
  spoonacularScore: number;
  summary: string;
  sustainable: boolean;
  title: string;
  vegan: boolean;
  vegetarian: boolean;
  veryHealthy: boolean;
  veryPopular: boolean;
  weightWatcherSmartPoints: number;
  winePairing: WinePairing;
}

// Nutrition Interfaces
export interface Nutrient {
  name: string;
  amount: number;
  unit: string;
  percentOfDailyNeeds: number;
}

export interface NutritionSummary {
  nutrients: Nutrient[];
}

// Recipe Value Interface
export interface RecipeValue {
  readyInMinutes: number;
  sourceUrl: string;
  servings: number;
  id: number;
  title: string;
  imageType: string;
}

// Recipe Item Interface
export interface RecipeItem {
  id: number;
  slot: number;
  position: number;
  type: string;
  value: any;
}

// Day Interface
export interface Day {
  date: number;
  day: string;
  items: RecipeItem[];
  nutritionSummary: NutritionSummary;
  nutritionSummaryBreakfast: NutritionSummary;
  nutritionSummaryLunch: NutritionSummary;
  nutritionSummaryDinner: NutritionSummary;
}

// Weekly Plan Interface
export interface WeeklyPlan {
  days: Day[];
}
