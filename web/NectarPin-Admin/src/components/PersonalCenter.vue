<script lang="ts" setup>
import eventBus from '@/plugin/event-bus/event-bus';
import {ref} from "vue";
import router from "@/router";
import dayjs from 'dayjs';
//接收父组件的UserInfoData数据
const props = defineProps(['userInfo']);
//传数据给父组件
/*设置侧边栏选择选项*/
let SelectedKeys: any = ref(["PersonalCenter"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);

const personalCenterHeardOnBack = () =>{
  router.push({ name: 'Dashboard' })
}
</script>

<template>

  <div class="personalCenter-heard-PC">
    <a-page-header
        title="个人中心"
        subtitle="PersonalCenter"
        :show-back="true"
        @back = personalCenterHeardOnBack
    >
      <template #breadcrumb>
        <a-breadcrumb>
          <a-breadcrumb-item>后台</a-breadcrumb-item>
          <a-breadcrumb-item>个人中心</a-breadcrumb-item>
        </a-breadcrumb>
      </template>

      <template #extra>
        <a-space>
          <a-button>编辑资料</a-button>
          <a-button>修改密码</a-button>
        </a-space>
      </template>
    </a-page-header>
  </div>
  <a-watermark
      :content="[props.userInfo.username,dayjs().format('YYYY-MM-DD H:mm:ss')]"
      :alpha="0.25"
      class="personalCenter-heardBox-PC">

    <div class="personalCenter-heardBox-avatar-PC">
      <a-avatar
          :size="128"
          trigger-type="mask"
      >
        <img
            alt="avatar"
            :src=props.userInfo.avater_url
        />
        <template #trigger-icon>
          <IconEdit />
        </template>
      </a-avatar>
      <h2>HanFengA7</h2>
      <h3>Stay Hungry, Stay Foolish.</h3>
    </div>

    <a-row class="personalCenter-heardBox-Card-row1-PC">
      <a-col :span="12">
        <div class="personalCenter-heardBox-Card1-PC">
          <p>用户名 : {{props.userInfo["username"]}}</p>
          <p>昵称 : {{props.userInfo["nickname"]}}</p>
          <p>邮箱 : {{props.userInfo["email"]}}</p>
          <p>权限 : {{props.userInfo["role"] == "1" ? "超级管理员" : "管理员"}}</p>
        </div>
      </a-col>
      <a-col :span="12">
        <div class="personalCenter-heardBox-Card2-PC">
          <p>上次登录时间 : {{props.userInfo["last_longin_date"]}}</p>
          <p>上次登录IP : {{props.userInfo["last_longin_ip_address"]}}</p>
        </div>
      </a-col>
    </a-row>

  </a-watermark>
</template>

<style scoped>
.personalCenter-heard-PC {
  background: #ffffff;
  padding: 0 50px 5px 25px;
  height: 100px;
}
.personalCenter-heardBox-PC{
  border-radius: 15px;
  margin: 40px;
  padding: 25px;
  background: #ffffff;
  height: 650px;
}

.personalCenter-heardBox-avatar-PC{
  text-align: center;
}
.personalCenter-heardBox-avatar-PC h2,h3{
  font-weight: lighter;
}

.personalCenter-heardBox-Card-row1-PC{
  margin: 35px;
}

.personalCenter-heardBox-Card1-PC{
  border-radius: 10px;
  margin-right: 25px;
  padding: 20px;
  background: #f2f3f5d4;
}
.personalCenter-heardBox-Card2-PC{
  border-radius: 10px;
  margin-left: 25px;
  padding: 20px;
  background: #f2f3f5d4;
}
</style>