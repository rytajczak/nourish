<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { user } = useUserSession();
// const { data: profile } = await useFetch(`/api/profiles/${user.id}`);
</script>

<template>
  <pre>user: {{ user }}</pre>
  <!-- <pre>data: {{ profile }}</pre> -->
</template>
