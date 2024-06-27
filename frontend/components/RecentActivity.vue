    
<script lang="ts" setup>
import { faker } from '@faker-js/faker';
import type { RoomStatus, TRoom } from '~/types/room';


const rooms: TRoom[] = [];

for (let i = 0; i < 3; i++) {
  const statusEnums = ['OPEN', 'CLOSED', 'FULL', 'PRIVATE']

  const room: TRoom = {
    room: {
      name: faker.hacker.adjective() + ' ' + faker.hacker.noun(),
      capacity: faker.number.int({ min: 1, max: 5 }),
      status: statusEnums[faker.number.int({ min: 0, max: 3 })] as RoomStatus,
      karma_points: Number(faker.number.int({ min: 0, max: 10 }).toPrecision(2)),
    },
    game: {
      name: faker.hacker.verb() + ' ' + faker.hacker.noun(),
      banner: faker.image.urlPicsumPhotos({
        width: 400,
        height: 100,
      }),
    }
  }

  rooms.push(room);
}

</script>

<template>
  <h1 class="w-full">Recent Activity</h1>
  <div class="py-4 flex flex-row gap-4">
    <div class="w-full grid sm:grid-cols-2 xl:grid-cols-3 auto-rows-fr gap-x-4 gap-y-8">
      <RoomCard v-for="room in rooms" :key="room.room.name" :room="room" />
    </div>

  </div>
</template>

<style></style>