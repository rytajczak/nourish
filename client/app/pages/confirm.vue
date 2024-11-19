<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: false,
});

const { getUser } = useUserStore();

onMounted(async () => {
  if (await getUser()) navigateTo("/dashboard");
  else navigateTo("/onboarding");
});
</script>

<template>
  <div>
    <div class="flex h-screen w-screen items-center justify-center">
      <Icon name="svg-spinners:ring-resize" class="h-16 w-16" />
    </div>
  </div>
</template>
