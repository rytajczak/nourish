<script setup lang="ts">
import type { Item } from "~~/types/types";

const props = defineProps<Item>();

const { recipeMap } = usePlannerStore();

const recipe = recipeMap?.get(Number(props.value.id));
const hours = Math.floor(recipe?.readyInMinutes! / 60);
const minutes = recipe?.readyInMinutes! % 60;
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
              class="h-32 rounded-lg"
            />
            <div class="absolute bottom-2 left-3">
              <span
                class="flex items-center rounded-2xl bg-white/70 px-2 py-1 text-xs shadow-xl backdrop-blur-md dark:bg-[#191919]/70"
              >
                <UIcon name="lucide:alarm-clock" class="mr-1" />
                <span>{{ hours }} hrs {{ minutes }} mins</span>
              </span>
            </div>
          </div>
          <div class="ml-4 flex flex-col">
            <h2 class="line-clamp-2 font-semibold">
              {{ recipe?.title }}
            </h2>
            <p class="text-muted flex items-center">
              <UIcon name="lucide:dollar-sign" />
              <span>
                {{ (recipe?.pricePerServing! / 100).toFixed(2) }} per serving •
                {{ recipe?.servings }} servings
              </span>
            </p>
          </div>
        </div>
        <div class="flex flex-1 items-start justify-end">
          <UButton variant="ghost" color="neutral" icon="lucide:info" />
          <UButton variant="ghost" color="secondary" icon="lucide:bookmark" />
          <UButton
            variant="ghost"
            color="error"
            icon="lucide:trash-2"
            @click="$emit('delete', id)"
          />
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
