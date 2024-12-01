<script setup lang="ts">
import type { Nutrient } from "~~/types/types";

const planner = usePlannerStore();
const showing = ref("selected day");

function getNutrient(name: string): Nutrient | undefined {
  return planner.selectedDay?.nutritionSummary.nutrients.find(
    (nutrient) => nutrient.name == name,
  );
}

const calories = computed(() => getNutrient("Calories"));
const protein = computed(() => getNutrient("Protein"));
const carbs = computed(() => getNutrient("Carbohydrates"));
const fat = computed(() => getNutrient("Fat"));
</script>

<template>
  <div>
    <div class="flex items-center justify-between">
      <h2 class="my-4 text-xl font-semibold">Macros</h2>
      <USelect
        v-model="showing"
        color="neutral"
        :items="['selected day', 'this week']"
        class="w-36"
      />
    </div>
    <div class="grid grid-cols-3 gap-4">
      <UCard class="col-span-3">
        <div class="flex flex-col">
          <span class="text-muted text-sm"
            >Total calories
            {{
              showing === "selected day" ? "selected day" : "this week"
            }}</span
          >
          <span class="text-xl font-semibold"
            >{{ Math.round(calories?.amount ?? 0) }} {{ calories?.unit }}</span
          >
        </div>
      </UCard>
      <UCard>
        <div class="flex flex-col">
          <span class="text-muted text-sm">Protein (g)</span>
          <span class="text-xl font-semibold"
            >{{ Math.round(protein?.amount ?? 0) }} {{ protein?.unit }}</span
          >
        </div>
      </UCard>
      <UCard>
        <div class="flex flex-col">
          <span class="text-muted text-sm">Carbs (g)</span>
          <span class="text-xl font-semibold"
            >{{ Math.round(carbs?.amount ?? 0) }} {{ carbs?.unit }}</span
          >
        </div>
      </UCard>
      <UCard>
        <div class="flex flex-col">
          <span class="text-muted text-sm">Fat (g)</span>
          <span class="text-xl font-semibold"
            >{{ Math.round(fat?.amount ?? 0) }} {{ fat?.unit }}</span
          >
        </div>
      </UCard>
    </div>
  </div>
</template>
