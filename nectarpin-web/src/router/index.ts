import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/admin/login',
      name: 'admin-login',
      component: () => import('../views/admin/LoginView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/admin',
      name: 'admin-dashboard',
      component: () => import('../views/admin/DashboardView.vue'),
      meta: { requiresAuth: false },
    },
    {
      path: '/',
      redirect: '/admin',
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const userStore = useUserStore()
  userStore.loadFromStorage()

  const isAuthenticated = !!userStore.token

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'admin-login', query: { redirect: to.fullPath } })
  } else if (to.name === 'admin-login' && isAuthenticated) {
    next({ name: 'admin-dashboard' })
  } else {
    next()
  }
})

export default router
