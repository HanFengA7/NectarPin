<script lang="ts" setup>
import {onBeforeMount, provide, reactive, ref} from 'vue';
import {Message, Notification} from '@arco-design/web-vue';
import {IconCaretLeft, IconCaretRight,} from '@arco-design/web-vue/es/icon';
import {useRouter} from 'vue-router';
import {GetUserInfo, TokenGetUserInfo} from '@/api/User/user'
import eventBus from "@/plugin/event-bus/event-bus";

const router = useRouter()
const selectedKeys = ref();
const collapsed = ref(false);
const onCollapse = () => {
  collapsed.value = !collapsed.value;
};

//点击侧边栏事件
function onClickMenuItem(key: any) {
  router.push({name: key})
}

//获取用户信息
const token = window.localStorage.getItem("token");
const UserInfoData = reactive({
  "id": "",
  "username": "",
  "nickname": "",
  "email": "",
  "role": "",
  "avater_url": "",
  "p_signatures":"",
  "this_longin_date": "",
  "this_longin_ip_address": "",
  "last_longin_date": "",
  "last_longin_ip_address": "",
})
const isUserInfoData = ref(false)
const LoadUserInfo = () => {
  TokenGetUserInfo(token).then((res: any) => {
    if (res.data.code == 200) {
      GetUserInfo(res.data["tokenData"].id).then((res: any) => {
        UserInfoData.id = res.data.data[0]["id"]
        UserInfoData.username = res.data.data[0]["username"]
        UserInfoData.nickname = res.data.data[0]["nickname"]
        UserInfoData.email = res.data.data[0]["email"]
        UserInfoData.role = res.data.data[0]["role"]
        UserInfoData.avater_url = res.data.data[0]["avater_url"]
        UserInfoData.p_signatures = res.data.data[0]["p_signatures"]
        UserInfoData.this_longin_date = res.data.data[0]["this_longin_date"]
        UserInfoData.this_longin_ip_address = res.data.data[0]["this_longin_ip_address"]
        UserInfoData.last_longin_date = res.data.data[0]["last_longin_date"]
        UserInfoData.last_longin_ip_address = res.data.data[0]["last_longin_ip_address"]
        isUserInfoData.value = true
      })
    } else {
      isUserInfoData.value = false
      Message.error({content: res.data.msg, showIcon: true});
    }
  })
}
LoadUserInfo()
//provide('UserInfoData', UserInfoData);
//onMounted(LoadUserInfoData);
// watch(
//     () => UserInfoData, // 要监视的表达式
//     (newValue) => {
//       console.log(newValue)
//       provide('UserInfoData', UserInfoData);
//     },
//     {deep: true} // 深度监听对象属性的变化
// )


//退出登录
const Logout = () => {
  window.localStorage.removeItem('token')
  router.push('/Login')
  Notification.success({
    content: "退出登录成功",
    duration: 1000,
    closable: true,
  })
}

const refreshKey = ref(new Date())

//接收子组件数据
onBeforeMount(() => {
  // 监听事件
  eventBus.on('child-data-selectedKeys', (newData: any) => {
    // 处理从子组件接收到的数据
    selectedKeys.value = newData.value;
  });

  eventBus.on('child-data-userInfo-refresh', (refreshKey: any) => {
    //数据更新DataKey
    isUserInfoData.value = false
    LoadUserInfo()
    isUserInfoData.value = true
    refreshKey.value = refreshKey
  });

});



</script>

<template>
  <a-layout class="layout">
    <a-layout-sider
        hide-trigger
        collapsible
        :collapsed="collapsed"
    >
      <div id="logo">
        <h1 style="text-align: center;font-weight: 200">{{ collapsed ? "" : "NectarPin" }}</h1>
      </div>
      <a-menu
          :defaultSelectedKeys="['Dashboard']"
          :selected-keys="selectedKeys"
          :style="{ width: '100%' }"
          @menuItemClick="onClickMenuItem"
      >
        <a-menu-item key="Dashboard">
          <icon-dashboard/>
          控制台
        </a-menu-item>
        <a-menu-item key="PersonalCenter">
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
        <a-button shape="round" style="margin-left: 15px" @click="router.go(0)">
          <icon-sync/>
        </a-button>

        <div style="float: right; margin-right: 20px">

          <a-dropdown position="br">
            <a-avatar>
              <img
                  alt="avatar"
                  :src="UserInfoData.avater_url"
              />
            </a-avatar>
            <template #content>
              <a-doption @click="Logout">
                <template #icon>
                  <icon-export/>
                </template>
                <template #default>退出登录</template>
              </a-doption>
              <a-doption @click="router.push({name: 'PersonalCenter'})">
                <template #icon>
                  <icon-idcard/>
                </template>
                <template #default>个人中心</template>
              </a-doption>
            </template>
          </a-dropdown>
        </div>
      </a-layout-header>
      <a-layout>

        <RouterView :userInfo="UserInfoData" v-if="isUserInfoData" :key="refreshKey"></RouterView>

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

.footer {
  padding: 20px;
  margin-top: 55px;
  display: flex;
  justify-content: center;
  box-sizing: border-box;
  width: 100%;
  background-color: #f7f8fa;
}

.footer-copyright {
  margin-left: 12px;
  color: var(--color-text-3);
  font-size: 12px;
  line-height: 12px;
}
</style>