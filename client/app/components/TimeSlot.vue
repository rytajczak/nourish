<script setup lang="ts">
const planner = usePlannerStore();
const props = defineProps<{
  time: "breakfast" | "lunch" | "dinner";
}>();
const mealtime = computed(() => {
  switch (props.time) {
    case "breakfast":
      return planner.breakfast;
    case "lunch":
      return planner.lunch;
    case "dinner":
      return planner.dinner;
  }
});
</script>

<template>
  <div class="mt-4">
    <div
      v-if="
        planner.showingMeals === 'all' || planner.showingMeals == props.time
      "
    >
      <h2 class="text-lg font-semibold capitalize">{{ props.time }}</h2>
      <UCard class="mt-2">
        <div v-if="mealtime?.length! > 0">
          <div v-for="meal in mealtime" :key="meal.id">
            {{ meal }}
          </div>
        </div>
        <span v-else class="text-muted flex items-center justify-center">
          No meals for this time
        </span>
      </UCard>
    </div>
  </div>
</template>
