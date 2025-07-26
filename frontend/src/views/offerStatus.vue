<template>
  <div class="min-h-screen w-full   pt-10 ">
    <Navbar />

    <div class="flex justify-between pb-10 md:px-72 sm:px-40 lg:px-72 pt-20 items-center mb-6">
      <div>
        <h1 class="text-2xl font-swap lg:px-12 font-semibold">Offer Status</h1>
      </div>
      <div class="px-12">
        <div class="bg-yellow-100 text-yellow-700 px-4 py-1  rounded-full text-sm">
          Pending
        </div>
      </div>
    </div>
    <div class="flex flex-col items-center sm:justify-items-center justify-center">
      <div class="grid grid-cols-1 md:grid-cols-2 gap-8 max-w-4xl w-full">

        <div class=" rounded-lg p-6 mb-6 w-full sm:items-center flex flex-col ">
          <h2 class="text-sm text-gray-600 font-swap pb-4">You’re Getting</h2>

          <div class="border rounded-lg  overflow-hidden w-full max-w-sm shadow-sm ">
            <img :src="offer.requestedItems[0]?.image" alt="Receiving Item"
              class="w-full px-3 pt-3 rounded-4xl  h-60 object-cover" />

            <div class="p-5">
              <h3 class="font-extrabold text-2xl font-swap">{{ offer.requestedItems[0]?.title }}</h3>
              <h4 class="text-gray-600 text-lg font-swap font-semibold">
                $ {{ offer.requestedItems[0]?.price }}
              </h4>
              <p class="text-gray-400 font-swap text-sm">
                {{ offer.requestedItems[0]?.category }}
              </p>
            </div>
          </div>
        </div>

        <!-- <div class=" rounded-lg p-6 mb-6 w-full sm:items-center  flex flex-col ">
          <h2 class="text-sm text-gray-600 font-swap pb-4">You’re Giving</h2>

          <div class="flex flex-wrap gap-6  justify-center">
            <div v-for="(item, index) in offer.offeredItems" :key="index"
              class="border rounded-lg  overflow-hidden w-full max-w-sm shadow-sm ">
              <img :src="item.image" alt="Offered Item" class="w-full h-60 px-3 pt-3 object-cover rounded-4xl" />
              <div class="p-5">
                <h3 class="font-extrabold text-2xl font-swap">{{ item.title }}</h3>
                <h4 class="text-gray-600 text-lg font-swap font-semibold">
                  $ {{ item.price }}
                </h4>
                <p class="text-gray-400 font-swap text-sm">
                  {{ item.category }}
                </p>
              </div>
            </div>
          </div>
        </div> -->
<div class="rounded-lg p-6 mb-6 w-full sm:items-center flex flex-col">
  <h2 class="text-sm text-gray-600 font-swap pb-4">You’re Giving</h2>

  <!-- Force the inner div to match left-side width -->
  <div class="w-full max-w-sm">
    <div
      class="flex flex-wrap gap-6 justify-center"
      :class="{ 'justify-start': offer.offeredItems.length === 1 }"
    >
      <div
        v-for="(item, index) in offer.offeredItems"
        :key="index"
        class="border rounded-lg overflow-hidden w-full shadow-sm"
      >
        <img :src="item.image" alt="Offered Item" class="w-full h-60 px-3 pt-3 object-cover rounded-4xl" />
        <div class="p-5">
          <h3 class="font-extrabold text-2xl font-swap">{{ item.title }}</h3>
          <h4 class="text-gray-600 text-lg font-swap font-semibold">
            $ {{ item.price }}
          </h4>
          <p class="text-gray-400 font-swap text-sm">
            {{ item.category }}
          </p>
        </div>
      </div>
    </div>
  </div>
</div>

      </div>

      <div class="flex flex-wrap justify-center pt-20 gap-4 mt-6">

        <button class="bg-gray-200 text-black  hover:bg-gray-600 hover:text-white px-6 py-3 rounded font-swap">
          Edit Offer
        </button>
        <button class="bg-white border text-red-600 px-6 py-3 hover:bg-red-700 hover:text-white  rounded font-swap">
          Cancel Offer
        </button>
        <button
          class="bg-black  px-6 hover:text-swapbase hover:border-swapbase hover:border-2 hover:bg-white text-white py-2 rounded">
          Track offer


        </button>
      </div>

      <div class="px-46  py-10">
        <div class="border-2 p-10 rounded-lg font-swap">
          <h2 class="font-semibold text-lg pb-4">What Happens Next?</h2>
          <div class="space-y-4">
            <div class="flex items-start gap-4">
              <div class="w-6 h-6 bg-[#F6E5DC]  rounded-full flex items-center justify-center text-sm font-semibold">1
              </div>
              <div>
                <strong>Owner Reviews Your Offer</strong>
                <p class="text-sm text-gray-600">They will be notified and can review the offer details.</p>
              </div>
            </div>
            <div class="flex items-start gap-4">
              <div class="w-6 h-6 bg-[#F6E5DC]  rounded-full flex items-center justify-center text-sm font-semibold">2
              </div>
              <div>
                <strong>Response & Negotiation</strong>
                <p class="text-sm text-gray-600">They can accept or decline offer.</p>
              </div>
            </div>
            <div class="flex items-start gap-4">
              <div class="w-6 h-6 bg-[#F6E5DC]  rounded-full flex items-center justify-center text-sm font-semibold">3
              </div>
              <div>
                <strong>Complete the Swap</strong>
                <p class="text-sm text-gray-600">Once accepted, you’ll coordinate the swap through the platform.</p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <FooterNav />
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import Navbar from '@/components/navbar/mainNav.vue'
import FooterNav from '@/components/footer/FooterNav.vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const offer = ref<any>({
  offeredItems: [],
  requestedItems: []
})

onMounted(() => {
  const data = localStorage.getItem('swapDetails')
  if (data) {
    offer.value = JSON.parse(data)
  }
})
</script>