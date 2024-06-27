<script lang="ts" setup>
import { cn } from '~/lib/utils';
import type { TRoom } from '~/types/room';
import { games as gamesList } from '~/lib/constants/games';

const { room } = defineProps<{
  room: TRoom
}>()

</script>

<template>
  <div class="flex flex-col">
    <img :src="gamesList[2].banner" :alt="room.game.name" class="aspect-video object-cover" />

    <div class="p-2 text-white bg-darkGray">
      <p>{{ room.game.name }}</p>
      <span class="text-sm -left-2 relative bg-black text-white p-3">{{ room.room.karma_points }} karma points</span>
      <div class="mt-4 flex flex-row items-end justify-between">
        <!-- Members Component -->
        <span class="text-sm">Members</span>
        <!--  -->
        <UButton>Join</UButton>
      </div>
    </div>

  </div>
</template>

<style></style>