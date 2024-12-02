<script setup lang="ts">
const { profile, diet, intolerances } = useUserStore();
const { data } = useOnboardingStore();
const planner = usePlannerStore();

const open = ref(false);
const emit = defineEmits(["refetch"]);

async function handleGenerateDay() {
  await planner.generateDay(profile.calories, diet, intolerances);
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
    {{ diet }}
  </UButton>
  <!-- <UModal v-model:open="open" title="Select a time frame">
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
  </UModal> -->
</template>
