<script setup lang="ts">
import LoginForm from '@/components/auth/LoginForm.vue';
import RegisterForm from '@/components/auth/RegisterForm.vue';
import type { TabsItem } from '@nuxt/ui';

const items = [
  {
    label: 'Login',
    slot: 'login' as const,
  },
  {
    label: 'Register',
    slot: 'register' as const,
  },
] satisfies TabsItem[];
</script>

<template>
  <div class="flex gap-4 flex-col min-w-md p-4 border border-secondary rounded-lg bg-white">
    <div class="max-w-120">
      <div class="flex items-center justify-center gap-2">
        <img src="/swap-cart-icon.svg" alt="Swap Cart Logo" />
        <h1 class="text-2xl">SwapCart</h1>
      </div>
      <p class="block pt-3 text-black-mute text-sm text-center">Join the community of swappers</p>
    </div>

    <UTabs color="neutral" :items="items" ui="{ trigger: 'grow' }" class="w-full">
      <template #login>
        <LoginForm
          class="[&_label]:text-black [&_label]:pb-1 [&_input]:bg-white [&_input]:text-black [&_div]:w-full"
        />
      </template>

      <template #register>
        <RegisterForm
          class="[&_label]:text-black [&_label]:pb-1 [&_input]:bg-white [&_input]:text-black [&_div]:w-full"
        />
      </template>
    </UTabs>
  </div>
</template>
