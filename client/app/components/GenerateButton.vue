<script setup lang="ts">
const planner = usePlannerStore();
const selectedDate = useState<Date>("selectedDate");

const open = ref(false);

const weekStart = new Date();
const weekEnd = computed(() => {
  return new Date(weekStart.getTime() + 6 * 24 * 60 * 60 * 1000);
});

async function handleGenerateDay() {
  await planner.generateDay(selectedDate.value);
  open.value = false;
}

async function handleGenerateWeek() {
  await planner.generateWeek();
  open.value = false;
}
</script>

<template>
  <UModal v-model:open="open" title="Select a time frame">
    <UButton color="neutral" variant="solid">Generate Meals</UButton>
    <template #body>
      <div class="flex items-center">
        <UButton
          @click="handleGenerateDay"
          color="neutral"
          variant="outline"
          class="flex-1 py-8"
        >
          <div class="flex w-full flex-col items-center justify-center">
            <span class="text-lg font-semibold">Selected day</span>
            <span class="text-muted text-sm">{{
              selectedDate.toLocaleDateString("en-US", {
                weekday: "long",
              })
            }}</span>
          </div>
        </UButton>
        <USeparator orientation="vertical" class="h-24 px-4">or</USeparator>
        <UButton
          @click=""
          color="neutral"
          variant="outline"
          class="flex-1 py-8"
        >
          <div class="flex w-full flex-col items-center justify-center">
            <span class="text-lg font-semibold">This week</span>
            <span class="text-muted text-sm"
              >{{
                weekStart.toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                })
              }}
              -
              {{
                weekEnd.toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                })
              }}</span
            >
          </div>
        </UButton>
      </div>
    </template>
  </UModal>
</template>
