
<script setup lang="ts">
defineProps<{
  modelValue: string
}>()

const emit = defineEmits<{
  (e: 'update:modelValue', value: string): void
}>()
</script>

<template>
  <div class="mb-6 w-full flex justify-center px-4">
    <div class="relative w-full max-w-md sm:max-w-lg md:max-w-xl lg:max-w-2xl">
      <img
        src="../../assets/search.png"
        alt="Search Icon"
        class="absolute left-3 top-1/2 transform -translate-y-1/2 w-5 h-5 opacity-60"
      />
      <input
        :value="modelValue"
        @input="event => emit('update:modelValue', (event.target as HTMLInputElement).value)"
        type="text"
        placeholder="Search for products..."
        class="border border-gray-300 font-swap rounded-full pl-10 pr-4 py-2 w-full
               text-sm sm:text-base focus:outline-none focus:ring-2 focus:ring-swapbase"
      />
    </div>
  </div>
</template>
