import type { Day, Item, Recipe, RecipeValue } from "~~/types/types";

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
  const selectedDate = ref<Date>(new Date());

  /**
   * Info regarding the currently selected day
   */
  const selectedDay = computed(() => {
    const weekday = selectedDate.value.toLocaleDateString("en-US", {
      weekday: "long",
    });
    return days.value.find((day) => day.day === weekday);
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

    // Check for any recipes
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
    console.log(`diet: ${diet}`);
    status.value = "pending";
    // Generate recipes
    const csv = exclude.length > 0 ? exclude.join(",") : "";
    const day = await $fetch<{ meals: RecipeValue[] }>(
      `/api/recipes/mealplans/generate`,
      {
        query: { timeFrame: "day", targetCalories, diet, exclude: csv },
      },
    );

    // Format generated recipes to entries
    const date = dateToTimestamp(selectedDate.value);
    const body = day.meals.map((recipe, index) => {
      return {
        date,
        slot: 1,
        position: index + 1,
        type: "RECIPE",
        value: recipe,
      };
    });

    // Update planner
    await clearDay();
    await $fetch(`/api/mealplanner/me/items`, {
      method: "POST",
      body,
    });

    await fetchWeek();
  }

  async function addItem(item: any) {
    status.value = "pending";
    const body = {
      date: dateToTimestamp(selectedDate.value),
      slot: 1,
      position: item.position,
      type: "RECIPE",
      value: item.value,
    };
    await $fetch(`/api/mealplanner/me/items`, {
      method: "POST",
      body,
    });
    await fetchWeek();
  }

  async function deleteItem(id: number) {
    status.value = "pending";
    if (selectedDay.value && selectedDay.value.items) {
      selectedDay.value.items = selectedDay.value.items.filter(
        (item) => item.id !== id,
      );
    }
    await $fetch(`/api/mealplanner/me/items/${id}`, {
      method: "DELETE",
    });
    await fetchWeek();
  }

  async function clearDay() {
    status.value = "pending";
    const date = dateToString(selectedDate.value);
    await $fetch(`/api/mealplanner/me/day/${date}`, {
      method: "DELETE",
    });
    await fetchWeek();
  }

  const status = ref<"idle" | "pending" | "success" | "error">("idle");

  return {
    days,
    recipeMap,
    selectedDate,
    selectedDay,

    // actions
    fetchWeek,
    generateDay,
    generateWeek,
    addItem,
    deleteItem,
    clearDay,

    // helpers
    weekStartDate,
    status,
  };
});
