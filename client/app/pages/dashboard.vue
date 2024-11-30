<script setup lang="ts">
import type { Item } from "~~/types/types";
import { dropOrSwap } from "@formkit/drag-and-drop";
import { useDragAndDrop } from "@formkit/drag-and-drop/vue";

definePageMeta({
  middleware: "auth",
});

const planner = usePlannerStore();
const onboarding = useOnboardingStore();
const showing = ref("all");

const [parent, items] = useDragAndDrop<Item>([], {
  onDragend: (event) => {
    console.log(event.values);
  },
  plugins: [dropOrSwap()],
});

async function handleDeleteItem(id: number) {
  items.value = items.value.filter((item) => item.id !== id);
  await planner.deleteItem(id);
}

onMounted(async () => {
  await planner.fetchWeek();
  items.value = planner.selectedDay?.items ?? [];
});
watch(
  () => planner.selectedDate,
  () => (items.value = planner.selectedDay?.items ?? []),
);
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Dashboard</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <GenerateButton />
    </div>
    <div class="grid grid-cols-1 xl:grid-cols-5 xl:gap-8">
      <div class="col-span-3">
        <DaySelector />
        <div class="my-6 flex items-center justify-between">
          <h2 class="text-xl font-semibold">Meals to prepare</h2>
          <USelect
            v-model="showing"
            color="neutral"
            :items="['all', 'breakfast', 'lunch', 'dinner']"
            class="w-32"
          />
        </div>
        <ul ref="parent">
          <PrepCard
            v-for="item in items"
            :key="item.id"
            v-bind="item"
            @delete="handleDeleteItem"
          />
        </ul>
      </div>
      <div class="col-span-1 xl:col-span-2">
        <NutritionInfo />
        <pre>{{ items }}</pre>
      </div>
    </div>
    <UModal v-model:open="onboarding.open" prevent-close>
      <template #content>
        <UCard :ui="{ header: 'p-0 sm:p-0 border-none' }">
          <UProgress
            color="neutral"
            v-model="onboarding.progress"
            class="mb-6"
          />
          <OnboardingWelcome v-if="onboarding.step == 0" />
          <OnboardingGoals v-if="onboarding.step == 1" />
          <OnboardingDiet v-if="onboarding.step == 2" />
          <OnboardingIntolerance v-if="onboarding.step == 3" />
        </UCard>
      </template>
    </UModal>
  </div>
</template>
