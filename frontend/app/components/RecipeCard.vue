<script setup lang="ts">
import type { RecipePreview } from "~~/server/api/recipes/index.get";
const props = defineProps<RecipePreview>();
</script>

<template>
  <UCard
    :ui="{ body: 'p-0 sm:p-0' }"
    class="max-w-72 bg-[#f5f5f5] shadow-xl dark:bg-[#262626]"
  >
    <div class="relative">
      <img :src="props.image" alt="" class="rounded-t-lg" />
      <UCard
        :ui="{ body: 'py-1 px-2 sm:py-1 sm:px-2' }"
        class="absolute bottom-0 left-0 mb-2 ml-6"
      >
        <div class="flex items-center text-sm">
          <Icon name="solar:alarm-outline" class="me-1" />
          <span>{{ props.readyInMinutes }} mins</span>
        </div>
      </UCard>
    </div>
    <template #footer>
      <div class="flex justify-between">
        <span class="line-clamp-2 h-12 pr-8 font-semibold">{{
          props.title
        }}</span>
        <div>
          <UButton
            class="px-0 pt-1 pb-0"
            size="lg"
            variant="link"
            icon="solar:bookmark-outline"
          />
        </div>
      </div>
      <div class="mt-4 flex items-center justify-between">
        <span
          >{{ Math.round(props.calories.amount) }}
          {{ props.calories.unit }}</span
        >
      </div>
    </template>
  </UCard>
</template>
