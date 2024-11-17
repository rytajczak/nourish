<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: false,
});

const { loadUser } = useUserStore();

await useFetch(`/api/users/me`, {
  onResponse({ response }) {
    switch (response.status) {
      case 404:
        return navigateTo("/onboarding");
      case 200:
        loadUser(response._data);
        return navigateTo("/dashboard");
      default:
        return navigateTo("/");
    }
  },
});
</script>

<template>
  <div>
    <div class="flex h-screen w-screen items-center justify-center">
      <Icon name="svg-spinners:ring-resize" class="h-16 w-16" />
    </div>
  </div>
</template>
