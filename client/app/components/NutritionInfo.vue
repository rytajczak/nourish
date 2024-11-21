<script setup lang="ts">
const showing = ref("selected day");
</script>

<template>
  <div class="flex items-center justify-between">
    <h2 class="my-4 text-xl font-semibold">Nutrition info</h2>
    <USelect
      v-model="showing"
      color="neutral"
      :items="['selected day', 'this week']"
      class="w-36"
    />
  </div>
  <div class="grid grid-cols-3 gap-4">
    <UCard class="col-span-3">
      <div class="flex flex-col">
        <span class="text-sm text-[var(--ui-text-muted)]"
          >Total calories
          {{ showing === "selected day" ? "selected day" : "this week" }}</span
        >
        <span class="text-xl font-semibold">2000 kcal</span>
      </div>
    </UCard>
    <UCard>
      <div class="flex flex-col">
        <span class="text-muted text-sm">Protein (g)</span>
        <span class="text-xl font-semibold">180 g</span>
      </div>
    </UCard>
    <UCard>
      <div class="flex flex-col">
        <span class="text-muted text-sm">Carbs (g)</span>
        <span class="text-xl font-semibold">180 g</span>
      </div>
    </UCard>
    <UCard>
      <div class="flex flex-col">
        <span class="text-muted text-sm">Fat (g)</span>
        <span class="text-xl font-semibold">180 g</span>
      </div>
    </UCard>
  </div>
</template>
