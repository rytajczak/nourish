<script setup lang="ts">
const { user } = useUserSession();
</script>

<template>
  <div class="flex h-20 items-center justify-between border-b-[1px] px-8">
    <RecipeSearch />
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
      <ULink to="/profile">
        <!-- <UAvatar size="lg" :src="user.picture" class="ms-4"></UAvatar> -->
      </ULink>
    </div>
  </div>
</template>
