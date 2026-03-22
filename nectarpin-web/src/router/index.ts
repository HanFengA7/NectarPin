import { createRouter, createWebHistory } from 'vue-router'

import { isAdminAuthenticated } from '@/lib/admin-auth'
import { applyDocumentTitleFromRoute } from '@/lib/page-title'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: () => import('@/components/index/IndexLayout.vue'),
      children: [
        {
          path: '',
          name: 'index',
          meta: { title: '首页' },
          component: () => import('@/pages/index/IndexView.vue'),
        },
        {
          path: 'posts/:slug',
          name: 'index-post-detail',
          meta: { titleFromPage: true },
          component: () => import('@/pages/index/PostDetailView.vue'),
        },
        {
          path: 'articles',
          name: 'index-articles',
          meta: { title: '文章' },
          component: () => import('@/pages/index/IndexView.vue'),
        },
        {
          path: 'projects',
          name: 'index-projects',
          meta: { title: '项目' },
          component: () => import('@/pages/index/IndexView.vue'),
        },
        {
          path: 'friends',
          name: 'index-friends',
          meta: { title: '友链' },
          component: () => import('@/pages/index/IndexFriendsView.vue'),
        },
        {
          path: 'about',
          name: 'index-about',
          meta: { title: '关于' },
          component: () => import('@/pages/index/IndexView.vue'),
        },
      ],
    },
    {
      path: '/admin/login',
      name: 'admin-login',
      component: () => import('@/pages/admin/AdminLoginView.vue'),
      meta: {
        guestOnly: true,
        title: '登录',
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
          meta: { title: '控制台' },
          component: () => import('@/pages/admin/DashboardView.vue'),
        },
        {
          path: 'profile',
          name: 'admin-profile',
          meta: { title: '个人资料' },
          component: () => import('@/pages/admin/AdminProfileView.vue'),
        },
        {
          path: 'articles',
          name: 'admin-articles',
          meta: { title: '文章' },
          component: () => import('@/pages/admin/ArticlesView.vue'),
        },
        {
          path: 'articles/new',
          name: 'admin-article-create',
          meta: { title: '写文章' },
          component: () => import('@/pages/admin/ArticleEditorView.vue'),
        },
        {
          path: 'articles/:id/edit',
          name: 'admin-article-edit',
          meta: { title: '编辑文章' },
          component: () => import('@/pages/admin/ArticleEditorView.vue'),
        },
        {
          path: 'articleCategories',
          name: 'admin-article-categories',
          meta: { title: '分类' },
          component: () => import('@/pages/admin/ArticleCategoriesView.vue'),
        },
        {
          path: 'articleTags',
          name: 'admin-article-tags',
          meta: { title: '标签' },
          component: () => import('@/pages/admin/ArticleTagsView.vue'),
        },
        {
          path: 'links',
          name: 'admin-friend-links',
          meta: { title: '友链列表' },
          component: () => import('@/pages/admin/FriendLinksAdminView.vue'),
        },
        {
          path: 'linkCategories',
          name: 'admin-link-page-settings',
          meta: { title: '页面设置' },
          component: () => import('@/pages/admin/LinkPageSettingsView.vue'),
        },
        {
          path: 'global/siteSettings/index',
          name: 'global-siteSettings-index',
          meta: { title: '站点首页' },
          component: () => import('@/pages/admin/SiteHomeSettingsView.vue'),
        },
        {
          path: 'global/siteSettings/footer',
          name: 'global-siteSettings-footer',
          meta: { title: '页脚设置' },
          component: () => import('@/pages/admin/SiteFooterSettingsView.vue'),
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

router.afterEach((to) => {
  applyDocumentTitleFromRoute(to)
})

export default router
