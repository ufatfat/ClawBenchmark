import { createRouter, createWebHistory } from 'vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    component: () => import('@/layouts/MainLayout.vue'),
    children: [
      { path: '', name: 'Home', component: () => import('@/views/HomeView.vue') },
      { path: 'benchmarks', name: 'Benchmarks', component: () => import('@/views/BenchmarkListView.vue') },
      { path: 'benchmarks/:id', name: 'BenchmarkDetail', component: () => import('@/views/BenchmarkDetailView.vue') },
      { path: 'leaderboard', name: 'Leaderboard', component: () => import('@/views/LeaderboardView.vue') },
      { path: 'models', name: 'Models', component: () => import('@/views/ModelListView.vue') },
      { path: 'reports', name: 'Reports', component: () => import('@/views/ReportView.vue') },
    ],
  },
  { path: '/login', name: 'Login', component: () => import('@/views/LoginView.vue') },
  { path: '/:pathMatch(.*)*', name: 'NotFound', component: () => import('@/views/NotFoundView.vue') },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  const publicRoutes = ['Login', 'NotFound']
  if (!publicRoutes.includes(to.name as string) && !auth.isLoggedIn()) {
    return { name: 'Login' }
  }
})

export default router
