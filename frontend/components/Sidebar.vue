<script lang="ts" setup>
import { faker } from '@faker-js/faker';

const dummydata: { name: string; email: string; icon: string }[] = [];
for (let i = 0; i < 3; i++) {
  const user = {
    name: faker.name.prefix(),
    email: faker.internet.email(),
    icon: faker.image.avatar(),
  };
  dummydata.push(user);
}
</script>

<template>
  <aside
    class="hidden fixed inset-0 z-50 w-16 h-screen overflow-y-auto overflow-x-hidden bg-fadeRed md:flex flex-col scrollbar-none py-4 px-2">
    <div class="space-y-3 mb-auto">
      <NuxtLink to="/" class="flex text-lightRed hover:text-darkRed">
        <HomeIcon class="w-6 h-6 m-auto" />
      </NuxtLink>
    </div>

    <hr class="w-3/4 mx-auto border border-lightRed my-3" />
    <div class="flex flex-col items-center justify-center mx-auto gap-3">
      <template v-for="user in dummydata" :key="user.id">
        <Avatar :src="user.icon" />
      </template>
      <SeeMoreButton href="/" />
    </div>
    <hr class="w-3/4 mx-auto border border-lightRed my-3" />
    <div class="flex flex-col items-center justify-center mx-auto gap-3">
      <template v-for="user in dummydata" :key="user.id">
        <Avatar :src="user.icon" />
      </template>
      <SeeMoreButton href="/" />
    </div>
    <hr class="w-3/4 mx-auto border border-lightRed my-3" />

    <div class="space-y-3 mt-auto">
      <NuxtLink to="/" class="flex text-lightRed hover:text-darkRed">
        <User2 class="w-6 h-6 m-auto" />
      </NuxtLink>
      <NuxtLink to="/" class="flex text-lightRed hover:text-darkRed">
        <LogOutIcon class="w-6 h-6 m-auto" />
      </NuxtLink>
    </div>
  </aside>
</template>

<style scoped></style>
