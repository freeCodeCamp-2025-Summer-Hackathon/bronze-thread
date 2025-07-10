<script setup lang="ts">
import type { FormSubmitEvent } from '@nuxt/ui';
import axios from 'axios';
import { reactive } from 'vue';
import * as z from 'zod';

const schema = z.object({
  username: z.string(),
  email: z.string().email('Invalid email'),
  password: z.string().min(2, 'Must be at least 2 characters'),
});

type Schema = z.output<typeof schema>;

const state = reactive<Partial<Schema>>({
  username: undefined,
  email: undefined,
  password: undefined,
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    const registeredUser = await axios.post(
      import.meta.env.VITE_API_BASE_URL + '/auth/register',
      JSON.stringify(event.data),
    );

    // router.push('/signin');
    console.log(registeredUser);
  } catch (error) {
    console.error(error || 'Something went wrong!');
  }
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
    <h1 class="text-white font-bold uppercase">Register Form</h1>

    <UFormField label="Username" name="username">
      <UInput v-model="state.username" class="w-full" />
    </UFormField>

    <UFormField label="Email" name="email">
      <UInput v-model="state.email" class="w-full" />
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" class="w-full" />
    </UFormField>

    <UButton type="submit"> Register </UButton>
  </UForm>
</template>
