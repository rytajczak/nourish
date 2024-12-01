<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: false,
});

const { begin } = useOnboardingStore();
const { getUser } = useUserStore();

onMounted(async () => {
  if (!(await getUser())) begin();
  navigateTo("/dashboard");
});
</script>

<template>
  <div>
    <div class="flex h-screen w-screen items-center justify-center">
      <Icon name="svg-spinners:ring-resize" class="h-16 w-16" />
    </div>
  </div>
</template>
