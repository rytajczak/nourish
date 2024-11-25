import type { Day } from "~~/server/utils/spoonacular";

interface PlannerEntry {
  date: number;
  slot: number;
  position: number;
  type: "RECIPE" | "MENU_ITEM" | "PRODUCT" | "INGREDIENTS";
  value: any;
}

export const usePlannerStore = defineStore("planner", () => {
  /**
   * the good stuff
   */
  const days = ref<Day[]>([]);

  /**
   * The current day selected
   */
  const selectedDay = ref<Date>(new Date());
  const showingMeals = ref("all");
  const breakfast = computed(() =>
    selectedDayInfo.value?.items.filter((item) => item.slot == 1),
  );
  const lunch = computed(() =>
    selectedDayInfo.value?.items.filter((item) => item.slot == 2),
  );
  const dinner = computed(() =>
    selectedDayInfo.value?.items.filter((item) => item.slot == 3),
  );

  /**
   * All the information for
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
   * The start date of the week
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
   * @param startDate The starting date for the week
   */
  async function fetchWeek() {
    const response = await $fetch(
      `/api/planner/me/week/${dateToString(weekStartDate.value)}`,
    );
    if (!response) return;

    days.value = response.days;
  }
  /**
   * Generate breakfast, lunch, and dinner for the current week
   */
  async function generateWeek() {}

  /**
   * Generate breakfast, lunch, and dinner for a date
   */
  async function generateDay() {}

  async function addEntry() {}

  return {
    days,
    selectedDay,
    selectedDayInfo,
    showingMeals,
    breakfast,
    lunch,
    dinner,
    weekStartDate,
    fetchWeek,
    generateDay,
    generateWeek,
    addEntry,
  };
});
