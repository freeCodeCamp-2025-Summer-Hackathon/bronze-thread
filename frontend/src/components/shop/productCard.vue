<template>
  <div class="border w-72 rounded-xl p-5 shadow-md hover:shadow-lg transition duration-200">
    <img
      :src="product.image"
      :alt="product.title"
      class="w-full h-60 object-cover rounded-lg mb-4"
    />
    <h3 class="text-xl font-swap text-center font-semibold mb-2">{{ product.title }}</h3>
    <p class="text-sm font-swap text-center text-gray-500 mb-2">{{ product.category }}</p>
    <p class="text-lg font-swap text-center font-bold mb-4">${{ product.price }}</p>
    <button
      @click="handleClick"
      class="w-full bg-black hover:text-swapbase hover:border-swapbase hover:border-2 hover:bg-white text-white py-2 rounded"
    >
      View Product
    </button>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();
const props = defineProps<{
  product: {
    id: number;
    title: string;
    category: string;
    price: number;
    image: string;
  };
}>();

function handleClick() {
  router.push(`/product/${props.product.id}`);
}
</script>
