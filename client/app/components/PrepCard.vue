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
    :ui="{ header: 'pl-0 sm:pl-2', body: 'p-0 sm:p-0' }"
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
          <NuxtImg
            :src="recipe?.image"
            class="h-28 w-40 rounded-md"
            preload
            :placeholder="[144, 128]"
          />
          <div class="ml-4 flex flex-col">
            <h2 class="line-clamp-2 font-semibold">
              {{ recipe?.title }}
            </h2>
            <p class="text-muted flex items-center">
              <UIcon name="lucide:clock" class="me-1" />
              <span>{{ hours }} hours and {{ minutes }} minutes</span>
            </p>
          </div>
        </div>
        <div class="flex flex-1 items-start justify-end">
          <UButton
            variant="ghost"
            color="error"
            icon="lucide:trash-2"
            @click="$emit('delete', id)"
          />
        </div>
      </div>
    </template>
  </UCard>
</template>
