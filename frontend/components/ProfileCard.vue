<script lang="ts" setup>
import { faker } from '@faker-js/faker/locale/en_US';
import UploadAvatar from './UploadAvatar.vue';
const user = {
  icon: faker.image.url(),
  name: faker.name.fullName(),
  description: faker.lorem.lines(2),
};

const isModalOpen = ref(false);
</script>

<template>
  <div class="w-full max-w-sm border border-lightRed">
    <img
      src="https://i0.wp.com/picjumbo.com/wp-content/uploads/dark-abstract-background-wallpaper-black-and-orange-free-photo.jpg?w=600&quality=80"
      alt="profile-background" class="aspect-video object-cover border-b border-lightRed" />
    <div class="aspect-video p-4 flex flex-col justify-end items-end relative bg-gradient-to-tr from-medRed to-darkRed">
      <!-- Avatar -->
      <div @click="isModalOpen = true"
        class="absolute w-20 h-20 top-0 right-4 -translate-y-1/2 group rotate-45 border-2 border-lightRed bg-lightRed overflow-hidden">
        <img :src="user.icon" alt="avatar" class="absolute inset-0 w-full h-full scale-150 object-cover -rotate-45" />
      </div>

      <UModal v-model="isModalOpen">
        <UploadAvatar />
      </UModal>

      <!-- Text Profile -->
      <h3 class="text-white pt-10">{{ user.name }}</h3>
      <i class="leading-none text-sm">{{ user.description }}</i>
    </div>
  </div>
</template>

<style scoped></style>
