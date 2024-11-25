import { WeeklyPlan } from "~~/server/utils/spoonacular";

export default defineEventHandler(async (event) => {
  try {
    const { secure } = await requireUserSession(event);
    const startDate = getRouterParam(event, "date");

    const response = await $fetch<WeeklyPlan>(
      `${spoonUrl}/mealplanner/${secure?.spoonName}/week/${startDate}`,
      {
        headers: spoonHeaders,
        query: { hash: secure?.spoonHash },
      },
    );
    return response;
  } catch (error) {
    return { days: [] };
  }
});
