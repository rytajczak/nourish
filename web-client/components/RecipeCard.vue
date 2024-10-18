<script setup lang="ts">
const props = defineProps<{
  id: number
  title: string
  image: string
  imageType: string
}>()
</script>

<template>
  <div>{{ props.title }}</div>
  <NuxtImg :src="image" />
</template>