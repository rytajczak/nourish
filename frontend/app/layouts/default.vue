<script setup lang="ts">
const { clear } = useUserSession();
</script>

<template>
  <aside
    class="fixed top-0 left-0 block h-screen w-64 border-r-[1px] bg-[#181818]"
  >
    <div class="flex flex-col">
      <div class="h-20 border-b-[1px]">
        <NuxtLink to="/" class="block px-6 py-[24px] text-2xl font-semibold">
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
          to="/"
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
          to="/"
          icon="solar:chef-hat-heart-outline"
          >My Favorites</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/my-recipes"
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
  <div class="ml-64">
    <div class="flex h-20 items-center justify-between border-b-[1px] px-8">
      <UInput
        color="neutral"
        size="xl"
        icon="i-heroicons-magnifying-glass"
        placeholder="What would you like to cook?"
      />
      <div class="flex items-center">
        <div class="border-r pr-2">
          <UButton
            color="neutral"
            icon="solar:bell-outline"
            size="lg"
            variant="link"
          ></UButton>
          <UButton
            color="neutral"
            icon="solar:settings-outline"
            size="lg"
            variant="link"
          ></UButton>
        </div>
        <UAvatar
          size="lg"
          src="https://github.com/benjamincanac.png"
          class="ms-4"
        ></UAvatar>
      </div>
    </div>
    <slot></slot>
  </div>
</template>

<style scoped>
* {
  border-color: #242424;
}
</style>
