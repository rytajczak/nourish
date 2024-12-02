<script setup lang="ts">
definePageMeta({
  layout: false,
});
</script>

<template>
  <div class="relative flex h-dvh w-dvw items-center justify-center">
    <div class="flex flex-col">
      <h1 class="text-5xl font-bold">A Meal Prep Tool for the Busy.</h1>
      <p class="text-muted text-lg">
        because time is important and so is your health
      </p>
      <div class="mt-8 flex justify-center">
        <UButton
          to="/api/auth/google"
          icon="fe:google"
          color="neutral"
          variant="subtle"
          external
          size="xl"
          class="px-8"
        >
          Sign in with google
        </UButton>
      </div>
    </div>
    <div class="absolute top-0 left-0">
      <NuxtLink class="flex items-center px-6 py-[24px] text-2xl font-bold">
        <Icon name="lucide:leaf" class="text-primary mr-2" />
        Nourish
      </NuxtLink>
    </div>
  </div>
</template>
