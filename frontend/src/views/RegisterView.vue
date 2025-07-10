<script setup lang="ts">
import RegisterForm from '@/components/auth/RegisterForm.vue';
</script>

<template>
  <div class="flex flex-col p-4 bg-gray-500 text-black gap-4">
    <RegisterForm class="flex-1" />
    <div>
      <span class="text-white">Already have an account? </span>
      <RouterLink to="/signin" class="mt-2">
        <UButton color="neutral" variant="outline">Signin here</UButton></RouterLink
      >
    </div>
  </div>
</template>
