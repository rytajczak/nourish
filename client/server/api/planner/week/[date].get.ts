import { WeeklyPlan } from "~~/server/utils/spoonacular";

export default defineEventHandler(async (event) => {
  const { secure } = await requireUserSession(event);
  const date = getRouterParam(event, "date");

  try {
    const response = await $fetch<WeeklyPlan>(
      `${spoonUrl}/mealplanner/${secure?.spoonName}/week/${date}`,
      {
        headers: spoonHeaders,
        query: { hash: secure?.spoonHash },
      },
    );
    return response;
  } catch (error) {
    return null;
  }
});
