<script setup lang="ts">
const { user } = useUserSession();
const { signOut } = useUserStore();
const { end, nextStep } = useOnboardingStore();

async function handleSignOut() {
  end();
  await signOut();
}
</script>

<template>
  <div>
    <span class="text-2xl">
      <UIcon name="lucide:leaf" class="text-primary mr-1" />
    </span>
    <div class="mt-2 mb-6">
      <h2 class="text-2xl font-semibold">Welcome, {{ user?.firstName }}!</h2>
      <h2 class="text-muted mt-1">
        Tell us a little about yourself to personalize your meal planning
        experience.
      </h2>
    </div>
    <div class="mt-6 grid grid-cols-2 gap-6">
      <UButton
        variant="outline"
        color="neutral"
        class="flex items-center justify-center py-3 font-bold"
        @click="handleSignOut"
      >
        <span>Cancel</span>
      </UButton>
      <UButton
        color="neutral"
        class="flex items-center justify-center py-3 font-bold"
        @click="nextStep"
      >
        <span>Get Started</span>
      </UButton>
    </div>
  </div>
</template>
