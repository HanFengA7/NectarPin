<script setup lang="ts">
import {reactive, ref} from 'vue';
import {debounce} from "@/plugin/debounce/debounce";
import {md5} from "js-md5";
import {Login} from "@/api/User/login";
import {Notification} from "@arco-design/web-vue";
import router from "@/router";
import moment from "moment";
moment.locale('zh-cn', {
  months: '一月_二月_三月_四月_五月_六月_七月_八月_九月_十月_十一月_十二月'.split('_'),
  monthsShort: '1月_2月_3月_4月_5月_6月_7月_8月_9月_10月_11月_12月'.split('_'),
  weekdays: '星期日_星期一_星期二_星期三_星期四_星期五_星期六'.split('_'),
  weekdaysShort: '周日_周一_周二_周三_周四_周五_周六'.split('_'),
  weekdaysMin: '日_一_二_三_四_五_六'.split('_'),
})

const layout = ref('vertical');
const form = reactive({
  username: '',
  password: '',
  this_longin_ip_address: '',
  this_longin_date: '',
});

fetch('https://api.ipify.org?format=json')
    .then(response => response.json())
    .then(json => {
      form.this_longin_ip_address = json.ip;
    })

interface handleSubmit {
  errors: any;
}

const handleSubmit = debounce(({errors}: handleSubmit) => {
  if (errors === undefined) {
    form.password = md5(form.password);
    form.this_longin_date = moment().format("YYYY-MM-DD dddd H:mm:ss");
    Login(form).then((res: any) => {
      if (res.data.code == 200) {
        //window.sessionStorage.setItem('token', res.data.token)
        window.localStorage.setItem('token', res.data.token)
        Notification.success({
          content: res.data.msg,
          duration: 4000,
          closable: true,
        })
        router.push('/Dashboard')
      }
      if (res.data.code == 500) {
        Notification.error({
          content: res.data.msg,
          duration: 7000,
          closable: true,
        })
      }
    })
  }
}, 900);


</script>

<template>
  <!--[User Body] [PC] [Start]-->
  <div class="Login-Body div-class-name">
    <div class="Login-Box">
      <a-row :gutter="24">
        <a-col :span="11">
          <div class="Login-Box-L">
            <h2>NectarPin - 身份认证</h2>
            <a-form :model="form" :layout="layout" @submit="handleSubmit">
              <a-form-item
                  field="username" label="用户名:"
                  :rules="[{required:true,message:'请输入用户名'}]"
                  :validate-trigger="['change','input']"
              >
                <a-input v-model="form.username"/>
              </a-form-item>
              <a-form-item
                  field="password" label="密码:"
                  :rules="[{required:true,message:'请输入密码'}]"
                  :validate-trigger="['change','input']"
              >
                <a-input v-model="form.password" type="password"/>
              </a-form-item>
              <div style="text-align: center;">
                <a-button class="Login-Box-L-but" html-type="submit">登录</a-button>
              </div>
            </a-form>
          </div>
        </a-col>
        <a-col :span="2">

        </a-col>
        <a-col :span="11">
          <div class="Login-Box-R">
            <h1>NectarPin</h1>
            <h3>A moment like nailing nectar !</h3>
          </div>
        </a-col>
      </a-row>
    </div>
  </div>
  <!--[User Body] [PC] [End]-->
</template>

<style scoped>
.Login-Body {
  background-image: url("/bj-2.jpg");
  height: 100vh;
  max-width: 100%;
  background-size: cover;
  background-attachment: fixed;
  background-position: 100%;
}

.Login-Box {
  position: absolute;
  top: 45%;
  left: 45%;
  transform: translate(-50%, -50%);
}

.Login-Box-L {
  background: rgba(255, 255, 255, 0.2);
  border-radius: 15px;
  backdrop-filter: blur(10px);
  padding: 20px;
  width: 400px;
  text-align: center;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.1);
  position: relative;
  overflow: hidden;
}


.Login-Box-L h2 {
  font-weight: 300;
}

.Login-Box-L-but {
  width: 40%;
  background-color: rgb(253 253 253 / 82%);
  border-radius: 5px;
  margin: 5px;
}

.arco-input-wrapper {
  background-color: rgb(251 251 252) !important;
  border-radius: 8px;
}

.Login-Box-R {
  padding: 5px;
  top: 50px;
  left: 100px;
  text-align: center;
  position: relative;
  overflow: hidden;
}

.Login-Box-R h1 {
  background: linear-gradient(to top right, #5F9DF7, #7360DF, #d7dde8, #efefef);
  -webkit-background-clip: text;
  color: transparent;
  font-size: 3em;
  font-weight: bold;
  margin: 0;
  padding: 10px;
  text-align: left;
}

.Login-Box-R h3 {
  background: #ffffffe8;
  -webkit-background-clip: text;
  color: transparent;
  font-weight: 400;
  text-align: left;
}
</style>