/**
 * 前端运行时配置入口。
 * 后端接口基址：默认 http://localhost:3001；本地覆盖请新增 app.config.local.ts（已 gitignore）。
 */
const localModules = import.meta.glob<{ apiBaseUrl: string }>('./app.config.local.ts', {
  eager: true,
})

const localMod = localModules['./app.config.local.ts']

export const apiBaseUrl =
  localMod && typeof localMod.apiBaseUrl === 'string' && localMod.apiBaseUrl.length > 0
    ? localMod.apiBaseUrl.replace(/\/$/, '')
    : 'http://localhost:3001'

/** axios 请求超时（毫秒） */
export const requestTimeoutMs = 10_000
