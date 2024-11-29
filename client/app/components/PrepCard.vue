<script setup lang="ts">
const props = defineProps<{ id: number }>();
const { recipeMap } = usePlannerStore();
const recipeInfo = recipeMap?.get(props.id);
</script>

<template>
  <UCard class="mb-4">
    <template #header>
      <div class="flex">
        <img :src="recipeInfo?.image" alt="" class="h-24 rounded-xl" />
        <div class="ml-4 flex flex-col">
          <h2 class="text-lg font-semibold">{{ recipeInfo?.title }}</h2>
          <div class="flex flex-row">
            <UBadge v-for="diet in recipeInfo?.diets">{{ diet }}</UBadge>
          </div>
        </div>
      </div>
    </template>
  </UCard>
</template>
