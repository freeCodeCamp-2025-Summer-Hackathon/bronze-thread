<template>
  <div class="flex justify-center w-full">
    <div
      class="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-3 gap-6 gap-x-20 md:justify-center md:gap-x-20 md:grid-col-2 max-w-7xl p-10 justify-items-center"
    >
      <ProductCard v-for="product in filteredProducts" :key="product.id" :product="product" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import ProductCard from '../../components/shop/productCard.vue';

const props = defineProps<{
  products: {
    id: number;
    title: string;
    category: string;
    price: number;
    image: string;
  }[];
  selectedCategories: string[];
  priceRange: {
    min: number;
    max: number;
  };
}>();

// const emit = defineEmits(['view'])

const filteredProducts = computed(() =>
  props.products.filter(
    (product) =>
      (!props.selectedCategories.length || props.selectedCategories.includes(product.category)) &&
      product.price >= props.priceRange.min &&
      product.price <= props.priceRange.max,
  ),
);
</script>
