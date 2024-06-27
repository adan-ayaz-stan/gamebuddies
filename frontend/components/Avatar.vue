<script lang="ts" setup>
import CornerBlock from '~/assets/svgs/stub.svg';

const props = defineProps({
  src: {
    type: String,
    default: "",
  },
});


</script>

<template>
  <div class="group relative w-10 h-10 aspect-square p-1 bg-black border border-slate-800 cursor-pointer">
    <!-- Avatar -->
    <img :src="props.src" alt="avatar"
      class="w-full h-full object-cover p-1 border border-slate-800 group-hover:bg-white transition-colors duration-300" />

    <!-- Corner Blocks -->
    <img :src="CornerBlock" alt="corner-block"
      class="group-hover:-translate-x-1 group-hover:-translate-y-1 absolute bottom-0 right-0 w-5 bg-transparent -m-1.5 transition-all duration-300 ease-out" />
    <img :src="CornerBlock" alt="corner-block"
      class="group-hover:translate-x-1 group-hover:translate-y-1 absolute top-0 left-0 rotate-180 w-5 bg-transparent -m-1.5 transition-all duration-300 ease-out" />

    <img :src="CornerBlock" alt="corner-block"
      class="group-hover:-translate-x-1 group-hover:translate-y-1 absolute top-0 right-0 -rotate-90 w-5 bg-transparent -m-1.5 transition-all duration-300 ease-out" />
    <img :src="CornerBlock" alt="corner-block"
      class="group-hover:translate-x-1 group-hover:-translate-y-1 absolute bottom-0 left-0 rotate-90 w-5 bg-transparent -m-1.5 transition-all duration-300 ease-out" />
  </div>
</template>

<style scoped></style>
