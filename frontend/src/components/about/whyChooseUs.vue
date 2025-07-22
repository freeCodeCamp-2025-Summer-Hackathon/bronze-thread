<template>
  <section class="bg-[#FFEBE2] py-16 px-6 text-center">
    <h2 class="text-3xl md:text-4xl font-bold text-gray-900 mb-4">Why Choose Our Collection?</h2>
    <p class="text-gray-700 mb-12 max-w-xl mx-auto">
      Every piece in our collection is thoughtfully designed with premium materials and attention to
      detail.
    </p>

    <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-5xl mx-auto">
      <div
        v-for="feature in features"
        :key="feature.title"
        class="bg-[#F6E5DC] rounded-xl shadow p-6"
      >
        <div
          class="text-3xl bg-orange-400 text-white w-12 h-12 flex items-center justify-center rounded-full mx-auto mb-4"
        >
          {{ feature.icon }}
        </div>
        <h3 class="text-lg font-semibold text-gray-800 mb-2">
          {{ feature.title }}
        </h3>
        <p class="text-gray-600 text-sm">
          {{ feature.description }}
        </p>
      </div>
    </div>
  </section>
</template>

<script lang="ts" setup>
interface Feature {
  icon: string;
  title: string;
  description: string;
}

const features: Feature[] = [
  {
    icon: '✔️',
    title: 'Premium Quality',
    description:
      'Crafted from the finest materials with meticulous attention to detail and durability.',
  },
  {
    icon: '🕰️',
    title: 'Timeless Design',
    description:
      'Classic styles that transcend trends, ensuring your investment lasts for years to come.',
  },
  {
    icon: '❤️',
    title: 'Customer Satisfaction',
    description:
      'We stand behind every product with exceptional customer service and satisfaction guarantee.',
  },
];
</script>
