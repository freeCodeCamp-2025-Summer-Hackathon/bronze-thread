import AuthView from '@/views/AuthView.vue';
import { createRouter, createWebHistory } from 'vue-router';
import AboutView from '../views/AboutView.vue';
import HomeView from '../views/HomeView.vue';
import ShopView from '@/views/ShopView.vue';

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
    },
    {
      path: '/shop',
      name: 'shop',
      component: ShopView,
    },
    {
      path: '/about',
      name: 'about',
      component: AboutView,
    },
    {
      path: '/signin',
      name: 'signin',
      component: AuthView,
    },

    {
      path: '/product/:id',
      name: 'product-detail',
      component: () => import('@/views/ProductDetailPage.vue'),
      props: true,
    },
    {
      path: '/makeOffer/:id',
      name: 'make-offer',
      component: () => import('@/views/MakeOfferPage.vue'),
      props: true,
    },
    {
      path: '/offerStatus',
      name: 'offerStatus',
      component: () => import('@/views/offerStatus.vue'),
    },
  ],
});

export default router;
