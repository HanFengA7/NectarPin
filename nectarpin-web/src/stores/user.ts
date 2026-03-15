import { defineStore } from 'pinia'
import { ref, computed } from 'vue'

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
    
    if (storedToken) {
      token.value = storedToken
      refreshToken.value = storedRefresh
    }
  }
  
  // 注销用户，清除所有令牌和用户信息
  function logout() {
    user.value = null
    token.value = null
    refreshToken.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('refresh_token')
    sessionStorage.removeItem('token')
    sessionStorage.removeItem('refresh_token')
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
    logout,
  }
})
