<script setup lang="ts">
const { createUser } = useUserStore();
const onboarding = useOnboardingStore();

const buttons = ref([
  { name: "dairy", icon: "lucide:milk" },
  { name: "egg", icon: "lucide:egg" },
  { name: "gluten", icon: "lucide:wheat" },
  { name: "grain", icon: "lucide:wheat" },
  { name: "peanut", icon: "lucide:nut" },
  { name: "seafood", icon: "lucide:fish" },
  { name: "sesame", icon: "lucide:mouse" },
  { name: "shellfish", icon: "lucide:fish" },
  { name: "soy", icon: "lucide:bean" },
  { name: "sulfite", icon: "lucide:wine" },
  { name: "tree nut", icon: "lucide:nut" },
  { name: "wheat", icon: "lucide:wheat" },
]);

const intolerances = ref<string[]>([]);
function handleToggle(name: string, value: boolean) {
  if (value) {
    if (!intolerances.value.includes(name)) {
      intolerances.value.push(name);
    }
  } else {
    intolerances.value = intolerances.value.filter((item) => item != name);
  }
}

async function handleCreateUser() {
  onboarding.setIntolerances(intolerances.value);
  if (await createUser(onboarding.data)) onboarding.end();
}
</script>

<template>
  <div>
    <span class="text-2xl">
      <UIcon name="lucide:heart" />
    </span>
    <div class="mt-2 mb-6">
      <h2 class="text-2xl font-semibold">Avoid Certain Foods</h2>
      <h2 class="text-muted mt-1">Tell us what to exclude.</h2>
    </div>
    <div class="grid grid-cols-3 gap-4">
      <OnboardingIntoleranceButton
        v-for="button in buttons"
        :key="button.name"
        :name="button.name"
        :icon="button.icon"
        @toggle="handleToggle"
      />
    </div>
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
        loading-auto
        @click="handleCreateUser"
      >
        <span>Submit</span>
      </UButton>
    </div>
  </div>
</template>
