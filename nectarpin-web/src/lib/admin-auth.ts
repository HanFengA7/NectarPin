export const ADMIN_SESSION_KEY = 'nectarpin.admin.session'

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
  /** 由 access 的 `expires_at` 推算的剩余秒数，供展示或调度；无则可为 0 */
  expires_in: number
}

/** 与登录/刷新接口返回的 token 字段一致（`expires_at` 为 RFC3339） */
export interface TokenBundleFromApi {
  access_token: string
  refresh_token: string
  expires_at: string
  token_type?: string
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

/** 将登录接口返回的 `user` + `token` 规范为本地会话（从 `expires_at` 推算 `expires_in`） */
export function buildAdminSessionFromLogin(data: {
  user: AdminUserProfile
  token: {
    access_token: string
    refresh_token: string
    expires_in?: number
    expires_at?: string
  }
}): AdminSession {
  const expMs = data.token.expires_at ? Date.parse(data.token.expires_at) : Number.NaN
  const expires_in =
    typeof data.token.expires_in === 'number' && Number.isFinite(data.token.expires_in)
      ? data.token.expires_in
      : Number.isFinite(expMs)
        ? Math.max(0, Math.floor((expMs - Date.now()) / 1000))
        : 0

  return {
    user: data.user,
    token: {
      access_token: data.token.access_token,
      refresh_token: data.token.refresh_token,
      expires_in,
    },
  }
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

/** 合并更新本地会话中的用户信息（保存资料后同步侧栏展示） */
export function patchAdminSessionUser(partial: Partial<AdminUserProfile>) {
  const session = getAdminSession()
  if (!session) {
    return
  }

  session.user = { ...session.user, ...partial }
  saveAdminSession(session)
}

export function patchAdminSessionTokens(payload: TokenBundleFromApi) {
  const session = getAdminSession()
  if (!session) {
    return
  }

  const expMs = Date.parse(payload.expires_at)
  const expires_in = Number.isFinite(expMs)
    ? Math.max(0, Math.floor((expMs - Date.now()) / 1000))
    : 0

  session.token = {
    access_token: payload.access_token,
    refresh_token: payload.refresh_token,
    expires_in,
  }
  saveAdminSession(session)
}

export function isAdminAuthenticated() {
  return Boolean(getAccessToken())
}

export function logoutAdmin() {
  clearAdminSession()
}
