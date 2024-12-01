<script setup lang="ts">
import type { FormSubmitEvent } from "@nuxt/ui";
import { z } from "zod";

const onboarding = useOnboardingStore();

const schema = z.object({
  calories: z.number().int(),
  protein: z.number().int(),
  carbs: z.number().int(),
  fat: z.number().int(),
});

type Schema = z.output<typeof schema>;

const state = reactive<Partial<Schema>>({
  calories: onboarding.data.profile.calories,
  protein: onboarding.data.profile.protein,
  carbs: onboarding.data.profile.carbs,
  fat: onboarding.data.profile.fat,
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  onboarding.setGoals(event.data);
  onboarding.nextStep();
}
</script>

<template>
  <UForm :schema="schema" :state="state" @submit="onSubmit">
    <span class="text-2xl">
      <UIcon name="lucide:target" />
    </span>
    <div class="mt-2 mb-6">
      <h2 class="text-2xl font-semibold">Set Your Daily Nutrition Goals</h2>
      <h2 class="text-muted mt-1">Add your calorie and macro targets.</h2>
    </div>
    <UFormField label="Calories" name="calories">
      <UInput
        color="secondary"
        size="xl"
        v-model="state.calories"
        type="number"
        class="w-full"
      />
    </UFormField>
    <div class="mt-4 grid grid-cols-3 gap-4">
      <UFormField label="Protein" name="protein">
        <UInput
          color="secondary"
          size="xl"
          v-model="state.protein"
          type="number"
        />
      </UFormField>
      <UFormField label="Carbs" name="carbs">
        <UInput
          color="secondary"
          size="xl"
          v-model="state.carbs"
          type="number"
        />
      </UFormField>
      <UFormField label="Fat" name="fat">
        <UInput color="secondary" size="xl" v-model="state.fat" type="number" />
      </UFormField>
    </div>
    <div class="mt-6 grid grid-cols-2 gap-6">
      <UButton
        variant="outline"
        color="neutral"
        class="flex items-center justify-center py-3 font-bold"
        @click="onboarding.previousStep"
      >
        <span>Go Back</span>
      </UButton>
      <UButton
        color="neutral"
        class="flex items-center justify-center py-3 font-bold"
        type="submit"
      >
        <span>Next Step: Diet</span>
      </UButton>
    </div>
  </UForm>
</template>
