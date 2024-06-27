<template>
  <div class="w-full space-y-4 md:w-2/3">
    <div class="flex flex-row justify-between items-center gap-6">
      <h2>Ride Private</h2>
      <!-- Separator -->
      <div class="flex-1 relative">
        <hr class="w-full border-lightRed border" />
        <!-- Squares -->
        <div class="w-2 h-2 bg-lightRed absolute top-1/2 left-0 -translate-y-1/2 rotate-45" />
        <div class="w-2 h-2 bg-lightRed absolute top-1/2 right-0 -translate-y-1/2 rotate-45" />
      </div>
      <!-- Button -->
      <UButton variant='secondary'>
        Join by Code
      </UButton>
    </div>

    <div class="w-full max-w-full overflow-x-auto flex flex-row gap-4 pb-1">
      <div v-for="friend in friendsOnline" :key="friend.name" class="h-fit bg-darkGray p-2">
        <Avatar :src="friend.icon" />
      </div>
      <div class="h-fit bg-darkGray p-2">
        <AddFriend />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faker } from '@faker-js/faker';

const friendsOnline: { name: string, icon: string }[] = Array.from({ length: 6 }).map((_, i) => ({
  name: faker.name.fullName(),
  icon: faker.image.avatar(),
}));

</script>

<style></style>