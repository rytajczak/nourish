<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
const { user } = useUserStore();
const dates = [
  "2024-11-15",
  "2024-11-16",
  "2024-11-17",
  "2024-11-18",
  "2024-11-19",
  "2024-11-20",
  "2024-11-21",
  "2024-11-22",
  "2024-11-23",
  "2024-11-24",
  "2024-11-25",
];
const selectedDate = ref(new Date().toISOString().split("T")[0]);
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Dashboard</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <UButton variant="subtle">Generate Plan</UButton>
    </div>
    <div class="flex">
      <div class="mr-4 flex-2">
        <UCard class="px-8">
          <template #header>
            <h2 class="text-xl font-semibold">Select a date</h2>
          </template>
          <!-- <UCarousel
            v-slot="{ item }"
            :items="dates"
            :ui="{
              item: 'basis-1/7',
            }"
            arrows
          >
            <div class="max-w-24 text-center">{{ item }}</div>
          </UCarousel> -->
        </UCard>
        <h2 class="mt-6 text-xl font-semibold">Meals to prepare</h2>
      </div>
      <div class="ml-4 flex-1">
        <h2 class="mb-4 text-xl font-semibold">Statistics</h2>
        <UCard>Today's nutrition</UCard>
      </div>
    </div>
  </div>
</template>
