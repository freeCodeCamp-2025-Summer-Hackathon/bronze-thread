<script setup lang="ts">
import type { FormSubmitEvent } from '@nuxt/ui';
import { useToast } from '@nuxt/ui/runtime/composables/useToast.js';
import axios from 'axios';
import { reactive } from 'vue';
import * as z from 'zod';

const schema = z.object({
  username: z.string().min(2, 'At least 2 char long.'),
  email: z.string().email(),
  password: z.string().min(4, 'At least 4 char long.'),
});

type Schema = z.output<typeof schema>;

const state = reactive<Partial<Schema>>({
  username: 'guest',
  email: 'guest@mail.com',
  password: 'guestPass',
});

const toast = useToast();
async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    const registeredUser = await axios.post(
      import.meta.env.VITE_API_BASE_URL + '/auth/register',
      event.data,
    );

    console.log(registeredUser);

    toast.add({
      title: 'Success',
      description: 'Registeration Successful.',
      color: 'success',
    });
  } catch (error) {
    if (axios.isAxiosError(error)) {
      toast.add({
        title: error.response?.data.error || error.message,
        description: error.message,
        color: 'error',
      });
    } else {
      toast.add({
        title: 'Failed',
        description: 'Something went wrong!',
        color: 'error',
      });
    }
  }
}
</script>

<template>
  <UForm :schema="schema" :state="state" class="space-y-4 flex flex-col gap-3" @submit="onSubmit">
    <UFormField label="Username" name="username">
      <UInput v-model="state.username" />
    </UFormField>

    <UFormField label="Email" name="email">
      <UInput v-model="state.email" />
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" />
    </UFormField>

    <UButton
      type="submit"
      label="Create Account"
      class="flex items-center justify-center bg-black hover:bg-black-mute text-primary"
    />
  </UForm>
</template>
