<script setup lang="ts">
const onboarding = useOnboardingStore();
</script>

<template>
  <div>
    <span class="text-2xl">
      <UIcon name="lucide:target" class="mr-1" />
    </span>
    <div class="my-4">
      <h2 class="text-2xl font-semibold">Set your daily goals</h2>
      <h2 class="text-muted mt-1">
        Tell us a little about yourself to personalize your meal planning
        experience.
      </h2>
    </div>
    <div class="mt-6 grid grid-cols-2 gap-6">
      <UButton
        variant="subtle"
        color="neutral"
        class="flex items-center justify-center py-3"
        @click="onboarding.previousStep"
      >
        <span>Go Back</span>
      </UButton>
      <UButton
        color="neutral"
        class="flex items-center justify-center py-3 font-bold"
        @click="onboarding.nextStep"
      >
        <span>Next Step: diet</span>
      </UButton>
    </div>
  </div>
</template>
