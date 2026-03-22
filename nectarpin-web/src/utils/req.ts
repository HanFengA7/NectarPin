import axios, { AxiosError, type AxiosRequestConfig } from 'axios'

import { apiBaseUrl, requestTimeoutMs } from '@/config'
import {
  clearAdminSession,
  getAccessToken,
  getRefreshToken,
  patchAdminSessionTokens,
  type TokenBundleFromApi,
} from '@/lib/admin-auth'

export const REQUEST_CONFIG = {
  baseURL: apiBaseUrl,
  timeout: requestTimeoutMs,
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

/** 仅用于刷新 token，避免与本实例的 401 拦截器形成递归 */
const refreshClient = axios.create({
  baseURL: REQUEST_CONFIG.baseURL,
  timeout: REQUEST_CONFIG.timeout,
})

let refreshPromise: Promise<string | null> | null = null

function performTokenRefresh(): Promise<string | null> {
  const rt = getRefreshToken()
  if (!rt) {
    return Promise.resolve(null)
  }

  if (!refreshPromise) {
    refreshPromise = (async () => {
      try {
        const res = await refreshClient.post<ApiResponse<TokenBundleFromApi>>(
          '/api/public/user/v1/token/refresh',
          { refresh_token: rt },
        )
        const body = res.data
        if (
          body.code === 200 &&
          body.data?.access_token &&
          body.data?.refresh_token &&
          body.data?.expires_at
        ) {
          patchAdminSessionTokens(body.data)
          return body.data.access_token
        }
        return null
      } catch {
        return null
      } finally {
        refreshPromise = null
      }
    })()
  }

  return refreshPromise
}

function rejectFromAxiosError(error: AxiosError<ApiErrorPayload>) {
  const payload = error.response?.data
  return Promise.reject(
    new RequestError(payload?.message ?? '请求失败，请稍后重试。', {
      code: payload?.code,
      status: error.response?.status,
      details: payload?.error,
    }),
  )
}

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
  async (error: AxiosError<ApiErrorPayload>) => {
    const payload = error.response?.data
    const originalRequest = error.config as (AxiosRequestConfig & { _retry?: boolean }) | undefined

    if (error.response?.status === 401 && originalRequest && !originalRequest._retry) {
      const url = originalRequest.url ?? ''

      if (
        url.includes('/api/public/user/v1/login') ||
        url.includes('/api/public/user/v1/register')
      ) {
        return rejectFromAxiosError(error)
      }

      if (url.includes('/api/public/user/v1/token/refresh')) {
        clearAdminSession()
        return rejectFromAxiosError(error)
      }

      originalRequest._retry = true
      const newAccess = await performTokenRefresh()
      if (newAccess) {
        originalRequest.headers = originalRequest.headers ?? {}
        originalRequest.headers.Authorization = `Bearer ${newAccess}`
        return service.request(originalRequest)
      }

      clearAdminSession()
    } else if (error.response?.status === 401) {
      clearAdminSession()
    }

    return rejectFromAxiosError(error)
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
