interface Recipe {
  readyInMinutes: number;
  sourceUrl: string;
  servings: number;
  id: number;
  title: string;
  imageType: string;
}

interface AddRecipeBody {
  date: number;
  slot: 1 | 2 | 3;
  position: number;
  type: "RECIPE";
  value: Recipe;
}

interface GenerateResponse {
  meals: Recipe[];
  nutrients: {
    fat: number;
    carbohydrates: number;
    calories: number;
    protein: number;
  };
}

interface AddResponse {
  status: string;
}

export default defineEventHandler(async (event) => {
  const { secure } = await requireUserSession(event);
  const query = getQuery(event);

  const GenerateResponse = await $fetch<GenerateResponse>(
    `${spoonUrl}/recipes/mealplans/generate`,
    {
      headers: spoonHeaders,
      params: { timeFrame: "day" },
    },
  );

  const body: AddRecipeBody[] = [];
  GenerateResponse.meals.forEach((recipe, index) => {
    body.push({
      date: parseInt(query.date as string),
      slot: (index + 1) as 1 | 2 | 3,
      position: 0,
      type: "RECIPE",
      value: recipe,
    });
  });

  const addResponse = await $fetch<AddResponse>(
    `${spoonUrl}/mealplanner/${secure?.spoonName}/items`,
    {
      method: "POST",
      headers: spoonHeaders,
      query: { hash: secure?.spoonHash },
      body: body,
    },
  );
  return addResponse;
});
