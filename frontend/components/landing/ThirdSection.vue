<template>
  <div class="min-h-screen py-10 flex flex-col gap-8">
    <div class="mx-auto flex flex-row gap-6 text-lightRed">
      <svg width="72" height="72" viewBox="0 0 1024 1089.841" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
        class="text-lightRed">
        <path
          d="M520.258 1056.92H503.76c0-134.61-50.306-252.21-150.918-352.824S134.628 553.178 0 553.178V536.68c134.628 0 252.23-50.306 352.842-151.145C453.454 284.677 503.76 167.075 503.76 32.938h16.498c0 134.61 50.306 252.21 150.918 352.824S889.39 536.68 1024 536.68v16.498c-134.61 0-252.21 50.306-352.824 150.918S520.258 922.31 520.258 1056.92z" />
      </svg>

      <h1 class="tracking-wider">Contact Me</h1>

      <svg width="72" height="72" viewBox="0 0 1024 1089.841" xmlns="http://www.w3.org/2000/svg" fill="currentColor"
        class="text-lightRed">
        <path
          d="M520.258 1056.92H503.76c0-134.61-50.306-252.21-150.918-352.824S134.628 553.178 0 553.178V536.68c134.628 0 252.23-50.306 352.842-151.145C453.454 284.677 503.76 167.075 503.76 32.938h16.498c0 134.61 50.306 252.21 150.918 352.824S889.39 536.68 1024 536.68v16.498c-134.61 0-252.21 50.306-352.824 150.918S520.258 922.31 520.258 1056.92z" />
      </svg>
    </div>

    <div class="flex-1 w-full grid grid-cols-2 auto-rows-fr gap-4 px-10">
      <div class="flex items-center justify-center">
        <UForm :schema="schema" :state="state" class="w-full max-w-[368px] space-y-4 bg-darkGray p-4 py-8"
          @submit="onSubmit">
          <UFormGroup label="Email" name="email">
            <UInput v-model="state.email" />
          </UFormGroup>

          <UFormGroup label="Message" name="message">
            <UTextarea v-model="state.message" />
          </UFormGroup>

          <UButton type="submit">
            Submit
          </UButton>
        </UForm>
      </div>
      <div class="flex flex-col gap-6 items-center justify-center">
        <!-- My info -->
        <div class="space-y-2 flex flex-col items-center">
          <div class="h-20 w-20 bg-lightRed rounded-full" />
          <h3 class="text-lightRed tracking-wider">Adan Ayaz</h3>
          <span class="text-sm text-center">
            Lorem ipsum dolor sit amet consectetur adipisicing elit. Expedita nulla dolorem dolores officiis facere,
            quia dignissimos similique nam porro unde veniam! Natus, rerum alias mollitia consectetur pariatur explicabo
            officiis laboriosam.
          </span>
        </div>

        <!-- Quick Feedback -->
        <div class="flex flex-row gap-1">
          <svg v-for="i in 5" xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24"
            class="text-lightRed">
            <path fill="currentColor"
              d="M7.625 6.4L12 .725L16.375 6.4l6.85 2.3l-4.325 6.125l.175 6.825L12 19.675L4.925 21.65L5.1 14.8L.8 8.7z" />
          </svg>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

const schema = object({
  email: string().email('Invalid email').required('Required'),
  message: string()
    .required('Required')
})

type Schema = InferType<typeof schema>

const state = reactive({
  email: undefined,
  message: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
  // Do something with event.data
  console.log(event.data)
}

</script>

<style></style>