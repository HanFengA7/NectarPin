<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { login, type LoginResult } from '@/api/user'
import { encryptPassword } from '@/utils'

const router = useRouter()

const username = ref('')
const password = ref('')
const rememberMe = ref(false)
const loading = ref(false)
const errorMessage = ref('')

const handleLogin = async () => {
  if (!username.value.trim()) {
    errorMessage.value = '请输入用户名'
    return
  }

  if (!password.value.trim()) {
    errorMessage.value = '请输入密码'
    return
  }

  loading.value = true
  errorMessage.value = ''

  try {
    const encryptedPassword = encryptPassword(password.value)
    const res = await login({ 
      username: username.value, 
      password: encryptedPassword 
    })
    const data = res.data as LoginResult
    
    const storage = rememberMe.value ? localStorage : sessionStorage
    storage.setItem('token', data.token.access_token)
    storage.setItem('refresh_token', data.token.refresh_token)
    storage.setItem('user', JSON.stringify(data.user))
    
    router.push('/admin')
  } catch (error: unknown) {
    const err = error as { response?: { data?: { message?: string } } }
    errorMessage.value = err.response?.data?.message || '登录失败，请重试'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-page">
    <div class="hero-section">
      <div class="brand">
        <div class="brand-logo"></div>
        <div class="brand-info">
          <span class="brand-name">NectarPin Admin</span>
          <span class="brand-tagline">Secure publishing workspace</span>
        </div>
      </div>

      <div class="hero-text">
        <h1 class="hero-title">登录后台，处理内容与运营决策。</h1>
        <p class="hero-subtitle">统一管理文章、评论和系统配置，所有关键数据在一个工作台内完成。</p>
      </div>

      <!-- <div class="hero-card">
        <h3 class="card-title">今日运维提醒</h3>
        <div class="card-item">
          <span class="dot dot-primary"></span>
          <span>8 篇文章待审核发布</span>
        </div>
        <div class="card-item">
          <span class="dot dot-warning"></span>
          <span>28 条评论进入审核队列</span>
        </div>
        <div class="card-item">
          <span class="dot dot-success"></span>
          <span>站点健康检查通过</span>
        </div>
      </div> -->
    </div>

    <div class="form-section">
      <div class="form-card">
        <div class="form-header">
          <h2 class="form-title">欢迎回来</h2>
          <p class="form-subtitle">使用管理员账号登录 Nectarpin 博客后台。</p>
        </div>

        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <label class="form-label">用户名</label>
            <input
              v-model="username"
              type="text"
              class="form-input"
              placeholder="请输入用户名"
              :disabled="loading"
            />
          </div>

          <div class="form-group">
            <label class="form-label">密码</label>
            <input
              v-model="password"
              type="password"
              class="form-input"
              placeholder="请输入密码"
              :disabled="loading"
            />
          </div>

          <div v-if="errorMessage" class="error-message">
            {{ errorMessage }}
          </div>

          <div class="form-row">
            <label class="remember-me">
              <input type="checkbox" v-model="rememberMe" :disabled="loading" />
              <span>记住登录</span>
            </label>
            <span class="sso-badge">服务端已连接</span>
          </div>

          <button type="submit" class="submit-btn" :disabled="loading">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
              <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4"/>
              <polyline points="10 17 15 12 10 7"/>
              <line x1="15" y1="12" x2="3" y2="12"/>
            </svg>
            <span>{{ loading ? '登录中...' : '登录后台' }}</span>
          </button>

          <div class="alt-link">
            <span>忘记密码？</span>
            <a href="#">前往重置</a>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<style scoped>
.login-page {
  display: flex;
  min-height: 100vh;
  background-color: #F4F8FF;
  font-family: 'IBM Plex Sans', sans-serif;
}

.hero-section {
  width: 620px;
  padding: 48px;
  display: flex;
  flex-direction: column;
  gap: 24px;
  background: linear-gradient(145deg, #EAF1FF 0%, #CFE0FF 100%);
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
}

.brand-logo {
  width: 44px;
  height: 44px;
  border-radius: 50%;
  background: linear-gradient(140deg, #2F6BFF 0%, #1D4ED8 100%);
}

.brand-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.brand-name {
  font-family: 'Manrope', sans-serif;
  font-size: 20px;
  font-weight: 700;
  color: #10233F;
}

.brand-tagline {
  font-size: 12px;
  color: #58708F;
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: 14px;
}

.hero-title {
  font-family: 'Manrope', sans-serif;
  font-size: 42px;
  font-weight: 700;
  color: #10233F;
  line-height: 1.2;
  max-width: 420px;
}

.hero-subtitle {
  font-size: 16px;
  color: #58708F;
  max-width: 420px;
  line-height: 1.5;
}

.hero-card {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 24px;
  background: rgba(255, 255, 255, 0.6);
  backdrop-filter: blur(14px);
  border-radius: 24px;
  width: 420px;
}

.card-title {
  font-family: 'Manrope', sans-serif;
  font-size: 18px;
  font-weight: 700;
  color: #10233F;
}

.card-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #10233F;
}

.dot {
  width: 10px;
  height: 10px;
  border-radius: 50%;
}

.dot-primary {
  background-color: #2F6BFF;
}

.dot-warning {
  background-color: #F59E0B;
}

.dot-success {
  background-color: #10B981;
}

.form-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 48px;
}

.form-card {
  width: 440px;
  padding: 32px;
  background: #FFFFFF;
  border-radius: 28px;
  border: 1px solid #D7E4FF;
  box-shadow: 0 18px 40px rgba(47, 107, 255, 0.1);
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.form-header {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-title {
  font-family: 'Manrope', sans-serif;
  font-size: 30px;
  font-weight: 700;
  color: #10233F;
}

.form-subtitle {
  font-size: 14px;
  color: #58708F;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 18px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-label {
  font-size: 13px;
  font-weight: 600;
  color: #58708F;
}

.form-input {
  width: 100%;
  padding: 14px 16px;
  font-size: 14px;
  color: #10233F;
  background-color: #EEF4FF;
  border: none;
  border-radius: 14px;
  outline: none;
  box-sizing: border-box;
}

.form-input::placeholder {
  color: #58708F;
}

.form-input:focus {
  box-shadow: 0 0 0 2px #2F6BFF;
}

.form-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.error-message {
  padding: 12px 16px;
  font-size: 14px;
  color: #EF4444;
  background-color: #FEE2E2;
  border-radius: 14px;
}

.form-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 14px;
  color: #58708F;
  cursor: pointer;
}

.remember-me input[type="checkbox"] {
  width: 16px;
  height: 16px;
  cursor: pointer;
  accent-color: #2F6BFF;
}

.sso-badge {
  padding: 8px 12px;
  font-size: 12px;
  font-weight: 600;
  color: #1D4ED8;
  background-color: #DDE9FF;
  border-radius: 999px;
}

.submit-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 10px;
  width: 100%;
  padding: 12px 18px;
  font-size: 14px;
  font-weight: 600;
  color: #FFFFFF;
  background-color: #2F6BFF;
  border: none;
  border-radius: 999px;
  cursor: pointer;
  transition: background-color 0.2s;
}

.submit-btn:hover:not(:disabled) {
  background-color: #1D4ED8;
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.alt-link {
  display: flex;
  justify-content: center;
  gap: 6px;
  font-size: 14px;
}

.alt-link span {
  color: #58708F;
}

.alt-link a {
  font-weight: 600;
  color: #1D4ED8;
  text-decoration: none;
}

.alt-link a:hover {
  text-decoration: underline;
}
</style>
