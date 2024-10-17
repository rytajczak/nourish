<script setup lang="ts">
const { data, status } = await useFetch("http://localhost:8080/recipes", {});
</script>

<template>
  <h1>dashboard</h1>

  <div v-if="status === 'success'">{{ data }}</div>
  <div v-else>loading...</div>
</template>
