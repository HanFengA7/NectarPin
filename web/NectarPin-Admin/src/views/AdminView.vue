<script setup>
import {ref} from 'vue';
import {Message, Notification} from '@arco-design/web-vue';
import {IconCaretLeft, IconCaretRight,} from '@arco-design/web-vue/es/icon';
import { useRouter } from 'vue-router';

const router = useRouter()

const collapsed = ref(false);
const onCollapse = () => {
  collapsed.value = !collapsed.value;
};

function onClickMenuItem(key) {
  Message.info({content: `You select ${key}`, showIcon: true});
}

//退出登录
const Logout = () =>{
  window.sessionStorage.removeItem('token')
  router.push('/Login')
  Notification.success({
    content: "退出登录成功",
    duration: 1000,
    closable: true,
  })
}
</script>

<template>
  <a-layout class="layout">
    <a-layout-sider
        hide-trigger
        collapsible
        :collapsed="collapsed"
    >
      <div id="logo">
        <h1 style="text-align: center;font-weight: 200">NectarPin</h1>
      </div>
      <a-menu
          :defaultSelectedKeys="['0_1']"
          :style="{ width: '100%' }"
          @menuItemClick="onClickMenuItem"
      >
        <a-menu-item key="0_1">
          <icon-dashboard/>
          控制台
        </a-menu-item>
        <a-menu-item key="0_2">
          <icon-idcard/>
          个人中心
        </a-menu-item>
        <a-menu-item key="0_3">
          <icon-pushpin/>
          开始写作
        </a-menu-item>
        <a-sub-menu key="1">
          <template #title>
            <span><icon-pen/>撰写</span>
          </template>
          <a-menu-item key="1_1">撰写文章</a-menu-item>
          <a-menu-item key="1_2">撰写说说</a-menu-item>
          <a-menu-item key="1_3">撰写页面</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="4">
          <template #title>
            <span><icon-relation/>管理</span>
          </template>
          <a-menu-item key="4_1">文章</a-menu-item>
          <a-menu-item key="4_2">页面</a-menu-item>
          <a-menu-item key="4_3">评论</a-menu-item>
          <a-menu-item key="4_4">分类</a-menu-item>
          <a-menu-item key="4_5">用户</a-menu-item>
          <a-menu-item key="4_6">友链</a-menu-item>
          <a-menu-item key="4_7">相册</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="5">
          <template #title>
            <span><icon-apps/>扩展</span>
          </template>
          <a-menu-item key="5_1">验证码</a-menu-item>
          <a-menu-item key="5_2">SEO优化</a-menu-item>
          <a-menu-item key="5_3">定时器</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="6">
          <template #title>
            <span><icon-palette/>主题</span>
          </template>
          <a-menu-item key="6_1">主题控制器</a-menu-item>
          <a-menu-item key="6_2">主题样式</a-menu-item>
        </a-sub-menu>

        <a-sub-menu key="7">
          <template #title>
            <span><icon-public/>全局</span>
          </template>
          <a-menu-item key="7_1">基本设置</a-menu-item>
          <a-menu-item key="7_2">导航设置</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="8">
          <template #title>
            <span><icon-bulb/>系统</span>
          </template>
          <a-menu-item key="8_1">关于</a-menu-item>
          <a-menu-item key="7_3">设置</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="padding-left: 20px;">
        <a-button shape="round" @click="onCollapse">
          <IconCaretRight v-if="collapsed"/>
          <IconCaretLeft v-else/>
        </a-button>
        <a-button shape="round" style="margin-left: 15px"><icon-sync /></a-button>

        <div style="float: right; margin-right: 20px">

          <a-dropdown position="br">
            <a-avatar>
              <img
                  alt="avatar"
                  src="https://q1.qlogo.cn/g?b=qq&nk=1091044631&s=640"
              />
            </a-avatar>
            <template #content>
              <a-doption @click="Logout">
                <template #icon>
                  <icon-export />
                </template>
                <template #default>退出登录</template>
              </a-doption>
              <a-doption>
                <template #icon>
                  <icon-idcard />
                </template>
                <template #default>个人中心</template>
              </a-doption>
            </template>
          </a-dropdown>
        </div>
      </a-layout-header>
      <a-layout>

        <RouterView></RouterView>

        <a-layout-footer>
          <div class="footer">
            <div class="footer-copyright">
              NectarPin © 2015-2024 LychApe
            </div>
          </div>
        </a-layout-footer>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<style scoped>
.layout {
  height: 100%;
  background: var(--color-fill-2);
  border: 1px solid var(--color-border);
}

.layout :deep(.arco-layout-sider) .logo {
  height: 32px;
  margin: 12px 8px;
  background: rgba(255, 255, 255, 0.2);
}

.layout :deep(.arco-layout-header) {
  height: 64px;
  line-height: 64px;
  background: var(--color-bg-3);
}

.layout :deep(.arco-layout-footer) {
  height: 48px;
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  line-height: 48px;
}

.layout :deep(.arco-layout-content) {
  color: var(--color-text-2);
  font-weight: 400;
  font-size: 14px;
  background: var(--color-bg-3);
}

.layout :deep(.arco-layout-footer),
.layout :deep(.arco-layout-content) {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
}

.footer{
  margin-top: 55px;
  display: flex;
  justify-content: center;
  box-sizing: border-box;
  width: 100%;
  padding: 20px 20px 0;
  background-color: #f7f8fa;
}
.footer-copyright {
  margin-left: 12px;
  color: var(--color-text-3);
  font-size: 12px;
  line-height: 12px;
}
</style>
