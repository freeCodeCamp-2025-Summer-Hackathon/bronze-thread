<template>
  <div class="border rounded-xl p-4 w-64 shadow-sm">
    <h2 class="text-lg font-semibold mb-3 font-swap">Price Range</h2>
    <div class="mb-2">
      <label class="text-sm font-swap">Min:</label>
      <input type="number" v-model.number="modelValue.min" class="w-full border p-2 rounded mt-1" />
    </div>
    <div>
      <label class="text-sm font-swap">Max:</label>
      <input type="number" v-model.number="modelValue.max" class="w-full border p-2 rounded mt-1" />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
  modelValue: {
    min: number;
    max: number;
  };
}>();
const emit = defineEmits(['update:modelValue']);

const modelValue = computed({
  get: () => props.modelValue,
  set: (val) => emit('update:modelValue', val),
});
</script>
