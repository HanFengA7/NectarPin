import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { logout as logoutApi, logoutByRefreshToken, revokeAllSessions } from '@/api/auth'
import { getProfile, type UserInfo as ApiUserInfo } from '@/api/user'

// 缓存配置
const CACHE_CONFIG = {
  USER_INFO_TTL: 60 * 60 * 1000, // 用户信息缓存有效期：1 小时
  SAVE_DEBOUNCE_DELAY: 300, // 保存防抖延迟：300ms
}

// 带缓存时间的用户信息类型
interface UserInfoWithCache extends ApiUserInfo {
  _cachedAt?: number
}

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
  
  // 防抖定时器
  let saveDebounceTimer: ReturnType<typeof setTimeout> | null = null

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
          const userData = JSON.parse(storedUser) as UserInfoWithCache
          // 检查缓存是否过期
          if (!isCacheExpired(userData)) {
            user.value = userData
          } else {
            // 缓存过期，清除本地用户信息
            console.log('用户信息缓存已过期，将重新加载')
            const storage = localStorage.getItem('token') ? localStorage : sessionStorage
            storage.removeItem('user')
          }
        } catch (e) {
          console.error('解析用户信息失败:', e)
        }
      }
    }
    
    return !!storedToken
  }
  
  // 检查缓存是否过期
  function isCacheExpired(userData: UserInfoWithCache): boolean {
    if (!userData._cachedAt) return true // 没有时间戳，视为过期
    
    const now = Date.now()
    const cachedAt = userData._cachedAt
    const ttl = CACHE_CONFIG.USER_INFO_TTL
    
    return now - cachedAt > ttl
  }
  
  // 清除过期缓存
  function clearExpiredCache() {
    const storedUser = localStorage.getItem('user') || sessionStorage.getItem('user')
    if (storedUser) {
      try {
        const userData = JSON.parse(storedUser) as UserInfoWithCache
        if (isCacheExpired(userData)) {
          const storage = localStorage.getItem('token') ? localStorage : sessionStorage
          storage.removeItem('user')
          console.log('已清除过期用户缓存')
        }
      } catch (e) {
        console.error('解析用户信息失败:', e)
      }
    }
  }
  
  // 手动清除用户缓存
  function clearUserCache() {
    localStorage.removeItem('user')
    sessionStorage.removeItem('user')
  }
  
  // 保存用户信息到本地存储（防抖优化）
  function saveToStorage() {
    // 清除之前的定时器
    if (saveDebounceTimer) {
      clearTimeout(saveDebounceTimer)
    }
    
    // 设置新的防抖定时器
    saveDebounceTimer = setTimeout(() => {
      const storage = localStorage.getItem('token') ? localStorage : sessionStorage
      if (user.value) {
        // 添加缓存时间戳
        const userDataWithCache: UserInfoWithCache = {
          ...user.value,
          _cachedAt: Date.now(),
        }
        storage.setItem('user', JSON.stringify(userDataWithCache))
      }
    }, CACHE_CONFIG.SAVE_DEBOUNCE_DELAY)
  }
  
  // 立即保存（不防抖）
  function saveToStorageImmediate() {
    if (saveDebounceTimer) {
      clearTimeout(saveDebounceTimer)
      saveDebounceTimer = null
    }
    
    const storage = localStorage.getItem('token') ? localStorage : sessionStorage
    if (user.value) {
      const userDataWithCache: UserInfoWithCache = {
        ...user.value,
        _cachedAt: Date.now(),
      }
      storage.setItem('user', JSON.stringify(userDataWithCache))
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
    // 清除防抖定时器
    if (saveDebounceTimer) {
      clearTimeout(saveDebounceTimer)
      saveDebounceTimer = null
    }
    
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
    clearExpiredCache,
    clearUserCache,
    saveToStorageImmediate,
    logout,
    logoutWithRefreshToken,
    logoutAllOtherSessions,
    clearAuth,
  }
})
