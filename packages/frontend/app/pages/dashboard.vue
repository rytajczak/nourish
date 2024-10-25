<script setup lang="ts">
// definePageMeta({
//   middleware: "auth",
// });

const query = ref("");
const previousQuery = ref("");

const { data, status, execute } = await useFetch<any[]>(
  `http://localhost:8081/search`,
  {
    query: { query },
    watch: false,
    immediate: false,
  }
);
async function onSubmit() {
  previousQuery.value = query.value;
  await execute();
}
</script>

<template>
  <div class="max-w-xl mx-auto mt-4">
    <label
      for="default-search"
      class="mb-2 text-sm font-medium text-gray-900 sr-only"
      >Search</label
    >
    <div class="relative">
      <div
        class="absolute inset-y-0 start-0 flex items-center ps-3 pointer-events-none"
      >
        <svg
          class="w-4 h-4 text-gray-500"
          aria-hidden="true"
          xmlns="http://www.w3.org/2000/svg"
          fill="none"
          viewBox="0 0 20 20"
        >
          <path
            stroke="currentColor"
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="m19 19-4-4m0-7A7 7 0 1 1 1 8a7 7 0 0 1 14 0Z"
          />
        </svg>
      </div>
      <input
        type="search"
        id="default-search"
        class="block w-full p-4 ps-10 text-sm text-gray-900 border rounded-full focus:ring-blue-500 focus:border-blue-500"
        placeholder="search for recipes and cuisines"
        required
        v-model="query"
      />
      <button
        @click="onSubmit"
        class="text-white absolute end-2.5 bottom-2.5 bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-full text-sm px-4 py-2"
      >
        Search
      </button>
    </div>
  </div>
  <div class="flex max-w-xl mx-auto flex-wrap mt-4">
    <span v-if="status === 'pending'">fetching data...</span>
    <div v-if="status === 'success'">
      <span>results for '{{ previousQuery }}'</span>
      <RecipeCard class="mt-4" v-for="recipe in data" v-bind="recipe" />
    </div>
    <div v-if="status === 'error'">failed to fetch data</div>
  </div>
</template>
