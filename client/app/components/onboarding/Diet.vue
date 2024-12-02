<script setup lang="ts">
const onboarding = useOnboardingStore();
const { profile } = useUserStore();

const items = ref([
  {
    label: "No Diet",
    description: "No specific dietary restrictions or preferences.",
    value: "none",
  },
  {
    label: "Pescetarian",
    description: "Includes fish and seafood but avoids other meats.",
    value: "pescetarian",
  },
  {
    label: "Lacto Vegetarian",
    description: "Includes dairy but avoids meat, fish, and eggs.",
    value: "lacto vegetarian",
  },
  {
    label: "Ovo Vegetarian",
    description: "Includes eggs but avoids meat, fish, and dairy.",
    value: "ovo vegetarian",
  },
  {
    label: "Vegan",
    description: "Excludes all animal products, including dairy and eggs.",
    value: "vegan",
  },
  {
    label: "Paleo",
    description:
      "Focuses on unprocessed foods, e.g., meat, fish, fruits, and vegetables.",
    value: "paleo",
  },
  {
    label: "Primal",
    description: "Similar to Paleo, with the addition of dairy.",
    value: "primal",
  },
  {
    label: "Vegetarian",
    description: "Excludes meat and fish but may include dairy and eggs.",
    value: "vegetarian",
  },
]);

const diet = ref("none");

function handleNextStep() {
  onboarding.setDiet(diet.value);
  profile.diet = diet.value;
  onboarding.nextStep();
}
</script>

<template>
  <div>
    <span class="text-2xl">
      <UIcon name="lucide:apple" />
    </span>
    <div class="mt-2 mb-6">
      <h2 class="text-2xl font-semibold">Choose Your Diet</h2>
      <h2 class="text-muted mt-1">Pick a diet that suits your lifestyle</h2>
    </div>
    <URadioGroup
      color="secondary"
      v-model="diet"
      :items="items"
      :ui="{
        item: 'my-1',
      }"
    />
    <div class="mt-7 grid grid-cols-2 gap-6">
      <UButton
        size="xl"
        variant="outline"
        color="neutral"
        class="flex items-center justify-center font-bold"
        @click="onboarding.previousStep"
      >
        <span>Go Back</span>
      </UButton>
      <UButton
        size="xl"
        color="neutral"
        class="flex items-center justify-center font-bold"
        @click="handleNextStep"
      >
        <span>Next Step: Intolerance</span>
      </UButton>
    </div>
  </div>
</template>
