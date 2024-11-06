<script setup lang="ts">
const route = useRoute();
const query = ref("");

async function search() {
  navigateTo(`/search-recipes?query=${query.value}`);
}
watch(
  () => route.path,
  () => {
    query.value = "";
  },
);
</script>

<template>
  <UInput
    color="neutral"
    size="xl"
    icon="i-heroicons-magnifying-glass"
    placeholder="Search Recipes"
    v-model="query"
    @keydown.enter="search"
  >
  </UInput>
</template>
