<script setup lang="ts">
const { user } = useUserSession();
const { profile, open } = useUserStore();
</script>

<template>
  <UModal v-model:open="open">
    <template #content>
      <UCard>
        <template #header>
          <div class="flex">
            <div><img :src="user?.picture" alt="" class="rounded-full" /></div>
            <div class="mt-4 ml-8">
              <h2 class="text-lg font-bold">{{ user?.firstName }}</h2>
              <span>{{ user?.email }}</span>
              <p class="mt-4">
                <span class="font-bold">Diet:</span>
                {{ profile.diet }}
              </p>
              <UButton color="error" variant="subtle" @click="closeModel">
                Close
              </UButton>
            </div>
          </div>
        </template>
      </UCard>
    </template>
  </UModal>
</template>
