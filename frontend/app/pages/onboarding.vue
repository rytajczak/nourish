<script setup>
definePageMeta({
  layout: false,
});
const vueform = ref({
  size: "md",
  displayErrors: false,
  validateOn: "step|change",
  steps: {
    page0: {
      elements: [
        "goal_header",
        "divider",
        "calories",
        "protein",
        "carbs",
        "fat",
        "divider_1",
      ],
      label: "Daily Goals",
    },
    page1: {
      elements: ["checkbox", "diet_header", "divider_2", "diet", "divider_3"],
      label: "Diet",
    },
    page2: {
      label: "Intolerances",
      elements: [
        "intolerance_header",
        "divider_5",
        "intolerances",
        "divider_4",
      ],
    },
  },
  schema: {
    goal_header: {
      type: "static",
      tag: "h1",
      content: "What is your daily nutrition goal?",
    },
    divider: {
      type: "static",
      tag: "hr",
    },
    calories: {
      type: "text",
      rules: ["required", "min:0", "integer"],
      autocomplete: "off",
      label: "Calories\n",
      default: "0",
      floating: false,
    },
    protein: {
      type: "text",
      rules: ["required", "min:0", "integer"],
      autocomplete: "off",
      label: "Protein",
      columns: {
        container: 4,
      },
      default: "0",
    },
    carbs: {
      type: "text",
      rules: ["required", "min:0", "integer"],
      autocomplete: "off",
      label: "Net Carbs",
      columns: {
        container: 4,
      },
      default: "0",
    },
    fat: {
      type: "text",
      rules: ["required", "min:0", "integer"],
      autocomplete: "off",
      label: "Fat",
      columns: {
        container: 4,
      },
      default: "0",
    },
    divider_1: {
      type: "static",
      tag: "hr",
    },
    diet_header: {
      type: "static",
      tag: "h1",
      content: "What does your diet look like?",
    },
    divider_2: {
      type: "static",
      tag: "hr",
    },
    diet: {
      type: "radiogroup",
      rules: ["required"],
      items: [
        {
          value: "none",
          label: "No Diet",
        },
        {
          value: "lacto vegetarian",
          label: "Lacto Vegetarian",
        },
        {
          value: "ovo vegetarian",
          label: "Ovo Vegetarian",
        },
        {
          value: "paleo",
          label: "Paleo",
        },
        {
          value: "primal",
          label: "Primal",
        },
        {
          value: "pescetarian",
          label: "Pescetarian",
        },
        {
          value: "vegetarian",
          label: "Vegetarian",
        },
      ],
      default: "none",
    },
    divider_3: {
      type: "static",
      tag: "hr",
    },
    intolerance_header: {
      type: "static",
      tag: "h1",
      content: "Do you have any allergies?",
    },
    divider_5: {
      type: "static",
      tag: "hr",
    },
    intolerances: {
      type: "checkboxgroup",
      view: "blocks",
      items: [
        {
          value: "dairy",
          label: "Dairy",
          description: "",
        },
        {
          value: "egg",
          label: "Egg",
          description: "",
        },
        {
          value: "gluten",
          label: "Gluten",
          description: null,
        },
        {
          value: "peanut",
          label: "Peanut",
          description: null,
        },
        {
          value: "sesame",
          label: "Sesame",
          description: null,
        },
        {
          value: "seafood",
          label: "Seafood",
          description: null,
        },
        {
          value: "shellfish",
          label: "Shellfish",
          description: null,
        },
        {
          value: "soy",
          label: "Soy",
          description: null,
        },
        {
          value: "sulfite",
          label: "Sulfite",
          description: null,
        },
        {
          value: "tree nut",
          label: "Tree Nut",
          description: null,
        },
        {
          value: "wheat",
          label: "Wheat",
          description: null,
        },
      ],
    },
    divider_4: {
      type: "static",
      tag: "hr",
    },
  },
});
</script>

<template>
  <div class="flex h-dvh w-dvw flex-col items-center justify-center">
    <UCard class="mt-16 shadow-xl" :ui="{ body: 'p-8 sm:p-8' }">
      <Vueform v-bind="vueform" />
    </UCard>
  </div>
</template>
