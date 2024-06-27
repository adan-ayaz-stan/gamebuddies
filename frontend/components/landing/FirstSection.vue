<template>
  <div class="min-h-screen flex flex-col justify-between pb-10">
    <!-- header -->
    <div class="w-full flex flex-row justify-between items-center gap-2 p-4">
      <div class="flex flex-row gap-4">
        <div class="h-12 w-12 aspect-square bg-darkGray border-2 border-white rounded-r-full rounded-b-full" />
        <div class="h-12 w-12 aspect-square bg-lightRed rounded-b-full" />
      </div>
      <!--  -->
      <div class="flex flex-row items-center gap-8">
        <UButton variant='link'>
          <h3 class="tracking-wider">Peek</h3>
        </UButton>

        <NuxtLink to="/api/login" external>
          <UButton variant='secondary'>
            Dive In
          </UButton>
        </NuxtLink>

        <UButton variant='link'>
          <h3 class="tracking-wider">About</h3>
        </UButton>
      </div>
      <!--  -->
      <div class="flex flex-row gap-4">
        <div class="h-12 w-12 aspect-square bg-lightRed rounded-b-full" />
        <div class="h-12 w-12 aspect-square bg-darkGray border-2 border-white rounded-l-full rounded-b-full" />
      </div>
      <!--  -->
    </div>

    <img :src="gamepadImage" alt="gamepad"
      class="w-40 h-40 absolute top-1/2 right-10 -translate-y-1/2 object-contain -rotate-45" />

    <!-- center -->
    <div class="w-full text-center relative">
      <h1 class="xl:text-9xl lg:text-8xl sm:text-5xl text-4xl font-bold">GameBuddies</h1>
      <p>A G-themed client</p>
    </div>

    <!-- footer -->
    <div class="relative flex flex-row gap-4 items-end justify-between">
      <!-- number of people on platform -->
      <div class="p-2 border-2 border-lightRed">
        <h2>2300 <sub>users</sub></h2>
      </div>

      <!-- Scroll down -->
      <div class="absolute bottom-4 left-1/2 -translate-x-1/2">
        <LandingScrollTo href="#about" />
      </div>

      <!-- social handles -->
      <div class="flex flex-col gap-2">
        <UButton variant='secondary'>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 20 20">
            <path fill="currentColor"
              d="M13.18 11.309c-.718 0-1.3.807-1.3 1.799c0 .994.582 1.801 1.3 1.801s1.3-.807 1.3-1.801c-.001-.992-.582-1.799-1.3-1.799m4.526-4.683c.149-.365.155-2.439-.635-4.426c0 0-1.811.199-4.551 2.08c-.575-.16-1.548-.238-2.519-.238c-.973 0-1.945.078-2.52.238C4.74 2.399 2.929 2.2 2.929 2.2c-.789 1.987-.781 4.061-.634 4.426C1.367 7.634.8 8.845.8 10.497c0 7.186 5.963 7.301 7.467 7.301l1.734.002l1.732-.002c1.506 0 7.467-.115 7.467-7.301c0-1.652-.566-2.863-1.494-3.871m-7.678 10.289h-.056c-3.771 0-6.709-.449-6.709-4.115c0-.879.31-1.693 1.047-2.369C5.537 9.304 7.615 9.9 9.972 9.9h.056c2.357 0 4.436-.596 5.664.531c.735.676 1.045 1.49 1.045 2.369c0 3.666-2.937 4.115-6.709 4.115m-3.207-5.606c-.718 0-1.3.807-1.3 1.799c0 .994.582 1.801 1.3 1.801c.719 0 1.301-.807 1.301-1.801c0-.992-.582-1.799-1.301-1.799" />
          </svg>
        </UButton>
        <UButton variant='secondary'>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="M18.205 2.25h3.308l-7.227 8.26l8.502 11.24H16.13l-5.214-6.817L4.95 21.75H1.64l7.73-8.835L1.215 2.25H8.04l4.713 6.231zm-1.161 17.52h1.833L7.045 4.126H5.078z" />
          </svg>
        </UButton>
        <UButton variant='secondary'>
          <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
            <path fill="currentColor"
              d="M19.27 5.33C17.94 4.71 16.5 4.26 15 4a.09.09 0 0 0-.07.03c-.18.33-.39.76-.53 1.09a16.09 16.09 0 0 0-4.8 0c-.14-.34-.35-.76-.54-1.09c-.01-.02-.04-.03-.07-.03c-1.5.26-2.93.71-4.27 1.33c-.01 0-.02.01-.03.02c-2.72 4.07-3.47 8.03-3.1 11.95c0 .02.01.04.03.05c1.8 1.32 3.53 2.12 5.24 2.65c.03.01.06 0 .07-.02c.4-.55.76-1.13 1.07-1.74c.02-.04 0-.08-.04-.09c-.57-.22-1.11-.48-1.64-.78c-.04-.02-.04-.08-.01-.11c.11-.08.22-.17.33-.25c.02-.02.05-.02.07-.01c3.44 1.57 7.15 1.57 10.55 0c.02-.01.05-.01.07.01c.11.09.22.17.33.26c.04.03.04.09-.01.11c-.52.31-1.07.56-1.64.78c-.04.01-.05.06-.04.09c.32.61.68 1.19 1.07 1.74c.03.01.06.02.09.01c1.72-.53 3.45-1.33 5.25-2.65c.02-.01.03-.03.03-.05c.44-4.53-.73-8.46-3.1-11.95c-.01-.01-.02-.02-.04-.02M8.52 14.91c-1.03 0-1.89-.95-1.89-2.12s.84-2.12 1.89-2.12c1.06 0 1.9.96 1.89 2.12c0 1.17-.84 2.12-1.89 2.12m6.97 0c-1.03 0-1.89-.95-1.89-2.12s.84-2.12 1.89-2.12c1.06 0 1.9.96 1.89 2.12c0 1.17-.83 2.12-1.89 2.12" />
          </svg>
        </UButton>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import gamepadImage from '~/public/asset/gamepad.png';

</script>

<style></style>