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
    console.log(response);
    return response;
  } catch (error) {
    console.log(error);
    return null;
  }
});
