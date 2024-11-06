<script setup lang="ts">
const { clear } = useUserSession();
</script>

<template>
  <aside
    class="fixed top-0 left-0 block h-screen w-64 -translate-x-full border-r-[1px] transition-transform md:translate-x-0"
  >
    <div class="flex flex-col">
      <div class="h-20 border-b-[1px]">
        <NuxtLink to="/" class="block px-6 py-[24px] text-2xl font-semibold">
          Nourish
        </NuxtLink>
      </div>
      <div class="flex flex-col px-6 py-4">
        <UButton
          color="primary"
          variant="subtle"
          icon="solar:add-circle-bold"
          class="items-center rounded-lg px-6 py-3"
          >Add a Recipe</UButton
        >
        <span class="ms-3 mt-8 mb-2 text-sm text-gray-700 dark:text-gray-400"
          >Menu</span
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/planner"
          icon="solar:calendar-linear"
          >Planner</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/saved-recipes"
          icon="solar:bookmark-outline"
          >Saved Recipes</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/grocery-list"
          icon="solar:bag-4-outline"
          >Grocery List</UButton
        >
        <span class="ms-3 mt-8 mb-2 text-sm text-gray-700 dark:text-gray-400"
          >Settings</span
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/profile"
          icon="solar:user-circle-outline"
          >Profile</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/custom-recipes"
          icon="solar:document-text-outline"
          >My Recipes</UButton
        >
        <UButton
          @click="clear"
          size="lg"
          color="neutral"
          variant="ghost"
          to="/"
          icon="solar:logout-2-outline"
          >Log Out</UButton
        >
      </div>
    </div>
  </aside>
</template>
