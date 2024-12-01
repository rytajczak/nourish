import { parse, stringify } from "zipson/lib";

interface Data {
  profile: {
    diet: string;
    calories: number;
    protein: number;
    carbs: number;
    fat: number;
  };
  intolerances: string[];
}

interface Goals {
  calories: number;
  protein: number;
  carbs: number;
  fat: number;
}

// ma look! good practices!
const TOTAL_STEPS = 3;

export const useOnboardingStore = defineStore(
  "onboarding",
  () => {
    const open = ref(false);
    const step = ref(0);
    const progress = computed(() => step.value * (100 / TOTAL_STEPS));
    const data = ref<Data>({
      profile: {
        diet: "",
        calories: 0,
        protein: 0,
        carbs: 0,
        fat: 0,
      },
      intolerances: [],
    });

    function begin() {
      open.value = true;
      step.value = 0;
      data.value.profile = {
        diet: "",
        calories: 0,
        protein: 0,
        carbs: 0,
        fat: 0,
      };
      data.value.intolerances = [];
    }

    function end() {
      open.value = false;
      step.value = 0;
      data.value.profile = {
        diet: "",
        calories: 0,
        protein: 0,
        carbs: 0,
        fat: 0,
      };
      data.value.intolerances = [];
    }

    function setGoals(goals: Goals) {
      data.value.profile.calories = goals.calories;
      data.value.profile.protein = goals.protein;
      data.value.profile.carbs = goals.carbs;
      data.value.profile.fat = goals.fat;
    }

    function setDiet(diet: string) {
      data.value.profile.diet = diet;
    }

    function setIntolerances(intolerances: string[]) {
      data.value.intolerances = intolerances;
    }

    function nextStep() {
      step.value += 1;
    }

    function previousStep() {
      step.value -= 1;
    }

    return {
      open,
      step,
      progress,
      data,
      begin,
      end,
      setGoals,
      setDiet,
      setIntolerances,
      nextStep,
      previousStep,
    };
  },
  {
    persist: {
      storage: sessionStorage,
      serializer: {
        serialize: (value) => stringify(value),
        deserialize: (value) => parse(value),
      },
    },
  },
);
