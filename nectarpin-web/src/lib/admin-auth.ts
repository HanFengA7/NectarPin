export const ADMIN_SESSION_KEY = 'nectarpin.admin.session'

export const ADMIN_LOGIN_CONFIG = {
  title: 'NectarPin 管理后台',
} as const

export interface AdminUserProfile {
  id: number
  username: string
  email: string
  nickname?: string
  avatar?: string
  status: number
  role: number
}

export interface AdminTokenSet {
  access_token: string
  refresh_token: string
  expires_in: number
}

export interface AdminSession {
  user: AdminUserProfile
  token: AdminTokenSet
}

export function getAdminSession() {
  if (typeof window === 'undefined') {
    return null
  }

  const raw = sessionStorage.getItem(ADMIN_SESSION_KEY)

  if (!raw) {
    return null
  }

  try {
    return JSON.parse(raw) as AdminSession
  } catch {
    sessionStorage.removeItem(ADMIN_SESSION_KEY)
    return null
  }
}

export function saveAdminSession(session: AdminSession) {
  if (typeof window === 'undefined') {
    return
  }

  sessionStorage.setItem(ADMIN_SESSION_KEY, JSON.stringify(session))
}

export function clearAdminSession() {
  if (typeof window === 'undefined') {
    return
  }

  sessionStorage.removeItem(ADMIN_SESSION_KEY)
}

export function getAccessToken() {
  return getAdminSession()?.token.access_token ?? ''
}

export function getRefreshToken() {
  return getAdminSession()?.token.refresh_token ?? ''
}

export function getAdminProfile() {
  return getAdminSession()?.user ?? null
}

export function isAdminAuthenticated() {
  return Boolean(getAccessToken())
}

export function logoutAdmin() {
  clearAdminSession()
}
