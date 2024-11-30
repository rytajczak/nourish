import type { Day, Recipe, Entry, RecipeValue } from "~~/types/types";

export const usePlannerStore = defineStore("planner", () => {
  /**
   * Array of days containing the recipes for each day
   */
  const days = ref<Day[]>([]);

  /**
   * Acts as a cache for the recipes that are in the current user's plan
   */
  const recipeMap = ref<Map<number, Recipe>>();

  /**
   * The currently selected Date
   */
  const selectedDay = ref<Date>(new Date());

  /**
   * Info regarding the currently selected day
   */
  const selectedDayInfo = computed(() => {
    const weekday = selectedDay.value.toLocaleDateString("en-US", {
      weekday: "long",
    });
    const result = days.value.find((day) => {
      return day.day === weekday;
    });
    return result;
  });

  /**
   * The start date for the week
   */
  const weekStartDate = computed<Date>(() => {
    const startDate = new Date();
    const day = startDate.getDay();
    const diff = day === 0 ? 6 : day - 1;
    startDate.setDate(startDate.getDate() - diff);
    startDate.setHours(0, 0, 0, 0);
    return startDate;
  });

  /**
   * Fetch the meals of the current week starting from today
   */
  async function fetchWeek() {
    status.value = "pending";
    const startDate = dateToString(weekStartDate.value);
    const week = await $fetch<{ days: Day[] }>(
      `/api/mealplanner/me/week/${startDate}`,
    );
    if (!week) {
      status.value = "error";
      return;
    }

    days.value = week.days;

    const recipeIds = [
      ...new Set(
        week.days.flatMap((day) => day.items.map((item) => item.value.id)),
      ),
    ];

    const csv = recipeIds.join(",");
    const recipes = await $fetch("/api/recipes/info-bulk", {
      query: { ids: csv },
    });
    if (!recipes) {
      status.value = "success";
      return;
    }

    const mappedRecipes = new Map<number, any>(
      recipes.map((recipe: any) => [recipe.id, recipe]),
    );

    recipeMap.value = mappedRecipes;
    status.value = "success";
  }

  /**
   * Generate breakfast, lunch, and dinner for the current week
   */
  async function generateWeek() {}

  /**
   * Generate breakfast, lunch, and dinner for the currently selected date
   */
  async function generateDay(
    targetCalories: number,
    diet: string,
    exclude: string[],
  ) {
    const csv = exclude.join(",");
    await clearDay();
    const response = await $fetch<{ meals: RecipeValue[] }>(
      `/api/recipes/mealplans/generate`,
      {
        query: { timeFrame: "day", targetCalories, diet, exclude: csv },
      },
    );

    const date = dateToTimestamp(selectedDay.value);
    const body = response.meals.map((recipe, index) => {
      return {
        date,
        slot: index + 1,
        position: index,
        type: "RECIPE",
        value: recipe,
      } as Entry;
    });

    console.log(body);
  }

  async function addEntry() {}

  async function deleteEntry() {}

  async function clearDay() {
    const date = dateToString(selectedDay.value);
    await $fetch(`/api/mealplanner/me/day/${date}`, {
      method: "DELETE",
    });
  }

  const status = ref<"idle" | "pending" | "success" | "error">("idle");

  return {
    days,
    recipeMap,
    selectedDay,
    selectedDayInfo,

    // actions
    fetchWeek,
    generateDay,
    generateWeek,
    addEntry,
    deleteEntry,
    clearDay,

    // helpers
    weekStartDate,
    status,
  };
});
