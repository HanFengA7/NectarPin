import { createRouter, createWebHistory } from 'vue-router'
import AdminView from '../views/AdminView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/Login',
      name: 'NectarPin - Login',
      meta: {
        title: 'NectarPin - 身份认证',
      },
      component:() => import('../views/LoginView.vue')
    },
    {
      path: '',
      name: 'NectarPin - Admin',
      meta: {
        title: 'NectarPin',
      },
      component: AdminView,
      children: [
        {
          path: '/Dashboard',
          name: 'NectarPin - Dashboard',
          meta: {
            title: 'NectarPin - 仪表盘',
          },
          component:() => import('../components/Dashboard.vue')
        }
      ],
    }
  ]
})

router.beforeEach((to, from, next) => {
  // 设置标题
  if (to.meta.title) {
    document.title = to.meta.title;
  }

  // 检查是否存在token
  const token = window.sessionStorage.getItem("token");
  // 检查token有效性

  // 登录页面重定向
  const redirectToIndex = () => {
    if (token) {
      next("/Index");
    } else {
      next("/Login");
    }
  };

  // 需要登录才能访问的页面
  const requiresAuth = to.matched.some(record => record.meta.requiresAuth);

  if (requiresAuth) {
    if (!token) {
      // 未登录，跳转到登录页面
      next("/login");
    } else {
      // 已登录，继续导航
      next();
    }
  } else if (to.path === "/login" || to.path === "/login/") {
    // 防止已登录用户访问登录页面
    redirectToIndex();
  } else if (to.path === "" || to.path === "/") {
    // 根路径重定向
    redirectToIndex();
  } else {
    // 其他情况直接导航
    next();
  }
});

export default router
