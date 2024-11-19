export const usePlannerStore = defineStore(
  "planner",
  () => {
    async function generateDailyPlan() {}

    async function generateWeeklyPlan() {}

    return {
      generateDailyPlan,
      generateWeeklyPlan,
    };
  },
  { persist: true },
);
