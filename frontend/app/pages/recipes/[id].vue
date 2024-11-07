<script setup lang="ts">
const route = useRoute();
const { data } = await useFetch(`/api/recipes/search/${route.params.id}`, {
  lazy: true,
});
</script>

<template>
  <pre>{{ data }}</pre>
</template>
