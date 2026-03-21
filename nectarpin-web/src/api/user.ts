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

export function logoutUser() {
  return post<null>('/api/protected/user/v1/logout')
}

export function logoutUserByRefreshToken(payload: LogoutByRefreshTokenPayload) {
  return post<null, LogoutByRefreshTokenPayload>('/api/protected/user/v1/logout/refresh-token', payload)
}
