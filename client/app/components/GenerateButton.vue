<script setup lang="ts">
const planner = usePlannerStore();

const open = ref(false);
const emit = defineEmits(["refetch"]);

async function handleGenerateDay() {
  await planner.generateDay();
  open.value = false;
  emit("refetch");
}
</script>

<template>
  <UButton
    color="neutral"
    variant="solid"
    @click="handleGenerateDay"
    loading-auto
  >
    Generate Meals
  </UButton>
</template>
