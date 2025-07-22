<template>
  <section
    class="relative h-[99vh] w-full flex flex-col justify-center items-center text-center px-6 text-white"
    :style="{
      backgroundImage: `url('${images[currentIndex]}')`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
      backgroundRepeat: 'no-repeat',
    }"
  >
    <div class="bg-white/30 absolute inset-0 z-0"></div>

    <div class="relative z-10 text-center px-4 py-10">
      <h1 class="text-6xl font-swap text-black sm:text-8xl font-extrabold mb-8">
        {{ titles[currentIndex] }}
      </h1>

      <p class="text-xl font-swap font-bold sm:text-2xl text-black mb-10">
        {{ descriptions[currentIndex] }}
      </p>
    </div>
    <div class="relative z-10 text-center mt-2px">
      <UButton
        class="text-white font-swap text-xl font-semibold p-4 px-7 hover:text-black bg-black"
        to="/shop"
      >
        Shop Now
      </UButton>
    </div>
  </section>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue';

const images = [
  new URL('@/assets/slide1.jpg', import.meta.url).href,
  new URL('@/assets/slide2.jpg', import.meta.url).href,
  new URL('@/assets/slide3.jpg', import.meta.url).href,
];

const titles = ['Welcome to SwapCart!', 'Find What You Love!', 'Swap and Save Today!'];

const descriptions = [
  'A smarter way to exchange your stuff with ease.',
  'Explore items you’ve always wanted — at no cost!',
  'Trade items, save money, and connect with your community.',
];

const currentIndex = ref(0);

let interval: ReturnType<typeof setInterval>;

onMounted(() => {
  interval = setInterval(() => {
    currentIndex.value = (currentIndex.value + 1) % images.length;
  }, 3000);
});

onUnmounted(() => {
  clearInterval(interval);
});
</script>
