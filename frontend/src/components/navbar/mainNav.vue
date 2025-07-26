<template>
  <header
    class="fixed top-0 right-0 w-full z-50 bg-[#F6E5DC] flex justify-between items-center px-6 py-4 border-b border-gray-200 shadow-sm"
  >
    <div class="flex items-center gap-2">
      <img src="../../assets/swapcartlogo.png" alt="Logo" class="w-8 h-8" />
      <h1 class="text-2xl font-extrabold font-swap text-black">SwapCart</h1>
    </div>

    <div class="hidden lg:flex items-center gap-8 w-full justify-end">
      <nav class="flex items-center space-x-10 text-sm font-swap font-medium text-black">
        <RouterLink
          v-for="btn in buttons"
          :key="btn.to"
          :to="btn.to"
          class="relative inline-flex pb-1 px-10 transition-colors duration-300"
          :class="{
            'text-black': isActive(btn.to),
            'hover:text-swapbase': true,
          }"
        >
          {{ btn.button }}
          <span
            v-if="isActive(btn.to)"
            class="absolute -bottom-0.5 left-1/2 w-1/2 h-[2px] bg-black transition-all duration-300 transform -translate-x-1/2"
          ></span>
        </RouterLink>
      </nav>

      <div>
        <RouterLink
          v-if="user && user.name"
          to="/profile"
          class="w-10 h-10 bg-pink-200 rounded-full px-4 flex items-center justify-center text-xs font-bold text-black"
        >
          {{ user.name }}
        </RouterLink>
        <RouterLink
          v-else
          to="/signin"
          class="w-12 h-12 bg-pink-200 rounded-full px-4 flex items-center justify-center text-xs font-bold text-black"
        >
          LogIn
        </RouterLink>
      </div>
    </div>

    <button @click="isMobileMenuOpen = !isMobileMenuOpen" class="lg:hidden focus:outline-none">
      <svg
        v-if="!isMobileMenuOpen"
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6 text-black"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
      <svg
        v-else
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6 text-black"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M6 18L18 6M6 6l12 12"
        />
      </svg>
    </button>

    <Transition name="slide">
      <div
        v-if="isMobileMenuOpen"
        class="fixed top-16 right-0 w-1/2 h-screen bg-[#F6E5DC] shadow-md lg:hidden flex flex-col px-6 py-2 z-40"
      >
        <div class="flex flex-col">
          <template v-for="(btn, index) in buttons" :key="btn.to">
            <RouterLink
              :to="btn.to"
              class="text-base font-normal hover:text-swapbase font-swap text-black py-8"
              :class="{ 'text-black font-semibold': isActive(btn.to) }"
            >
              {{ btn.button }}
            </RouterLink>
            <hr v-if="index < buttons.length - 1" class="border-gray-300" />
          </template>
        </div>

        <div class="mt-8 border-t border-gray-300 pt-8">
          <RouterLink
            v-if="user && user.name"
            to="/profile"
            class="block font-swap text-md text-black hover:text-amber-700"
            @click="isMobileMenuOpen = false"
          >
            {{ user.name }}
          </RouterLink>
          <RouterLink
            v-else
            to="/signin"
            class="block font-swap text-md text-black hover:text-amber-700"
            @click="isMobileMenuOpen = false"
          >
            <button>Log in</button>
          </RouterLink>
        </div>
      </div>
    </Transition>
  </header>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRoute } from 'vue-router';

const isMobileMenuOpen = ref(false);
const route = useRoute();

const user = ref<Record<string, unknown> | null>(null);

if (typeof window !== 'undefined') {
  const storedUser = localStorage.getItem('user');
  if (storedUser) {
    user.value = JSON.parse(storedUser);
  }
}

const buttons = [
  { button: 'Home', to: '/' },
  { button: 'Shop', to: '/shop' },
  { button: 'About', to: '/about' },
  { button: 'Contact', to: '/contact' },
];


const isActive = (path: string) => route.path === path || route.path.startsWith(path + '/shop');


</script>

<style scoped>
.slide-enter-active,
.slide-leave-active {
  transition: all 0.3s ease;
}
.slide-enter-from {
  transform: translateX(-100%);
  opacity: 0;
}
.slide-enter-to {
  transform: translateX(0);
  opacity: 1;
}
.slide-leave-from {
  transform: translateX(0);
  opacity: 1;
}
.slide-leave-to {
  transform: translateX(-100%);
  opacity: 0;
}
</style>
