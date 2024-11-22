export const useOnboardingStore = defineStore(
  "onboarding",
  () => {
    const open = ref<boolean>(false);
    const progress = ref<number>(0);
    const data = ref({});

    function onboardUser() {
      open.value = true;
      progress.value = 0;
      data.value = {};
    }

    function nextStep() {
      progress.value += 25;
    }

    function previousStep() {
      progress.value -= 25;
    }

    return {
      open,
      progress,
      data,
      onboardUser,
      nextStep,
      previousStep,
    };
  },
  { persist: true },
);
