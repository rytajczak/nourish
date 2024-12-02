<script setup lang="ts">
import type { Item, Nutrient, AnalyzedInstruction } from "~~/types/types";
import { Chart, registerables } from "chart.js";
import { DoughnutChart, useDoughnutChart } from "vue-chart-3";

const props = defineProps<Item>();
Chart.register(...registerables);

const { recipeMap } = usePlannerStore();
const { profile, saveRecipe, removeSavedRecipe } = useUserStore();
const img = useImage();

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
    [Math.round(calories?.amount!), profile.calories],
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
  const corrected = recipe?.pricePerServing! / 100;
  if (corrected <= 3) {
    return "text-green-500";
  } else if (corrected <= 5) {
    return "text-orange-500";
  } else {
    return "text-red-500";
  }
});
function getUniqueEquipmentNames(
  instructions: AnalyzedInstruction[],
): string[] {
  const equipmentSet = new Set<string>();

  for (const instruction of instructions) {
    for (const step of instruction.steps) {
      for (const equipment of step.equipment) {
        equipmentSet.add(equipment.name);
      }
    }
  }

  return Array.from(equipmentSet);
}
const equipment = getUniqueEquipmentNames(recipe?.analyzedInstructions ?? []);
</script>

<template>
  <UCard
    as="li"
    class="mb-6"
    :ui="{ header: 'pl-0 sm:pl-2 sm:pr-6 sm:m-0', body: 'p-0 sm:p-0' }"
  >
    <template #header>
      <div class="flex">
        <div class="flex items-center">
          <UIcon
            size="28"
            name="material-symbols:drag-indicator"
            class="entry-handle text-muted mr-2"
          />
        </div>
        <img :src="recipe?.image" class="h-32 w-40 rounded-lg" />
        <div class="ml-4 flex w-full flex-col">
          <div class="flex justify-between">
            <h2 class="line-clamp-2 font-semibold">
              {{ recipe?.title }}
            </h2>
            <div>
              <div class="flex">
                <UButton variant="ghost" color="neutral" icon="lucide:info" />
                <UButton
                  variant="ghost"
                  color="secondary"
                  icon="lucide:bookmark"
                />
                <UButton
                  variant="ghost"
                  color="error"
                  icon="lucide:trash-2"
                  @click="$emit('delete', id)"
                />
              </div>
            </div>
          </div>
          <div class="mt-2 flex items-end justify-between">
            <div class="flex">
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Calories</span>
                <div class="relative flex items-center justify-center">
                  <DoughnutChart v-bind="calorieProps" class="mt-1 w-16" />
                  <span class="absolute text-sm font-semibold">
                    {{ calories?.amount.toFixed(0) }}
                  </span>
                </div>
              </div>
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Protein</span>
                <div class="relative flex items-center justify-center">
                  <DoughnutChart v-bind="proteinProps" class="mt-1 w-12" />
                  <span class="absolute text-sm font-semibold">
                    {{ protein?.amount.toFixed(0) }}
                  </span>
                </div>
              </div>
              <div class="mr-4 flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Carbs</span>
                <div class="relative flex items-center justify-center">
                  <DoughnutChart v-bind="carbProps" class="mt-1 w-12" />
                  <span class="absolute text-sm font-semibold">
                    {{ carbs?.amount.toFixed(0) }}
                  </span>
                </div>
              </div>
              <div class="flex flex-col items-center">
                <span class="text-muted text-sm font-semibold">Fat</span>
                <div class="relative flex items-center justify-center">
                  <DoughnutChart v-bind="fatProps" class="mt-1 w-12" />
                  <span class="absolute text-sm font-semibold">
                    {{ fat?.amount.toFixed(0) }}
                  </span>
                </div>
              </div>
            </div>
            <div class="flex flex-col items-end">
              <span class="text-muted">{{ recipe?.servings }} servings</span>
              <span class="flex items-center text-lg font-semibold">
                <UIcon name="lucide:dollar-sign" />
                <span>
                  {{ (recipe?.pricePerServing! / 100).toFixed(2) }} per serving
                </span>
              </span>
            </div>
          </div>
        </div>
      </div>
    </template>
    <UCollapsible class="flex flex-col gap-2">
      <UButton
        label="Details"
        color="neutral"
        variant="link"
        trailing-icon="lucide:chevron-down"
        block
      />
      <template #content>
        <div class="flex justify-between px-6 pt-2 pb-6">
          <UCard class="mr-6 flex flex-1">
            <h2 class="text-lg font-semibold">Ingredients</h2>
            <ul>
              <li v-for="ingredient in recipe?.extendedIngredients">
                <span class="mr-1 flex flex-row items-center">
                  <span class="text-muted">
                    {{ ingredient.amount }} {{ ingredient.name }}
                    {{ ingredient.unit }}
                  </span>
                </span>
              </li>
            </ul>
          </UCard>
          <UCard class="flex-1">
            <h2 class="text-lg font-semibold">Equipment</h2>
            <ul>
              <li v-for="tool in equipment" class="text-muted">{{ tool }}</li>
            </ul>
          </UCard>
          <UCard class="ml-6 flex-1">
            <h2 class="text-lg font-semibold">Tags</h2>
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
            <h2 class="mt-4 text-lg font-semibold">Diets</h2>
            <ul>
              <li v-for="diet in recipe?.diets" class="text-muted">
                {{ diet }}
              </li>
            </ul>
          </UCard>
        </div>
      </template>
    </UCollapsible>
  </UCard>
</template>
