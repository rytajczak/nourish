<template>
  <Sidebar />
  <div class="md:pl-64">
    <div class="flex h-20 items-center justify-between border-b-[1px] px-8">
      <UInput
        color="neutral"
        size="xl"
        icon="i-heroicons-magnifying-glass"
        placeholder="What would you like to cook?"
      />
      <div class="flex items-center">
        <div class="border-r pr-2">
          <UButton
            color="neutral"
            icon="solar:bell-outline"
            size="lg"
            variant="link"
          ></UButton>
          <UButton
            color="neutral"
            icon="solar:settings-outline"
            size="lg"
            variant="link"
          ></UButton>
        </div>
        <UAvatar
          size="lg"
          src="https://github.com/benjamincanac.png"
          class="ms-4"
        ></UAvatar>
      </div>
    </div>
    <slot></slot>
  </div>
</template>

<style scoped>
* {
  border-color: #242424;
}
</style>
