<script setup lang="ts">
import {reactive, ref} from 'vue';

const layout = ref('vertical');
const form = reactive({
  name: '',
  post: '',
});
</script>

<template>
  <!--[Login Body] [PC] [Start]-->
  <div class="Login-Body div-class-name">
    <div class="Login-Box">
      <a-row :gutter="24">
        <a-col :span="11">
          <div class="Login-Box-L">
            <div class="Login-Box-L-overlay"></div>
            <h2>NectarPin - 身份认证</h2>
            <a-form :model="form" :layout="layout">
              <a-form-item field="name" label="用户名:">
                <a-input v-model="form.name"/>
              </a-form-item>
              <a-form-item field="post" label="密码:">
                <a-input v-model="form.post" type="password"/>
              </a-form-item>
              <div style="text-align: center;">
                <a-button class="Login-Box-L-but">登录</a-button>
              </div>
            </a-form>
          </div>
        </a-col>
        <a-col :span="2">

        </a-col>
        <a-col :span="11">
          <div class="Login-Box-R">
            <h1>NectarPin</h1>
            <h3>A moment like nailing nectar!</h3>
          </div>
        </a-col>
      </a-row>
    </div>
  </div>
  <!--[Login Body] [PC] [End]-->
</template>

<style scoped>
.Login-Body {
  background-image: url("https://cloudcache.tencent-cloud.com/qcloud/ui/static/tc_portal/c3f153dd-80a0-40a8-9a54-a3b7d6717ab8.jpg");
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

.Login-Box-L-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: repeating-linear-gradient(45deg, rgba(255, 255, 255, 0.05), rgba(255, 255, 255, 0.05) 10px, rgba(0, 0, 0, 0) 10px, rgba(0, 0, 0, 0) 20px);
  pointer-events: none;
}

.Login-Box-L h2 {
  font-weight: 300;
}
.Login-Box-L-but{
  width: 80%;
  background-color: rgb(253 253 253 / 82%);
  border-radius: 8px;
}
.arco-input-wrapper{
  background-color: rgb(251 251 252)!important;
  border-radius: 8px;
}
.Login-divider{
  border-left: 1px solid #1b1b1f; /* 宽度可调整 */
  height: 500px;
}

.Login-Box-R {
  padding: 5px;
  top:50px;
  left: 100px;
  text-align: center;
  position: relative;
  overflow: hidden;
}
.Login-Box-R h1{
  background: linear-gradient(to right, #5d72e7, #e1e1e8);
  -webkit-background-clip: text;
  color: transparent;
  font-size: 3em;
  font-weight: bold;
  margin: 0;
  padding: 10px;
  text-align: left;
}
.Login-Box-R h3{
  background: #b5b5b9db;
  -webkit-background-clip: text;
  color: transparent;
  font-weight: 400;
  text-align: left;
}
</style>