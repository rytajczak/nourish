<script setup lang="ts">
useSeoMeta({
  title: "Nourish Meal Planner",
  description: "A meal planning app to save you time",
  ogTitle: "Nourish Meal Planner",
  ogDescription: "A meal planning app to save you time",
  ogImage: "[og:image]",
  ogUrl: "[og:url]",
  twitterTitle: "[twitter:title]",
  twitterDescription: "[twitter:description]",
  twitterImage: "[twitter:image]",
  twitterCard: "summary",
});

useHead({
  htmlAttrs: {
    lang: "en",
  },
  link: [
    {
      rel: "icon",
      type: "image/png",
      href: "/favicon.png",
    },
  ],
});
</script>

<template>
  <div>
    <ul>
      <li>
        <NuxtLink to="/">landing</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/signup">Sign up</NuxtLink>
      </li>
      <li>
        <NuxtLink to="/login">Login</NuxtLink>
      </li>
    </ul>
  </div>
  <NuxtPage />
</template>
