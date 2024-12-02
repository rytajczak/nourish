<script setup lang="ts">
const { signOut } = useUserStore();
const { user } = useUserSession();
const { profile } = useUserStore();
const open = ref(false);
</script>

<template>
  <aside
    class="border-border bg-elevated fixed top-0 left-0 z-50 block h-screen w-64 -translate-x-full border-r-[1px] transition-transform md:translate-x-0"
  >
    <div class="flex flex-col">
      <div class="h-20">
        <NuxtLink
          to="/dashboard"
          class="flex items-center px-6 py-[24px] text-2xl font-bold"
        >
          <Icon name="lucide:leaf" class="text-primary mr-2" />
          Nourish
        </NuxtLink>
      </div>
      <div class="flex flex-col px-6 py-6">
        <span class="mb-2 text-sm text-gray-700 dark:text-gray-400">Menu</span>
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/dashboard"
          icon="lucide:layout-dashboard"
        >
          Dashboard
        </UButton>
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/recipes/saved"
          icon="lucide:book-marked"
        >
          Saved Recipes
        </UButton>
        <UButton
          size="lg"
          color="neutral"
          variant="ghost"
          to="/grocery-list"
          icon="lucide:shopping-basket"
        >
          Grocery List
        </UButton>
        <span class="mt-8 mb-2 text-sm text-gray-700 dark:text-gray-400">
          Settings
        </span>

        <UModal v-model:open="open">
          <UButton
            size="lg"
            color="neutral"
            variant="ghost"
            icon="lucide:circle-user-round"
            class="hover:cursor-pointer"
          >
            Profile
          </UButton>
          <template #content>
            <UCard>
              <template #header>
                <div class="flex">
                  <div>
                    <img :src="user?.picture" alt="" class="rounded-full" />
                  </div>
                  <div class="mt-4 ml-8">
                    <h2 class="text-lg font-bold">{{ user?.firstName }}</h2>
                    <span>{{ user?.email }}</span>
                    <p class="mt-4">
                      <span class="font-bold">Diet:</span>
                      {{ profile.diet }}
                    </p>
                  </div>
                </div>
              </template>
            </UCard>
          </template>
        </UModal>
        <UButton
          size="lg"
          color="error"
          variant="ghost"
          icon="lucide:log-out"
          @click="signOut"
        >
          Log Out
        </UButton>
      </div>
    </div>
  </aside>
</template>
