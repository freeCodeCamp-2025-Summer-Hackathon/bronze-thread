<script setup lang="ts">
import { useRoute } from 'vue-router';
import Navbar from '@/components/navbar/mainNav.vue';
import { products } from '@/data/productsData.ts';
import { computed, watch } from 'vue';
import FooterNav from '@/components/footer/FooterNav.vue';
import { ref } from 'vue';

const route = useRoute();
const productId = Number(route.params.id);
const activeTab = ref<'description' | 'specs'>('description');

const product = computed(() => products.find((p) => p.id === productId));
const mainImage = ref(product.value?.image || '');
import { useRouter } from 'vue-router';

const router = useRouter();

function goToOfferPage() {
  if (product.value) {
    router.push(`/makeOffer/${product.value.id}`);
  }
}

watch(product, (newProduct) => {
  mainImage.value = newProduct?.image || '';
});
</script>

<template>
  <div class="min-h-screen w-full flex flex-col">
    <Navbar />
    <div class="flex-1 sm:pt-24 sm:p-10 lg:pt-24 lg:p-26 md:pt-24 md:p-26">
      <div v-if="product" class="max-w-7xl mx-auto p-6">
        <div class="grid grid-cols-1 lg:grid-cols-2 gap-10">
          <div class="flex flex-col gap-y-5">
            <div class="mb-10 border-2 p-2 rounded-lg">
              <img
                :src="mainImage"
                :alt="product.title"
                class="rounded-lg shadow-md object-center object-cover mb-10 h-96 w-full"
              />
            </div>

            <div class="flex flex-wrap gap-4 justify-center mb-10 mt-10">
              <img
                v-for="(thumb, index) in product.thumbnails"
                :key="index"
                :src="thumb"
                @click="mainImage = thumb"
                class="w-32 sm:w-20 md:w-28 h-auto border hover:scale-105 transition rounded cursor-pointer"
              />
            </div>
          </div>

          <div class="flex flex-col gap-y-2">
            <h1 class="text-4xl font-swap font-bold mb-2">{{ product.title }}</h1>
            <p class="text-xl font-swap font-bold text-gray-600 mb-4">${{ product.price }}</p>

            <div class="border-swapbase border-2 bg-[#F6E5DC] rounded-2xl p-5">
              <div class="flex justify-between items-center">
                <div class="flex items-start gap-3">
                  <div
                    class="bg-[#FCCEE8] text-white w-10 h-10 rounded-full flex items-center justify-center text-sm font-bold"
                  >
                    {{ product.seller.initials }}
                  </div>

                  <div class="flex flex-col">
                    <p class="font-bold font-swap text-sm">{{ product.seller.name }}</p>

                    <div class="flex items-center text-sm text-gray-700 gap-3 mt-1">
                      <div class="flex items-center gap-1 text-black">
                        <span class="text-yellow-500">★ ★ ★ ★ ★</span>
                        <span class="text-black font-medium">{{ product.seller.rating }}</span>
                      </div>

                      <div class="flex items-center gap-1">
                        <svg
                          xmlns="http://www.w3.org/2000/svg"
                          class="h-4 w-4 text-gray-600"
                          viewBox="0 0 20 20"
                          fill="currentColor"
                        >
                          <path
                            fill-rule="evenodd"
                            d="M10 2a6 6 0 016 6c0 4.418-6 10-6 10S4 12.418 4 8a6 6 0 016-6zm0 8a2 2 0 100-4 2 2 0 000 4z"
                            clip-rule="evenodd"
                          />
                        </svg>
                        <span>{{ product.seller.location }}</span>
                      </div>
                    </div>
                  </div>
                </div>
                <button
                  class="border border-black bg-white font-swap px-4 py-2 rounded hover:bg-swapbase hover:text-white"
                >
                  View Profile
                </button>
              </div>
            </div>
            <div class="border-[#FFE082] mt-10 bg-[#FFF8E1] p-4 rounded mb-6">
              <h3 class="font-extrabold text-xl font-swap mb-2">Looking to swap for:</h3>
              <ul class="list-disc list-inside font-swap text-sm text-gray-700">
                <li v-for="(item, index) in product.highlights" :key="index" class="pt-2">
                  {{ item }}
                </li>
              </ul>
            </div>
            <div class="flex gap-4 mb-6 flex-wrap">
              <button
                @click="goToOfferPage"
                class="bg-black text-white font-swap px-24 py-2 rounded hover:border-swapbase hover:border-2 hover:text-swapbase hover:bg-white"
              >
                Make an Offer
              </button>

              <button
                class="border border-black font-swap px-24 py-2 rounded hover:text-white hover:bg-swapbase"
              >
                Send Message
              </button>
            </div>
          </div>
        </div>

        <div class="pt-10 mt-10">
          <div class="flex gap-6 border-b pb-2 mb-4 cursor-pointer text-sm">
            <span
              :class="[
                activeTab === 'description'
                  ? 'font-semibold font-swap  text-black border-b-2 border-black'
                  : 'text-gray-400',
              ]"
              @click="activeTab = 'description'"
            >
              Description
            </span>
            <span
              :class="[
                activeTab === 'specs'
                  ? 'font-semibold font-swap text-black border-b-2 border-black'
                  : 'text-gray-400',
              ]"
              @click="activeTab = 'specs'"
            >
              Specifications
            </span>
          </div>
          <div v-if="activeTab === 'description'" class="text-gray-700">
            {{ product.description }}
          </div>
          <div v-else class="text-gray-700">
            <ul class="list-disc list-inside space-y-1">
              <li v-for="(spec, index) in product.specs" :key="index">
                {{ spec }}
              </li>
            </ul>
          </div>
        </div>
      </div>
    </div>
    <FooterNav class="mt-10" />
  </div>
</template>
