<script setup lang="ts">
definePageMeta({
  middleware: "queryprotection",
});
const query = useState("query");
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Results for '{{ query }}'</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <!-- <span class="text-blue-400">{{ data?.totalResults }} results</span> -->
    </div>
  </div>
</template>
