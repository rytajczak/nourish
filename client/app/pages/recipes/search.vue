<script setup lang="ts">
definePageMeta({
  middleware: "queryprotection",
});
const query = useState("query");

const { data } = await useFetch(`/api/recipes/search`, {
  query: { query },
});
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Results for '{{ query }}'</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span class="text-blue-400">{{ data?.totalResults }} results</span>
    </div>
    <div class="grid grid-cols-5 gap-x-6 gap-y-8">
      <RecipeCard v-for="recipe of data?.results" v-bind="recipe" />
    </div>
  </div>
</template>
