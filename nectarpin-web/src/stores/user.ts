import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { logout as logoutApi, logoutByRefreshToken, revokeAllSessions } from '@/api/auth'
import { getProfile } from '@/api/user'

export interface UserInfo {
  id: number
  username: string
  email: string
  nickname: string
  avatar: string
  status: number
  role: number
}

export const useUserStore = defineStore('user', () => {
  const user = ref<UserInfo | null>(null)
  const token = ref<string | null>(null)
  const refreshToken = ref<string | null>(null)

  const isLoggedIn = computed(() => !!token.value)
  const isAdmin = computed(() => user.value?.role === 1)

  // 设置用户信息
  function setUser(userData: UserInfo) {
    user.value = userData
    // 保存到本地存储
    saveToStorage()
  }

  // 设置访问令牌和刷新令牌
  function setTokens(accessToken: string, refresh: string) {
    token.value = accessToken
    refreshToken.value = refresh
  }

  // 从本地存储加载用户信息和令牌
  function loadFromStorage() {
    const storedToken = localStorage.getItem('token') || sessionStorage.getItem('token')
    const storedRefresh = localStorage.getItem('refresh_token') || sessionStorage.getItem('refresh_token')
    const storedUser = localStorage.getItem('user') || sessionStorage.getItem('user')
    
    if (storedToken) {
      token.value = storedToken
      refreshToken.value = storedRefresh
      // 尝试从本地存储加载用户信息
      if (storedUser) {
        try {
          user.value = JSON.parse(storedUser)
        } catch (e) {
          console.error('解析用户信息失败:', e)
        }
      }
    }
    
    return !!storedToken
  }
  
  // 保存用户信息到本地存储
  function saveToStorage() {
    const storage = localStorage.getItem('token') ? localStorage : sessionStorage
    if (user.value) {
      storage.setItem('user', JSON.stringify(user.value))
    }
  }

  // 从服务端加载用户信息
  async function loadUserProfile() {
    try {
      const response = await getProfile()
      // response 已经是 ApiResponse<UserInfo> 格式，需要提取 data 字段
      const userData = response.data as UserInfo
      user.value = userData
      // 保存到本地存储
      saveToStorage()
      return userData
    } catch (error) {
      console.error('加载用户资料失败:', error)
      throw error
    }
  }

  // 本地清除令牌和用户信息
  function clearAuth() {
    user.value = null
    token.value = null
    refreshToken.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('refresh_token')
    localStorage.removeItem('user')
    sessionStorage.removeItem('token')
    sessionStorage.removeItem('refresh_token')
    sessionStorage.removeItem('user')
  }

  // 调用服务端退出登录接口，并清除本地数据
  async function logout() {
    // 保存 token 的引用，因为稍后可能会清除
    const currentToken = token.value
    const currentRefreshToken = refreshToken.value

    try {
      // 先尝试标准退出（此时 token 还存在，请求拦截器会自动添加）
      await logoutApi()
      console.log('服务端退出成功')
    } catch (error) {
      // Token 可能已过期，尝试使用 Refresh Token 退出
      console.warn('标准退出失败（可能 Token 已过期），尝试使用 Refresh Token 退出')
      if (currentRefreshToken) {
        try {
          await logoutByRefreshToken(currentRefreshToken)
          console.log('Refresh Token 退出成功')
        } catch (refreshError) {
          console.error('Refresh Token 退出也失败:', refreshError)
        }
      }
    } finally {
      // 无论成功失败，都要清除本地数据
      clearAuth()
    }
  }

  // 通过 Refresh Token 退出登录
  async function logoutWithRefreshToken() {
    if (!refreshToken.value) {
      clearAuth()
      return
    }

    try {
      await logoutByRefreshToken(refreshToken.value)
    } catch (error) {
      console.error('退出登录失败:', error)
    } finally {
      clearAuth()
    }
  }

  // 撤销所有其他设备的会话
  async function logoutAllOtherSessions() {
    try {
      await revokeAllSessions()
    } catch (error) {
      console.error('撤销会话失败:', error)
      throw error
    }
  }

  return {
    user,
    token,
    refreshToken,
    isLoggedIn,
    isAdmin,
    setUser,
    setTokens,
    loadFromStorage,
    loadUserProfile,
    logout,
    logoutWithRefreshToken,
    logoutAllOtherSessions,
    clearAuth,
  }
})
