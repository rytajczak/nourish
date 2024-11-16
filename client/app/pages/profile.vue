<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { user, intolerances } = useUserStore();
</script>

<template>
  <div>
    <pre>{{ user }}</pre>
    <pre>{{ intolerances }}</pre>
  </div>
</template>
