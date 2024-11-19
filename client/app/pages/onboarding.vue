<script setup>
definePageMeta({
  middleware: "auth",
  layout: false,
});

async function handleResponse() {
  navigateTo("/confirm");
}
</script>

<template>
  <div class="flex h-screen items-center justify-center">
    <UCard>
      <Vueform
        size="md"
        :display-errors="false"
        validate-on="step|change"
        endpoint="/api/users/signup"
        method="post"
        @response="handleResponse"
      >
        <template #empty>
          <FormSteps>
            <FormStep
              name="daily_goals"
              label="Daily Goals"
              :elements="[
                'goal_header',
                'divider',
                'calories',
                'protein',
                'carbs',
                'fat',
                'divider_1',
              ]"
            />
            <FormStep
              name="diet"
              label="Diet"
              :elements="[
                'checkbox',
                'diet_header',
                'divider_2',
                'diet',
                'divider_3',
              ]"
            />
            <FormStep
              name="intolerances"
              label="Intolerances"
              :elements="[
                'intolerance_header',
                'divider_5',
                'intolerances',
                'divider_4',
                'submit',
              ]"
            />
          </FormSteps>
          <FormElements>
            <StaticElement
              name="goal_header"
              tag="h1"
              content="What is your daily nutrition goal?"
            />
            <StaticElement name="divider" tag="hr" />
            <TextElement
              name="calories"
              :rules="['required', 'min:0', 'integer']"
              autocomplete="off"
              label="Calories
      "
              default="0"
              :floating="false"
            />
            <TextElement
              name="protein"
              :rules="['required', 'min:0', 'integer']"
              autocomplete="off"
              label="Protein"
              :columns="{
                container: 4,
              }"
              default="0"
            />
            <TextElement
              name="carbs"
              :rules="['required', 'min:0', 'integer']"
              autocomplete="off"
              label="Net Carbs"
              :columns="{
                container: 4,
              }"
              default="0"
            />
            <TextElement
              name="fat"
              :rules="['required', 'min:0', 'integer']"
              autocomplete="off"
              label="Fat"
              :columns="{
                container: 4,
              }"
              default="0"
            />
            <StaticElement name="divider_1" tag="hr" />
            <StaticElement
              name="diet_header"
              tag="h1"
              content="What does your diet look like?"
            />
            <StaticElement name="divider_2" tag="hr" />
            <RadiogroupElement
              name="diet"
              :rules="['required']"
              :items="[
                {
                  value: 'none',
                  label: 'No Diet',
                },
                {
                  value: 'lacto vegetarian',
                  label: 'Lacto Vegetarian',
                },
                {
                  value: 'ovo vegetarian',
                  label: 'Ovo Vegetarian',
                },
                {
                  value: 'paleo',
                  label: 'Paleo',
                },
                {
                  value: 'primal',
                  label: 'Primal',
                },
                {
                  value: 'pescetarian',
                  label: 'Pescetarian',
                },
                {
                  value: 'vegetarian',
                  label: 'Vegetarian',
                },
              ]"
              default="none"
            />
            <StaticElement name="divider_3" tag="hr" />
            <StaticElement
              name="intolerance_header"
              tag="h1"
              content="Do you have any allergies?"
            />
            <StaticElement name="divider_5" tag="hr" />
            <CheckboxgroupElement
              name="intolerances"
              view="blocks"
              :items="[
                {
                  value: 'dairy',
                  label: 'Dairy',
                  description: '',
                },
                {
                  value: 'egg',
                  label: 'Egg',
                  description: '',
                },
                {
                  value: 'gluten',
                  label: 'Gluten',
                  description: null,
                },
                {
                  value: 'peanut',
                  label: 'Peanut',
                  description: null,
                },
                {
                  value: 'sesame',
                  label: 'Sesame',
                  description: null,
                },
                {
                  value: 'seafood',
                  label: 'Seafood',
                  description: null,
                },
                {
                  value: 'shellfish',
                  label: 'Shellfish',
                  description: null,
                },
                {
                  value: 'soy',
                  label: 'Soy',
                  description: null,
                },
                {
                  value: 'sulfite',
                  label: 'Sulfite',
                  description: null,
                },
                {
                  value: 'tree nut',
                  label: 'Tree Nut',
                  description: null,
                },
                {
                  value: 'wheat',
                  label: 'Wheat',
                  description: null,
                },
              ]"
            />
            <StaticElement name="divider_4" tag="hr" />
          </FormElements>
          <FormStepsControls />
        </template>
      </Vueform>
    </UCard>
  </div>
</template>
