<script setup lang="ts">
definePageMeta({
  middleware: "queryprotection",
});

const query = useState<string>("query");
const { data, execute } = await useFetch("/api/recipes/search", {
  lazy: true,
  immediate: false,
  watch: false,
  query: { query },
});

watch(query, async () => {
  await execute();
});

const route = useRoute();
onMounted(async () => {
  if (!query.value) query.value = route.query.query?.toString() ?? "";
  await execute();
});
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Results for '{{ query }}'</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span class="text-blue-400">{{ data?.totalResults ?? 0 }} results</span>
    </div>
    <div class="grid grid-cols-1 gap-4 sm:grid-cols-2 lg:grid-cols-6">
      <RecipeCard v-for="recipe in data?.results" v-bind="recipe" />
    </div>
  </div>
</template>
