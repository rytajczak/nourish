<script setup lang="ts">
import type { Recipe } from "~~/server/api/recipes/search.get";
const props = defineProps<Recipe>();
</script>

<template>
  <UCard
    :ui="{
      root: 'shadow-xl rounded-xl',
      header: 'p-0 sm:p-0',
      body: 'p-4 sm:p-4',
      footer: 'p-4 sm:p-4',
    }"
    class="cursor-pointer transition-all duration-300 hover:scale-105 dark:bg-[#191919]"
  >
    <template #header>
      <div class="relative">
        <img
          :src="props.image"
          :alt="props.title"
          class="w-full rounded-t-xl"
        />
        <div class="absolute bottom-2 left-3">
          <span
            class="flex items-center rounded-2xl bg-white/70 px-2 py-1 text-sm shadow-xl backdrop-blur-md dark:bg-[#191919]/70"
          >
            <UIcon name="solar:alarm-outline" class="mr-1" />
            {{ props.readyInMinutes }} min
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
          size="lg"
          variant="ghost"
          icon="solar:bookmark-outline"
          class="p-1"
        />
      </div>
    </div>
    <div
      class="mt-4 flex items-center justify-between text-gray-500 dark:text-gray-400"
    >
      <span
        >{{ props.nutrition.nutrients[0]?.amount }}
        {{ props.nutrition.nutrients[0]?.unit }}
        •
        {{ props.servings }} servings</span
      >
    </div>
  </UCard>
</template>
