import axios, { AxiosError, type AxiosRequestConfig } from 'axios'

import { clearAdminSession, getAccessToken } from '@/lib/admin-auth'

export const REQUEST_CONFIG = {
  baseURL: 'http://localhost:3001',
  timeout: 10000,
} as const

export interface ApiResponse<T> {
  code: number
  message: string
  data: T
}

export interface ApiErrorPayload {
  code?: number
  message?: string
  error?: string
}

export class RequestError extends Error {
  code?: number
  status?: number
  details?: string

  constructor(message: string, options?: { code?: number; status?: number; details?: string }) {
    super(message)
    this.name = 'RequestError'
    this.code = options?.code
    this.status = options?.status
    this.details = options?.details
  }
}

const service = axios.create({
  baseURL: REQUEST_CONFIG.baseURL,
  timeout: REQUEST_CONFIG.timeout,
})

service.interceptors.request.use((config) => {
  const token = getAccessToken()

  if (token) {
    config.headers = config.headers ?? {}
    config.headers.Authorization = `Bearer ${token}`
  }

  return config
})

service.interceptors.response.use(
  (response) => response.data,
  (error: AxiosError<ApiErrorPayload>) => {
    const payload = error.response?.data

    if (error.response?.status === 401) {
      clearAdminSession()
    }

    return Promise.reject(
      new RequestError(payload?.message ?? '请求失败，请稍后重试。', {
        code: payload?.code,
        status: error.response?.status,
        details: payload?.error,
      }),
    )
  },
)

export function request<T>(config: AxiosRequestConfig) {
  return service.request<ApiResponse<T>, ApiResponse<T>>(config)
}

export function get<T>(url: string, config?: AxiosRequestConfig) {
  return request<T>({
    ...config,
    method: 'get',
    url,
  })
}

export function post<TResponse, TBody = unknown>(
  url: string,
  data?: TBody,
  config?: AxiosRequestConfig<TBody>,
) {
  return request<TResponse>({
    ...config,
    method: 'post',
    url,
    data,
  })
}

export default service
