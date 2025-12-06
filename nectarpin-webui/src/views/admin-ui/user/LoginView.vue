<template>
  <div class="login-page">
    <div class="background-decoration">
      <div class="decoration-circle circle-1"></div>
      <div class="decoration-circle circle-2"></div>
      <div class="decoration-circle circle-3"></div>
    </div>

    <n-grid cols="3" item-responsive class="login-grid">
      <!-- 宣传栏 -->
      <n-grid-item span="0 600:1 800:2">
        <div class="promo-section">
          <div class="promo-content">
            <div class="brand-logo">
              <div class="logo-icon">🍯</div>
            </div>
            <h1 class="brand-title">NectarPin</h1>
            <p class="brand-slogan">钉住花蜜一般的瞬间！</p>
          </div>
        </div>
      </n-grid-item>

      <!-- 登录框 -->
      <n-grid-item>
        <div class="login-container">
          <div class="login-card">
            <div class="login-header">
              <h2 class="login-title">欢迎回来</h2>
              <p class="login-subtitle">请登录您的账户以继续</p>
            </div>

            <n-form
              ref="LoginFormRef"
              :model="LoginFormData"
              :rules="rules"
              :show-label="false"
              size="large"
              class="login-form"
            >
              <n-form-item path="username">
                <n-input
                  v-model:value="LoginFormData.username"
                  placeholder="用户名"
                  :bordered="false"
                  class="custom-input"
                >
                  <template #prefix>
                    <n-icon :component="PersonIcon" class="input-icon" />
                  </template>
                </n-input>
              </n-form-item>

              <n-form-item path="password">
                <n-input
                  v-model:value="LoginFormData.password"
                  type="password"
                  placeholder="密码"
                  show-password-on="click"
                  :bordered="false"
                  class="custom-input"
                >
                  <template #prefix>
                    <n-icon :component="LockClosedIcon" class="input-icon" />
                  </template>
                </n-input>
              </n-form-item>

              <div class="form-options">
                <n-checkbox v-model:checked="rememberMe">记住我</n-checkbox>
                <a href="#" class="forgot-link">忘记密码？</a>
              </div>

              <n-form-item>
                <n-button
                  type="primary"
                  attr-type="button"
                  @click="handleLogin"
                  :loading="loading"
                  block
                  size="large"
                  class="login-button"
                >
                  登录
                </n-button>
              </n-form-item>
            </n-form>
          </div>
        </div>
      </n-grid-item>
    </n-grid>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import {
  NGrid,
  NGridItem,
  NForm,
  NFormItem,
  NInput,
  NButton,
  NCheckbox,
  NIcon,
  useMessage,
} from 'naive-ui'
import type { FormInst, FormRules } from 'naive-ui'
import { PersonOutline as PersonIcon, LockClosedOutline as LockClosedIcon } from '@vicons/ionicons5'

const message = useMessage()

//登录表单引用
const LoginFormRef = ref<FormInst | null>(null)

//登录表单数据
const LoginFormData = ref({
  username: '',
  password: '',
})

//记住我
const rememberMe = ref(false)

//加载状态
const loading = ref(false)

//表单验证规则
const rules: FormRules = {
  username: [{ required: true, message: '请输入用户名', trigger: 'blur' }],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' },
    { min: 6, message: '密码长度不能少于6位', trigger: 'blur' },
  ],
}

//登录处理
const handleLogin = async () => {
  try {
    await LoginFormRef.value?.validate()
    loading.value = true

    // 模拟登录请求
    setTimeout(() => {
      loading.value = false
      message.success('登录成功！')
      // 这里可以添加实际的登录逻辑
    }, 1500)
  } catch (error) {
    console.error('表单验证失败:', error)
  }
}
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
  background: linear-gradient(135deg, #1e3a8a 0%, #3b82f6 50%, #60a5fa 100%);
  background-size: 200% 200%;
  animation: gradientShift 15s ease infinite;
  overflow: hidden;
}

@keyframes gradientShift {
  0% {
    background-position: 0 50%;
  }
  50% {
    background-position: 100% 50%;
  }
  100% {
    background-position: 0 50%;
  }
}

//背景装饰
.background-decoration {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  pointer-events: none;
  z-index: 0;
}

.decoration-circle {
  position: absolute;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  animation: float 20s ease-in-out infinite;
}

.circle-1 {
  width: 300px;
  height: 300px;
  top: -100px;
  left: -100px;
  animation-delay: 0s;
}

.circle-2 {
  width: 200px;
  height: 200px;
  bottom: -50px;
  right: 10%;
  animation-delay: 5s;
}

.circle-3 {
  width: 150px;
  height: 150px;
  top: 50%;
  left: 5%;
  animation-delay: 10s;
}

@keyframes float {
  0%,
  100% {
    transform: translate(0, 0) scale(1);
  }
  33% {
    transform: translate(30px, -30px) scale(1.1);
  }
  66% {
    transform: translate(-20px, 20px) scale(0.9);
  }
}

.login-grid {
  position: relative;
  z-index: 1;
  height: 100%;
}

//宣传栏
.promo-section {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 40px;

  .promo-content {
    text-align: center;
    color: #fff;
    animation: fadeInUp 0.8s ease-out;

    .brand-logo {
      margin-bottom: 30px;

      .logo-icon {
        font-size: 5rem;
        display: inline-block;
        animation: bounce 2s ease-in-out infinite;
      }
    }

    .brand-title {
      font-size: 4.5rem;
      font-weight: 800;
      margin: 0 0 20px 0;
      text-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
      letter-spacing: 3px;
      background: linear-gradient(135deg, #ffffff 0%, #dbeafe 50%, #bfdbfe 100%);
      -webkit-background-clip: text;
      -webkit-text-fill-color: transparent;
      background-clip: text;
    }

    .brand-slogan {
      font-size: 1.5rem;
      font-weight: 300;
      margin: 0 0 40px 0;
      opacity: 0.95;
      text-shadow: 0 2px 10px rgba(0, 0, 0, 0.2);
      letter-spacing: 1px;
    }
  }
}

@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(30px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

@keyframes bounce {
  0%,
  100% {
    transform: translateY(0);
  }
  50% {
    transform: translateY(-20px);
  }
}

//登录容器
.login-container {
  height: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 20px;
  position: relative;
  z-index: 1;
}

//登录卡片
.login-card {
  width: 100%;
  max-width: 420px;
  background: rgba(255, 255, 255, 0.95);
  backdrop-filter: blur(20px);
  border-radius: 24px;
  padding: 48px 40px;
  box-shadow:
    0 20px 60px rgba(0, 0, 0, 0.3),
    0 0 0 1px rgba(255, 255, 255, 0.5) inset;
  animation: slideInRight 0.6s ease-out;
  transition:
    transform 0.3s ease,
    box-shadow 0.3s ease;

  &:hover {
    transform: translateY(-5px);
    box-shadow:
      0 25px 70px rgba(0, 0, 0, 0.35),
      0 0 0 1px rgba(255, 255, 255, 0.5) inset;
  }
}

@keyframes slideInRight {
  from {
    opacity: 0;
    transform: translateX(30px);
  }
  to {
    opacity: 1;
    transform: translateX(0);
  }
}

.login-header {
  text-align: center;
  margin-bottom: 36px;

  .login-title {
    font-size: 2rem;
    font-weight: 700;
    margin: 0 0 8px 0;
    color: #1a1a1a;
    letter-spacing: -0.5px;
  }

  .login-subtitle {
    font-size: 0.95rem;
    color: #666;
    margin: 0;
  }
}

.login-form {
  .custom-input {
    background: #f5f5f5;
    border-radius: 12px;
    transition: all 0.3s ease;

    .input-icon {
      font-size: 18px;
      color: #999;
      display: flex;
      align-items: center;
    }

    &:hover {
      background: #eeeeee;
    }

    &:focus-within {
      background: #fff;
      box-shadow: 0 0 0 3px rgba(59, 130, 246, 0.15);
    }
  }

  .form-options {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 24px;
    font-size: 14px;

    .forgot-link {
      color: #3b82f6;
      text-decoration: none;
      font-weight: 500;
      transition: color 0.2s ease;

      &:hover {
        color: #2563eb;
      }
    }
  }

  .login-button {
    height: 50px;
    font-size: 16px;
    font-weight: 600;
    border-radius: 12px;
    background: linear-gradient(135deg, #3b82f6 0%, #2563eb 100%);
    border: none;
    box-shadow: 0 4px 15px rgba(59, 130, 246, 0.3);
    transition: all 0.3s ease;

    &:hover:not(:disabled) {
      transform: translateY(-2px);
      box-shadow: 0 6px 20px rgba(59, 130, 246, 0.4);
      background: linear-gradient(135deg, #60a5fa 0%, #3b82f6 100%);
    }

    &:active:not(:disabled) {
      transform: translateY(0);
    }
  }
}

//响应式设计
@media (max-width: 800px) {
  .promo-section {
    display: none;
  }

  .login-card {
    max-width: 100%;
    padding: 40px 32px;
  }

  .brand-title {
    font-size: 3.5rem !important;
  }
}

@media (max-width: 600px) {
  .login-card {
    padding: 32px 24px;
    border-radius: 20px;
  }

  .login-header .login-title {
    font-size: 1.75rem;
  }
}
</style>
