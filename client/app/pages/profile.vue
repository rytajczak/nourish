<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { user } = useUserSession();
const { profile, intolerances, dislikedIngredients, savedRecipes } =
  useUserStore();
</script>

<template>
  <div>
    <pre>user: {{ user }}</pre>
    <pre>profile: {{ profile }}</pre>
    <pre>intolerances: {{ intolerances }}</pre>
    <pre>dislikedIngredients: {{ dislikedIngredients }}</pre>
    <pre>savedRecipes: {{ savedRecipes }}</pre>
  </div>
</template>
