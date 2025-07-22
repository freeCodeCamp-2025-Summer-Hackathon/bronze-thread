<template>
  <div
    class="relative border p-4 rounded-xl transition hover:shadow-md cursor-pointer"
    :class="selected ? 'border-black' : 'border-gray-300'"
    @click="emit('toggle')"
  >
    <img :src="item.image" class="rounded-md mb-3 w-full h-40 object-cover" />
    <div class=" font-swap font-semibold">{{ item.title }}</div>
    <div class="font-swap text-sm mt-1">${{ item.price }}</div>

    <div
      v-if="selected"
      class="absolute top-2 right-2 bg-black text-white w-6 h-6 flex items-center justify-center rounded-full"
    >
      ✓
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  item: {
    image: string
    title: string
    price: number
  }
  selected: boolean
}>()

const emit = defineEmits(['toggle'])
</script>
