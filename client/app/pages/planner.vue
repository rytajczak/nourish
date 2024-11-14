<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="flex items-center">
      <h1 class="text-3xl font-semibold">Weekly Planner</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <UButton variant="subtle">Generate Plan</UButton>
    </div>
  </div>
</template>
