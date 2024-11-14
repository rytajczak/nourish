<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { user } = useUserStore();
</script>

<template>
  <div>
    <pre>{{ user }}</pre>
  </div>
</template>
