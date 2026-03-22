import { get, post } from '@/utils/req'

export interface SiteSocialLink {
  label: string
  href: string
}

export interface SiteTechStackItem {
  name: string
  class_name: string
}

/** 状态旁图标：圆点或 Lucide 示意图标 */
export type SiteStatusIcon = 'dot' | 'circle_dot' | 'radio' | 'activity' | 'none'

/** 状态图标主色（Tailwind 语义色） */
export type SiteStatusIconTone = 'emerald' | 'sky' | 'blue' | 'violet' | 'amber' | 'rose' | 'zinc'

/** 头像右下角叠放小圆标内的图标（与图中相机角标同一位置） */
export type SiteAvatarBadgeIcon =
  | 'camera'
  | 'sparkles'
  | 'coffee'
  | 'heart'
  | 'pen'
  | 'smile'
  | 'none'

export interface SiteHomePayload {
  /** 全站名称（浏览器标题、顶栏/后台品牌等） */
  site_name: string
  /** 头像图片 URL，留空则显示缩写 */
  avatar_url: string
  /** 无图时在圆圈内展示的缩写，1～4 字；留空则前台用展示名称推断 */
  avatar_initials: string
  /** 头像右下角小圆标内图标（叠在头像上，如相机） */
  avatar_badge_icon: SiteAvatarBadgeIcon
  /** 小圆标背景色（与状态图标颜色选项一致） */
  avatar_badge_tone: SiteStatusIconTone | ''
  /** 头像下方状态一行文案（如「正在折腾本站」） */
  status_text: string
  /** 状态文案左侧图标类型 */
  status_icon: SiteStatusIcon
  /** 状态图标颜色 */
  status_icon_tone: SiteStatusIconTone | ''
  hero_name: string
  /** 首页标题下简介，前台 v-html 渲染（HTML 片段） */
  hero_bio: string
  /** 琥珀提示区，Markdown */
  callout_text: string
  social_links: SiteSocialLink[]
  tech_stack: SiteTechStackItem[]
  /** GitHub 登录名（可含大写）；留空则首页不展示贡献图。图表来自 https://ghchart.rshah.org/ */
  github_username: string
  /** 贡献图主题色：6 位 hex，不含 #；非空时请求 …/ghchart.rshah.org/<hex>/<username> */
  github_chart_hex: string
  /** 备案号展示文案；可空隐藏 */
  footer_icp_text: string
  /** 备案号链接（https://beian.miit.gov.cn/…）；可空则不可点 */
  footer_icp_href: string
  /** 站点上线日 YYYY-MM-DD，用于「本站已运行 n 天」；可空隐藏 */
  footer_since: string
  /** 公安（公网）备案号文案；可空隐藏 */
  footer_psb_text: string
  /** 公安备案查询页链接；可空则不可点 */
  footer_psb_href: string
}

/** 前台读取首页站点配置（无需登录） */
export function getPublicSiteHome() {
  return get<SiteHomePayload>('/api/public/site/v1/home')
}

/** 后台保存首页站点配置（需 Bearer） */
export function saveProtectedSiteHome(body: SiteHomePayload) {
  return post<SiteHomePayload>('/api/protected/site/v1/home', body)
}
