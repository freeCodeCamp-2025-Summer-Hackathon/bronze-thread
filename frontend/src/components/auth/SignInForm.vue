<script setup lang="ts">
import router from '@/router';
import type { FormSubmitEvent } from '@nuxt/ui';
import axios from 'axios';
import { reactive } from 'vue';
import * as z from 'zod';

const schema = z.object({
  email: z.string().email('Invalid email'),
  password: z.string().min(2, 'Must be at least 2 characters'),
});

type Schema = z.output<typeof schema>;

const state = reactive<Partial<Schema>>({
  email: undefined,
  password: undefined,
});

async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    const signedInUser = await axios.post(
      import.meta.env.VITE_API_BASE_URL + '/auth/signin',
      JSON.stringify(event.data),
    );

    if (!signedInUser) throw 'Failed to sign in.';

    router.push('/');
    console.log(signedInUser);
  } catch (error) {
    console.error(error || 'Something went wrong!');
  }
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4" @submit="onSubmit">
    <h1 class="text-white font-bold uppercase">Signin Form</h1>

    <UFormField label="Email" name="email">
      <UInput v-model="state.email" class="w-full" />
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" class="w-full" />
    </UFormField>

    <UButton type="submit"> Signin </UButton>
  </UForm>
</template>
