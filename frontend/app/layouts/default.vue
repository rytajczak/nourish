<template>
  <aside class="fixed top-0 left-0 block h-screen w-64 border-r-[1px]">
    <div class="flex flex-col">
      <div class="h-16 border-b-[1px]">
        <NuxtLink to="/" class="block px-6 py-[14px] text-2xl font-semibold">
          Nourish
        </NuxtLink>
      </div>
      <div class="flex flex-col px-6 py-4">
        <UButton
          color="neutral"
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
          to="/home"
          icon="solar:calendar-linear"
          >Weekly Prep</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/home"
          icon="solar:bookmark-outline"
          >Saved Recipes</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/home"
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
          to="/home"
          icon="solar:medal-ribbon-star-outline"
          >My Plan</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/home"
          icon="solar:chef-hat-heart-outline"
          >My Favorites</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/home"
          icon="solar:document-text-outline"
          >My Recipes</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/home"
          icon="solar:logout-2-outline"
          >Log Out</UButton
        >
      </div>
    </div>
  </aside>
  <div class="ml-64">
    <div class="h-16 border-b-[1px]"></div>
    <slot></slot>
  </div>
</template>

<style scoped>
* {
  border-color: #242424;
}
</style>
