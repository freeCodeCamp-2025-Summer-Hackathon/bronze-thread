<!-- <script setup lang="ts"></script>

<template>
  <main>This is home page.</main>
  <ULink as="button" to="/signin" class="bg-black text-primary px-2 py-1 rounded-lg"
    >Go to Auth Page</ULink
  >
</template>  -->
<script setup lang="ts">
import Navbar from '@/components/navbar/mainNav.vue';
import HeroCarousel from '@/components/home/HeroSection.vue';
import InfoCards from '@/components/home/InfoCards.vue';
import ExploreItems from '@/components/home/ExploreItems.vue';
import FooterNav from '@/components/footer/FooterNav.vue';
</script>
<template>
  <div class="flex w-full h-auto flex-col gap-2">
    <Navbar />
    <HeroCarousel />
    <InfoCards />
    <ExploreItems />
    <FooterNav />
  </div>
</template>
