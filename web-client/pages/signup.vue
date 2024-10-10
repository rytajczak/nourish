<script setup lang="ts">
const { signIn, session } = useAuth();
const email = ref("");
const password = ref("");
</script>

<template>
  <h1>sign up</h1>
  <label for="email">email</label><br />
  <input type="text" v-model="email" /><br />
  <label for="email">password</label><br />
  <input type="text" v-model="password" /><br />
  <button @click="signIn(`credentials`)">sign up</button>
  <div>{{ session?.user }}</div>
</template>
