<script setup lang="ts">
const username = ref("");
const password = ref("");

async function submitLogin() {}
</script>

<template>
  <input type="text" v-model="username" /><br />
  <input type="text" v-model="password" /><br />
  <button @click="submitLogin">login</button>
</template>
