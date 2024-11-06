<script setup lang="ts">
const route = useRoute();
const query = ref(route.query.query);

const { data, status, execute } = await useFetch("/api/recipes/search", {
  lazy: true,
  query: { query },
});

watch(
  () => route.query.query,
  (newQuery) => {
    query.value = newQuery;
    execute();
  },
);
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="flex items-center">
      <h1 class="text-3xl font-semibold">
        Results for '{{ route.query.query }}'
      </h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span v-if="status === 'success'" class="text-blue-500 dark:text-blue-400"
        >{{ data?.totalResults }} recipes</span
      >
      <USkeleton v-if="status === 'pending'" class="h-6 w-24"></USkeleton>
      <span v-if="status === 'error'" class="text-red-400"
        >Failed to get recipes</span
      >
    </div>
    <div class="mt-8 grid gap-8 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-6">
      <RecipeCard
        v-if="status === 'success'"
        v-for="preview in data?.recipePreviews"
        v-bind="preview"
      ></RecipeCard>
      <UCard
        v-if="status === 'pending'"
        v-for="i in 30"
        :ui="{
          body: 'p-0 sm:p-0',
        }"
      >
        <USkeleton class="h-48 w-full max-w-72"></USkeleton>
        <template #footer>
          <USkeleton class="h-6 w-24"></USkeleton>
          <USkeleton class="mt-2 h-6 w-14"></USkeleton>
        </template>
      </UCard>
    </div>
  </div>
</template>
