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
  background-color: var(--color-bg);
  font-family: var(--font-sans);
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
  background: rgba(255, 255, 255, 0.05);
  top: -200px;
  left: -200px;
}

.deco-2 {
  width: 320px;
  height: 320px;
  background: rgba(255, 255, 255, 0.04);
  bottom: 180px;
  right: -140px;
}

.brand {
  display: flex;
  align-items: center;
  gap: var(--space-3);
  position: relative;
  z-index: 1;
}

.brand-logo {
  width: 48px;
  height: 48px;
  border-radius: 50%;
  background: rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(8px);
}

.brand-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.brand-name {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: #ffffff;
}

.brand-tagline {
  font-size: var(--font-size-xs);
  color: rgba(255, 255, 255, 0.7);
}

.hero-text {
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
  margin-top: 76px;
  position: relative;
  z-index: 1;
}

.hero-title {
  font-family: var(--font-heading);
  font-size: 48px;
  font-weight: var(--font-weight-extrabold);
  color: #ffffff;
  line-height: 1.15;
  letter-spacing: -0.02em;
  max-width: 500px;
}

.hero-subtitle {
  font-size: var(--font-size-md);
  color: rgba(255, 255, 255, 0.8);
  max-width: 500px;
  line-height: var(--line-height-relaxed);
}

.hero-card {
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
  padding: var(--space-6);
  background: rgba(255, 255, 255, 0.12);
  backdrop-filter: blur(14px);
  border-radius: var(--radius-2xl);
  border: 1px solid rgba(255, 255, 255, 0.2);
  width: 508px;
  margin-top: auto;
  position: relative;
  z-index: 1;
}

.card-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-bold);
  color: #ffffff;
  letter-spacing: 0.3px;
}

.card-item {
  display: flex;
  align-items: center;
  gap: var(--space-3);
  font-size: var(--font-size-sm);
  color: rgba(255, 255, 255, 0.9);
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
  padding: var(--space-16);
}

.form-card {
  width: 440px;
  padding: var(--space-10);
  background: var(--color-surface);
  border-radius: var(--radius-2xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-xl);
  display: flex;
  flex-direction: column;
  gap: var(--space-6);
}

.form-header {
  display: flex;
  flex-direction: column;
  gap: var(--space-2);
}

.form-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-3xl);
  font-weight: var(--font-weight-extrabold);
  color: var(--color-text-primary);
  letter-spacing: var(--letter-spacing-tight);
}

.form-subtitle {
  font-size: var(--font-size-base);
  color: var(--color-text-muted);
  line-height: var(--line-height-normal);
}

.login-form {
  display: flex;
  flex-direction: column;
  gap: var(--space-6);
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: var(--space-2);
}

.label-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.form-label {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-secondary);
  letter-spacing: var(--letter-spacing-wide);
}

.input-type-badge {
  padding: 3px 10px;
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-bold);
  color: var(--color-primary);
  background-color: var(--color-primary-soft);
  border-radius: var(--radius-full);
  letter-spacing: 0.3px;
}

.forgot-link {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
  color: var(--color-primary);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.forgot-link:hover {
  color: var(--color-primary-hover);
  text-decoration: underline;
}

.form-input {
  width: 100%;
  padding: var(--space-3) var(--space-4);
  font-size: var(--font-size-base);
  color: var(--color-text-primary);
  background-color: var(--color-surface-alt);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  outline: none;
  box-sizing: border-box;
  transition: border-color var(--transition-fast), box-shadow var(--transition-fast);
}

.form-input::placeholder {
  color: var(--color-text-placeholder);
}

.form-input:focus {
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px var(--color-primary-soft);
}

.form-input:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.form-input.input-error {
  border-color: var(--color-danger);
  box-shadow: 0 0 0 3px var(--color-danger-soft);
}

.input-hint {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
  min-height: 18px;
  transition: color var(--transition-fast);
}

.input-hint.error {
  color: var(--color-danger);
}

.error-message {
  padding: var(--space-3) var(--space-4);
  font-size: var(--font-size-sm);
  color: var(--color-danger);
  background-color: var(--color-danger-soft);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-danger-muted);
}

.form-row {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.remember-me {
  display: flex;
  align-items: center;
  gap: var(--space-3);
  font-size: var(--font-size-sm);
  color: var(--color-text-primary);
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
  border-radius: var(--radius-sm);
  background: var(--color-primary);
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  transition: background var(--transition-fast);
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
  background: var(--color-surface-alt);
  border: 1px solid var(--color-border);
}

.remember-me input[type='checkbox']:not(:checked) + .checkbox-custom::after {
  display: none;
}

.remember-me input[type='checkbox']:disabled + .checkbox-custom {
  opacity: 0.6;
  cursor: not-allowed;
}

.remember-text {
  color: var(--color-text-primary);
}

.status-badge {
  padding: var(--space-1) var(--space-3);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-bold);
  color: var(--color-primary);
  background-color: var(--color-primary-soft);
  border-radius: var(--radius-full);
  letter-spacing: 0.4px;
}

.submit-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-2);
  width: 100%;
  padding: var(--space-3) var(--space-5);
  font-size: var(--font-size-base);
  font-weight: var(--font-weight-semibold);
  color: #ffffff;
  background-color: var(--color-primary);
  border: none;
  border-radius: var(--radius-full);
  cursor: pointer;
  transition: background-color var(--transition-fast), transform var(--transition-fast), box-shadow var(--transition-fast);
  letter-spacing: 0.2px;
}

.submit-btn:hover:not(:disabled) {
  background-color: var(--color-primary-hover);
  box-shadow: var(--shadow-md);
  transform: translateY(-1px);
}

.submit-btn:active:not(:disabled) {
  transform: translateY(0);
}

.submit-btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

.alt-link {
  display: flex;
  justify-content: center;
  gap: var(--space-2);
  font-size: var(--font-size-sm);
}

.alt-link span {
  color: var(--color-text-muted);
}

.alt-link a {
  font-weight: var(--font-weight-semibold);
  color: var(--color-primary);
  text-decoration: none;
  transition: color var(--transition-fast);
}

.alt-link a:hover {
  color: var(--color-primary-hover);
  text-decoration: underline;
}

@media (max-width: 1200px) {
  .hero-section {
    width: 500px;
    padding: var(--space-12);
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
    padding: var(--space-10) var(--space-8);
  }

  .hero-text {
    margin-top: var(--space-10);
  }

  .hero-title {
    font-size: var(--font-size-4xl);
    max-width: 100%;
  }

  .hero-subtitle {
    max-width: 100%;
  }

  .hero-card {
    width: 100%;
    max-width: 508px;
    margin-top: var(--space-8);
  }

  .form-section {
    padding: var(--space-12) var(--space-8);
  }

  .form-card {
    width: 100%;
    max-width: 440px;
  }
}

@media (max-width: 640px) {
  .hero-section {
    padding: var(--space-8) var(--space-6);
  }

  .hero-title {
    font-size: var(--font-size-3xl);
  }

  .hero-card {
    padding: var(--space-5);
  }

  .form-section {
    padding: var(--space-8) var(--space-5);
  }

  .form-card {
    padding: var(--space-6);
    border-radius: var(--radius-xl);
  }

  .form-title {
    font-size: var(--font-size-2xl);
  }
}
</style>
