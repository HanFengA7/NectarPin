import md5 from 'crypto-js/md5'

import { get, post } from '@/utils/req'

export interface UserProfile {
  id: number
  username: string
  email: string
  nickname?: string
  avatar?: string
  status: number
  role: number
}

export interface UserToken {
  access_token: string
  refresh_token: string
  expires_in: number
}

export interface LoginPayload {
  account: string
  password: string
}

export interface LoginData {
  user: UserProfile
  token: UserToken
}

export interface LogoutByRefreshTokenPayload {
  refresh_token: string
}

function encryptPassword(password: string) {
  return md5(password).toString()
}

export function loginUser(payload: LoginPayload) {
  return post<LoginData, LoginPayload>('/api/public/user/v1/login', {
    account: payload.account.trim(),
    password: encryptPassword(payload.password),
  })
}

export function getCurrentUserProfile() {
  return get<UserProfile>('/api/protected/user/v1/profile')
}

export interface UpdateProfilePayload {
  nickname?: string
  email: string
  avatar?: string
}

/** POST /api/protected/user/v1/profile — 与 docs/openapi.json 中 updateProfile 一致 */
export function updateCurrentUserProfile(payload: UpdateProfilePayload) {
  return post<UserProfile, UpdateProfilePayload>('/api/protected/user/v1/profile', payload)
}

export interface ChangePasswordPayload {
  old_password: string
  new_password: string
}

/** POST /api/protected/user/v1/password — 与 docs/openapi.json 中 changePassword 一致（字段为 MD5 hex） */
export function changeUserPassword(payload: ChangePasswordPayload) {
  return post<null, ChangePasswordPayload>('/api/protected/user/v1/password', payload)
}

export function logoutUser() {
  return post<null>('/api/protected/user/v1/logout')
}

export function logoutUserByRefreshToken(payload: LogoutByRefreshTokenPayload) {
  return post<null, LogoutByRefreshTokenPayload>('/api/protected/user/v1/logout/refresh-token', payload)
}

export interface UserSessionItem {
  id: number
  created_at: string
  expires_at: string
  is_current: boolean
}

/** GET /sessions；传入 refresh 时通过 X-Refresh-Token 标记 is_current */
export function listUserSessions(refreshToken?: string | null) {
  const headers: Record<string, string> = {}
  if (refreshToken) {
    headers['X-Refresh-Token'] = refreshToken
  }
  return get<{ items: UserSessionItem[] }>('/api/protected/user/v1/sessions', {
    headers,
  })
}

/** POST /sessions/revoke */
export function revokeUserSession(sessionId: number) {
  return post<null, { session_id: number }>('/api/protected/user/v1/sessions/revoke', {
    session_id: sessionId,
  })
}

/** POST /sessions/all — 保留当前 refresh，撤销其余会话 */
export function revokeOtherUserSessions(refreshToken: string) {
  return post<null, { refresh_token: string }>('/api/protected/user/v1/sessions/all', {
    refresh_token: refreshToken,
  })
}
