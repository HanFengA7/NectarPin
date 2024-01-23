<script>
import { defineComponent, ref } from 'vue';
import { Message} from '@arco-design/web-vue';
import {
  IconCaretRight,
  IconCaretLeft,
  IconHome,
  IconCalendar,
} from '@arco-design/web-vue/es/icon';

export default defineComponent({
  components: {
    IconCaretRight,
    IconCaretLeft,
    IconHome,
    IconCalendar,
  },
  setup() {
    const collapsed = ref(false);
    const onCollapse = () => {
      collapsed.value = !collapsed.value;
    };
    return {
      collapsed,
      onCollapse,
      onClickMenuItem(key) {
        Message.info({ content: `You select ${key}`, showIcon: true });
      }
    };
  },
});
</script>

<template>
  <a-layout class="layout">
    <a-layout-sider
        hide-trigger
        collapsible
        :collapsed="collapsed"
    >
      <div>
        <h1 style="text-align: center;font-weight: 200">NectarPin</h1>
      </div>
      <a-menu
          :defaultSelectedKeys="['0_1']"
          :style="{ width: '100%' }"
          @menuItemClick="onClickMenuItem"
      >
        <a-menu-item key="0_1" >
          <icon-dashboard />
          控制台
        </a-menu-item>
        <a-menu-item key="0_2">
          <icon-idcard />
          个人中心
        </a-menu-item>
        <a-menu-item key="0_3">
          <icon-pushpin />
          开始写作
        </a-menu-item>
        <a-sub-menu key="1">
          <template #title>
            <span><icon-pen />撰写</span>
          </template>
          <a-menu-item key="1_1">撰写文章</a-menu-item>
          <a-menu-item key="1_2">撰写说说</a-menu-item>
          <a-menu-item key="1_3">撰写页面</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="4">
          <template #title>
            <span><icon-relation />管理</span>
          </template>
          <a-menu-item key="4_1">文章</a-menu-item>
          <a-menu-item key="4_2">页面</a-menu-item>
          <a-menu-item key="4_3">评论</a-menu-item>
          <a-menu-item key="4_4">分类</a-menu-item>
          <a-menu-item key="4_5">友链</a-menu-item>
          <a-menu-item key="4_6">相册</a-menu-item>
        </a-sub-menu>
        <a-sub-menu key="5">
          <template #title>
            <span><icon-public />全局设置</span>
          </template>
          <a-menu-item key="5_1">基本设置</a-menu-item>
          <a-menu-item key="5_2">SEO优化</a-menu-item>
        </a-sub-menu>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="padding-left: 20px;">
        <a-button shape="round" @click="onCollapse">
          <IconCaretRight v-if="collapsed" />
          <IconCaretLeft v-else />
        </a-button>
      </a-layout-header>
      <a-layout style="padding: 0 24px;">

          <RouterView></RouterView>

        <a-layout-footer>Footer</a-layout-footer>
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
.layout :deep(.arco-layout-header)  {
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
.layout :deep(.arco-layout-content)  {
  display: flex;
  flex-direction: column;
  justify-content: center;
  color: var(--color-white);
  font-size: 16px;
  font-stretch: condensed;
  text-align: center;
}
</style>
