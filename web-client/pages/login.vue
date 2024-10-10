<script setup lang="ts">
const { signIn, session } = useAuth();
</script>

<template>
  <h1>login</h1>
  <button @click="signIn(`credentials`)">login</button>
  <div>{{ session?.user }}</div>
</template>
