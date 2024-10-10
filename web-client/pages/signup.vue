<script setup lang="ts">
const { signIn, session } = useAuth();
</script>

<template>
  <h1>sign up</h1>
  <button @click="signIn(`credentials`)">sign up</button>
  <div>{{ session?.user }}</div>
</template>
