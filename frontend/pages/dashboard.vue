<script lang="ts" setup>
definePageMeta({
  // middleware: ['auth-logged-in'],
  middleware: ["auth-server"],
})

import Sidebar from '@/components/Sidebar.vue';
import ProfileCard from '@/components/ProfileCard.vue';


</script>

<template>
  <Sidebar />
  <div class="min-h-screen flex flex-col items-center md:pl-16">
    <div class="p-4 md:p-8 space-y-8">
      <!-- Top Bar -->
      <div class="flex flex-col-reverse md:grid md:grid-cols-3 auto-rows-fr gap-4 max-w-5xl">
        <div class="col-span-2">
          <UserHistory />
        </div>
        <ProfileCard />
      </div>

      <!-- Recent Activity -->
      <RecentActivity />
    </div>
  </div>

</template>

<style scoped></style>
