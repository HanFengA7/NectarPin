/**
 * 用户相关 API 接口
 * @module api/user
 */

import { post, get } from '@/utils/request'

/**
 * 登录请求参数
 */
export interface LoginParams {
  /** 账号（邮箱或用户名） */
  account: string
  /** 密码（MD5加密后的32位字符串） */
  password: string
}

/**
 * 令牌响应
 */
export interface TokenResponse {
  /** 访问令牌 */
  access_token: string
  /** 刷新令牌 */
  refresh_token: string
  /** 过期时间 */
  expires_at: string
  /** 令牌类型 */
  token_type: string
}

/**
 * 用户信息
 */
export interface UserInfo {
  /** 用户 ID */
  id: number
  /** 用户名 */
  username: string
  /** 邮箱 */
  email: string
  /** 昵称 */
  nickname: string
  /** 头像 URL */
  avatar: string
  /** 账户状态 */
  status: number
  /** 用户角色 */
  role: number
}

/**
 * 登录返回结果
 */
export interface LoginResult {
  /** 用户信息 */
  user: UserInfo
  /** 令牌信息 */
  token: TokenResponse
}

/**
 * 注册请求参数
 */
export interface RegisterParams {
  /** 用户名，必填，3-50字符 */
  username: string
  /** 密码（MD5加密后的32位字符串） */
  password: string
  /** 邮箱，必填，邮箱格式 */
  email: string
  /** 昵称，可选，最大100字符 */
  nickname?: string
}

/**
 * 用户登录
 * @description POST /api/public/user/v1/login
 * @param params - 登录参数（account: 邮箱或用户名, password: MD5加密密码）
 * @returns 登录结果，包含用户信息和令牌
 */
export function login(params: LoginParams) {
  return post<LoginResult>('/public/user/v1/login', params)
}

/**
 * 用户注册
 * @description POST /api/public/user/v1/register
 * @param params - 注册参数
 * @returns 用户信息
 */
export function register(params: RegisterParams) {
  return post<UserInfo>('/public/user/v1/register', params)
}

/**
 * 获取当前登录用户资料
 * @description GET /api/protected/user/v1/profile
 * @returns 用户资料
 */
export function getProfile() {
  return get<{ user_id: number }>('/protected/user/v1/profile')
}
