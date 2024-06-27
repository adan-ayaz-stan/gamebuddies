<template>
  <div @click="isModalOpen = true"
    class="group relative w-10 h-10 aspect-square p-1 bg-black border border-slate-800 cursor-pointer text-white flex">
    <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" class="m-auto">
      <path fill="currentColor" d="M6 17v-3H3v-2h3V9h2v3h3v2H8v3zm9.75 2V8.05l-2.3 1.65l-1.15-1.75L16.4 5H18v14z" />
    </svg>
  </div>

  <UModal v-model="isModalOpen">
    <div class="p-4">
      Add friend form is shown here
    </div>
  </UModal>
</template>

<script lang="ts" setup>
const isModalOpen = ref(false);

</script>

<style></style>