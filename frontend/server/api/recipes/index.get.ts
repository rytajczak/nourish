type Macro = {
  amount: number;
  unit: string;
};

export type RecipePreview = {
  id: number;
  title: string;
  image: string;
  calories: Macro;
  protein: Macro;
  fat: Macro;
  carbs: Macro;
  readyInMinutes: number;
};

type Resp = {
  offset: number;
  number: number;
  recipePreviews: RecipePreview[];
  totalResults: number;
};

const findNutrient = (nutrients: any[], name: string): Macro => {
  const nutrient = nutrients.find((n) => n.name === name);
  return nutrient
    ? { amount: Math.round(nutrient.amount), unit: nutrient.unit }
    : { amount: 0, unit: "" };
};

export default defineEventHandler(async (event) => {
  const res = await searchRecipes(event, "chicken");

  const recipePreviews = res.results.map((recipe) => ({
    id: recipe.id,
    title: recipe.title,
    image: recipe.image,
    calories: findNutrient(recipe.nutrition.nutrients, "Calories"),
    protein: findNutrient(recipe.nutrition.nutrients, "Protein"),
    fat: findNutrient(recipe.nutrition.nutrients, "Fat"),
    carbs: findNutrient(recipe.nutrition.nutrients, "Carbohydrates"),
    readyInMinutes: recipe.readyInMinutes,
  }));

  const result: Resp = {
    offset: res.offset,
    number: res.number,
    recipePreviews,
    totalResults: res.totalResults,
  };

  console.log(result.recipePreviews[0]);

  return result;
});
