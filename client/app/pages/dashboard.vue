<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});

const planner = usePlannerStore();
const onboarding = useOnboardingStore();

const showingMeals = ref("all");

onMounted(() => {
  planner.fetchWeek();
});
</script>

<template>
  <div class="mx-8 mt-6">
    <div class="mb-6 flex items-center">
      <h1 class="text-3xl font-semibold">Dashboard</h1>
      <USeparator orientation="vertical" class="h-8 px-4" />
      <GenerateButton />
    </div>
    <div class="grid grid-cols-1 xl:grid-cols-5 xl:gap-8">
      <div class="col-span-3">
        <DaySelector />
        <div class="mt-6 flex items-center justify-between">
          <h2 class="text-xl font-semibold">Meals to prepare</h2>
          <USelect
            v-model="showingMeals"
            color="neutral"
            :items="['all', 'breakfast', 'lunch', 'dinner']"
            class="w-32"
          />
        </div>
        <pre>{{ planner.selectedDayInfo?.items }}</pre>
        <div>
          <UButton color="neutral" @click="planner.addEntry">Add entry</UButton>
        </div>
      </div>
      <div class="col-span-1 xl:col-span-2">
        <NutritionInfo />
      </div>
      <UModal v-model:open="onboarding.open" prevent-close>
        <template #content>
          <UProgress
            color="neutral"
            v-model="onboarding.progress"
            :ui="{ base: 'rounded-b' }"
          />
          <UCard
            :ui="{ root: 'rounded-none rounded-b-[calc(var(--ui-radius)*2)]' }"
          >
            <OnboardingWelcome v-if="onboarding.progress == 0" />
            <OnboardingGoals v-if="onboarding.progress == 33" />
            <OnboardingDiet v-if="onboarding.progress == 66" />
            <OnboardingIntolerance v-if="onboarding.progress == 99" />
          </UCard>
        </template>
      </UModal>
    </div>
  </div>
</template>
