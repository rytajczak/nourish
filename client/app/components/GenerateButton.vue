<script setup lang="ts">
const { profile, intolerances } = useUserStore();
const planner = usePlannerStore();

const open = ref(false);
const emit = defineEmits(["refetch"]);

const weekEndDate = computed(() => {
  return new Date(planner.weekStartDate.getTime() + 6 * 24 * 60 * 60 * 1000);
});

async function handleGenerateDay() {
  await planner.generateDay(profile.calories, profile.diet, intolerances);
  open.value = false;
  emit("refetch");
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
          loading-auto
        >
          <div class="flex w-full flex-col items-center justify-center">
            <span class="text-lg font-semibold">Selected day</span>
            <span class="text-muted text-sm">
              {{
                planner.selectedDate.toLocaleDateString("en-US", {
                  weekday: "long",
                })
              }}
            </span>
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
            <span class="text-muted text-sm">
              {{
                planner.weekStartDate.toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                })
              }}
              -
              {{
                weekEndDate.toLocaleDateString("en-US", {
                  month: "short",
                  day: "numeric",
                })
              }}
            </span>
          </div>
        </UButton>
      </div>
    </template>
  </UModal>
</template>
