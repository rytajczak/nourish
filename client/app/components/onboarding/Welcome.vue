<script setup lang="ts">
const { logOut } = useUserStore();
const onboarding = useOnboardingStore();
async function handleLogOut() {
  navigateTo("/");
  await logOut();
}
</script>

<template>
  <div>
    <span class="text-2xl">
      <UIcon name="lucide:leaf" class="text-primary mr-1" />
      Nourish
    </span>
    <div class="my-4">
      <h2 class="text-2xl font-semibold">Welcome</h2>
      <h2 class="text-muted mt-1">
        Tell us a little about yourself to personalize your meal planning
        experience.
      </h2>
    </div>
    <div class="mt-6 grid grid-cols-2 gap-6">
      <UButton
        variant="outline"
        color="neutral"
        class="flex items-center justify-center py-3"
        @click="handleLogOut"
      >
        <span>Cancel</span>
      </UButton>
      <UButton
        color="neutral"
        class="flex items-center justify-center py-3"
        @click="onboarding.nextStep"
      >
        <span>Get Started</span>
      </UButton>
    </div>
  </div>
</template>
