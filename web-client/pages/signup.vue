<script setup lang="ts">
const username = ref("");
const password = ref("");

async function submitRegister() {}
</script>

<template>
  <input type="text" v-model="username" /><br />
  <input type="text" v-model="password" /><br />
  <button @click="submitRegister">login</button>
</template>
