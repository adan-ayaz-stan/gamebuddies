<template>
  <div class="min-h-screen p-4 grid grid-cols-2 auto-rows-fr gap-2">
    <!-- Left side -->
    <div class="flex flex-col gap-4">
      <!-- Globe -->
      <GlobalGlobe />

      <!-- Global Activity -->
      <GlobalMessageActivity />
    </div>

    <!-- Right side -->
    <div class="flex flex-col border border-fadeRed">
      <!-- Window of messages -->
      <div class="flex-1 bg-fadeRed">

      </div>

      <!-- Message form -->
      <GlobalMessageChatForm />
    </div>
  </div>
</template>

<script lang="ts" setup>

</script>

<style></style>