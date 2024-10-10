<script setup lang="ts">
definePageMeta({ middleware: "auth", auth: { guestRedirectTo: "/login" } });
</script>

<template>
  <div>an interface for creating new recipes</div>
</template>
