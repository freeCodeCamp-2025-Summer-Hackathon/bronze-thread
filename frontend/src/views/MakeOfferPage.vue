<script setup lang="ts">
import { useRoute } from 'vue-router';
import { products } from '@/data/productsData.ts';
import FixedItemCard from '@/components/offer/fixedItemCard.vue';
import UserItemCard from '@/components/offer/userItemCard.vue';
import { ref } from 'vue';
import Navbar from '@/components/navbar/mainNav.vue';
import FooterNav from '@/components/footer/FooterNav.vue';
import SwapMethodModal from '@/components/shop/swapMethodModal.vue';

const route = useRoute();

const productId = Number(route.params.id);

const wantedItem = products.find((item) => item.id === productId);

const userItems = products.filter((item) => item.id !== productId);

const selectedItems = ref<number[]>([]);

const toggleSelection = (index: number) => {
  const i = selectedItems.value.indexOf(index);
  if (i === -1) selectedItems.value.push(index);
  else selectedItems.value.splice(i, 1);
};
const showSwapModal = ref(false);

const goToSwapMethod = () => {
  const selectedProductIds = selectedItems.value.map((index) => userItems[index].id);
  localStorage.setItem('selectedOfferItems', JSON.stringify(selectedProductIds));
  localStorage.setItem('wantedItemId', productId.toString());
  showSwapModal.value = true;
};
</script>

<template>
  <div class="min-h-screen w-full flex flex-col">
    <Navbar />
    <div class="flex flex-col pt-30 items-center justify-center text-center px-4">
      <h1 class="text-3xl font-swap font-bold mb-4">Make an Offer</h1>
      <p class="font-swap text-gray-500 text-lg max-w-xl">
        Select one or more of your items to offer in exchange for this item.
      </p>
    </div>
    <div class="max-w-7xl mx-auto px-4 py-10">
      <div class="flex flex-col lg:flex-row gap-10">
        <div class="flex justify-center lg:block">
          <FixedItemCard :item="wantedItem!" />
        </div>

        <div class="flex-1">
          <h2 class="text-xl font-swap font-semibold mb-4">Select Items to Offer</h2>
          <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-5">
            <UserItemCard
              v-for="(item, index) in userItems"
              :key="item.id"
              :item="item"
              :selected="selectedItems.includes(index)"
              @toggle="toggleSelection(index)"
            />
          </div>

          <div class="pt-10 flex flex-col items-center">
            <button
              class="bg-black font-swap hover:bg-white hover:border-swapbase hover:border-2 hover:text-swapbase text-white px-6 py-3 rounded-lg disabled:opacity-50"
              :disabled="selectedItems.length === 0"
              @click="goToSwapMethod"
            >
              Continue with Offer
            </button>
            <SwapMethodModal v-model:show="showSwapModal" />

            <p class="text-gray-500 font-swap text-sm mt-3">
              Select one or more items from your collection to make an offer.
            </p>
          </div>
        </div>
      </div>
    </div>
    <FooterNav />
  </div>
</template>
