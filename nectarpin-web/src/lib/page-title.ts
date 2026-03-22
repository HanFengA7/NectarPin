import { useSiteStore } from '@/stores/site'

/** 仅描述设置标题所需字段，避免依赖 vue-router 各版本导出名差异 */
export interface RouteSnapshotForTitle {
  path: string
  matched: ReadonlyArray<{
    meta: Record<string, unknown>
  }>
}

export function formatDocumentTitle(options: {
  pageTitle?: string
  siteName: string
  variant: 'public' | 'admin'
}) {
  const { pageTitle, siteName, variant } = options
  const suffix = variant === 'public' ? siteName : `${siteName} 管理后台`
  if (!pageTitle?.trim()) {
    return suffix
  }
  return `${pageTitle.trim()} · ${suffix}`
}

export function applyDocumentTitleFromRoute(to: RouteSnapshotForTitle) {
  const siteStore = useSiteStore()
  const siteName = siteStore.siteName

  const deepest = to.matched[to.matched.length - 1]
  const meta = deepest?.meta
  if (meta?.titleFromPage === true) {
    return
  }

  const pageTitle = typeof meta?.title === 'string' ? meta.title : undefined
  const variant: 'public' | 'admin' = to.path.startsWith('/admin') ? 'admin' : 'public'

  document.title = formatDocumentTitle({ pageTitle, siteName, variant })
}
