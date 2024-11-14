<script setup lang="ts">
const searchString = ref("");
const query = useState("query", () => "");

async function search() {
  if (!searchString.value) return;
  query.value = searchString.value;
  navigateTo(`/recipes/search?query=${query.value}`);
}
</script>

<template>
  <UInput
    v-model="searchString"
    color="neutral"
    size="xl"
    icon="i-heroicons-magnifying-glass"
    placeholder="What would you like to cook?"
    @keydown.enter="search"
  />
</template>
