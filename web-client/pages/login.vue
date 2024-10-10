<script setup lang="ts">
const { signIn, session } = useAuth();

const email = ref("");
const password = ref("");
</script>

<template>
  <h1>login</h1>
  <label for="email">email</label><br />
  <input type="text" v-model="email" /><br />
  <label for="email">password</label><br />
  <input type="text" v-model="password" /><br />
  <button @click="signIn(`credentials`, { email, password })">login</button>
  <div>or</div>
  <button @click="signIn(`google`)">login with oauth (google)</button>
</template>
