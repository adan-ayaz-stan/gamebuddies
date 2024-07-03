<script setup lang="ts">
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

definePageMeta({
    middleware: ["auth-server"]
})

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
    // Send the form data to server route "http://localhost:8080/api/v1/login"
    const { data, error } = await useFetch("http://localhost:8080/api/v1/login", {
        method: "POST",
        credentials: "include",
        body: JSON.stringify({
            email: event.data.email,
            encrypted_password: event.data.password,
            redirect_url: "/dashboard",
            origin_url: "/auth/login"
        }),
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
        }
    });

    console.log(data.value, error.value);
}

async function checkServerStatus() {
    const { data, error } = await useFetch("http://localhost:8080/api/v1/heartbeat", {
        credentials: "include",
    });

    console.log(data.value, error.value);
}
</script>

<template>
    <div class="min-h-screen bg-black flex flex-row items-stretch">
        <!-- Left Globe Side -->
        <div class="hidden md:block h-screen w-full bg-darkGray flex-1">
            Globe goes here
        </div>

        <!-- Right Login Form -->
        <div class="p-2 min-w-60 w-full md:w-fit flex flex-col items-center">
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

            <UButton type="button" @click="checkServerStatus">
                Check Server Status
            </UButton>
        </div>
    </div>
</template>