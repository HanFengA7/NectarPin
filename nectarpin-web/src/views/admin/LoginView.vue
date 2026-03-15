<script setup lang="ts">
import { ref, computed, watch } from 'vue'
import { useRouter } from 'vue-router'
import { login, type LoginResult } from '@/api/user'
import { encryptPassword } from '@/utils'
import { useUserStore } from '@/stores/user'

const router = useRouter()
const userStore = useUserStore()

const account = ref('')
const password = ref('')
const rememberMe = ref(false)
const loading = ref(false)
const errorMessage = ref('')
const inputError = ref('')

const isEmail = computed(() => {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
  return emailRegex.test(account.value.trim())
})

const inputType = computed(() => {
  if (!account.value.trim()) return ''
  return isEmail.value ? 'email' : 'username'
})

const inputPlaceholder = computed(() => {
  return '请输入邮箱或用户名'
})

const inputHint = computed(() => {
  if (!account.value.trim()) {
    return '支持邮箱格式（如 admin@example.com）或用户名格式'
  }
  if (isEmail.value) {
    return '检测到邮箱格式登录'
  }
  return '检测到用户名格式登录'
})

const validateInput = (): boolean => {
  const trimmedAccount = account.value.trim()

  if (!trimmedAccount) {
    inputError.value = '请输入邮箱或用户名'
    return false
  }

  if (isEmail.value) {
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/
    if (!emailRegex.test(trimmedAccount)) {
      inputError.value = '邮箱格式不正确'
      return false
    }
  } else {
    if (trimmedAccount.length < 3) {
      inputError.value = '用户名至少需要 3 个字符'
      return false
    }
    if (trimmedAccount.length > 50) {
      inputError.value = '用户名不能超过 50 个字符'
      return false
    }
    const usernameRegex = /^[a-zA-Z0-9_\u4e00-\u9fa5]+$/
    if (!usernameRegex.test(trimmedAccount)) {
      inputError.value = '用户名只能包含字母、数字、下划线或中文'
      return false
    }
  }

  inputError.value = ''
  return true
}

watch(account, () => {
  if (inputError.value) {
    validateInput()
  }
})

const handleLogin = async () => {
  if (!validateInput()) {
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
      account: account.value.trim(),
      password: encryptedPassword,
    })
    const data = res.data as LoginResult

    const storage = rememberMe.value ? localStorage : sessionStorage
    storage.setItem('token', data.token.access_token)
    storage.setItem('refresh_token', data.token.refresh_token)

    userStore.setTokens(data.token.access_token, data.token.refresh_token)
    userStore.setUser(data.user)

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
      <div class="deco deco-1"></div>
      <div class="deco deco-2"></div>

      <div class="brand">
        <div class="brand-logo"></div>
        <div class="brand-info">
          <span class="brand-name">NectarPin Admin</span>
          <span class="brand-tagline">Secure publishing workspace</span>
        </div>
      </div>

      <div class="hero-text">
        <h1 class="hero-title">钉住花蜜一般的瞬间！</h1>
        <p class="hero-subtitle">统一管理文章、评论、系统配置，所有关键数据在一个工作台内完成。</p>
      </div>

      <div class="hero-card">
        <h3 class="card-title">今日运维提醒</h3>
        <div class="card-item">
          <span class="dot dot-blue"></span>
          <span>8 篇文章待审核发布</span>
        </div>
        <div class="card-item">
          <span class="dot dot-yellow"></span>
          <span>28 条评论进入审核队列</span>
        </div>
        <div class="card-item">
          <span class="dot dot-green"></span>
          <span>站点健康检查通过</span>
        </div>
      </div>
    </div>

    <div class="form-section">
      <div class="form-card">
        <div class="form-header">
          <h2 class="form-title">欢迎回来</h2>
          <p class="form-subtitle">使用管理员账号登录 Nectarpin 后台。</p>
        </div>

        <form @submit.prevent="handleLogin" class="login-form">
          <div class="form-group">
            <div class="label-row">
              <label class="form-label">邮箱 / 用户名</label>
              <span class="input-type-badge" v-if="inputType">
                {{ inputType === 'email' ? '邮箱登录' : '用户名登录' }}
              </span>
            </div>
            <input
              v-model="account"
              type="text"
              class="form-input"
              :class="{ 'input-error': inputError }"
              :placeholder="inputPlaceholder"
              :disabled="loading"
              @blur="validateInput"
            />
            <div class="input-hint" :class="{ error: inputError }">
              {{ inputError || inputHint }}
            </div>
          </div>

          <div class="form-group">
            <div class="label-row">
              <label class="form-label">密码</label>
              <a href="#" class="forgot-link">忘记密码？</a>
            </div>
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
              <span class="checkbox-custom"></span>
              <span class="remember-text">保持登录 7 天</span>
            </label>
            <span class="status-badge">服务端 已连接</span>
          </div>

          <button type="submit" class="submit-btn" :disabled="loading">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4" />
              <polyline points="10 17 15 12 10 7" />
              <line x1="15" y1="12" x2="3" y2="12" />
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
  background-color: #f4f8ff;
  font-family: 'IBM Plex Sans', sans-serif;
}

.hero-section {
  width: 620px;
  padding: 56px;
  display: flex;
  flex-direction: column;
  position: relative;
  overflow: hidden;
  background: linear-gradient(135deg, #1a3b8c 0%, #2f6bff 55%, #5b8fff 100%);
}

.deco {
  position: absolute;
  border-radius: 50%;
  pointer-events: none;
}

.deco-1 {
  width: 480px;
  height: 480px;
  background: #ffffff08;
  top: -200px;
  left: -200px;
}

.deco-2 {
  width: 320px;
  height: 320px;
  background: #ffffff06;
  bottom: 180px;
  right: -140px;
}

.brand {
  display: flex;
  align-items: center;
  gap: 12px;
  position: relative;
  z-index: 1;
}

.brand-logo {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: #ffffff30;
}

.brand-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.brand-name {
  font-family: 'Manrope', sans-serif;
  font-size: 18px;
  font-weight: 700;
  color: #ffffff;
}

.brand-tagline {
  font-size: 12px;
  color: #ffffffaa;
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: 14px;
  margin-top: 76px;
  position: relative;
  z-index: 1;
}

.hero-title {
  font-family: 'Manrope', sans-serif;
  font-size: 48px;
  font-weight: 800;
  color: #ffffff;
  line-height: 1.15;
  letter-spacing: -1.5px;
  max-width: 500px;
}

.hero-subtitle {
  font-size: 16px;
  color: #ffffffcc;
  max-width: 500px;
  line-height: 1.6;
}

.hero-card {
  display: flex;
  flex-direction: column;
  gap: 14px;
  padding: 24px;
  background: #ffffff18;
  backdrop-filter: blur(14px);
  border-radius: 20px;
  border: 1px solid #ffffff30;
  width: 508px;
  margin-top: auto;
  position: relative;
  z-index: 1;
}

.card-title {
  font-family: 'Manrope', sans-serif;
  font-size: 14px;
  font-weight: 700;
  color: #ffffff;
  letter-spacing: 0.3px;
}

.card-item {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #ffffffdd;
}

.dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.dot-blue {
  background-color: #60a5fa;
}

.dot-yellow {
  background-color: #fcd34d;
}

.dot-green {
  background-color: #34d399;
}

.form-section {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 64px;
}

.form-card {
  width: 440px;
  padding: 40px;
  background: #ffffff;
  border-radius: 24px;
  border: 1px solid #d7e4ff;
  box-shadow: 0 20px 60px rgba(26, 63, 112, 0.125);
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-header {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.form-title {
  font-family: 'Manrope', sans-serif;
  font-size: 28px;
  font-weight: 800;
  color: #10233f;
  letter-spacing: -0.5px;
}

.form-subtitle {
  font-size: 14px;
  color: #58708f;
  line-height: 1.5;
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: 24px;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.form-label {
  font-size: 12px;
  font-weight: 700;
  color: #58708f;
  letter-spacing: 0.4px;
}

.input-type-badge {
  padding: 3px 10px;
  font-size: 10px;
  font-weight: 700;
  color: #1d4ed8;
  background-color: #dde9ff;
  border-radius: 999px;
  letter-spacing: 0.3px;
}

.forgot-link {
  font-size: 13px;
  font-weight: 600;
  color: #1d4ed8;
  text-decoration: none;
}

.forgot-link:hover {
  text-decoration: underline;
}

.form-input {
  width: 100%;
  padding: 14px 16px;
  font-size: 14px;
  color: #10233f;
  background-color: #eef4ff;
  border: 1px solid #d7e4ff;
  border-radius: 12px;
  outline: none;
  box-sizing: border-box;
  transition:
    border-color 0.2s,
    box-shadow 0.2s;
}

.form-input::placeholder {
  color: #58708f;
}

.form-input:focus {
  border-color: #2f6bff;
  box-shadow: 0 0 0 3px rgba(47, 107, 255, 0.15);
}

.form-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.form-input.input-error {
  border-color: #ef4444;
  box-shadow: 0 0 0 3px rgba(239, 68, 68, 0.15);
}

.input-hint {
  font-size: 12px;
  color: #58708f;
  min-height: 18px;
  transition: color 0.2s;
}

.input-hint.error {
  color: #ef4444;
}

.error-message {
  padding: 12px 16px;
  font-size: 14px;
  color: #ef4444;
  background-color: #fee2e2;
  border-radius: 12px;
}

.form-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: 10px;
  font-size: 14px;
  color: #10233f;
  cursor: pointer;
  position: relative;
}

.remember-me input[type='checkbox'] {
  position: absolute;
  opacity: 0;
  width: 18px;
  height: 18px;
  cursor: pointer;
}

.checkbox-custom {
  width: 18px;
  height: 18px;
  border-radius: 6px;
  background: #2f6bff;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.checkbox-custom::after {
  content: '';
  width: 10px;
  height: 6px;
  border-left: 2px solid #ffffff;
  border-bottom: 2px solid #ffffff;
  transform: rotate(-45deg) translateY(-1px);
}

.remember-me input[type='checkbox']:not(:checked) + .checkbox-custom {
  background: #eef4ff;
  border: 1px solid #d7e4ff;
}

.remember-me input[type='checkbox']:not(:checked) + .checkbox-custom::after {
  display: none;
}

.remember-me input[type='checkbox']:disabled + .checkbox-custom {
  opacity: 0.6;
  cursor: not-allowed;
}

.remember-text {
  color: #10233f;
}

.status-badge {
  padding: 5px 12px;
  font-size: 11px;
  font-weight: 700;
  color: #1d4ed8;
  background-color: #dde9ff;
  border-radius: 999px;
  letter-spacing: 0.4px;
}

.submit-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  width: 100%;
  padding: 11px 20px;
  font-size: 14px;
  font-weight: 600;
  color: #ffffff;
  background-color: #2f6bff;
  border: none;
  border-radius: 999px;
  cursor: pointer;
  transition: background-color 0.2s;
  letter-spacing: 0.2px;
}

.submit-btn:hover:not(:disabled) {
  background-color: #1d4ed8;
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
  color: #58708f;
}

.alt-link a {
  font-weight: 600;
  color: #1d4ed8;
  text-decoration: none;
}

.alt-link a:hover {
  text-decoration: underline;
}

@media (max-width: 1200px) {
  .hero-section {
    width: 500px;
    padding: 48px;
  }

  .hero-title {
    font-size: 40px;
  }

  .hero-card {
    width: 420px;
  }
}

@media (max-width: 1024px) {
  .login-page {
    flex-direction: column;
  }

  .hero-section {
    width: 100%;
    min-height: auto;
    padding: 40px 32px;
  }

  .hero-text {
    margin-top: 40px;
  }

  .hero-title {
    font-size: 36px;
    max-width: 100%;
  }

  .hero-subtitle {
    max-width: 100%;
  }

  .hero-card {
    width: 100%;
    max-width: 508px;
    margin-top: 32px;
  }

  .form-section {
    padding: 48px 32px;
  }

  .form-card {
    width: 100%;
    max-width: 440px;
  }
}

@media (max-width: 640px) {
  .hero-section {
    padding: 32px 24px;
  }

  .hero-title {
    font-size: 28px;
  }

  .hero-card {
    padding: 20px;
  }

  .form-section {
    padding: 32px 20px;
  }

  .form-card {
    padding: 28px 24px;
  }

  .form-title {
    font-size: 24px;
  }
}
</style>
