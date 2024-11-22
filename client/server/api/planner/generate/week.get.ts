interface MealPlanItem {
  slot: number;
  position: number;
  mealPlanId: number;
  type: string;
  day: number;
  value: string;
}

interface Response {
  name: string;
  publishAsPublic: boolean;
  items: MealPlanItem[];
}

export default defineEventHandler(async (event) => {
  const response = await $fetch<Response>(
    `${spoonUrl}/recipes/mealplans/generate`,
    {
      headers: spoonHeaders,
      params: { timeFrame: "week" },
    },
  );
  return response;
});
