<script setup lang="ts">
import type { Component } from 'vue'
import { computed, onMounted, ref, shallowRef, watch } from 'vue'
import { RouterLink, useRoute } from 'vue-router'
import { useDark } from '@vueuse/core'
import { MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { Activity, Camera, CircleDot, Coffee, Github, Heart, Mail, PenLine, Radio, Smile, Sparkles } from 'lucide-vue-next'
import { listPublicArticles, type ArticleItem } from '@/api/article'
import { getPublicSiteHome, type SiteAvatarBadgeIcon, type SiteStatusIcon, type SiteStatusIconTone } from '@/api/site'
import { cn } from '@/lib/utils'
import { RequestError } from '@/utils/req'

/** 公开贡献图嵌入，见 https://ghchart.rshah.org/ */
const GH_CONTRIBUTION_CHART_ORIGIN = 'https://ghchart.rshah.org'

const isDark = useDark()
const mdPreviewTheme = computed(() => (isDark.value ? 'dark' : 'light'))

const route = useRoute()

const isHome = computed(() => route.name === 'index')

const pageTitle = computed(() => {
  const map: Record<string, string> = {
    'index-articles': '文章',
    'index-projects': '项目',
    'index-friends': '朋友',
    'index-about': '关于',
  }
  return map[route.name as string] ?? '页面'
})

const avatarUrl = ref('')
const avatarInitialsConfigured = ref('')
const statusText = ref('')
const statusIcon = ref<SiteStatusIcon>('dot')
const statusIconTone = ref<SiteStatusIconTone>('emerald')
const avatarBadgeIcon = ref<SiteAvatarBadgeIcon>('camera')
const avatarBadgeTone = ref<SiteStatusIconTone>('blue')
const avatarLoadFailed = ref(false)

const VALID_STATUS_TONES: SiteStatusIconTone[] = ['emerald', 'sky', 'blue', 'violet', 'amber', 'rose', 'zinc']

const VALID_AVATAR_BADGE_ICONS: SiteAvatarBadgeIcon[] = ['camera', 'sparkles', 'coffee', 'heart', 'pen', 'smile', 'none']

function parseAvatarBadgeIcon(raw: unknown): SiteAvatarBadgeIcon {
  const s = typeof raw === 'string' ? raw.toLowerCase().trim() : ''
  return VALID_AVATAR_BADGE_ICONS.includes(s as SiteAvatarBadgeIcon) ? (s as SiteAvatarBadgeIcon) : 'camera'
}

function parseAvatarBadgeTone(raw: unknown, badge: SiteAvatarBadgeIcon): SiteStatusIconTone {
  if (badge === 'none')
    return 'blue'
  const s = typeof raw === 'string' ? raw.toLowerCase().trim() : ''
  return VALID_STATUS_TONES.includes(s as SiteStatusIconTone) ? (s as SiteStatusIconTone) : 'blue'
}

const AVATAR_BADGE_BG: Record<SiteStatusIconTone, string> = {
  emerald: 'bg-emerald-500',
  sky: 'bg-sky-500',
  blue: 'bg-blue-500',
  violet: 'bg-violet-500',
  amber: 'bg-amber-500',
  rose: 'bg-rose-500',
  zinc: 'bg-zinc-600',
}

const avatarBadgeGlyph = computed(() => {
  switch (avatarBadgeIcon.value) {
    case 'camera':
      return Camera
    case 'sparkles':
      return Sparkles
    case 'coffee':
      return Coffee
    case 'heart':
      return Heart
    case 'pen':
      return PenLine
    case 'smile':
      return Smile
    default:
      return null
  }
})

const VALID_STATUS_ICONS: SiteStatusIcon[] = ['dot', 'circle_dot', 'radio', 'activity', 'none']

function parseStatusIcon(raw: unknown): SiteStatusIcon {
  const s = typeof raw === 'string' ? raw.toLowerCase().trim() : ''
  return VALID_STATUS_ICONS.includes(s as SiteStatusIcon) ? (s as SiteStatusIcon) : 'dot'
}

function parseStatusTone(raw: unknown, icon: SiteStatusIcon): SiteStatusIconTone {
  if (icon === 'none')
    return 'emerald'
  const s = typeof raw === 'string' ? raw.toLowerCase().trim() : ''
  return VALID_STATUS_TONES.includes(s as SiteStatusIconTone) ? (s as SiteStatusIconTone) : 'emerald'
}

const STATUS_DOT_RING: Record<SiteStatusIconTone, string> = {
  emerald: 'bg-emerald-500 ring-emerald-500/25 dark:ring-emerald-400/20',
  sky: 'bg-sky-500 ring-sky-500/25 dark:ring-sky-400/20',
  blue: 'bg-blue-500 ring-blue-500/25 dark:ring-blue-400/20',
  violet: 'bg-violet-500 ring-violet-500/25 dark:ring-violet-400/25',
  amber: 'bg-amber-500 ring-amber-500/25 dark:ring-amber-400/25',
  rose: 'bg-rose-500 ring-rose-500/25 dark:ring-rose-400/25',
  zinc: 'bg-zinc-500 ring-zinc-400/30 dark:ring-zinc-500/40',
}

const STATUS_LUCIDE_CLASS: Record<SiteStatusIconTone, string> = {
  emerald: 'size-4 shrink-0 text-emerald-500',
  sky: 'size-4 shrink-0 text-sky-500',
  blue: 'size-4 shrink-0 text-blue-500',
  violet: 'size-4 shrink-0 text-violet-500',
  amber: 'size-4 shrink-0 text-amber-500',
  rose: 'size-4 shrink-0 text-rose-500',
  zinc: 'size-4 shrink-0 text-zinc-500',
}

const statusGlyph = computed(() => {
  switch (statusIcon.value) {
    case 'circle_dot':
      return CircleDot
    case 'radio':
      return Radio
    case 'activity':
      return Activity
    default:
      return null
  }
})

watch(avatarUrl, () => {
  avatarLoadFailed.value = false
})

const avatarFallbackText = computed(() => {
  const cfg = avatarInitialsConfigured.value.trim()
  if (cfg)
    return Array.from(cfg).slice(0, 4).join('')
  const n = heroName.value.trim()
  const chars = Array.from(n)
  if (chars.length >= 2)
    return chars.slice(0, 2).join('')
  if (chars.length === 1)
    return chars[0] ?? ''
  return 'NP'
})

const heroName = ref('NectarPin')
const heroBio = ref(
  '这里记录学习与工程实践，关注基础设施与前后端协作，也折腾文档、工具链与一点点设计。欢迎随便看看。',
)
const calloutText = ref(
  '喜欢把想法写成文字：技术笔记、随笔与编程相关的小结；也关注人工智能与效率工具。订阅更新可使用 RSS（即将提供）。',
)

function socialIconForLabel(label: string): Component | undefined {
  const l = label.toLowerCase()
  if (l.includes('github'))
    return Github
  if (l.includes('mail') || l.includes('email') || label.includes('邮箱'))
    return Mail
  return undefined
}

const socialLinks = ref<{ label: string; href: string; icon?: Component }[]>([
  { label: 'GitHub', href: 'https://github.com', icon: Github },
  { label: 'Email', href: 'mailto:hello@example.com', icon: Mail },
  { label: 'QQ', href: '#' },
])

const techStack = ref<{ name: string; className: string }[]>([
  { name: 'Vue', className: 'bg-emerald-600' },
  { name: 'TypeScript', className: 'bg-blue-600' },
  { name: 'Vite', className: 'bg-violet-600' },
  { name: 'Tailwind CSS', className: 'bg-sky-600' },
  { name: 'Go', className: 'bg-cyan-700' },
  { name: 'Docker', className: 'bg-sky-500' },
  { name: 'Node.js', className: 'bg-green-700' },
  { name: 'Markdown', className: 'bg-zinc-600' },
])

const githubUsername = ref('')
/** 6 位小写 hex，无 #；与 https://ghchart.rshah.org/<HEX>/username 一致 */
const githubChartHex = ref('')
const ghContributionChartSrc = computed(() => {
  const u = githubUsername.value.trim()
  if (!u)
    return ''
  const hex = githubChartHex.value.trim().toLowerCase().replace(/^#/, '')
  if (/^[0-9a-f]{6}$/.test(hex))
    return `${GH_CONTRIBUTION_CHART_ORIGIN}/${encodeURIComponent(hex)}/${encodeURIComponent(u)}`
  return `${GH_CONTRIBUTION_CHART_ORIGIN}/${encodeURIComponent(u)}`
})
const githubProfileHref = computed(() => {
  const u = githubUsername.value.trim()
  if (!u)
    return ''
  return `https://github.com/${encodeURIComponent(u)}`
})

/** 后台未填写或清空保存后不展示对应区块 */
const showHeroBio = computed(() => heroBio.value.trim().length > 0)
const showCallout = computed(() => calloutText.value.trim().length > 0)

async function loadSiteHomeFromApi() {
  if (route.name !== 'index')
    return
  try {
    const { data } = await getPublicSiteHome()
    avatarUrl.value = (data.avatar_url ?? '').trim()
    avatarInitialsConfigured.value = data.avatar_initials ?? ''
    {
      const b = parseAvatarBadgeIcon(data.avatar_badge_icon)
      avatarBadgeIcon.value = b
      avatarBadgeTone.value = parseAvatarBadgeTone(data.avatar_badge_tone, b)
    }
    statusText.value = (data.status_text ?? '').trim()
    {
      const ic = parseStatusIcon(data.status_icon)
      statusIcon.value = ic
      statusIconTone.value = parseStatusTone(data.status_icon_tone, ic)
    }
    if (data.hero_name)
      heroName.value = data.hero_name
    heroBio.value = data.hero_bio ?? ''
    calloutText.value = data.callout_text ?? ''
    socialLinks.value = (data.social_links ?? []).map(s => ({
      label: s.label,
      href: s.href,
      icon: socialIconForLabel(s.label),
    }))
    techStack.value = (data.tech_stack ?? []).map(t => ({
      name: t.name,
      className: t.class_name || 'bg-zinc-600',
    }))
    githubUsername.value = (data.github_username ?? '').trim()
    githubChartHex.value = (data.github_chart_hex ?? '').trim().replace(/^#/, '').toLowerCase()
  }
  catch {
    // 接口不可用时保留本地默认文案
  }
}

watch(
  () => route.name,
  (name) => {
    if (name === 'index')
      void loadSiteHomeFromApi()
  },
  { immediate: true },
)

const TIMELINE_PAGE_SIZE = 12

const DATE_FORMATTER = new Intl.DateTimeFormat('zh-CN', {
  year: 'numeric',
  month: 'long',
  day: 'numeric',
})

interface TimelineEntry {
  id: number
  date: string
  title: string
  slug: string
}

function formatArticleDate(a: ArticleItem) {
  const raw = a.published_at ?? a.created_at
  const d = new Date(raw)
  if (Number.isNaN(d.getTime()))
    return raw
  return DATE_FORMATTER.format(d)
}

const timeline = shallowRef<TimelineEntry[]>([])
const timelineLoading = ref(true)
const timelineError = ref('')

onMounted(async () => {
  timelineLoading.value = true
  timelineError.value = ''
  try {
    const res = await listPublicArticles({
      page: 1,
      page_size: TIMELINE_PAGE_SIZE,
      status: 1,
    })
    timeline.value = res.data.items.map(a => ({
      id: a.id,
      date: formatArticleDate(a),
      title: a.title,
      slug: a.slug,
    }))
  }
  catch (e) {
    timeline.value = []
    timelineError.value =
      e instanceof RequestError ? e.message : '无法加载文章列表'
  }
  finally {
    timelineLoading.value = false
  }
})

</script>

<template>
  <div v-if="!isHome" class="w-full px-4 py-16 sm:px-6 lg:px-10">
    <div class="mx-auto max-w-3xl">
      <h1 class="text-2xl font-semibold tracking-tight">
        {{ pageTitle }}
      </h1>
      <p class="mt-2 text-muted-foreground">
        内容建设中
      </p>
    </div>
  </div>

  <div
    v-else
    class="w-full px-4 pb-16 pt-8 sm:px-6 sm:pt-10 lg:px-10 lg:pt-12"
  >
    <div
      class="mx-auto flex w-full max-w-6xl flex-col gap-10 lg:flex-row lg:gap-12 xl:max-w-7xl xl:gap-16"
    >
      <!-- 左侧栏：头像 + 更新日志 -->
      <aside
        class="flex w-full shrink-0 flex-col items-center lg:sticky lg:top-14 lg:w-[min(100%,280px)] lg:items-start xl:w-[300px]"
      >
        <div class="flex w-full shrink-0 flex-col items-center gap-3 lg:items-start">
          <div class="relative shrink-0">
            <div
              class="size-36 overflow-hidden rounded-full border-4 border-background bg-gradient-to-br from-blue-100 to-violet-200 shadow-lg ring-2 ring-border/60 dark:from-blue-950 dark:to-violet-950 dark:ring-border sm:size-40"
            >
              <img
                v-if="avatarUrl && !avatarLoadFailed"
                :src="avatarUrl"
                :alt="heroName"
                class="size-full object-cover"
                @error="avatarLoadFailed = true"
              >
              <div
                v-else
                class="flex size-full items-center justify-center text-3xl font-bold tracking-tight text-blue-700/80 dark:text-blue-200/90"
              >
                {{ avatarFallbackText }}
              </div>
            </div>
            <div
              v-if="avatarBadgeIcon !== 'none' && avatarBadgeGlyph"
              class="absolute bottom-1 right-1 flex size-9 items-center justify-center rounded-full border-2 border-background text-white shadow-md dark:border-zinc-950"
              :class="AVATAR_BADGE_BG[avatarBadgeTone]"
              aria-hidden="true"
            >
              <component
                :is="avatarBadgeGlyph"
                class="size-4"
                :stroke-width="2"
              />
            </div>
          </div>

          <div
            v-if="statusText"
            class="flex max-w-full items-center justify-center gap-2 text-sm text-muted-foreground lg:justify-start"
          >
            <span
              v-if="statusIcon === 'dot'"
              class="size-2 shrink-0 rounded-full shadow-sm ring-2"
              :class="STATUS_DOT_RING[statusIconTone]"
              aria-hidden="true"
            />
            <component
              :is="statusGlyph"
              v-else-if="statusGlyph"
              :class="STATUS_LUCIDE_CLASS[statusIconTone]"
              :stroke-width="2"
              aria-hidden="true"
            />
            <span class="max-w-[min(100%,16rem)] text-center leading-snug lg:max-w-none lg:text-left">
              {{ statusText }}
            </span>
          </div>
        </div>

        <section class="mt-8 w-full max-w-md lg:max-w-none">
          <h2 class="mb-4 text-base font-semibold tracking-tight">
            更新日志
          </h2>
          <!-- 轴线与圆点共用 left-2 + -translate-x-1/2，保证 1px 竖线穿过空心圆中心 -->
          <div class="relative">
            <div
              v-if="!timelineLoading && timeline.length > 0"
              class="pointer-events-none absolute top-2 bottom-2 left-2 w-px -translate-x-1/2 bg-border"
              aria-hidden="true"
            />
            <p
              v-if="timelineLoading"
              class="pl-1 text-sm text-muted-foreground"
            >
              加载中…
            </p>
            <p
              v-else-if="timelineError"
              class="pl-1 text-sm text-destructive"
            >
              {{ timelineError }}
            </p>
            <p
              v-else-if="timeline.length === 0"
              class="pl-1 text-sm text-muted-foreground"
            >
              暂无已发布文章
            </p>
            <ul v-else class="space-y-6">
              <li
                v-for="item in timeline"
                :key="item.id"
                class="relative pl-8"
              >
                <span
                  class="absolute top-1.5 left-2 size-4 -translate-x-1/2 rounded-full border-2 border-blue-500 bg-background ring-2 ring-background dark:bg-background dark:ring-background"
                  aria-hidden="true"
                />
                <time
                  class="block text-sm font-medium text-foreground"
                >{{ item.date }}</time>
                <RouterLink
                  :to="{ name: 'index-post-detail', params: { slug: item.slug } }"
                  class="mt-1 block text-sm leading-relaxed text-muted-foreground transition-colors hover:text-blue-600 dark:hover:text-blue-400"
                >
                  {{ item.title }}
                </RouterLink>
              </li>
            </ul>
          </div>
        </section>
      </aside>

      <!-- 右侧主内容 -->
      <main class="min-w-0 flex-1 space-y-12">
        <section class="space-y-5">
          <h1
            class="text-3xl font-semibold tracking-tight sm:text-4xl lg:text-[2.25rem] lg:leading-tight"
          >
            <span class="text-foreground">Hello, I'm </span>
            <span
              class="bg-gradient-to-r from-blue-600 via-violet-600 to-blue-500 bg-clip-text font-bold text-transparent dark:from-blue-400 dark:via-violet-400 dark:to-sky-400"
            >{{ heroName }}</span>
            <span class="ml-1 inline-block" aria-hidden="true">👋</span>
          </h1>
          <!-- 简介由后台配置为 HTML，仅管理员可写 -->
          <div
            v-if="showHeroBio"
            class="hero-bio-html max-w-2xl text-base leading-7 text-muted-foreground sm:text-[1.05rem] sm:leading-8 [&_a]:text-blue-600 [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-blue-700 dark:[&_a]:text-blue-400 dark:hover:[&_a]:text-blue-300 [&_p]:mb-3 [&_p:last-child]:mb-0 [&_ul]:my-3 [&_ol]:my-3 [&_li]:my-1 [&_strong]:font-semibold [&_code]:rounded [&_code]:bg-muted/80 [&_code]:px-1 [&_code]:py-0.5 [&_code]:text-[0.9em]"
            v-html="heroBio"
          />
          <div v-if="socialLinks.length" class="flex flex-wrap gap-3">
            <a
              v-for="link in socialLinks"
              :key="link.label"
              :href="link.href"
              class="inline-flex items-center gap-2 rounded-full border border-border/80 bg-muted/50 px-4 py-2 text-sm font-medium text-foreground/90 shadow-sm transition-colors hover:bg-muted hover:text-foreground dark:border-border dark:bg-muted/30 dark:hover:bg-muted/50"
              :target="link.href.startsWith('http') ? '_blank' : undefined"
              :rel="link.href.startsWith('http') ? 'noopener noreferrer' : undefined"
            >
              <component
                :is="link.icon"
                v-if="link.icon"
                class="size-4 shrink-0 opacity-80"
              />
              <span
                v-else
                class="flex size-4 shrink-0 items-center justify-center rounded bg-sky-500 text-[10px] font-bold text-white"
              >Q</span>
              {{ link.label }}
            </a>
          </div>
        </section>

        <div
          v-if="showCallout"
          class="rounded-2xl border border-amber-200/80 bg-amber-50/90 p-5 shadow-sm dark:border-amber-900/40 dark:bg-amber-950/35 sm:p-6 sm:rounded-3xl"
        >
          <div class="flex gap-3">
            <span class="text-xl leading-none" aria-hidden="true">📝</span>
            <div
              class="site-home-callout-md min-w-0 flex-1 [--md-bk-color:transparent] [&_.md-editor-preview]:text-sm [&_.md-editor-preview]:leading-7 [&_.md-editor-preview]:text-amber-950/90 dark:[&_.md-editor-preview]:text-amber-100/85 sm:[&_.md-editor-preview]:text-[0.9375rem] [&_.md-editor-preview_a]:text-amber-900 dark:[&_.md-editor-preview_a]:text-amber-200"
            >
              <MdPreview
                id="nectarpin-index-callout"
                class="!bg-transparent [--md-bk-color:transparent]"
                :model-value="calloutText"
                :theme="mdPreviewTheme"
                language="zh-CN"
                preview-theme="github"
                code-theme="atom"
                :show-code-row-number="false"
              />
            </div>
          </div>
        </div>

        <section v-if="githubUsername.trim()" class="space-y-4">
          <div class="flex flex-wrap items-baseline justify-between gap-2">
            <h2 class="text-lg font-semibold tracking-tight">
              GitHub 贡献
            </h2>
            <a
              :href="githubProfileHref"
              class="text-sm text-muted-foreground underline-offset-4 transition-colors hover:text-foreground hover:underline"
              target="_blank"
              rel="noopener noreferrer"
            >
              @{{ githubUsername.trim() }} · GitHub
            </a>
          </div>
          <div class="overflow-x-auto rounded-xl border border-border/60 bg-card/30 p-4 dark:bg-card/20">
            <a
              :href="githubProfileHref"
              class="block w-full min-w-0 rounded-md outline-none ring-offset-background focus-visible:ring-2 focus-visible:ring-ring"
              target="_blank"
              rel="noopener noreferrer"
            >
              <img
                :src="ghContributionChartSrc"
                :alt="`${githubUsername.trim()} 的 GitHub 贡献图`"
                class="block h-auto w-full min-w-0 max-w-full rounded-md bg-background"
                loading="lazy"
                decoding="async"
              >
            </a>
          </div>
        </section>

        <section v-if="techStack.length" class="space-y-4">
          <h2 class="text-lg font-semibold tracking-tight">
            技术栈
          </h2>
          <div class="flex flex-wrap gap-2.5">
            <span
              v-for="t in techStack"
              :key="t.name"
              :class="cn('inline-flex items-center gap-1.5 rounded-full px-3 py-1.5 text-xs font-medium text-white shadow-sm', t.className)"
            >
              {{ t.name }}
            </span>
          </div>
        </section>
      </main>
    </div>
  </div>
</template>

<style scoped>
/*
 * MdPreview（previewOnly）…主题 CSS 会给预览块白色背景，需强制透明以透出琥珀提示框底色。
 */
.site-home-callout-md :deep(.md-editor),
.site-home-callout-md :deep(.md-editor-preview),
.site-home-callout-md :deep(.md-editor-preview-wrapper) {
  background: transparent !important;
  background-color: transparent !important;
}

.site-home-callout-md :deep(.md-editor-preview) {
  min-height: 0 !important;
}
</style>
