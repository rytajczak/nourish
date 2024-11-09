<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
const { profile } = useUserStore();

useSeoMeta({
  title: `${profile?.displayName} | Profile`,
});
</script>

<template>
  <div>
    <NuxtImg :src="profile?.picture ?? ''" />
    <pre>{{ profile }}</pre>
  </div>
</template>
