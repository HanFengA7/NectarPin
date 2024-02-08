<script lang="ts" setup>
import eventBus from '@/plugin/event-bus/event-bus';
import {reactive, ref} from "vue";
import router from "@/router";
import dayjs from 'dayjs';
import {EditUserInfo} from '@/api/User/user'
import {Message} from "@arco-design/web-vue";

/*
接收父组件的数据
[1][UserInfoData]
 */
const props = defineProps(['userInfo']);

/*
传数据给父组件
[1]设置侧边栏选择选项 [child-data-selectedKeys]
*/
let SelectedKeys: any = ref(["PersonalCenter"]);
eventBus.emit("child-data-selectedKeys", SelectedKeys);

//返回Dashboard的函数
const personalCenterHeardOnBack = () =>{
  router.push({ name: 'Dashboard' })
}


/*
编辑资料模态框 [Start]
 */
//编辑资料模态框的函数[公共]
//定义editInfo_Form
const editInfo_Form = reactive({
  username: props.userInfo["username"],
  nickname: props.userInfo["nickname"],
  email: props.userInfo["email"],
  avater_url: props.userInfo["avater_url"],
});
//编辑资料模态框的函数[PC]
const editInfo_Visible_PC = ref(false);
const editInfo_handleClickPC = () => {
  editInfo_Visible_PC.value = true;
};
const editInfo_handleBeforeOk_PC = (done:any) => {
  EditUserInfo(props.userInfo["id"],editInfo_Form).then((res:any) => {
    if (res.data.code == 200){
      window.setTimeout(() => {
        //刷新UserInfo数据 [child-data-userInfo-refresh]
        eventBus.emit("child-data-userInfo-refresh", new Date());
        done()
        Message.success({content: res.data.msg, showIcon: true});
      }, 3000)
    }else {
      window.setTimeout(() => {
        editInfo_Form.username = props.userInfo["username"]
        editInfo_Form.nickname = props.userInfo["nickname"]
        editInfo_Form.email = props.userInfo["email"]
        editInfo_Form.avater_url = props.userInfo["avater_url"]
        done()
        Message.error({content: res.data.msg, showIcon: true});
      }, 3000)
    }
  })
};
const editInfo_handleCancel_PC = () => {
  editInfo_Visible_PC.value = false;
}
/*
编辑资料模态框 [End]
 */
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
          <a-button @click="editInfo_handleClickPC">编辑资料</a-button>
          <a-button>修改密码</a-button>
        </a-space>
      </template>
    </a-page-header>
  </div>

  <!--[编辑资料][模态框][PC]-->
  <a-modal v-model:visible="editInfo_Visible_PC" title="编辑资料" @cancel="editInfo_handleCancel_PC" @before-ok="editInfo_handleBeforeOk_PC">
    <a-form :model="editInfo_Form">
      <a-form-item field="username" label="用户名">
        <a-input v-model="editInfo_Form.username" />
      </a-form-item>
      <a-form-item field="nickname" label="昵称">
        <a-input v-model="editInfo_Form.nickname" />
      </a-form-item>
      <a-form-item field="email" label="邮箱">
        <a-input v-model="editInfo_Form.email" />
      </a-form-item>
      <a-form-item field="avater_url" label="头像外链">
        <a-input v-model="editInfo_Form.avater_url" />
      </a-form-item>
    </a-form>
  </a-modal>

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
          <p>权限 : {{props.userInfo["role"] == "0" ? "管理员" : "编辑者"}}</p>
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
  margin: 50px;
  padding: 25px;
  background: #ffffff;
  height: 550px;
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