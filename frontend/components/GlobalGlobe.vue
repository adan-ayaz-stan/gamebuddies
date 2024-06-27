<template>
  <div class="relative flex-1 flex items-center justify-center border border-fadeRed">
    <TresCanvas clear-color="#2b2d42" preset="realistic">
      <TresPerspectiveCamera :position="[0, 5, 5]" :look-at="[0, 0, 0]" />
      <TresMesh>
        <TresTorusGeometry :args="[1, 0.2, 16, 32]" />
        <TresMeshBasicMaterial :color="'red'" />
      </TresMesh>
      <GlobalGlobeMeshGlobe />
      <TresAmbientLight :intensity="1" />
    </TresCanvas>
  </div>
</template>

<script lang="ts" setup>

</script>

<style></style>