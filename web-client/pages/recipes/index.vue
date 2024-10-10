<script setup lang="ts">
definePageMeta({ middleware: "auth", auth: { guestRedirectTo: "/login" } });
</script>

<template>
  <div>a searchable list of recipes</div>
</template>
