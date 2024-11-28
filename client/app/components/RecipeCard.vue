<script setup lang="ts">
import type { Nutrient } from "~~/server/utils/bff";
const { savedRecipes, saveRecipe } = useUserStore();

const props = defineProps<{
  id: number;
  image: string;
  title: string;
  readyInMinutes: number;
  nutrition: {
    nutrients: Nutrient[];
  };
}>();

function getNutrient(name: string): Nutrient | undefined {
  return props.nutrition.nutrients.find((nutrient) => nutrient.name == name);
}

const calories = computed(() => getNutrient("Calories"));
const protein = computed(() => getNutrient("Protein"));
const carbs = computed(() => getNutrient("Carbohydrates"));
const fat = computed(() => getNutrient("Fat"));

const [value, toggle] = useToggle(savedRecipes.includes(props.id));
async function handleSave() {
  toggle();
  saveRecipe();
}
</script>

<template>
  <UCard
    :ui="{
      header: 'p-0 sm:p-0 border-none',
    }"
    class="cursor-pointer transition-all duration-300 hover:scale-105"
  >
    <template #header>
      <div class="relative">
        <img :src="props.image" :alt="props.title" class="w-full rounded-lg" />
        <div class="absolute bottom-2 left-3">
          <span
            class="flex items-center rounded-2xl bg-white/70 px-2 py-1 text-sm shadow-xl backdrop-blur-md dark:bg-[#191919]/70"
          >
            <UIcon name="lucide:alarm-clock" class="mr-1" />
            {{ props.readyInMinutes }} min
            <DevOnly>{{ props.id }}</DevOnly>
          </span>
        </div>
      </div>
    </template>
    <div class="space-between flex">
      <span class="mr-4 line-clamp-2 min-h-[3rem] flex-3 font-semibold">{{
        props.title
      }}</span>
      <div class="flex flex-1 items-start justify-end">
        <UButton
          color="secondary"
          size="xl"
          variant="ghost"
          :icon="value ? 'lucide:bookmark-check' : 'lucide:bookmark'"
          class="p-0"
          v-on:click="handleSave"
        />
      </div>
    </div>
    <div class="mt-2 flex flex-wrap justify-between">
      <div class="flex flex-col items-center text-sm font-semibold">
        <span class="text-muted">Calories</span>
        <span class="text-orange-500 dark:text-orange-400">{{
          Math.round(calories?.amount!)
        }}</span>
      </div>
      <div class="flex flex-col items-center text-sm font-semibold">
        <span class="text-muted">Protein</span>
        <span class="text-purple-500 dark:text-purple-400">{{
          Math.round(protein?.amount!)
        }}</span>
      </div>
      <div class="flex flex-col items-center text-sm font-semibold">
        <span class="text-muted">Carbs</span>
        <span class="text-teal-500 dark:text-teal-400">{{
          Math.round(carbs?.amount!)
        }}</span>
      </div>
      <div class="flex flex-col items-center text-sm font-semibold">
        <span class="text-muted">Fat</span>
        <span class="text-green-500 dark:text-green-400">{{
          Math.round(fat?.amount!)
        }}</span>
      </div>
    </div>
  </UCard>
</template>
