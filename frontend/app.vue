<script setup lang="ts">
const { signIn } = useAuth();
</script>

<template>
  <div>
    <NuxtRouteAnnouncer />
    <button @click="signIn(`google`)">sign in with google</button>
  </div>
</template>
