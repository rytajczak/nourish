<script setup lang="ts">
const {
  data: days,
  status,
  execute,
} = useFetch("/api/mealplanner/generate", {
  immediate: false,
});
</script>

<template>
  <div class="mx-5 mt-5">
    <div class="flex items-center">
      <h1 class="text-3xl">Meal Planner</h1>
    </div>
    <UButton @click="execute()" color="secondary" variant="subtle" class="mt-4"
      >Generate Plan</UButton
    >
    <div>status: {{ status }}</div>
  </div>
</template>
