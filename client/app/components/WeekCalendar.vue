<script setup lang="ts">
const planner = usePlannerStore();

const selectedDate = useState<Date>("selectedDate", () => new Date());
const days = ref<Date[]>([]);

onMounted(async () => {
  for (let i = 0; i < 7; i++) {
    const nextDate = new Date(selectedDate.value);
    nextDate.setDate(selectedDate.value.getDate() + i);
    days.value.push(nextDate);
  }
  await planner.fetchWeek(new Date());
});

watch(selectedDate, async () => {
  await planner.fetchWeek(new Date());
});
</script>

<template>
  <UCard :ui="{ body: 'px-3 sm:px-3' }">
    <template #header>
      <div class="flex items-center justify-center">
        <span class="text-xl font-semibold"
          >{{
            days[0]?.toLocaleDateString("en-US", {
              month: "long",
              day: "numeric",
            })
          }}
          -
          {{
            days[6]?.toLocaleDateString("en-US", {
              month: "long",
              day: "numeric",
            })
          }}</span
        >
      </div>
    </template>
    <div class="flex items-center justify-center">
      <UCard
        v-for="day in days"
        class="mx-3 hidden flex-1 cursor-pointer transition-all duration-150 hover:scale-105 lg:block"
        :class="{
          'bg-foreground/90 text-background':
            selectedDate.getDate() == day.getDate(),
        }"
        :ui="{ body: 'p-3 sm:p-3' }"
        @click="selectedDate = day"
      >
        <div class="flex flex-col items-center justify-center">
          <span class="text-xl font-semibold">{{
            day.toLocaleDateString("en-US", { day: "numeric" })
          }}</span>
          <span
            :class="{
              'text-neutral': selectedDate.getDate() == day.getDate(),
            }"
            >{{ day.toLocaleDateString("en-US", { weekday: "short" }) }}</span
          >
        </div>
      </UCard>
      <div class="flex items-center justify-center lg:hidden">
        <UButton variant="ghost" icon="i-heroicons-chevron-left" />
        <span class="text-xl font-semibold">
          {{ selectedDate.toLocaleDateString("en-US", { weekday: "long" }) }}
        </span>
        <UButton variant="ghost" icon="i-heroicons-chevron-right" />
      </div>
    </div>
  </UCard>
</template>
