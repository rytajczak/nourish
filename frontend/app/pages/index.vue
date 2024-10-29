<template>
  <div class="flex w-screen p-4">
    <div></div>
    <UPopover>
      <UButton color="neutral" label="Open">Get Started</UButton>
      <template #content>
        <UCard>
          <template #header
            ><span class="font-bold">Login via provider</span></template
          >
          <div class="flex flex-col items-center">
            <UButton
              color="neutral"
              icon="logos:google-icon"
              to="/api/auth/google"
              label="Sign up with Google"
              external
              class="mb-4"
              >Sign up with Google</UButton
            >
            <UButton
              color="neutral"
              icon="i-simple-icons-github"
              to="/api/auth/google"
              label="Sign up with Google"
              external
              >Sign up with Google</UButton
            >
          </div>
        </UCard>
      </template>
    </UPopover>
  </div>
</template>
