<script setup lang="ts">
const { logOut } = useUserStore();
</script>

<template>
  <aside
    class="fixed top-0 left-0 z-50 block h-screen w-64 -translate-x-full border-r-[1px] border-[var(--ui-border)] bg-[var(--ui-bg-elevated)] transition-transform md:translate-x-0"
  >
    <div class="flex flex-col">
      <div class="h-20">
        <NuxtLink
          to="/dashboard"
          class="flex items-center px-6 py-[24px] text-2xl font-bold"
        >
          <Icon name="lucide:leaf" class="mr-2 text-[var(--ui-primary)]" />
          Nourish
        </NuxtLink>
      </div>
      <div class="flex flex-col px-6 py-6">
        <span class="ms-3 mb-2 text-sm text-gray-700 dark:text-gray-400"
          >Menu</span
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/dashboard"
          icon="solar:calendar-linear"
          >dashboard</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/recipes/saved"
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
          to="/recipes/custom"
          icon="solar:document-text-outline"
          >My Recipes</UButton
        >
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          icon="solar:logout-2-outline"
          @click="logOut"
          >Log Out</UButton
        >
      </div>
    </div>
  </aside>
</template>
