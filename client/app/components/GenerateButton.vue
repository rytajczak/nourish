<script setup>
const selectedDate = useState("selectedDate");
</script>

<template>
  <UModal title="Select a time frame">
    <UButton>Generate Meals</UButton>
    <template #body>
      <div class="flex items-center">
        <UButton
          @click="generateDailyPlan()"
          color="neutral"
          variant="outline"
          class="flex-1 py-8"
          ><div class="flex w-full flex-col items-center justify-center">
            <span class="text-lg font-semibold">Today</span>
            <span class="text-sm text-gray-500">{{
              selectedDate.getTime()
            }}</span>
          </div></UButton
        >
        <USeparator orientation="vertical" class="h-16 px-4">or</USeparator>
        <UButton
          @click="generateWeeklyPlan()"
          color="neutral"
          variant="outline"
          class="flex-1 py-8"
          ><div class="flex w-full flex-col items-center justify-center">
            <span class="text-lg font-semibold">Next 7 days</span>
            <span class="text-sm text-gray-500">{{}}</span>
          </div></UButton
        >
      </div>
    </template>
  </UModal>
</template>
