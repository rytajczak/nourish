<script setup lang="ts">
const { savedRecipes } = useRecipes();
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="flex items-center">
      <h1 class="text-3xl font-semibold">Saved Recipes</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span class="text-blue-500 dark:text-blue-400">recipes found</span>
    </div>
    <div class="mt-8 grid gap-8 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-6">
      <RecipeCard
        v-for="recipe in savedRecipes"
        :key="recipe.id"
        v-bind="recipe"
      ></RecipeCard>
    </div>
  </div>
</template>
