<script setup lang="ts">
const props = defineProps<{
  name: string;
  icon: string;
}>();

const iconOff = props.icon + "-off";

const [value, toggle] = useToggle(false);

const emit = defineEmits<{
  (event: "toggle", name: string, value: boolean): void;
}>();

function toggleAndEmit() {
  toggle();
  emit("toggle", props.name, value.value);
}
</script>

<template>
  <UButton
    color="neutral"
    variant="outline"
    class="border-border flex h-12 items-center justify-center rounded-[calc(var(--ui-radius)*2)] border"
    @click="toggleAndEmit"
  >
    <span class="text-lg">
      <UIcon
        :name="value ? iconOff : icon"
        :class="value ? 'text-red-500' : 'text-foreground'"
      />
      {{ name }}
    </span>
  </UButton>
</template>
