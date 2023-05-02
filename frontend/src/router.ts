import { RouteRecordRaw, createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    component: () => import('./pages/Connection.vue'),
  },
  {
    path: '/home',
    component: () => import('./pages/Home.vue'),
  },
  {
    path: '/query',
    component: () => import('./pages/Query.vue'),
  },{
    path: '/connections',
    component: () => import('./pages/ListConnection.vue'),
  }
] as RouteRecordRaw[];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
