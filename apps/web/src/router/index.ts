import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    component: () => import('@/layouts/DefaultLayout.vue'),
    children: [
      { path: '', name: 'Home', component: () => import('@/pages/HomePage.vue') },
      { path: 'agents', name: 'AgentList', component: () => import('@/pages/AgentListPage.vue') },
      { path: 'agents/:slug', name: 'AgentDetail', component: () => import('@/pages/AgentDetailPage.vue') },
      { path: 'leaderboard', name: 'Leaderboard', component: () => import('@/pages/LeaderboardPage.vue') },
      { path: 'compare', name: 'Compare', component: () => import('@/pages/ComparePage.vue') },
      { path: 'login', name: 'Login', component: () => import('@/pages/LoginPage.vue') },
      { path: 'profile', name: 'Profile', component: () => import('@/pages/ProfilePage.vue') },
    ],
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes,
})

export default router
