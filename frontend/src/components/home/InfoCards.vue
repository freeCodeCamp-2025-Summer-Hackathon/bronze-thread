<template>
  <div class="w-full py-16 px-6 sm:px-12 lg:px-20">
    <div
      class="max-w-7xl mx-auto flex flex-col lg:flex-row items-center sm:justify-between lg:justify-around gap-10"
    >
      <div class="w-full flex justify-center lg:justify-start max-w-sm">
        <img src="../../assets/cta.png" class="w-80 h-90 sm:mr-10 rounded-lg shadow-md" />
      </div>

      <div class="text-center lg:text-left space-x-2 max-w-xl">
        <h2 class="text-3xl sm:text-4xl font-swap font-extrabold text-gray-900">Stylish Bags</h2>
        <p class="text-lg pt-5 font-swap sm:text-xl text-gray-700">
          Elevate your look with our collection of stylish bags. From sleek designs to vibrant
          patterns, find the perfect one to match your personal style. Carry confidence wherever you
          go. Make every outfit unforgettable.
        </p>

        <div class="flex justify-center pt-10 lg:justify-start gap-4">
          <UButton
            class="text-white font-swap text-xl px-6 py-3 rounded-md font-medium hover:bg-white hover:text-swapbase transition bg-black"
            to="/shop"
          >
            Explore
          </UButton>
        </div>
      </div>
    </div>
  </div>
</template>
