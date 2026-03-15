/**
 * 退出登录 API 接口
 * @module api/auth
 */

import { post } from '@/utils/request'

/**
 * 用户退出登录
 * @description POST /api/protected/user/v1/logout
 * @returns 退出结果
 */
export function logout() {
  return post('/protected/user/v1/logout')
}

/**
 * 通过 Refresh Token 退出登录
 * @description POST /api/protected/user/v1/logout/refresh-token
 * @param refreshToken - 刷新令牌
 * @returns 退出结果
 */
export function logoutByRefreshToken(refreshToken: string) {
  return post('/protected/user/v1/logout/refresh-token', {
    refresh_token: refreshToken,
  })
}

/**
 * 撤销所有其他设备的会话
 * @description POST /api/protected/user/v1/sessions/all
 * @returns 撤销结果
 */
export function revokeAllSessions() {
  return post('/protected/user/v1/sessions/all')
}
