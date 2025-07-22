<template>
  <div class="w-full">
    <navbar />
    <heroHeading />

    <SearchBar v-model="searchQuery" />

    <div class="container mx-auto px-4 mt-6">
      <div class="flex flex-col lg:flex-row gap-8 px-4">
        <div class="w-full lg:w-auto flex justify-center lg:justify-start">
          <div
            class="flex flex-col items-center md:flex-row md:items-start lg:flex-col gap-6 pt-10"
          >
            <CategoryFilter v-model="selectedCategories" :products="products" />
            <PriceFilter v-model="priceRange" />
          </div>
        </div>

        <div class="w-full lg:w-4/6 flex justify-center">
          <ProductGrid
            :products="filteredProducts"
            :selected-categories="selectedCategories"
            :price-range="priceRange"
            @view="handleViewProduct"
          />
        </div>
      </div>
    </div>
    <FooterNav />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import navbar from '@/components/navbar/mainNav.vue';
import SearchBar from '@/components/shop/ProductSearch.vue';
import PriceFilter from '@/components/shop/priceFilter.vue';
import CategoryFilter from '@/components/shop/categoryFilter.vue';
import ProductGrid from '@/components/shop/productGrid.vue';
import FooterNav from '@/components/footer/FooterNav.vue';
import heroHeading from '@/components/shop/heroHeading.vue';
import { products } from '@/data/productsData';

const searchQuery = ref('');
const selectedCategories = ref<string[]>([]);
const priceRange = ref({ min: 0, max: 10000 });

const filteredProducts = computed(() => {
  return products.filter((product) => {
    const matchesSearch = product.title.toLowerCase().includes(searchQuery.value.toLowerCase());
    return matchesSearch;
  });
});

function handleViewProduct(id: number) {
  console.log('View product', id);
}
</script>
