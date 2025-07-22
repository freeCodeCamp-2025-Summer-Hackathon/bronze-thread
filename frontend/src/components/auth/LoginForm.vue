<script setup lang="ts">
import type { FormSubmitEvent } from '@nuxt/ui';
import { useToast } from '@nuxt/ui/runtime/composables/useToast.js';
import axios from 'axios';
import { reactive } from 'vue';
import { useRouter } from 'vue-router';
import * as z from 'zod';

const router = useRouter();

const schema = z.object({
  email: z.string().email(),
  password: z.string().min(4, 'At least 4 char long.'),
});

type Schema = z.output<typeof schema>;

const state = reactive<Partial<Schema>>({
  email: 'guest@mail.com',
  password: 'guestPass',
});

const toast = useToast();
async function onSubmit(event: FormSubmitEvent<Schema>) {
  try {
    const loggedInUser = await axios.post(
      import.meta.env.VITE_API_BASE_URL + '/auth/signin',
      event.data,
    );

    console.log(loggedInUser);

    toast.add({
      title: 'Success',
      description: 'Login Successful.',
      color: 'success',
    });

    setTimeout(() => {
      router.replace('/');
    }, 5000);
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
    <UFormField label="Email" name="email">
      <UInput v-model="state.email" />
    </UFormField>

    <UFormField label="Password" name="password">
      <UInput v-model="state.password" type="password" />
    </UFormField>

    <UButton
      type="submit"
      label="Login"
      class="flex items-center justify-center bg-black hover:bg-black-mute text-primary"
    />
  </UForm>
</template>
