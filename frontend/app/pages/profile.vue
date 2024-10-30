<script setup lang="ts">
const { user } = useUserSession();
const { data: profile } = await useFetch("/api/profiles");
</script>

<template>
  <pre>user: {{ user }}</pre>
  <pre>data: {{ profile }}</pre>
</template>
