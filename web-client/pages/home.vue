<script setup lang="ts">
definePageMeta({ middleware: "auth", auth: { guestRedirectTo: "/login" } });
const { session, signOut } = useAuth();
</script>

<template>
  <h1>home</h1>
  <div>user: {{ session?.user }}</div>
  <div>expires: {{ session?.expires }}</div>
  <button @click="signOut()">sign out</button>
</template>
