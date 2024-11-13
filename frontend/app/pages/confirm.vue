<script setup lang="ts">
definePageMeta({
  middleware: "auth",
  layout: false,
});

const { user } = useUserSession();

const { data } = await useFetch(`/api/users/${user.value?.id}`, {
  lazy: true,
});

// watch(data, () => {
//   if (data.value?.profile && user.value?.id == data.value?.profile.userId) {
//     setUser(data.value);
//     return navigateTo("/planner");
//   }
// });
</script>

<template>
  <div>
    <DevOnly>
      <pre>{{ user }}</pre>
      <pre>{{ data }}</pre>
    </DevOnly>
    <div class="flex h-screen w-screen items-center justify-center">
      <Icon name="svg-spinners:ring-resize" class="h-16 w-16" />
    </div>
  </div>
</template>
