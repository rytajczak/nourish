<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const { loggedIn, user, session, fetch, clear } = useUserSession();

async function signOut() {
  await clear();
  navigateTo("/login");
}
</script>

<template>
  <h1>my profile</h1>
  <div>{{ (loggedIn, user, session) }}</div>
  <button @click="signOut">sign out</button>
</template>
