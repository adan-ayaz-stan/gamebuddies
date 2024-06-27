<template>
  <aside style="width: 440px;" class="fixed w-2/3 right-0 z-50 h-screen overflow-y-hidden p-4">
    <hr class="border-2 border-lightRed" />
    <div class="p-4 w-full max-h-full overflow-y-auto bg-fadeRed md:flex flex-col h-full">
      <!--  -->
      <h2 class="mx-auto">Filter</h2>
      <!--  -->
      <h3 class="mx-auto mt-2">== Games ==</h3>

      <div class="space-y-4 my-4">
        <!--  -->
        <div>
          <h5 class="bg-lightRed w-fit px-1">Selected</h5>
          <div class="p-2 border-lightRed bg-fadeRed border-2 space-y-2">
            <!-- If no games are selected -->
            <div v-if="selectedGames.length === 0" class="text-center">No games selected</div>

            <!-- If games are present -->
            <div v-for="game in gamesList.filter(game => selectedGames.includes(game.id))" :key="game.title"
              @click="toggleGame(game)"
              :class="cn('w-full h-20 flex items-center justify-center relative border-2 border-lightRed cursor-pointer')">
              <img :src="game.banner" :alt="game.title"
                :class="cn('w-full h-full object-cover', selectedGames.includes(game.id) && 'opacity-70')" />

              <!-- Badge -->
              <div class="absolute bottom-0 right-0 flex flex-row gap-1 items-center px-1 bg-lightRed text-white">
                <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
                  <path fill="currentColor"
                    d="m10.6 15.508l6.396-6.396l-.708-.708l-5.688 5.688l-2.85-2.85l-.708.708zM4 20V4h16v16z" />
                </svg>
                <span class="text-sm">{{ game.title }}</span>
              </div>
            </div>
          </div>
        </div>
        <div v-for="game in gamesList.filter(game => !selectedGames.includes(game.id))" :key="game.title"
          @click="toggleGame(game)"
          :class="cn('w-full h-20 flex items-center justify-center relative border-2 border-lightRed cursor-pointer', selectedGames.includes(game.id) && 'opacity-70')">
          <img :src="game.banner" :alt="game.title" class="w-full h-full object-cover" />
        </div>
      </div>

    </div>
  </aside>
</template>

<script lang="ts" setup>

import { games, type Game } from '~/lib/constants/games';
import { cn } from '~/lib/utils';

const { onToggle } = defineProps<{
  onToggle: (games: string[]) => void;
}>();

const gamesList = games;

const selectedGames = ref<string[]>([]);

const unselectedGames = computed(() => gamesList.filter(game => !selectedGames.value.includes(game.id)));

function toggleGame(game: Game) {
  if (!selectedGames.value.includes(game.id)) {
    selectedGames.value.push(game.id);
  } else {
    selectedGames.value = selectedGames.value.filter(g => g !== game.id);
  }
  onToggle(selectedGames.value);
}
</script>

<style></style>