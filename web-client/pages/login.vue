<script setup lang="ts">
const username = ref("");
const password = ref("");

async function submitLogin() {}
</script>

<template>
  <h1>login</h1>
  <label for="username">username:</label>
  <input type="text" v-model="username" /><br />
  <label for="password">password:</label>
  <input type="text" v-model="password" /><br />
  <button @click="submitLogin">login</button>
</template>
