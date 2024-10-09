<script setup lang="ts">
const username = ref("");
const password = ref("");

async function submitRegister() {}
</script>

<template>
  <h1>sign up</h1>
  <label for="email">email:</label><br />
  <input type="text" v-model="username" /><br />
  <label for="password">password:</label><br />
  <input type="text" v-model="password" /><br />
  <button @click="submitRegister">Sign Up</button>
</template>
