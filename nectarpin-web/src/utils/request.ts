/**
 * HTTP 请求封装
 * @module utils/request
 * @description 基于 axios 封装的请求工具，支持请求/响应拦截、自动添加 token、统一错误处理
 */

import axios, {
  type AxiosInstance,
  type AxiosRequestConfig,
  type AxiosResponse,
  type InternalAxiosRequestConfig,
} from 'axios'

/**
 * Axios 实例
 * @description 配置了基础 URL、超时时间和默认请求头
 */
const request: AxiosInstance = axios.create({
  /** API 基础地址，优先从环境变量读取 */
  baseURL: import.meta.env.VITE_API_BASE_URL || '/api',
  /** 请求超时时间（毫秒） */
  timeout: 10000,
  /** 默认请求头 */
  headers: {
    'Content-Type': 'application/json',
  },
})

/**
 * 请求拦截器
 * @description 自动在请求头中添加 Authorization token
 */
request.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    const token = localStorage.getItem('token')
    if (token && config.headers) {
      config.headers.Authorization = `Bearer ${token}`
    }
    return config
  },
  (error) => {
    return Promise.reject(error)
  },
)

/** 公开接口列表，这些接口的 401 错误不触发跳转登录页 */
const publicEndpoints = ['/public/user/v1/login', '/public/user/v1/register']

/**
 * 响应拦截器
 * @description 统一处理响应数据，401 状态自动跳转登录页（公开接口除外）
 */
request.interceptors.response.use(
  (response: AxiosResponse) => {
    return response.data
  },
  (error) => {
    if (error.response) {
      const { status, config } = error.response
      const isPublicEndpoint = publicEndpoints.some((endpoint) => config?.url?.includes(endpoint))

      /** 401 未授权，清除 token 并跳转登录页（公开接口除外） */
      if (status === 401 && !isPublicEndpoint) {
        localStorage.removeItem('token')
        localStorage.removeItem('refresh_token')
        localStorage.removeItem('user')
        sessionStorage.removeItem('token')
        sessionStorage.removeItem('refresh_token')
        sessionStorage.removeItem('user')
        window.location.href = '/admin/login'
      }
    }
    return Promise.reject(error)
  },
)

/**
 * 统一 API 响应格式
 * @template T - 响应数据类型
 */
export interface ApiResponse<T = unknown> {
  /** 状态码 */
  code: number
  /** 响应消息 */
  message: string
  /** 响应数据 */
  data: T
}

/**
 * 发送 GET 请求
 * @template T - 响应数据类型
 * @param url - 请求地址
 * @param config - 请求配置
 * @returns Promise<ApiResponse<T>>
 */
export function get<T>(url: string, config?: AxiosRequestConfig): Promise<ApiResponse<T>> {
  return request.get(url, config)
}

/**
 * 发送 POST 请求
 * @template T - 响应数据类型
 * @param url - 请求地址
 * @param data - 请求体数据
 * @param config - 请求配置
 * @returns Promise<ApiResponse<T>>
 */
export function post<T>(
  url: string,
  data?: unknown,
  config?: AxiosRequestConfig,
): Promise<ApiResponse<T>> {
  return request.post(url, data, config)
}

/**
 * 发送 PUT 请求
 * @template T - 响应数据类型
 * @param url - 请求地址
 * @param data - 请求体数据
 * @param config - 请求配置
 * @returns Promise<ApiResponse<T>>
 */
export function put<T>(
  url: string,
  data?: unknown,
  config?: AxiosRequestConfig,
): Promise<ApiResponse<T>> {
  return request.put(url, data, config)
}

/**
 * 发送 DELETE 请求
 * @template T - 响应数据类型
 * @param url - 请求地址
 * @param config - 请求配置
 * @returns Promise<ApiResponse<T>>
 */
export function del<T>(url: string, config?: AxiosRequestConfig): Promise<ApiResponse<T>> {
  return request.delete(url, config)
}

export default request
