<script setup lang="ts">
const {
  data: recipePreviews,
  status,
  execute,
} = useFetch("/api/recipes", {
  immediate: false,
  watch: false,
});
const searchString = ref("");
</script>

<template>
  <UInput
    color="neutral"
    size="xl"
    icon="i-heroicons-magnifying-glass"
    placeholder="Search Recipes"
    v-model="searchString"
  >
    <template #trailing>
      <UModal
        :title="`${recipePreviews?.totalResults} results for '${searchString}'`"
      >
        <UButton
          @click="execute()"
          :disabled="searchString.length < 1 || status === 'pending'"
          variant="subtle"
          size="sm"
          >Search</UButton
        >
        <template #body>
          <div v-if="status === 'success'"></div>
          <div v-else>pending</div>
        </template>
      </UModal>
    </template>
  </UInput>
</template>
