<script lang="ts" setup>
import eventBus from '@/plugin/event-bus/event-bus';
import {reactive, ref} from "vue";
import router from "@/router";
import dayjs from 'dayjs';
import {EditUserInfo, EditUserPwd} from '@/api/User/user'
import {Message} from "@arco-design/web-vue";
import {md5} from "js-md5";

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
const personalCenterHeardOnBack = () => {
  router.push({name: 'Dashboard'})
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
  p_signatures: props.userInfo["p_signatures"],
});
//编辑资料模态框的函数[PC]
const editInfo_Visible_PC = ref(false);
const editInfo_handleClickPC = () => {
  editInfo_Visible_PC.value = true;
};
const editInfo_handleBeforeOk_PC = (done: any) => {
  EditUserInfo(props.userInfo["id"], editInfo_Form).then((res: any) => {
    if (res.data.code == 200) {
      window.setTimeout(() => {
        //刷新UserInfo数据 [child-data-userInfo-refresh]
        eventBus.emit("child-data-userInfo-refresh", new Date());
        done()
        Message.success({content: res.data.msg, showIcon: true});
      }, 3000)
    } else {
      window.setTimeout(() => {
        editInfo_Form.username = props.userInfo["username"]
        editInfo_Form.nickname = props.userInfo["nickname"]
        editInfo_Form.email = props.userInfo["email"]
        editInfo_Form.avater_url = props.userInfo["avater_url"]
        editInfo_Form.p_signatures = props.userInfo["p_signatures"]
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



/*
修改密码模态框 [Start]
 */

/*
  编辑资料模态框的函数[公共]
  [1]定义editPwd_Form
  [2]修改密码表单规则
 */
const editPwd_Form = reactive({
  old_password: '',
  new_password: '',
  confirm_password: '',
});
const editPwd_Form_Rules = {
  old_password: [
    {
      required: true,
      message: '请输入原始密码',
    },
  ],
  new_password: [
    {
      required: true,
      message: '请输入新密码',
    },
    {
      validator: (value: any, cb: any) => {
        if (value === editPwd_Form.old_password) {
          cb('不能与原始密码相同')
        } else {
          cb()
        }
      }
    }
  ],
  confirm_password: [
    {
      required: true,
      message: '请输入确认密码',
    },
    {
      validator: (value: any, cb: any) => {
        if (value !== editPwd_Form.new_password) {
          cb('两次密码不一致')
        } else {
          cb()
        }
      }
    }
  ],
}
const reset_editPwd_Form = () => {
  editPwd_Form.old_password = '';
  editPwd_Form.new_password = '';
  editPwd_Form.confirm_password = '';
};

//修改密码模态框的函数[PC]

const editPwd_Visible_PC = ref(false);
const editPwd_handleClickPC = () => {
  editPwd_Visible_PC.value = true;
};
const editPwd_handleBeforeOk_PC = (done: any) => {
  if (editPwd_Form.old_password != "" && editPwd_Form.new_password != "" && editPwd_Form.confirm_password != "") {
    if (editPwd_Form.old_password === editPwd_Form.new_password) {
      Message.error({content: "不能与原始密码相同", showIcon: true});
      done(false)
      return
    }
    if (editPwd_Form.new_password !== editPwd_Form.confirm_password) {
      Message.error({content: "两次密码不一致", showIcon: true});
      done(false)
      return
    }

    editPwd_Form.old_password = md5(editPwd_Form.old_password)
    editPwd_Form.new_password = md5(editPwd_Form.new_password)
    editPwd_Form.confirm_password = md5(editPwd_Form.confirm_password)
    EditUserPwd(props.userInfo["username"], editPwd_Form.old_password, editPwd_Form.new_password).then((res: any) => {
      if (res.data.code === 200) {
        window.setTimeout(() => {
          Message.success({content: res.data.msg, showIcon: true});
          done()
          reset_editPwd_Form()
          return
        }, 2000)
      } else {
        window.setTimeout(() => {
          Message.error({content: res.data.msg, showIcon: true});
          console.log(editPwd_Form)
          done(false)
          reset_editPwd_Form()
          return
        }, 2000)
      }
    })

  } else {
    Message.error({content: "请勿提交空值", showIcon: true});
    done(false)
    return
  }

};
const editPwd_handleCancel_PC = () => {
  editPwd_Visible_PC.value = false;
  reset_editPwd_Form()
}

/*
修改密码模态框 [End]
 */
</script>

<template>

  <div class="personalCenter-heard-PC">
    <a-page-header
        title="个人中心"
        subtitle="PersonalCenter"
        :show-back="true"
        @back=personalCenterHeardOnBack
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
          <a-button @click="editPwd_handleClickPC">修改密码</a-button>
        </a-space>
      </template>
    </a-page-header>
  </div>

  <!--[编辑资料][模态框][PC][Start]-->
  <a-modal v-model:visible="editInfo_Visible_PC" title="编辑资料" @cancel="editInfo_handleCancel_PC"
           @before-ok="editInfo_handleBeforeOk_PC">
    <a-form :model="editInfo_Form">
      <a-form-item field="username" label="用户名">
        <a-input v-model="editInfo_Form.username"/>
      </a-form-item>
      <a-form-item field="nickname" label="昵称">
        <a-input v-model="editInfo_Form.nickname"/>
      </a-form-item>
      <a-form-item field="email" label="邮箱">
        <a-input v-model="editInfo_Form.email"/>
      </a-form-item>
      <a-form-item field="p_signatures" label="个性签名">
        <a-input v-model="editInfo_Form.p_signatures"/>
      </a-form-item>
      <a-form-item field="avater_url" label="头像外链">
        <a-input v-model="editInfo_Form.avater_url"/>
      </a-form-item>
    </a-form>
  </a-modal>
  <!--[编辑资料][模态框][PC][End]-->

  <!--[修改密码][模态框][PC][Start]-->
  <a-modal
      v-model:visible="editPwd_Visible_PC"
      title="修改密码"
      @cancel="editPwd_handleCancel_PC"
      @before-ok="editPwd_handleBeforeOk_PC"
      :closable=false
  >
    <a-form :model="editPwd_Form" :rules="editPwd_Form_Rules">
      <a-form-item field="old_password" label="原始密码" validate-trigger="blur">
        <a-input-password v-model="editPwd_Form.old_password"/>
      </a-form-item>
      <a-form-item field="new_password" label="新密码" validate-trigger="blur">
        <a-input-password v-model="editPwd_Form.new_password"/>
      </a-form-item>
      <a-form-item field="confirm_password" label="确认密码" validate-trigger="blur">
        <a-input-password v-model="editPwd_Form.confirm_password"/>
      </a-form-item>
    </a-form>
  </a-modal>
  <!--[修改密码][模态框][PC][End]-->

  <a-watermark
      :content="[props.userInfo.username,dayjs().format('YYYY-MM-DD H:mm:ss')]"
      :alpha="0.25"
      class="personalCenter-heardBox-PC">

    <div class="personalCenter-heardBox-avatar-PC">
      <a-avatar
          :size="128"
          trigger-type="mask"
          @click="editInfo_handleClickPC"
      >
        <img
            alt="avatar"
            :src=props.userInfo.avater_url
        />
        <template #trigger-icon>
          <IconEdit/>
        </template>
      </a-avatar>
      <h2>{{ props.userInfo["nickname"] }}</h2>
      <h3>{{ (props.userInfo["p_signatures"]) == "" ? "这个人很懒，什么都没写！" : props.userInfo["p_signatures"] }}</h3>
    </div>

    <a-row class="personalCenter-heardBox-Card-row1-PC">
      <a-col :span="12">
        <div class="personalCenter-heardBox-Card1-PC">
          <p>ID : {{ props.userInfo["id"] }}</p>
          <p>用户名 : {{ props.userInfo["username"] }}</p>
          <p>昵称 : {{ props.userInfo["nickname"] }}</p>
          <p>邮箱 : {{ props.userInfo["email"] }}</p>
          <p>权限 : {{ props.userInfo["role"] == "0" ? "管理员" : "编辑者" }}</p>
        </div>
      </a-col>
      <a-col :span="12">
        <div class="personalCenter-heardBox-Card2-PC">
          <p>上次登录时间 : {{ props.userInfo["last_longin_date"] }}</p>
          <p>上次登录IP : {{ props.userInfo["last_longin_ip_address"] }}</p>
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

.personalCenter-heardBox-PC {
  border-radius: 15px;
  margin: 50px;
  padding: 25px;
  background: #ffffff;
  height: 550px;
}

.personalCenter-heardBox-avatar-PC {
  text-align: center;
}

.personalCenter-heardBox-avatar-PC h2, h3 {
  font-weight: lighter;
}

.personalCenter-heardBox-Card-row1-PC {
  margin: 35px;
}

.personalCenter-heardBox-Card1-PC {
  border-radius: 10px;
  margin-right: 25px;
  padding: 20px;
  background: #f2f3f5d4;
}

.personalCenter-heardBox-Card2-PC {
  border-radius: 10px;
  margin-left: 25px;
  padding: 20px;
  background: #f2f3f5d4;
}
</style>