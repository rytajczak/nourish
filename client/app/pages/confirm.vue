<script setup lang="ts">
definePageMeta({
  layout: false,
});

const { data } = await useFetch(`/api/users/me`);

// watch(data, () => {
//   if (data.value?.profile && user.value?.id == data.value?.profile.userId) {
//     setUser(data.value);
//     return navigateTo("/dashboard");
//   }
// });
</script>

<template>
  <div>
    <DevOnly>
      <pre>{{ data }}</pre>
    </DevOnly>
    <div class="flex h-screen w-screen items-center justify-center">
      <Icon name="svg-spinners:ring-resize" class="h-16 w-16" />
    </div>
  </div>
</template>
