import {createRouter, createWebHistory} from 'vue-router'
import AdminView from '../views/AdminView.vue'
import {Login_CheckToken} from "@/api/User/login";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/Login',
            name: 'Login',
            meta: {
                title: 'NectarPin - 身份认证',
            },
            component: () => import('../views/LoginView.vue')
        },
        {
            path: '',
            name: 'Admin',
            meta: {
                title: 'NectarPin',
            },
            component: AdminView,
            //props: route => ({ userInfo: route.params.userInfo }),
            children: [
                {
                    path: '/Dashboard',
                    name: 'Dashboard',
                    meta: {
                        title: 'NectarPin - 仪表盘',
                    },
                    component: () => import('../components/Dashboard.vue')
                },
                {
                    path: '/PersonalCenter',
                    name: 'PersonalCenter',
                    meta:{
                        title: 'NectarPin - 个人中心'
                    },
                    component: () => import('../components/PersonalCenter.vue')
                },
                // [Article]
                {
                    path:'/Article/add',
                    name: 'Article/add',
                    meta:{
                        title: 'NectarPin - 撰写文章'
                    },
                    component: () => import('../components/Article/add.vue')
                },
                {
                    path:'/Article',
                    name: 'Article',
                    meta:{
                        title: 'NectarPin - 文章管理'
                    },
                    component: () => import('../components/Article/article.vue')
                },
            ],
        }
    ]
})

router.beforeEach(async (to, from) => {
    // 设置标题
    if (to.meta.title) {
        document.title = to.meta.title;
    }

    // 检查token是否存在  && 检查token有效性
    const token: string = window.localStorage.getItem("token");
    let tokenIF: boolean = false;
    if (token) {
        try {
            const res = await Login_CheckToken(token);
            tokenIF = res.data["tokenBool"];
        } catch {
            tokenIF = false;
        }
    }

    // 重定向->Login [未登录]
    if (!tokenIF && to.name !== 'Login') {
        return {name: 'Login'}
    }
    // 重定向->Dashboard [已登录]
    if (tokenIF && to.name === 'Login' || tokenIF && to.name === 'Admin') {
        return {name: 'Dashboard'}
    }
})


export default router