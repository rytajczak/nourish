<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const session = useUserSession();
const user = useUserStore();
</script>

<template>
  <div>
    <pre>user: {{ session.user }}</pre>
    <pre>profile: {{ user.profile }}</pre>
    <pre>intolerances: {{ user.intolerances }}</pre>
    <pre>savedRecipes: {{ user.savedRecipes }}</pre>
  </div>
</template>
