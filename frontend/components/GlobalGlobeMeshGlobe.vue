<script setup>
import { useLoop } from '@tresjs/core'

const boxRef = ref()

const { onBeforeRender } = useLoop()

onBeforeRender(({ delta }) => {
  boxRef.value.rotation.y += delta
})
</script>

<template>
  <TresMesh ref="boxRef">
    <TresBoxGeometry />
    <TresMeshBasicMaterial color="teal" />
  </TresMesh>
</template>