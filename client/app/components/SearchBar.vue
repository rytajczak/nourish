<script setup lang="ts">
const query = useState("query", () => "");
const searchString = ref("");

function search() {
  if (!searchString.value) return;
  query.value = searchString.value;
  navigateTo(`/recipes/search?query=${query.value}`);
  searchString.value = "";
}
</script>

<template>
  <UInput
    v-model="searchString"
    color="neutral"
    size="xl"
    icon="lucide:search"
    placeholder="What would you like to cook?"
    @keydown.enter="search"
  />
</template>
