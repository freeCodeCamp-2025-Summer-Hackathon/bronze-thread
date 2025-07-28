<template>
  <section class="px-4 py-12">
    <h2 class="text-5xl mb-8 font-swap font-extrabold text-center">Explore Items</h2>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 justify-center p-10 md:grid-cols-3 lg:flex lg:flex-wrap lg:gap-6 gap-8"
    >
      <div
        v-for="item in items"
        :key="item.id"
        class="flex flex-col justify-between w-full lg:w-[calc(25%-1rem)]"
      >
        <img :src="item.image" :alt="item.name" class="w-full h-auto object-cover rounded-lg" />
        <div class="border-2 rounded-lg">
          <h3 class="text-lg text-center pb-5 font-bold">{{ item.name }}</h3>
          <p class="text-sm text-center text-gray-600">{{ item.category }}</p>
        </div>
      </div>
    </div>
  </section>
</template>
<script setup lang="ts">
// const { data: items, pending, error } = await useFetch('/api/items')
import bag1 from '../../assets/cta.png';
const items = [
  { id: 1, name: 'Leather Tote', category: 'Fashion Bag', image: bag1 },
  { id: 2, name: 'Mini Purse', category: 'Vintage', image: bag1 },
  { id: 3, name: 'Backpack', category: 'Unique and Verstalie', image: bag1 },
  { id: 4, name: 'Clutch', category: 'Stylish Storage', image: bag1 },
  { id: 5, name: 'Office Bag', category: 'Everyday Carry', image: bag1 },
  { id: 6, name: 'Evening Purse', category: 'Fashion Bag', image: bag1 },
  { id: 7, name: 'Shopping Bag', category: 'Vintage', image: bag1 },
  { id: 8, name: 'Tan Pouch', category: 'Stylish Storage', image: bag1 },
  { id: 8, name: 'Tan Pouch', category: 'Everday Carry', image: bag1 },
];
</script>
