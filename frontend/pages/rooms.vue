<template>
  <Sidebar />
  <div class="min-h-screen flex items-stretch md:pl-16">
    <div class="flex-1 p-4 md:p-8 space-y-8 border-x-2 border-lightRed">
      <JoinPrivate />
      <RecentRooms />
      <ExploreRooms />
    </div>
    <!-- The filter bar -->
    <RoomFilter :onToggle="(selectedGames: string[]) => selectedGames = selectedGames" />
  </div>

</template>

<script lang="ts" setup>

const selectedGames = ref<string[]>([]);
</script>

<style></style>