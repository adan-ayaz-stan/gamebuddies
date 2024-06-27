<script lang="ts" setup>
const props = defineProps({
  href: {
    type: String,
    default: "",
  },
})
</script>

<template>
  <div class="group relative h-10 w-10">
    <a href="{{ href }}"
      class="absolute origin-bottom group-hover:-rotate-12 inset-0 aspect-square h-full w-full text-center flex items-center justify-center bg-medRed text-black transition-all duration-300"
      style="clip-path: polygon(50% 0%, 100% 50%, 50% 100%, 0% 50%);">
      <Plus class="w-6 h-6" />
    </a>
    <a href="{{ href }}" id="mid-card-add-more"
      class="absolute origin-bottom group-hover:-rotate-6 inset-0 aspect-square h-full w-full text-center flex items-center justify-center bg-white text-black transition-all duration-300"
      style="clip-path: polygon(50% 0%, 100% 50%, 50% 100%, 0% 50%);">
      <Plus class="w-6 h-6" />
    </a>
    <a href="{{ href }}"
      class="aspect-square h-full w-full origin-bottom group-hover:rotate-6 text-center flex items-center justify-center bg-lightRed text-black transition-all duration-300"
      style="clip-path: polygon(50% 0%, 100% 50%, 50% 100%, 0% 50%);">
      <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24">
        <path fill="white" d="M11 16h2v-3h3v-2h-3V8h-2v3H8v2h3zm1.025 6.85L1.225 12l10.8-10.825L22.775 12z" />
      </svg>
    </a>
  </div>

</template>

<style scoped></style>
