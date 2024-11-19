<script setup lang="ts">
const planner = usePlannerStore();

const selectedDate = useState<Date>("selectedDate", () => new Date());
const days = ref<Date[]>([]);

onMounted(() => {
  for (let i = 0; i < 7; i++) {
    const nextDate = new Date(selectedDate.value);
    nextDate.setDate(selectedDate.value.getDate() + i);
    days.value.push(nextDate);
  }
});
</script>

<template>
  <UCard :ui="{ body: 'p-0 sm:px-4 sm:py-0' }">
    <template #header>
      {{ Math.floor(selectedDate.getTime() / 1000) }}
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
        class="m-4 hidden flex-1 cursor-pointer transition-all duration-300 hover:scale-105 xl:block"
        :ui="{ body: 'p-4 sm:p-4' }"
        :class="{
          'bg-[#14ae4e] text-white': selectedDate.getDate() == day.getDate(),
        }"
        @click="selectedDate = day"
      >
        <div class="flex flex-col items-center justify-center">
          <span class="text-xl font-semibold">{{
            day.toLocaleDateString("en-US", { day: "numeric" })
          }}</span>
          {{ day.toLocaleDateString("en-US", { weekday: "short" }) }}
        </div>
      </UCard>
      <div class="flex items-center justify-center xl:hidden">
        <UButton variant="ghost" icon="i-heroicons-chevron-left" />
        <span class="text-xl font-semibold">
          {{ selectedDate.toLocaleDateString("en-US", { day: "numeric" }) }}
        </span>
        <UButton variant="ghost" icon="i-heroicons-chevron-right" />
      </div>
    </div>
  </UCard>
</template>
