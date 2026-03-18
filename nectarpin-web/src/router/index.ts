import { createRouter, createWebHistory } from 'vue-router'
import { useUserStore } from '@/stores/user'
import AdminLayout from '@/components/admin/AdminLayout.vue'

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
      component: AdminLayout,
      meta: { requiresAuth: true },
      children: [
        {
          path: '',
          name: 'admin-dashboard',
          component: () => import('../views/admin/DashboardView.vue'),
        },
        {
          path: 'articles',
          name: 'admin-articles',
          component: () => import('../views/admin/ArticleListView.vue'),
        },
        {
          path: 'articles/add',
          name: 'admin-article-add',
          component: () => import('../views/admin/ArticleEditorView.vue'),
        },
        {
          path: 'articles/edit',
          name: 'admin-article-edit',
          component: () => import('../views/admin/ArticleEditorView.vue'),
        },
      ],
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

  if (userStore.user) {
    next()
    return
  }

  if (hasToken && isAuthenticated) {
    next()

    userStore.loadUserProfile().catch((error) => {
      console.error('后台加载用户信息失败:', error)
      userStore.clearAuth()
    })
  } else if (!isAuthenticated && to.meta.requiresAuth) {
    next({ name: 'admin-login', query: { redirect: to.fullPath } })
  } else {
    next()
  }
})

export default router
