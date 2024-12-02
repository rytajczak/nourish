<script setup lang="ts">
import type { Recipe } from "~~/types/types";
const { savedRecipes } = useUserStore();

const open = ref(false);
const searchString = ref("");

const csv = ref("");
if (savedRecipes.length > 0) {
  csv.value = savedRecipes.join(",");
}
const { data: recipes, status } = await useFetch<Recipe[]>(
  `/api/recipes/info-bulk`,
  {
    lazy: true,
    query: { ids: csv },
  },
);

const emit = defineEmits(["add"]);
function handleAdd(id: number) {
  emit("add", id);
  open.value = false;
}
</script>

<template>
  <USlideover v-model:open="open" title="Add a saved recipe">
    <UButton color="neutral">Add Recipe</UButton>
    <template #header>
      <UInput
        v-model="searchString"
        color="neutral"
        size="xl"
        icon="lucide:search"
        placeholder="Search Saved Recipes"
        :ui="{
          base: 'bg-elevated',
        }"
      />
    </template>
    <template #body>
      <UCard
        v-for="recipe in recipes"
        :ui="{ header: 'p-0 sm:p-0' }"
        class="mb-6"
      >
        <template #header>
          <img :src="recipe.image" alt="" class="rounded-xl" />
        </template>
        <span class="mr-4 line-clamp-2 min-h-[3rem] flex-3 font-semibold">
          {{ recipe.title }}
        </span>
        <UButton
          class="mt-4"
          color="neutral"
          loading-auto
          @click="handleAdd(recipe.id)"
        >
          Add to Day
        </UButton>
      </UCard>
    </template>
  </USlideover>
</template>
