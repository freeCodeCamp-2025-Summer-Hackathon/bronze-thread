<template>
  <footer class="bg-[#F6E5DC] text-gray-700 px-6 py-12">
    <div class="max-w-7xl mx-auto grid grid-cols-1 pl-10 sm:grid-cols-2 lg:grid-cols-4 gap-10">
      <div class="flex flex-col items-start gap-4">
        <div class="w-10 h-10">
          <img
            src="../../assets/swapcartlogo.png"
            alt="Logo"
            class="w-full h-full object-contain"
          />
        </div>
        <p class="text-sm text-gray-600">© 2025 SwapCart.<br />All rights reserved.</p>
      </div>

      <div>
        <h4 class="font-extrabold text-lg font-swap mb-2">Quick Links</h4>
        <ul class="font-swap space-y-1">
          <li><a href="#" class="hover:underline">Home</a></li>
          <li><a href="/about" class="hover:underline">About</a></li>
          <li><a href="/shop" class="hover:underline">Shop</a></li>
          <li><a href="/contact" class="hover:underline">Contact</a></li>
        </ul>
      </div>

      <div>
        <h4 class="font-extrabold text-lg font-swap mb-2">Resources</h4>
        <ul class="font-swap space-y-1">
          <li><a href="/faq" class="hover:underline">FAQ</a></li>
          <li><a href="/sellerGuide" class="hover:underline">Seller Guide</a></li>
          <li><a href="/returnPolicy" class="hover:underline">Return Policy</a></li>
          <li><a href="/privacyPolicy" class="hover:underline">Privacy Policy</a></li>
        </ul>
      </div>

      <div class="font-bold mb-2">
        <form class="flex flex-col sm:flex-row items-center gap-3">
          <input
            type="email"
            placeholder="Enter your email"
            required
            class="w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-gray-500"
          />
          <button
            type="submit"
            class="px-6 py-2 bg-black text-white text-sm rounded-md hover:bg-white hover:text-swapbase transition"
          >
            Subscribe
          </button>
        </form>
      </div>
    </div>
  </footer>
</template>
