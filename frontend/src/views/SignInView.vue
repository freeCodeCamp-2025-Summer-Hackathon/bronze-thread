<script setup lang="ts">
import SignInForm from '@/components/auth/SignInForm.vue';
</script>

<template>
  <div class="flex flex-col p-4 bg-gray-500 text-black gap-4">
    <SignInForm />
    <div>
      <span class="text-white">Don't have an account? </span>
      <RouterLink to="/register" class="mt-2">
        <UButton color="neutral" variant="outline">Register here</UButton></RouterLink
      >
    </div>
  </div>
</template>
