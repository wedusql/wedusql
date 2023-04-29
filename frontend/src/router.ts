import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    component: () => import('./pages/Connection.vue'),
  },
  {
    path: '/query',
    component: () => import('./pages/Query.vue'),
  },
] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
