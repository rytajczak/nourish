<script setup lang="ts">
const planner = usePlannerStore();
const days = ref<Date[]>([]);

onMounted(async () => {
  for (let i = 0; i < 7; i++) {
    const date = new Date(planner.weekStartDate);
    date.setDate(planner.weekStartDate.getDate() + i);
    days.value.push(date);
  }
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
        class="bg-elevated mx-3 hidden flex-1 cursor-pointer transition-all duration-150 hover:scale-105 lg:block"
        :class="{
          'bg-black text-white dark:bg-white dark:text-black':
            planner.selectedDate.getDate() == day.getDate(),
        }"
        :ui="{ body: 'p-3 sm:p-3 rounded-xl' }"
        @click="planner.selectedDate = day"
      >
        <div class="flex flex-col items-center justify-center">
          <span class="text-xl font-semibold">{{
            day.toLocaleDateString("en-US", { day: "numeric" })
          }}</span>
          <span
            :class="{
              'text-neutral': planner.selectedDate.getDate() == day.getDate(),
            }"
            >{{ day.toLocaleDateString("en-US", { weekday: "short" }) }}</span
          >
        </div>
      </UCard>
      <div class="flex items-center justify-center lg:hidden">
        <UButton variant="ghost" icon="lucide:arrow-left" />
        <span class="text-xl font-semibold">
          {{
            planner.selectedDate.toLocaleDateString("en-US", {
              weekday: "long",
            })
          }}
        </span>
        <UButton variant="ghost" icon="lucide:arrow-right" />
      </div>
    </div>
  </UCard>
</template>
