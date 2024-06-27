<script lang="ts" setup>
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

const schema = object({
  file: string(),
})

type Schema = InferType<typeof schema>

const state = reactive({
  file: undefined,
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  // Do something with event.data
  console.log(event.data)
}

</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4 p-4 bg-fadeRed rounded-lg" @submit="onSubmit">
    <UFormGroup label="File" name="file">
      <UInput type="file" accept="image/*" v-model="state.file" />
    </UFormGroup>

    <UButton type="submit">
      Submit
    </UButton>
  </UForm>
</template>

<style></style>
