<template>
  <section class="px-4 py-12">
    <h2 class="text-5xl mb-8 font-swap font-extrabold text-center">Explore Items</h2>

    <div
      class="grid grid-cols-1 sm:grid-cols-2 justify-center p-10 md:grid-cols-3 lg:flex lg:flex-wrap lg:gap-6 gap-8"
    >
      <div
        v-for="product in displayedProducts"
        :key="product.id"
        class="flex flex-col justify-between w-full lg:w-[calc(25%-1rem)] cursor-pointer"
      >
        <img
          :src="product.image"
          :alt="product.title"
          class="w-full h-auto object-cover rounded-lg"
          @click="goToDetail(product.id)"
        />

        <div class="border-2 rounded-lg text-center py-3">
          <h3 class="text-lg font-bold text-[#222] hover:underline" @click="goToDetail(product.id)">
            {{ product.title }}
          </h3>
          <p class="text-sm text-gray-600">{{ product.category }}</p>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
import { products } from '@/data/productsData';
const router = useRouter();

const displayedProducts = products.slice(0, 9);

const goToDetail = (id: number) => {
  router.push(`/product/${id}`);
};
</script>
