<script setup lang="ts">
import type { Item, Nutrient } from "~~/types/types";
import { Chart, registerables } from "chart.js";
import { DoughnutChart, useDoughnutChart } from "vue-chart-3";

const props = defineProps<Item>();
Chart.register(...registerables);

const { recipeMap } = usePlannerStore();
const { profile, saveRecipe, removeSavedRecipe } = useUserStore();

const recipe = recipeMap?.get(Number(props.value.id));

function getNutrient(name: string): Nutrient | undefined {
  return recipe?.nutrition.nutrients.find((nutrient) => nutrient.name == name);
}

const calories = getNutrient("Calories");
const protein = getNutrient("Protein");
const carbs = getNutrient("Carbohydrates");
const fat = getNutrient("Fat");

const { doughnutChartProps: calorieProps } = useDoughnutChart(
  createDoughnutChartProps(
    "Calories",
    [calories?.amount!, profile.calories],
    "#ff8904",
  ),
);

const { doughnutChartProps: proteinProps } = useDoughnutChart(
  createDoughnutChartProps(
    "Protein",
    [protein?.amount!, profile.protein],
    "#be77f9",
  ),
);

const { doughnutChartProps: carbProps } = useDoughnutChart(
  createDoughnutChartProps(
    "Carbohydrates",
    [carbs?.amount!, profile.carbs],
    "#00d5be",
  ),
);

const { doughnutChartProps: fatProps } = useDoughnutChart(
  createDoughnutChartProps("Fat", [fat?.amount!, profile.fat], "#05df72"),
);

const priceThermometerOfHell = computed(() => {
  const corrected = recipe?.pricePerServing / 100;
  if (corrected <= 3) {
    return "LOW";
  } else if (corrected <= 5) {
    return "MEDIUM";
  } else {
    return "HIGH";
  }
});
</script>

<template>
  <UCard
    as="li"
    class="mb-4"
    :ui="{ header: 'pl-0 sm:pl-2 sm:pr-6 sm:m-0', body: 'p-0 sm:p-0' }"
  >
    <template #header>
      <div class="flex">
        <div class="flex flex-2">
          <div class="flex items-center">
            <UIcon
              size="28"
              name="material-symbols:drag-indicator"
              class="entry-handle text-muted mr-2"
            />
          </div>
          <div class="relative">
            <img
              :src="recipe?.image"
              :alt="recipe?.title"
              class="h-32 w-48 rounded-lg"
            />
          </div>
          <div class="ml-4 flex flex-col">
            <h2 class="line-clamp-2 font-semibold">
              {{ recipe?.title }}
            </h2>
            <div class="mt-2 flex">
              <UBadge v-if="recipe?.vegetarian" variant="soft" class="mr-2">
                <UIcon name="lucide:carrot" class="mr-1" />
                <span>Vegetarian</span>
              </UBadge>
              <UBadge v-if="recipe?.vegan" variant="soft" class="mr-2">
                <UIcon name="lucide:carrot" class="mr-1" />
                <span>Vegan</span>
              </UBadge>
              <UBadge v-if="recipe?.glutenFree" variant="soft">
                <UIcon name="lucide:wheat" class="mr-1" />
                <span>GF</span>
              </UBadge>
            </div>
            <div class="mt-4 flex">
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Calories</span>
                <DoughnutChart v-bind="calorieProps" class="mt-1 w-12" />
              </div>
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Protein</span>
                <DoughnutChart v-bind="proteinProps" class="mt-1 w-12" />
              </div>
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Carbs</span>
                <DoughnutChart v-bind="carbProps" class="mt-1 w-12" />
              </div>
              <div class="flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Fat</span>
                <DoughnutChart v-bind="fatProps" class="mt-1 w-12" />
              </div>
            </div>
          </div>
        </div>
        <div class="flex flex-1 flex-col justify-between">
          <div class="flex justify-end">
            <UButton variant="ghost" color="neutral" icon="lucide:info" />
            <UButton variant="ghost" color="secondary" icon="lucide:bookmark" />
            <UButton
              variant="ghost"
              color="error"
              icon="lucide:trash-2"
              @click="$emit('delete', id)"
            />
          </div>
          <div class="flex justify-end">
            <span class="flex items-center text-xl">
              <UIcon name="lucide:dollar-sign" />
              <span>
                {{ (recipe?.pricePerServing / 100).toFixed(2) }} per serving
              </span>
            </span>
          </div>
        </div>
      </div>
    </template>
    <UCollapsible class="flex flex-col gap-2">
      <UButton
        label="See Details"
        color="neutral"
        variant="link"
        trailing-icon="lucide:chevron-down"
        block
      />

      <template #content>
        <div class="px-6 pb-6">
          <h2 class="text-lg font-semibold">Ingredients</h2>
          <ul>
            <li v-for="ingredient in recipe?.extendedIngredients">
              <span class="mr-1">{{ ingredient.amount }}</span>
              <span class="mr-1">{{ ingredient.unit }}</span>
              <span class="mr-1">{{ ingredient.name }}</span>
            </li>
          </ul>
        </div>
      </template>
    </UCollapsible>
  </UCard>
</template>
