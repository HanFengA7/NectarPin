<template>
  <div class="login-page">
    <n-grid cols="3" item-responsive>
      <!-- 宣传栏 -->
      <n-grid-item span="0 600:1 800:2">
        <div class="light-green"></div>
      </n-grid-item>
      <!-- 登录框 -->
      <n-grid-item>
        <div class="login-box">
          <n-form ref="LoginFormRef">
            <n-form-item label="用户名" path="username">
              <n-input v-model:value="LoginFormData.username" placeholder="输入用户名" />
            </n-form-item>
            <n-form-item label="密码" path="password">
              <n-input v-model:value="LoginFormData.password" placeholder="输入密码" />
            </n-form-item>
            <n-form-item>
              <n-button 
                attr-type="button" 
                @click=""
                block 
                circle> 登录 </n-button>
            </n-form-item>
          </n-form>
        </div>
      </n-grid-item>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { NGrid, NGridItem, NForm, NFormItem, NInput, NButton } from 'naive-ui'
import type { FormInst } from 'naive-ui'

//登录表单引用
const LoginFormRef = ref<FormInst | null>(null)

//登录表单数据
const LoginFormData = ref({
    username: '',
    password: '',
})
</script>

<style scoped lang="scss">
//登录页面
.login-page {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  width: 100vw;
  height: 100vh;
  margin: 0;
  padding: 0;
  background-image: linear-gradient(120deg, #a1c4fd 0%, #c2e9fb 100%);
  background-repeat: no-repeat;
  background-size: cover;
  background-attachment: fixed;

  :deep(.n-grid) {
    height: 100%;
  }
}

//登录框
.login-box {
  height: 100%;
  background-color: #fff;
  padding: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  
  :deep(.n-form) {
    width: 100%;
    max-width: 400px;
  }
}
</style>
