<template>
  <div class="border rounded-xl p-4 w-64 shadow-sm">
    <h2 class="text-lg font-semibold mb-3 font-swap">Categories</h2>
    <div v-for="category in uniqueCategories" :key="category" class="mb-2">
      <label class="flex items-center space-x-2 cursor-pointer">
        <input
          type="checkbox"
          :value="category"
          v-model="modelValue"
          class="appearance-none w-5 h-5 border border-gray-300 rounded-sm bg-white checked:bg-swapbase checked:border-swapbase relative checked:after:content-['✔'] checked:after:text-white checked:after:absolute checked:after:top-[-2px] checked:after:left-[3px] checked:after:text-sm checked:after:font-bold"
        />

        <span class="text-gray-700 pl-2 font-swap">{{ category }}</span>
      </label>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  modelValue: string[];
  products: { category: string }[];
}>();
const emit = defineEmits(['update:modelValue']);

const uniqueCategories = computed(() => {
  const set = new Set(props.products.map((p) => p.category));
  return [...set];
});

const modelValue = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
});
</script>
