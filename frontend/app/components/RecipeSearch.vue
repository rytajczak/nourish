<script setup lang="ts">
const route = useRoute();
const searchString = ref("");
const query = useState("query", () => "");

async function search() {
  query.value = searchString.value;
  navigateTo(`/recipes/search?query=${query.value}`);
  searchString.value = "";
}
</script>

<template>
  <UInput
    color="neutral"
    size="xl"
    icon="i-heroicons-magnifying-glass"
    placeholder="Search Recipes"
    v-model="searchString"
    @keydown.enter="search"
  >
  </UInput>
</template>
