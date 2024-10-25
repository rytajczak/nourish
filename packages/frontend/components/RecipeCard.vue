<script setup lang="ts">
const props = defineProps<{
  id: number;
  title: string;
  image: string;
  imageType: string;
}>();
</script>

<template>
  <div class="flex rounded-3xl bg-[#f7f7f9] p-2">
    <NuxtImg
      :src="image"
      alt="image failed to load"
      class="rounded-3xl"
    ></NuxtImg>
    <p class="p-4">{{ title }}</p>
  </div>
</template>
