<script setup lang="ts">
const { user } = useUserSession();
</script>

<template>
  <Sidebar />
  <div class="md:pl-64">
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
          <UPopover>
            <UButton
              color="neutral"
              icon="solar:settings-outline"
              size="lg"
              variant="link"
            ></UButton>
            <template #content></template>
          </UPopover>
          <UPopover>
            <UButton
              color="neutral"
              icon="solar:sun-2-outline"
              size="lg"
              variant="link"
            ></UButton>
            <template #content>
              <UCard>
                <URadioGroup
                  legend="theme"
                  v-model="$colorMode.preference"
                  :items="['system', 'dark', 'light', 'sepia']"
                ></URadioGroup>
              </UCard>
            </template>
          </UPopover>
        </div>
        <UAvatar size="lg" :src="user?.picture" class="ms-4"></UAvatar>
      </div>
    </div>
    <slot></slot>
  </div>
</template>
