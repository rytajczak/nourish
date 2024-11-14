<script setup lang="ts">
const query = useState("query");
</script>

<template>
  <!-- <div class="mx-8 mt-6">
    <pre>{{ data }}</pre>
    <div class="flex items-center">
      <h1 class="text-3xl font-semibold">Results for '{{ query }}'</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <span v-if="status === 'success'" class="text-blue-500 dark:text-blue-400"
        >{{ data?.totalResults }} recipes</span
      >
      <USkeleton v-if="status === 'pending'" class="h-6 w-24" />
      <span v-if="status === 'error'" class="text-red-500 dark:text-red-400"
        >Failed to get recipes</span
      >
    </div>
    <div class="mt-8 grid gap-8 md:grid-cols-2 lg:grid-cols-3 2xl:grid-cols-6">
      <div v-if="status === 'success'">
        <RecipeCard
          v-for="preview in data?.recipePreviews"
          :key="preview.id"
          v-bind="preview"
        />
      </div>
      <div v-if="status === 'pending'">
        <UCard
          v-for="i in 30"
          :key="i"
          :ui="{
            body: 'p-0 sm:p-0',
          }"
        >
          <USkeleton class="h-48 w-full max-w-72" />
          <template #footer>
            <USkeleton 4lass="h-6 w-24" />
            <USkeleton class="mt-2 h-6 w-14" />
          </template>
        </UCard>
      </div>
    </div>
  </div> -->
  <div>
    <pre>{{ query }}</pre>
  </div>
</template>
