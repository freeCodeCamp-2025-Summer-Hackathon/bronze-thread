
<template>
  <div v-if="show" class="fixed inset-0 z-50 flex items-center justify-center backdrop-blur-sm bg-black/30">
    <div class="bg-[#FDF5F1] max-w-3xl w-full p-6 rounded-lg shadow-lg overflow-y-auto max-h-[90vh] relative">

      <button @click="close" class="absolute top-3 right-3 text-black text-2xl hover:text-red-600">
        &times;
      </button>

      <template v-if="!success">
        <h2 class="text-xl font-swap font-semibold mb-4 text-center">Choose Swap Method</h2>

        <div class="flex gap-4 mb-6">
          <div class="w-1/2 p-4 border rounded-lg cursor-pointer"
            :class="selectedMethod === 'inPerson' ? 'border-orange-400' : 'border-gray-300'"
            @click="selectMethod('inPerson')">
            <div class="font-semibold font-swap mb-2">📍 Meet in Person</div>
            <ul class="text-sm font-swap text-gray-500 list-disc list-inside">
              <li>Free and immediate</li>
              <li>Inspect items before swap</li>
              <li>Safe public locations</li>
            </ul>
          </div>

          <div class="w-1/2 p-4 border rounded-lg cursor-pointer"
            :class="selectedMethod === 'delivery' ? 'border-orange-400' : 'border-gray-300'"
            @click="selectMethod('delivery')">
            <div class="font-semibold font-swap  mb-2">🚚 Delivery Service</div>
            <ul class="text-sm font-swap text-gray-500 list-disc list-inside">
              <li>$15 delivery fee (split)</li>
              <li>Insured and tracked</li>
              <li>Convenient scheduling</li>
            </ul>
          </div>
        </div>

        <div v-if="selectedMethod === 'inPerson'" class="mb-6">
          <label class="block font-medium mb-1">Meetup Location</label>
          <textarea v-model="inPersonLocation" rows="3" :class="inputClass(errors.inPersonLocation)"
            placeholder="Enter meetup point" />
          <p v-if="errors.inPersonLocation" class="text-red-500 font-swap text-sm mt-1">
            Meetup location is required.
          </p>
          <div class="bg-[#FFF8E1] border border-orange-300 font-swap p-4 mt-4 rounded text-sm text-gray-700">
            <p>- Meet in public, well-lit locations during daylight hours</p>
            <p>- Inspect items thoroughly before completing the swap</p>
            <p>- Trust your instincts - if something feels wrong, don't proceed</p>
            <p>- Bring a friend if possible, especially for high-value items</p>
          </div>
        </div>

        <div v-if="selectedMethod === 'delivery'" class="mb-6">
          <label class="block font-swap font-medium mb-1">Delivery Address</label>
          <textarea v-model="deliveryAddress" rows="3" :class="inputClass(errors.deliveryAddress)"
            placeholder="Enter full delivery address" />
          <p v-if="errors.deliveryAddress" class="text-red-500 font-swap text-sm mt-1">
            Delivery address is required.
          </p>
          <div class="bg-[#FFF8E1] border font-swap border-orange-300 p-4 mt-4 rounded text-sm text-gray-700">
            <p>- Delivery fee: $15 (you pay $7.50, they pay $7.50)</p>
            <p>- Estimated delivery: 2-3 business days</p>
            <p>- Tracking info provided</p>
          </div>
        </div>

        <button class="mt-4 w-full bg-black text-white px-12 py-4 rounded hover:bg-white hover:border-2 hover:border-swapbase hover:text-swapbase" @click="submitSwapDetails">
          CONFIRM SWAP DETAILS
        </button>
      </template>

      <div v-else class="text-center">
        <p class="text-green-600 font-swap font-semibold text-3xl pb-10">
          ✅ Offer sent successfully!
        </p>
        <div class="flex flex-row justify-center gap-x-10">
        <button class="bg-black  text-white px-14 py-4 rounded hover:bg-white hover:border-2 hover:border-swapbase hover:text-swapbase" @click="close">
          OK
        </button>
         <button class="bg-white  border-2 text-black px-6 py-4 rounded hover:bg-swapbase hover:text-white" @click="goToStatus">
          Check Status
        </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch, defineProps, defineEmits } from 'vue'
import { useRouter } from 'vue-router'

const props = defineProps<{ show: boolean }>()
const emit = defineEmits(['update:show'])
const router = useRouter()

const selectedMethod = ref<'inPerson' | 'delivery' | null>(null)
const inPersonLocation = ref('')
const deliveryAddress = ref('')
const success = ref(false)

const errors = ref({
  inPersonLocation: false,
  deliveryAddress: false,
})
interface Product {
  id: number
  title: string
  image: string
  price: number
  category: string
}
interface MethodData {
  method: 'In-Person Meetup' | 'Delivery'
  meetupLocation?: string
  address?: string
}

const selectMethod = (method: 'inPerson' | 'delivery') => {
  selectedMethod.value = 
  method
}

const clearErrors = () => {
  errors.value.inPersonLocation = false
  errors.value.deliveryAddress = false
}
const goToStatus = () => {
  router.push('/offerStatus')
}
const inputClass = (hasError: boolean) =>
  `w-full p-2 border rounded ${hasError ? 'border-red-500' : 'border-gray-300'}`
const submitSwapDetails = () => {
  clearErrors()
  if (!selectedMethod.value) {
    alert('Please choose a swap method.')
    return
  }

  const methodData: MethodData  = {
    method: selectedMethod.value === 'inPerson' ? 'In-Person Meetup' : 'Delivery'
  }

  if (selectedMethod.value === 'inPerson') {
    if (!inPersonLocation.value.trim()) {
      errors.value.inPersonLocation = true
      return
    }
    methodData.meetupLocation = inPersonLocation.value
  }

  if (selectedMethod.value === 'delivery') {
    if (!deliveryAddress.value.trim()) {
      errors.value.deliveryAddress = true
      return
    }
    methodData.address = deliveryAddress.value
  }

  const offeredIds = JSON.parse(localStorage.getItem('selectedOfferItems') || '[]')
  const wantedId = Number(localStorage.getItem('wantedItemId'))
const allProducts: Product[] = JSON.parse(localStorage.getItem('allProducts') || '[]')
const offeredItems = allProducts.filter(p => offeredIds.includes(p.id))
const wantedItem = allProducts.find(p => p.id === wantedId)
 

  const swapDetails = {
    ...methodData,
    offeredItems,
    requestedItems: wantedItem ? [wantedItem] : []
  }

  localStorage.setItem('swapDetails', JSON.stringify(swapDetails))
  success.value = true
}


watch(() => props.show, (newVal) => {
  if (newVal) {
    selectedMethod.value = 'inPerson'
  }
})

const close = () => {
  emit('update:show', false)
  success.value = false
  selectedMethod.value = null
  inPersonLocation.value = ''
  deliveryAddress.value = ''
}
</script>
