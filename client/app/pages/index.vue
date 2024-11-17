<script setup lang="ts">
definePageMeta({
  layout: false,
});
</script>

<template>
  <div class="mt-8 mx-auto">
    <div class="flex flex-col">
      <h1 class="text-4xl">A MEAL PREP TOOL FOR THE BUSY</h1>
      <p class="mt-4 dark:text-gray-400">no bullshit. no bloat.</p>
      <p class="dark:text-gray-400">because your health is important</p>
      <div class="mt-4">
        <span class="flex items-center">
          <Icon name="solar:login-2-outline" class="me-2" />
          login
        </span>
        <div class="mt-2">
          <UButton
            to="/api/auth/google"
            icon="fe:google"
            color="neutral"
            variant="soft"
            external
            class="px-8"
          />
          <UButton
            to="/api/auth/github"
            icon="fe:github"
            color="neutral"
            variant="soft"
            external
            class="ms-2 px-8"
          />
        </div>
      </div>
    </div>
  </div>
</template>
