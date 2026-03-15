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

router.beforeEach(async (to, _from, next) => {
  const userStore = useUserStore()
  const hasToken = userStore.loadFromStorage()

  const isAuthenticated = !!userStore.token

  // 如果有 token 但没有用户信息，尝试从服务端加载
  if (hasToken && !userStore.user && to.meta.requiresAuth !== false) {
    try {
      await userStore.loadUserProfile()
    } catch (error) {
      console.error('加载用户信息失败:', error)
      // Token 可能已过期，清除本地数据
      userStore.clearAuth()
      if (to.meta.requiresAuth !== false) {
        next({ name: 'admin-login', query: { redirect: to.fullPath } })
        return
      }
    }
  }

  if (to.meta.requiresAuth && !isAuthenticated) {
    next({ name: 'admin-login', query: { redirect: to.fullPath } })
  } else if (to.name === 'admin-login' && isAuthenticated) {
    next({ name: 'admin-dashboard' })
  } else {
    next()
  }
})

export default router
