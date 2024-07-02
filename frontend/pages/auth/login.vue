<script setup lang="ts">
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

const schema = object({
    email: string().email('Invalid email').required('Required'),
    password: string()
        .required('Required')
})

type Schema = InferType<typeof schema>

const state = reactive({
    email: undefined,
    password: undefined
})

async function onSubmit(event: FormSubmitEvent<Schema>) {
    // Do something with event.data
    window.alert(JSON.stringify(event.data))
}
</script>

<template>
    <div class="min-h-screen bg-black flex flex-row items-stretch">
        <!-- Left Globe Side -->
        <div class="h-screen w-full bg-darkGray flex-1">
            Globe goes here
        </div>

        <!-- Right Login Form -->
        <div class="p-2 min-w-60">
            <h2>Login</h2>

            <UForm :schema="schema" :state="state" class="w-full space-y-4 bg-darkGray p-4 py-8" @submit="onSubmit">
                <UFormGroup label="Email" name="email">
                    <UInput v-model="state.email" />
                </UFormGroup>

                <UFormGroup label="Password" name="password">
                    <UInput v-model="state.password" />
                </UFormGroup>

                <UButton type="submit" class="w-full text-center">
                    Submit
                </UButton>

                <ULink class="block text-xs" to="/auth/register">
                    Don't have an account?
                </ULink>
            </UForm>
        </div>
    </div>
</template>