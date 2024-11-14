<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { profile, intolerances } = useUserStore();

useSeoMeta({
  title: `${profile?.username} | Profile`,
  description: `${profile?.username} | Profile`,
  ogTitle: `${profile?.username} | Profile`,
  ogDescription: `${profile?.username} | Profile`,
  ogImage: `${profile?.picture}`,
  ogUrl: `${profile?.username} | Profile`,
  twitterTitle: `${profile?.username} | Profile`,
  twitterDescription: `${profile?.username} | Profile`,
  twitterImage: `${profile?.picture}`,
  twitterCard: "summary",
});
</script>

<template>
  <div>
    <NuxtImg :src="profile?.picture ?? ''" />
    <pre>{{ profile }}</pre>
    <pre>{{ intolerances }}</pre>
  </div>
</template>
