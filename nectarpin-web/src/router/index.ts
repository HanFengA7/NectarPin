import { createRouter, createWebHistory } from 'vue-router'

import { isAdminAuthenticated } from '@/lib/admin-auth'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      redirect: '/admin/login',
    },
    {
      path: '/admin/login',
      name: 'admin-login',
      component: () => import('@/pages/admin/AdminLoginView.vue'),
      meta: {
        guestOnly: true,
      },
    },
    {
      path: '/admin',
      component: () => import('@/components/admin/AdminLayout.vue'),
      meta: {
        requiresAuth: true,
      },
      children: [
        {
          path: '',
          redirect: '/admin/dashboard',
        },
        {
          path: 'dashboard',
          name: 'admin-dashboard',
          component: () => import('@/pages/admin/DashboardView.vue'),
        },
        {
          path: 'articles',
          name: 'admin-articles',
          component: () => import('@/pages/admin/ArticlesView.vue'),
        },
        {
          path: 'articles/new',
          name: 'admin-article-create',
          component: () => import('@/pages/admin/ArticleEditorView.vue'),
        },
        {
          path: 'articles/:id/edit',
          name: 'admin-article-edit',
          component: () => import('@/pages/admin/ArticleEditorView.vue'),
        },
      ],
    },
  ],
})

router.beforeEach((to) => {
  const authenticated = isAdminAuthenticated()

  if (to.meta.requiresAuth && !authenticated) {
    return {
      name: 'admin-login',
      query: {
        redirect: to.fullPath,
      },
    }
  }

  if (to.meta.guestOnly && authenticated) {
    return {
      name: 'admin-dashboard',
    }
  }

  return true
})

export default router
