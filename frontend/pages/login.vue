<script setup lang="ts">
const email = ref("");
const password = ref("");

async function handleSumbit() {
  console.log(email.value);
  console.log(password.value);
}
</script>

<template>
  <h1>Login</h1>
  <form @submit.prevent="handleSumbit">
    <input type="email" v-model="email" /><br />
    <input type="password" v-model="password" /><br />
    <button type="submit">login</button>
  </form>
  <a href="/api/auth/google">login with google</a>
</template>
