<template>
  <div class="border border-gray-200 bg-[#F6E5DC] rounded-2xl p-5 w-full max-w-sm">
  
    <img :src="item.image" alt="Item Image" class="rounded-xl  mb-4 w-full h-auto object-cover" />
    <div class="text-lg font-swap font-semibold mb-1">{{ item.title }}</div>
    <div class="text-gray-600 font-swap mb-2">${{ item.price }}</div>
    <div class="flex items-center gap-3 mt-3">
      <div>
        <div class="text-sm font-swap font-medium">{{ item.seller.name }}</div>
        <div class="text-xs font-swap text-gray-500">⭐ {{ item.seller.rating }} · {{ item.listed }}</div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
defineProps<{
  item: {
    image: string
    title: string
    price: number
   
    seller: { name:string
      initials: string
      rating: string
    }
    listed: string
  }
}>()
</script>
