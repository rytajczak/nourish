<script setup lang="ts">
import type { RecipePreview } from "~~/server/api/recipes/search.get";
const props = defineProps<RecipePreview>();
</script>

<template>
  <div>{{ props.title }}</div>
  <div>{{ props.image }}</div>
  <div>{{ props.calories }}</div>
  <div>{{ props.protein }}</div>
  <div>{{ props.carbs }}</div>
  <div>{{ props.fat }}</div>
  <div>{{ props.readyInMinutes }}</div>
</template>
