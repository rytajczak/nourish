<script setup lang="ts">
import type { Item, Recipe } from "~~/types/types";
import { animations } from "@formkit/drag-and-drop";
import { useDragAndDrop } from "@formkit/drag-and-drop/vue";
import DaySelector from "~/components/DaySelector.vue";

definePageMeta({
  middleware: "auth",
});

const planner = usePlannerStore();
const onboarding = useOnboardingStore();
const showing = ref("all");

const [parent, items] = useDragAndDrop<Item>([], {
  group: "items",
  onSort(data) {
    //@ts-ignore
    planner.selectedDay.items = data.values;
  },
  plugins: [animations()],
});

async function handleAddItem(id: number) {
  const response = await $fetch<Recipe>(`/api/recipes/${id}/info`);
  planner.recipeMap?.set(id, response);
  const highestPosition = items.value.reduce(
    (max, obj) => (obj.position > max ? obj.position : max),
    0,
  );
  const item = {
    slot: 1,
    position: highestPosition,
    type: "RECIPE",
    value: {
      servings: 1,
      id,
      title: response.title,
      imageType: response.imageType,
    },
  } as Item;
  items.value.push(item);
  await planner.addItem(item);
}

async function handleDeleteItem(id: number) {
  items.value = items.value.filter((item) => item.id !== id);
  await planner.deleteItem(id);
}

async function handleClearDay() {
  items.value = [];
  await planner.clearDay();
}

onMounted(async () => {
  try {
    await planner.fetchWeek();
    items.value = planner.selectedDay?.items ?? [];
  } catch (error) {
    console.warn("user probably not onboarded");
    planner.status = "idle";
  }
});
watch(
  () => planner.selectedDate,
  () => (items.value = planner.selectedDay?.items ?? []),
);

async function handleRefetch() {
  items.value = planner.selectedDay?.items ?? [];
}
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Dashboard</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <GenerateButton @refetch="handleRefetch" />
    </div>
    <div class="grid grid-cols-1 xl:grid-cols-5 xl:gap-8">
      <div class="col-span-3 mb-8">
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
        <UCard
          v-if="planner.status === 'success' && items.length == 0"
          class="mb-6"
        >
          <span class="text-muted flex justify-center">no items for today</span>
        </UCard>
        <div
          v-if="planner.status === 'pending'"
          class="mb-6 flex justify-center"
        >
          <UIcon name="svg-spinners:ring-resize" size="40" />
        </div>
        <div class="flex">
          <AddButton @add="handleAddItem" />
          <UButton
            variant="subtle"
            color="error"
            class="ml-2"
            @click="handleClearDay"
          >
            Clear Day
          </UButton>
        </div>
      </div>
      <div class="col-span-1 xl:col-span-2">
        <NutritionInfo />
      </div>
    </div>
    <UModal v-model:open="onboarding.open" prevent-close>
      <template #content>
        <UCard :ui="{ header: 'p-0 sm:p-0 border-none' }">
          <template #header>
            <UProgress
              color="neutral"
              size="lg"
              v-model="onboarding.progress"
              :ui="{ base: 'rounded-none rounded-t-xl' }"
            />
          </template>
          <OnboardingWelcome v-if="onboarding.step == 0" />
          <OnboardingGoals v-if="onboarding.step == 1" />
          <OnboardingDiet v-if="onboarding.step == 2" />
          <OnboardingIntolerance v-if="onboarding.step == 3" />
        </UCard>
      </template>
    </UModal>
  </div>
</template>
