<script setup lang="ts">
definePageMeta({
  middleware: "auth",
});
const { loggedIn, user, session, fetch, clear } = useUserSession();
</script>

<template>
  <ClientOnly>
    <UContainer class="flex">
      <UCard>
        <template #header>Monday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Tuesday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Wednesday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Thursday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Friday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Saturday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
      <UCard>
        <template #header>Sunday</template>
        put meal stuff here
        <template #footer>
          <UButton color="neutral">Edit</UButton>
        </template>
      </UCard>
    </UContainer>
  </ClientOnly>
</template>
