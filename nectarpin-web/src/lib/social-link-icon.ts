/** 页脚 / 首页社交按钮图标归类（按名称与链接启发式判断） */
export type SocialLinkIconKind = 'github' | 'mail' | 'qq' | 'generic'

export function socialLinkIconKind(label: string, href: string): SocialLinkIconKind {
  const l = label.toLowerCase()
  const h = href.toLowerCase()
  if (l.includes('github') || h.includes('github.com')) return 'github'
  if (
    l.includes('mail') ||
    l.includes('email') ||
    label.includes('邮箱') ||
    h.startsWith('mailto:')
  ) {
    return 'mail'
  }
  if (l.includes('qq') || h.includes('qq.com') || h.includes('qm.qq.com')) return 'qq'
  return 'generic'
}
