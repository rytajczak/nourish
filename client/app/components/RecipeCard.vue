<script setup lang="ts">
import type { Recipe } from "~~/server/api/recipes/search.get";
const props = defineProps<Recipe>();
</script>

<template>
  <UCard :ui="{ header: 'p-0 sm:p-0' }">
    <template #header>
      <NuxtImg
        :src="props.image"
        sizes="100vw md:50vw lg:33vw"
        placeholder
        preload
        class="rounded-lg"
      />
    </template>
  </UCard>
</template>
