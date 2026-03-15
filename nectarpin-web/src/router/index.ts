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
      meta: { requiresAuth: true },
    },
    {
      path: '/',
      redirect: '/admin',
    },
  ],
})

router.beforeEach((to, _from, next) => {
  const userStore = useUserStore()
  const hasToken = userStore.loadFromStorage()
  const isAuthenticated = !!userStore.token

  // 优先检查：已有用户信息，立即放行（最快路径）
  if (userStore.user) {
    next()
    return
  }

  // 有 token 但无用户信息：先放行，后台异步加载（非阻塞）
  if (hasToken && isAuthenticated) {
    next() // ← 立即放行，不阻塞路由导航

    // 后台异步加载用户信息
    userStore.loadUserProfile().catch((error) => {
      console.error('后台加载用户信息失败:', error)
      // Token 可能已过期，清除本地数据
      userStore.clearAuth()
      // 可选：如果当前页面需要认证，跳转到登录页
      // 这里不自动跳转，避免干扰用户操作
    })
  } else if (!isAuthenticated && to.meta.requiresAuth) {
    // 未认证且需要认证，跳转到登录页
    next({ name: 'admin-login', query: { redirect: to.fullPath } })
  } else {
    // 其他情况直接放行
    next()
  }
})

export default router
