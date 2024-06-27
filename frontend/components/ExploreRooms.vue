<template>
  <div class="w-full space-y-4 md:w-2/3">
    <div class="flex flex-row justify-between items-center gap-6">
      <h2>Explore</h2>
      <!-- Separator -->
      <div class="flex-1 relative">
        <hr class="w-full border-lightRed border" />
        <!-- Squares -->
        <div class="w-2 h-2 bg-lightRed absolute top-1/2 left-0 -translate-y-1/2 rotate-45" />
        <div class="w-2 h-2 bg-lightRed absolute top-1/2 right-0 -translate-y-1/2 rotate-45" />
      </div>
      <!-- Button -->
      <UButton variant='secondary'>
        See More
      </UButton>
    </div>

    <!-- Rooms Display / Render -->
    <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <RoomCard v-for="room in rooms" :key="room.room.name" :room="room" />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { faker } from '@faker-js/faker';
import type { TRoom } from '~/types/room';

const rooms: TRoom[] = Array.from({ length: 6 }).map((_, i) => ({
  room: {
    name: faker.hacker.adjective() + ' ' + faker.hacker.noun(),
    karma_points: Number(faker.number.int({ min: 0, max: 10 }).toPrecision(2)),
    capacity: faker.number.int({ min: 1, max: 5 }),
    status: faker.helpers.arrayElement(['OPEN', 'CLOSED', 'FULL', 'PRIVATE'] as const),
  },
  game: {
    name: faker.hacker.verb() + ' ' + faker.hacker.noun(),
    banner: faker.image.urlPicsumPhotos({
      width: 400,
      height: 100,
    }),
  }
}));


</script>

<style></style>