<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const showing = ref("all");
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Dashboard</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <GenerateButton />
    </div>
    <div class="flex flex-col 2xl:flex-row">
      <div class="flex-3 2xl:mr-4">
        <WeekCarousel />
        <div class="mt-6 flex items-center justify-between">
          <h2 class="text-xl font-semibold">Meals to prepare</h2>
          <USelect
            v-model="showing"
            color="neutral"
            :items="['all', 'breakfast', 'lunch', 'dinner']"
            class="w-28"
          />
        </div>
      </div>
      <div class="flex-2 2xl:ml-4">
        <h2 class="my-4 text-xl font-semibold">Statistics</h2>
        <div class="grid grid-cols-3 gap-4">
          <UCard class="col-span-3">
            <div class="flex flex-col">
              <span class="text-sm text-[var(--ui-text-muted)]"
                >Total calories today</span
              >
              <span class="text-xl font-semibold">2000 kcal</span>
            </div>
          </UCard>
          <UCard>
            <div class="flex flex-col">
              <span class="text-sm text-[var(--ui-text-muted)]"
                >Protein (g)</span
              >
              <span class="text-xl font-semibold">180 g</span>
            </div>
          </UCard>
          <UCard>
            <div class="flex flex-col">
              <span class="text-sm text-[var(--ui-text-muted)]">Carbs (g)</span>
              <span class="text-xl font-semibold">180 g</span>
            </div>
          </UCard>
          <UCard>
            <div class="flex flex-col">
              <span class="text-sm text-gray-400">Fat (g)</span>
              <span class="text-xl font-semibold">180 g</span>
            </div>
          </UCard>
        </div>
      </div>
    </div>
  </div>
</template>
