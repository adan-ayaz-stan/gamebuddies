<template>
  <div id="about" class="min-h-screen py-10 flex flex-col gap-8">
    <div class="mx-auto flex flex-row gap-6 text-lightRed">
      <svg width="72" height="72" viewBox="0 0 1024 1089.841" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
        class="text-lightRed">
        <path
          d="M520.258 1056.92H503.76c0-134.61-50.306-252.21-150.918-352.824S134.628 553.178 0 553.178V536.68c134.628 0 252.23-50.306 352.842-151.145C453.454 284.677 503.76 167.075 503.76 32.938h16.498c0 134.61 50.306 252.21 150.918 352.824S889.39 536.68 1024 536.68v16.498c-134.61 0-252.21 50.306-352.824 150.918S520.258 922.31 520.258 1056.92z" />
      </svg>

      <h1 class="tracking-wider">Inclusions</h1>

      <svg width="72" height="72" viewBox="0 0 1024 1089.841" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
        class="text-lightRed">
        <path
          d="M520.258 1056.92H503.76c0-134.61-50.306-252.21-150.918-352.824S134.628 553.178 0 553.178V536.68c134.628 0 252.23-50.306 352.842-151.145C453.454 284.677 503.76 167.075 503.76 32.938h16.498c0 134.61 50.306 252.21 150.918 352.824S889.39 536.68 1024 536.68v16.498c-134.61 0-252.21 50.306-352.824 150.918S520.258 922.31 520.258 1056.92z" />
      </svg>
    </div>

    <!-- Slider goes here -->
    <UCarousel ref="carouselRef" :items="items" :ui="{ item: 'basis-full snap-end' }" arrows>
      <template #default="{ item }">
        <div class="w-full h-full">
          <div class="flex items-center justify-center p-2">
            <div class="w-fit h-full border border-lightRed p-6">
              <img :src="item.img" alt="" class="w-full h-full object-cover aspect-video" />
              <h2 class="mt-4">{{ item.title }}</h2>
              <span>{{ item.description }}</span>
            </div>
          </div>
        </div>
      </template>

      <template #prev="{ onClick, disabled }">
        <button :disabled="disabled"
          :class="cn('absolute top-1/2 left-16 flex flex-row text-lightRed', disabled && 'opacity-50')"
          @click="() => { onClick(); updatePage(false); }">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-50 rotate-180">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-75 translate-x-0.5 rotate-180">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-90 translate-x-1 rotate-180">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-[2.5] translate-x-2 rotate-180">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
        </button>
      </template>

      <template #next="{ onClick, disabled }">
        <button :disabled="disabled"
          :class="cn('absolute top-1/2 right-8 flex flex-row text-lightRed', disabled && 'opacity-50')"
          @click="() => { onClick(); updatePage(true); }">
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" class="-mx-4 scale-[2.5]">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-90 translate-x-0.5">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-75 translate-x-1">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="-mx-4 scale-50 translate-x-2">
            <path fill="currentColor" d="M6.23 20.23L8 22l10-10L8 2L6.23 3.77L14.46 12z" />
          </svg>
        </button>
      </template>
    </UCarousel>

    <div class="relative px-12 -mb-4">
      <div
        class="absolute top-1/2 left-1/2 -translate-x-1/2 -translate-y-1/2 h-3 w-3 bg-gradient-to-br from-lightRed to-darkRed rotate-45" />
      <hr class="border border-lightRed" />
    </div>

    <UCarousel ref="headingCarouselRef" v-slot="{ item }" :items="items"
      :ui="{ item: 'basis-full md:basis-1/2 lg:basis-1/3 snap-end px-12 pointer-events-none' }" draggable="false"
      class="pointer-events-none">
      <h2 class="text-center w-full">
        {{ item.title }}
      </h2>
    </UCarousel>

  </div>
</template>

<script lang="ts" setup>
import { cn } from '~/lib/utils';

const carouselRef = ref();
const headingCarouselRef = ref();

const items: {
  img: string;
  title: string;
  description: string;
}[] = [
    {
      img: 'https://picsum.photos/600/800?random=1',
      title: 'Title 1',
      description: 'lorem ipsum',
    },
    {
      img: 'https://picsum.photos/600/800?random=2',
      title: 'Title 2',
      description: 'lorem ipsum',
    },
    {
      img: 'https://picsum.photos/600/800?random=3',
      title: 'Title 3',
      description: 'lorem ipsum',
    },
    {
      img: 'https://picsum.photos/600/800?random=4',
      title: 'Title 4',
      description: 'lorem ipsum',
    },
    {
      img: 'https://picsum.photos/600/800?random=5',
      title: 'Title 5',
      description: 'lorem ipsum',
    },
    {
      img: 'https://picsum.photos/600/800?random=6',
      title: 'Title 6',
      description: 'lorem ipsum',
    },
  ];

function updatePage(isSlidingNext: boolean = false) {
  headingCarouselRef.value.select(carouselRef.value.page + (isSlidingNext ? 1 : -1));
}

</script>

<style></style>