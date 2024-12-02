<script setup lang="ts">
const { user } = useUserSession();
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
  <div
    class="border-border bg-background flex h-20 items-center justify-between border-b-[1px] px-8"
  >
    <SearchBar />
    <div class="flex items-center">
      <div>
        <UPopover>
          <UButton
            color="neutral"
            icon="lucide:settings"
            size="lg"
            variant="link"
          />
        </UPopover>
        <UPopover class="relative right-24">
          <UButton color="neutral" icon="lucide:sun" size="lg" variant="link" />
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
      <USeparator orientation="vertical" class="mx-3 h-8" />
      <UPopover class="relative right-9">
        <UAvatar size="lg" :src="user?.picture ?? ''" />
        <template #content>
          <UCard>
            <template #header>
              <div class="flex items-center">
                <UAvatar size="xl" :src="user?.picture ?? ''" class="mr-2" />
                <div class="flex flex-col">
                  <span class="text-sm font-semibold">
                    {{ user?.username }}
                  </span>
                  <span class="text-sm text-[var(--ui-text-muted)]">
                    {{ user?.email }}
                  </span>
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
                icon="lucide:circle-user-round"
              >
                Profile
              </UButton>
              <UButton
                size="lg"
                color="neutral"
                variant="ghost"
                class="mt-4 p-0"
                to="/recipes/saved"
                icon="lucide:book-marked"
              >
                Saved Recipes
              </UButton>
            </div>
            <template #footer>
              <UButton
                size="lg"
                color="error"
                variant="ghost"
                class="p-0"
                icon="lucide:log-out"
              >
                Log Out
              </UButton>
            </template>
          </UCard>
        </template>
      </UPopover>
    </div>
  </div>
</template>
