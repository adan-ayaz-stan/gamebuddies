<script lang="ts" setup>
import { faker } from '@faker-js/faker';

const profileData = {
  hours: faker.number.int({ min: 0, max: 100 }),
  rooms: faker.number.int({ min: 0, max: 100 }),
  karma: faker.number.int({ min: 100, max: 500 }),
  games: faker.number.int({ min: 0, max: 23 }),
}

</script>

<template>

  <div class="w-full h-full grid grid-cols-1 sm:grid-cols-2 auto-rows-fr gap-4 justify-evenly xl:hidden">
    <div class="group flex flex-row gap-2 p-4 border-lightRed border-2 rounded-lg hover:bg-medRed hover:text-black">
      <h1 class="group-hover:text-black">{{ profileData.karma }}</h1> karma
    </div>

    <div class="group flex flex-row gap-2 p-4 border-lightRed border-2 rounded-lg hover:bg-medRed hover:text-black">
      <h1 class="group-hover:text-black">{{ profileData.games }}</h1> games
    </div>

    <div class="group flex flex-row gap-2 p-4 border-lightRed border-2 rounded-lg hover:bg-medRed hover:text-black">
      <h1 class="group-hover:text-black">{{ profileData.rooms }}</h1> rooms
    </div>

    <div class="group flex flex-row gap-2 p-4 border-lightRed border-2 rounded-lg hover:bg-medRed hover:text-black">
      <h1 class="group-hover:text-black">{{ profileData.hours }}</h1> hours
    </div>

  </div>

  <div class="hidden xl:block w-full h-full relative">
    <!-- Box Up -->
    <div
      class="aspect-square w-[25%] rotate-45 absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-full bg-lightRed/80 rounded-2xl">
      <div class="-rotate-45 text-center absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
        <h2 class="text-white">{{ profileData.hours }}</h2>
        <h4>hours</h4>
      </div>
    </div>
    <!-- Box Down -->
    <div class="aspect-square w-[25%] rotate-45 absolute top-1/2 left-1/2 -translate-x-1/2 bg-lightRed/80 rounded-2xl">
      <div class="-rotate-45 text-center absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
        <h2 class="text-white">{{ profileData.rooms }}</h2>
        <h4>rooms</h4>
      </div>
    </div>

    <!-- X Axis -->
    <!-- Box Right -->
    <div
      class="aspect-square w-[25%] rotate-45 absolute top-1/2 left-1/2 translate-x-1/2 -translate-y-1/2 bg-lightRed/80 rounded-2xl">
      <div class="-rotate-45 text-center absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
        <h2 class="text-white">{{ profileData.games }}</h2>
        <h4>games</h4>
      </div>
    </div>
    <!-- Box Left -->
    <div
      class="aspect-square w-[25%] rotate-45 absolute top-1/2 left-1/2 -translate-x-[150%] -translate-y-1/2 bg-lightRed/80 rounded-2xl">
      <div class="-rotate-45 text-center absolute top-1/2 left-1/2 -translate-y-1/2 -translate-x-1/2">
        <h2 class="text-white">{{ profileData.karma }}</h2>
        <h4>karma</h4>
      </div>
    </div>

  </div>
</template>

<style></style>