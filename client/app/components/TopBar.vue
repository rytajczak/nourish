<script setup lang="ts">
const { user } = useUserStore();
const themes = ref([
  {
    label: "System",
    description: "Match system theme",
    value: "system",
  },
  {
    label: "Dark",
    description: "Looks good in the night",
    value: "dark",
  },
  {
    label: "Light",
    description: "Looks good in the day",
    value: "light",
  },
]);
</script>

<template>
  <div class="flex h-20 items-center justify-between border-b-[1px] px-8">
    <SearchBar />
    <div class="flex items-center">
      <div class="border-r pr-2">
        <UPopover>
          <UButton
            color="neutral"
            icon="solar:settings-outline"
            size="lg"
            variant="link"
          />
        </UPopover>
        <UPopover class="relative right-24">
          <UButton
            color="neutral"
            icon="solar:sun-2-outline"
            size="lg"
            variant="link"
          />
          <template #content>
            <UCard>
              <URadioGroup
                v-model="$colorMode.preference"
                color="neutral"
                legend="Theme"
                :items="themes"
              />
            </UCard>
          </template>
        </UPopover>
      </div>
      <UPopover class="relative right-9">
        <UAvatar size="lg" :src="user?.picture ?? ''" class="ms-4" />
        <template #content>
          <UCard>
            <template #header>
              <div class="flex items-center">
                <UAvatar size="xl" :src="user?.picture ?? ''" class="mr-2" />
                <div class="flex flex-col">
                  <span class="text-sm font-semibold">{{
                    user?.username
                  }}</span>
                  <span class="text-sm text-gray-600 dark:text-gray-500">{{
                    user?.email
                  }}</span>
                </div>
              </div>
            </template>
            <div class="flex flex-col">
              <UButton
                size="lg"
                color="neutral"
                variant="ghost"
                class="p-0"
                to="/profile"
                icon="solar:user-circle-outline"
                >Profile</UButton
              >
              <UButton
                size="lg"
                color="neutral"
                variant="ghost"
                class="mt-4 p-0"
                to="/recipes/custom"
                icon="solar:document-text-outline"
                >My Recipes</UButton
              >
            </div>
            <template #footer>
              <UButton
                size="lg"
                color="neutral"
                variant="ghost"
                class="p-0"
                icon="solar:logout-2-outline"
                >Log Out</UButton
              >
            </template>
          </UCard>
        </template>
      </UPopover>
    </div>
  </div>
</template>
