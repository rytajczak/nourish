const config = useRuntimeConfig();

export interface RecipePreview {
  id: string;
  title: string;
  image: string;
  calories: number;
  protein: number;
  carbs: number;
  fat: number;
  readyInMinutes: number;
}

export interface RecipeSearchResponse {
  number: number;
  offset: number;
  results: RecipePreview[];
  totalResults: number;
}

export default defineCachedEventHandler(async (event) => {
  const query = getQuery(event);
  const resp = await $fetch<RecipeSearchResponse>(
    `${config.public.apiUrl}/recipes/search`,
    {
      query: {
        query: query.query,
      },
    },
  );
  const recipePreviews = resp.results.map((recipe) => ({
    id: recipe.id,
    title: recipe.title,
    image: recipe.image,
    calories: recipe.calories,
    protein: recipe.protein,
    carbs: recipe.carbs,
    fat: recipe.fat,
    readyInMinutes: recipe.readyInMinutes,
  }));

  return {
    number: resp.number,
    offset: resp.offset,
    results: recipePreviews,
    totalResults: resp.totalResults,
  };
});
