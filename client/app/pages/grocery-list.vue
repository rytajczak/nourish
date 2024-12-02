<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { weekStartDate } = usePlannerStore();

const startDate = dateToString(weekStartDate);
const endDate = dateToString(
  new Date(weekStartDate.getTime() + 6 * 24 * 60 * 60 * 1000),
);

const response = await $fetch(
  `/api/mealplanner/me/shopping-list/${startDate}/${endDate}`,
  { method: "POST", lazy: true },
);
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Grocery List</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span class="text-blue-400">{{ startDate }} - {{ endDate }}</span>
    </div>
    <div>
      <UCard v-for="aisle in response.aisles" class="mb-6">
        <template #header>
          <h2 class="flex items-center">
            <span class="text-lg font-bold">{{ aisle.aisle }} Aisle</span>
            <USeparator orientation="vertical" class="h-6 px-4" />
            <span class="text-muted">{{ aisle.items.length }} item(s)</span>
          </h2>
        </template>
        <ul>
          <li v-for="item in aisle.items" class="mb-6">
            <span class="me-2">{{ item.name }}</span>
            <span class="text-muted">${{ (item.cost / 100).toFixed(2) }}</span>
          </li>
        </ul>
      </UCard>
    </div>
  </div>
</template>
