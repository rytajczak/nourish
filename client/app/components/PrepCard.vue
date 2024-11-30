<script setup lang="ts">
const props = defineProps<{ id: number }>();
const { recipeMap } = usePlannerStore();
const recipeInfo = recipeMap?.get(props.id);

function getNutrient(name: string) {
  return recipeInfo?.nutrition?.nutrients.find(
    (nutrient) => nutrient.name == name,
  );
}

const calories = getNutrient("Calories");
const protein = getNutrient("Protein");
const carbs = getNutrient("Carbohydrates");
const fat = getNutrient("Fat");
const hours = Math.floor(recipeInfo?.readyInMinutes! / 60);
const minutes = recipeInfo?.readyInMinutes! % 60;
</script>

<template>
  <UCard class="mb-4" :ui="{ header: 'pl-0 sm:pl-2' }">
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
          <img :src="recipeInfo?.image" alt="" class="h-28 rounded-xl" />
          <div class="ml-4 flex flex-col">
            <h2 class="line-clamp-2 font-semibold">
              {{ recipeInfo?.title }}
            </h2>
            <p class="text-muted flex items-center">
              <UIcon name="lucide:clock" class="me-1" />
              <span>{{ hours }} hours and {{ minutes }} minutes</span>
            </p>
          </div>
        </div>
        <div class="flex flex-1 items-center justify-center"></div>
      </div>
      <!-- <pre>${{ (recipeInfo?.pricePerServing / 100).toFixed(2) }} per serving</pre> -->
    </template>
  </UCard>
</template>
