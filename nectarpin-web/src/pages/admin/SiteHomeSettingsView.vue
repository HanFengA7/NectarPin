<script setup lang="ts">
import { ChevronDown, ChevronUp, LoaderCircle, Plus, Trash2 } from 'lucide-vue-next'
import { defineAsyncComponent, onMounted, reactive, ref } from 'vue'

import {
  getPublicSiteHome,
  saveProtectedSiteHome,
  type SiteAvatarBadgeIcon,
  type SiteHomePayload,
  type SiteSocialLink,
  type SiteStatusIcon,
  type SiteStatusIconTone,
  type SiteTechStackItem,
} from '@/api/site'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { cn } from '@/lib/utils'
import { RequestError } from '@/utils/req'
import { useSiteStore } from '@/stores/site'

const MarkdownEditor = defineAsyncComponent(() => import('@/components/admin/VditorEditor.vue'))

const loading = ref(true)
const saving = ref(false)
const loadError = ref('')
const saveError = ref('')
const saveSuccess = ref('')
const siteStore = useSiteStore()

const STATUS_ICON_OPTIONS: { value: SiteStatusIcon; label: string }[] = [
  { value: 'dot', label: '实心圆点' },
  { value: 'circle_dot', label: '圆环加点（Circle dot）' },
  { value: 'radio', label: '信号（Radio）' },
  { value: 'activity', label: '活动（Activity）' },
  { value: 'none', label: '不显示图标' },
]

const STATUS_TONE_OPTIONS: { value: SiteStatusIconTone; label: string }[] = [
  { value: 'emerald', label: '翠绿' },
  { value: 'sky', label: '天蓝' },
  { value: 'blue', label: '蓝' },
  { value: 'violet', label: '紫' },
  { value: 'amber', label: '琥珀' },
  { value: 'rose', label: '玫红' },
  { value: 'zinc', label: '灰' },
]

const AVATAR_BADGE_ICON_OPTIONS: { value: SiteAvatarBadgeIcon; label: string }[] = [
  { value: 'camera', label: '相机（默认）' },
  { value: 'sparkles', label: '闪光' },
  { value: 'coffee', label: '咖啡' },
  { value: 'heart', label: '心形' },
  { value: 'pen', label: '钢笔' },
  { value: 'smile', label: '笑脸' },
  { value: 'none', label: '不显示角标' },
]

/** 技术栈标签底色预设（点击即可，无需手写 Tailwind） */
const TECH_STACK_COLOR_PRESETS: { class_name: string; label: string }[] = [
  { class_name: 'bg-emerald-600', label: '翠绿' },
  { class_name: 'bg-green-600', label: '绿色' },
  { class_name: 'bg-green-700', label: '深绿' },
  { class_name: 'bg-teal-600', label: '青色' },
  { class_name: 'bg-cyan-600', label: '亮青' },
  { class_name: 'bg-cyan-700', label: '深青' },
  { class_name: 'bg-sky-600', label: '天蓝' },
  { class_name: 'bg-sky-500', label: '浅天蓝' },
  { class_name: 'bg-blue-600', label: '蓝色' },
  { class_name: 'bg-indigo-600', label: '靛蓝' },
  { class_name: 'bg-violet-600', label: '紫罗兰' },
  { class_name: 'bg-fuchsia-600', label: '洋红' },
  { class_name: 'bg-rose-600', label: '玫红' },
  { class_name: 'bg-amber-600', label: '琥珀' },
  { class_name: 'bg-orange-600', label: '橙色' },
  { class_name: 'bg-red-600', label: '红色' },
  { class_name: 'bg-zinc-600', label: '锌灰' },
  { class_name: 'bg-slate-600', label: '石板灰' },
]

function isCustomTechClass(className: string) {
  return !TECH_STACK_COLOR_PRESETS.some(p => p.class_name === className)
}

function techCustomClassSummary(className: string) {
  const s = className.trim()
  if (!s)
    return '自定义类名'
  return s.length > 26 ? `${s.slice(0, 26)}…` : s
}

const form = reactive<SiteHomePayload>({
  site_name: 'NectarPin',
  avatar_url: '',
  avatar_initials: '',
  avatar_badge_icon: 'camera',
  avatar_badge_tone: 'blue',
  status_text: '',
  status_icon: 'dot',
  status_icon_tone: 'emerald',
  hero_name: '',
  hero_bio: '',
  callout_text: '',
  social_links: [],
  tech_stack: [],
  github_username: '',
  github_chart_hex: '',
})

function emptySocial(): SiteSocialLink {
  return { label: '', href: '' }
}

function emptyTech(): SiteTechStackItem {
  return { name: '', class_name: 'bg-zinc-600' }
}

function addSocial() {
  form.social_links.push(emptySocial())
}

function removeSocial(i: number) {
  form.social_links.splice(i, 1)
}

function addTech() {
  form.tech_stack.push(emptyTech())
}

function removeTech(i: number) {
  form.tech_stack.splice(i, 1)
}

function moveTech(i: number, dir: -1 | 1) {
  const j = i + dir
  if (j < 0 || j >= form.tech_stack.length)
    return
  const arr = form.tech_stack
  const [row] = arr.splice(i, 1)
  arr.splice(j, 0, row!)
}

function applyPayload(p: SiteHomePayload) {
  form.site_name = p.site_name?.trim() ? p.site_name : 'NectarPin'
  form.avatar_url = p.avatar_url ?? ''
  form.avatar_initials = p.avatar_initials ?? ''
  form.avatar_badge_icon = p.avatar_badge_icon ?? 'camera'
  form.avatar_badge_tone = p.avatar_badge_icon === 'none' ? '' : (p.avatar_badge_tone ?? 'blue')
  form.status_text = p.status_text ?? ''
  form.status_icon = p.status_icon ?? 'dot'
  form.status_icon_tone = p.status_icon === 'none' ? '' : (p.status_icon_tone ?? 'emerald')
  form.hero_name = p.hero_name
  form.hero_bio = p.hero_bio
  form.callout_text = p.callout_text
  form.social_links = p.social_links?.length ? p.social_links.map(s => ({ ...s })) : [emptySocial()]
  form.tech_stack = p.tech_stack?.length ? p.tech_stack.map(t => ({ ...t })) : [emptyTech()]
  form.github_username = p.github_username ?? ''
}

async function load() {
  loadError.value = ''
  loading.value = true
  try {
    const res = await getPublicSiteHome()
    applyPayload(res.data)
    siteStore.applyFromHomePayload(res.data)
  }
  catch (e) {
    loadError.value =
      e instanceof RequestError ? e.message : '无法加载首页配置'
  }
  finally {
    loading.value = false
  }
}

async function handleSave() {
  saveError.value = ''
  saveSuccess.value = ''
  saving.value = true
  const payload: SiteHomePayload = {
    site_name: form.site_name.trim() || 'NectarPin',
    avatar_url: form.avatar_url.trim(),
    avatar_initials: form.avatar_initials.trim(),
    avatar_badge_icon: form.avatar_badge_icon,
    avatar_badge_tone: form.avatar_badge_icon === 'none' ? '' : form.avatar_badge_tone,
    status_text: form.status_text.trim(),
    status_icon: form.status_icon,
    status_icon_tone: form.status_icon === 'none' ? '' : form.status_icon_tone,
    hero_name: form.hero_name,
    hero_bio: form.hero_bio,
    callout_text: form.callout_text,
    social_links: form.social_links.filter(s => s.label.trim() || s.href.trim()),
    tech_stack: form.tech_stack.filter(t => t.name.trim()),
    github_username: form.github_username.trim(),
    github_chart_hex: form.github_chart_hex.trim(),
  }
  try {
    const res = await saveProtectedSiteHome(payload)
    applyPayload(res.data)
    siteStore.applyFromHomePayload(res.data)
    saveSuccess.value = '已保存'
  }
  catch (e) {
    saveError.value =
      e instanceof RequestError ? e.message : '保存失败，请稍后重试'
  }
  finally {
    saving.value = false
  }
}

onMounted(() => {
  void load()
})

/** 表单内原生 select 统一样式 */
const selectBaseClass = cn(
  'border-input h-9 w-full rounded-md border bg-transparent px-3 text-sm shadow-xs outline-none',
  'focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]',
)

const heroBioHtmlTextareaClass = cn(
  'placeholder:text-muted-foreground border-input w-full min-w-0 rounded-md border bg-transparent px-3 py-2 font-mono text-sm shadow-xs transition-[color,box-shadow] outline-none',
  'focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px]',
)

</script>

<template>
  <section class="min-h-svh w-full bg-muted/25">
    <!-- 仅用 pt，避免 py 在吸底条下方再垫一层 padding 导致滚到底仍有空白 -->
    <div class="w-full px-4 pt-6 pb-0 sm:px-6 sm:pt-8 lg:px-8 lg:pt-10">
      <header class="mb-6 text-left sm:mb-8">
        <p class="text-xs font-medium uppercase tracking-wider text-muted-foreground">
          全局 · 站点管理
        </p>
        <h1 class="mt-2 text-2xl font-semibold tracking-tight sm:text-3xl">
          首页设置
        </h1>
        <p class="mt-2 text-sm leading-relaxed text-muted-foreground">
          分块配置前台首页与品牌信息；改完后请保存。
        </p>
      </header>

      <div
        v-if="loading"
        class="flex items-center gap-3 rounded-xl border bg-background px-5 py-6 text-sm text-muted-foreground shadow-sm"
      >
        <LoaderCircle class="size-5 shrink-0 animate-spin" />
        <span>正在加载配置…</span>
      </div>

      <p
        v-else-if="loadError"
        class="rounded-xl border border-destructive/35 bg-destructive/5 px-4 py-3 text-sm text-destructive"
      >
        {{ loadError }}
      </p>

      <form
        v-else
        class="flex flex-col gap-5 sm:gap-6"
        @submit.prevent="handleSave"
      >
        <div class="grid gap-5 sm:gap-6 lg:grid-cols-2 lg:items-start">
          <!-- 左列：品牌与侧栏相关 -->
          <div class="flex min-w-0 flex-col gap-5 sm:gap-6">
        <!-- 站点与品牌 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">
              站点与品牌
            </h2>
            <p class="mt-1 text-xs text-muted-foreground">
              网站名称会用于页签、顶栏与后台侧栏等。
            </p>
          </div>
          <div class="space-y-1.5">
            <Label for="site_name">网站名称</Label>
            <Input
              id="site_name"
              v-model="form.site_name"
              maxlength="80"
              placeholder="显示在浏览器标题、后台侧栏与登录页"
              autocomplete="organization"
            />
            <p class="text-xs text-muted-foreground">
              页签示例：「名称」「名称 · 文章」；文章页：「标题 · 名称」。
            </p>
          </div>
        </div>

        <!-- GitHub 贡献图（外链图片） -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">
              GitHub 贡献图
            </h2>
            <p class="mt-1 text-xs text-muted-foreground">
              首页填写登录名后显示GitHub贡献图,留空则不显示。
            </p>
          </div>
          <div class="space-y-1.5">
            <Label for="github_username">GitHub 用户名</Label>
            <Input
              id="github_username"
              v-model="form.github_username"
              maxlength="39"
              placeholder="例如 octocat，勿含 @"
              autocomplete="username"
            />
            <p class="text-xs text-muted-foreground">
              大小写均可，按你填写原样用于图表与主页链接。
            </p>
          </div>
          <div class="space-y-1.5">
            <Label for="github_chart_hex">贡献图主题色（可选）</Label>
            <Input
              id="github_chart_hex"
              v-model="form.github_chart_hex"
              maxlength="7"
              placeholder="例如 409ba5 或 #409ba5"
              autocomplete="off"
            />
            <p class="text-xs text-muted-foreground">
              留空为默认绿色系；保存时自动去掉 # 并转为小写，须恰好 6 位十六进制字符。
            </p>
          </div>
        </div>

        <!-- 侧栏头像 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">
              侧栏头像与角标
            </h2>
            <p class="mt-1 text-xs text-muted-foreground">
              头像图片、无图缩写，以及叠在头像右下角的小圆标。
            </p>
          </div>
          <div class="space-y-4">
            <div class="space-y-1.5">
              <Label for="avatar_url">头像图片地址（可选）</Label>
              <Input
                id="avatar_url"
                v-model="form.avatar_url"
                type="url"
                placeholder="https://… 留空则显示缩写"
                autocomplete="off"
              />
              <p class="text-xs text-muted-foreground">
                须为 http 或 https。
              </p>
            </div>
            <div class="space-y-1.5">
              <Label for="avatar_initials">头像缩写（无图时）</Label>
              <Input
                id="avatar_initials"
                v-model="form.avatar_initials"
                maxlength="4"
                placeholder="如 NP；留空则用展示名称前 1～2 字"
                autocomplete="off"
              />
            </div>
            <div class="grid gap-4 sm:grid-cols-2">
              <div class="space-y-1.5">
                <Label for="avatar_badge_icon">右下角角标</Label>
                <select
                  id="avatar_badge_icon"
                  v-model="form.avatar_badge_icon"
                  :class="selectBaseClass"
                >
                  <option
                    v-for="opt in AVATAR_BADGE_ICON_OPTIONS"
                    :key="opt.value"
                    :value="opt.value"
                  >
                    {{ opt.label }}
                  </option>
                </select>
                <p class="text-xs text-muted-foreground">
                  装饰用，非按钮。
                </p>
              </div>
              <div class="space-y-1.5">
                <Label for="avatar_badge_tone">角标底色</Label>
                <select
                  id="avatar_badge_tone"
                  v-model="form.avatar_badge_tone"
                  :disabled="form.avatar_badge_icon === 'none'"
                  :class="cn(selectBaseClass, 'disabled:cursor-not-allowed disabled:opacity-50')"
                >
                  <option
                    v-for="opt in STATUS_TONE_OPTIONS"
                    :key="opt.value"
                    :value="opt.value"
                  >
                    {{ opt.label }}
                  </option>
                </select>
              </div>
            </div>
          </div>
        </div>

        <!-- 状态行 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">
              头像下方状态行
            </h2>
            <p class="mt-1 text-xs text-muted-foreground">
              文案与左侧小图标（圆点或 Lucide 示意）。
            </p>
          </div>
          <div class="space-y-4">
            <div class="space-y-1.5">
              <Label for="status_text">状态文案</Label>
              <Input
                id="status_text"
                v-model="form.status_text"
                maxlength="120"
                placeholder="例如：正在折腾本站"
                autocomplete="off"
              />
            </div>
            <div class="grid gap-4 sm:grid-cols-2">
              <div class="space-y-1.5">
                <Label for="status_icon">状态旁图标</Label>
                <select
                  id="status_icon"
                  v-model="form.status_icon"
                  :class="selectBaseClass"
                >
                  <option
                    v-for="opt in STATUS_ICON_OPTIONS"
                    :key="opt.value"
                    :value="opt.value"
                  >
                    {{ opt.label }}
                  </option>
                </select>
              </div>
              <div class="space-y-1.5">
                <Label for="status_icon_tone">图标颜色</Label>
                <select
                  id="status_icon_tone"
                  v-model="form.status_icon_tone"
                  :disabled="form.status_icon === 'none'"
                  :class="cn(selectBaseClass, 'disabled:cursor-not-allowed disabled:opacity-50')"
                >
                  <option
                    v-for="opt in STATUS_TONE_OPTIONS"
                    :key="opt.value"
                    :value="opt.value"
                  >
                    {{ opt.label }}
                  </option>
                </select>
                <p class="text-xs text-muted-foreground">
                  「不显示图标」时颜色无效。
                </p>
              </div>
            </div>
          </div>
        </div>
          </div>

          <!-- 右列：主区文案与列表 -->
          <div class="flex min-w-0 flex-col gap-5 sm:gap-6">
        <!-- 首页主文案 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">
              首页主文案
            </h2>
            <p class="mt-1 text-xs text-muted-foreground">
              展示名为纯文本；简介使用 HTML（标题下方区域）；提示框仍为 Markdown。
            </p>
          </div>
          <div class="space-y-4">
            <div class="space-y-1.5">
              <Label for="hero_name">展示名称</Label>
              <Input
                id="hero_name"
                v-model="form.hero_name"
                placeholder="例如 NectarPin"
                autocomplete="off"
              />
            </div>
            <div class="space-y-1.5">
              <Label for="hero_bio">首页简介（HTML）</Label>
              <textarea
                id="hero_bio"
                v-model="form.hero_bio"
                rows="8"
                :class="heroBioHtmlTextareaClass"
                placeholder="例如 &lt;p&gt;段落&lt;/p&gt; 或 &lt;br&gt; 换行；仅管理员填写，请注意 XSS 风险"
                spellcheck="false"
              />
              <p class="text-xs text-muted-foreground">
                前台使用 v-html 渲染；勿粘贴不可信来源的 HTML。
              </p>
            </div>
            <div class="space-y-1.5">
              <Label for="callout_text">提示框文案</Label>
              <MarkdownEditor
                id="callout_text"
                v-model="form.callout_text"
                placeholder="琥珀色提示区域内的文字（Markdown）"
                :min-height="220"
                :default-preview-open="false"
              />
            </div>
          </div>
        </div>

        <!-- 社交链接 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 flex flex-wrap items-end justify-between gap-3 border-b border-border/60 pb-3">
            <div>
              <h2 class="text-base font-semibold tracking-tight">
                社交链接
              </h2>
              <p class="mt-1 text-xs text-muted-foreground">
                首页主区按钮式外链。
              </p>
            </div>
            <Button type="button" variant="outline" size="sm" class="h-8 shrink-0" @click="addSocial">
              <Plus class="size-3.5" />
              添加
            </Button>
          </div>
          <ul class="space-y-2">
            <li
              v-for="(item, i) in form.social_links"
              :key="i"
              class="flex flex-col gap-3 rounded-lg border border-border/70 bg-muted/20 p-3 sm:flex-row sm:items-end"
            >
              <div class="grid min-w-0 flex-1 gap-3 sm:grid-cols-2">
                <div class="space-y-1.5">
                  <Label :for="`soc_label_${i}`">名称</Label>
                  <Input
                    :id="`soc_label_${i}`"
                    v-model="item.label"
                    placeholder="GitHub"
                  />
                </div>
                <div class="space-y-1.5 sm:col-span-1">
                  <Label :for="`soc_href_${i}`">链接</Label>
                  <Input
                    :id="`soc_href_${i}`"
                    v-model="item.href"
                    placeholder="https:// 或 mailto:"
                  />
                </div>
              </div>
              <Button
                type="button"
                variant="outline"
                size="icon"
                class="size-9 shrink-0 self-end text-muted-foreground hover:border-destructive/40 hover:text-destructive sm:self-auto"
                :disabled="form.social_links.length <= 1"
                title="删除"
                @click="removeSocial(i)"
              >
                <Trash2 class="size-4" />
              </Button>
            </li>
          </ul>
        </div>

        <!-- 技术栈 -->
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 flex flex-wrap items-end justify-between gap-3 border-b border-border/60 pb-3">
            <div>
              <h2 class="text-base font-semibold tracking-tight">
                技术栈标签
              </h2>
              <p class="mt-1 text-xs text-muted-foreground">
                下拉选底色；右侧调整顺序或删除。
              </p>
            </div>
            <Button type="button" variant="outline" size="sm" class="h-8 shrink-0" @click="addTech">
              <Plus class="size-3.5" />
              添加
            </Button>
          </div>
          <ul class="space-y-2">
            <li
              v-for="(item, i) in form.tech_stack"
              :key="i"
              class="rounded-lg border border-border/70 bg-muted/15 px-3 py-2.5"
            >
              <div class="flex w-full flex-wrap items-center justify-between gap-3">
                <div class="flex min-w-0 flex-wrap items-end gap-4">
                  <div class="space-y-1.5">
                    <Label :for="`tech_name_${i}`" class="text-xs">名称</Label>
                    <Input
                      :id="`tech_name_${i}`"
                      v-model="item.name"
                      placeholder="例如 Vue"
                      class="h-9 w-[7.5rem] sm:w-36"
                    />
                  </div>
                  <div class="space-y-1.5">
                    <span class="text-xs font-medium text-foreground">底色</span>
                    <div class="flex items-center gap-2">
                      <span
                        class="size-7 shrink-0 rounded-full border border-border/60 shadow-inner"
                        :class="item.class_name || 'bg-zinc-600'"
                        aria-hidden="true"
                      />
                      <select
                        :id="`tech_color_${i}`"
                        v-model="item.class_name"
                        :class="cn(
                          selectBaseClass,
                          'w-[min(100%,9.5rem)] px-2 sm:w-40',
                        )"
                        aria-label="标签底色"
                      >
                        <option
                          v-if="isCustomTechClass(item.class_name)"
                          :value="item.class_name"
                        >
                          自定义：{{ techCustomClassSummary(item.class_name) }}
                        </option>
                        <option
                          v-for="preset in TECH_STACK_COLOR_PRESETS"
                          :key="preset.class_name"
                          :value="preset.class_name"
                        >
                          {{ preset.label }}
                        </option>
                      </select>
                    </div>
                  </div>
                </div>
                <div class="flex shrink-0 items-center gap-1">
                  <Button
                    type="button"
                    variant="outline"
                    size="icon"
                    class="size-8"
                    :disabled="i === 0"
                    title="上移"
                    @click="moveTech(i, -1)"
                  >
                    <ChevronUp class="size-3.5" />
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    size="icon"
                    class="size-8"
                    :disabled="i === form.tech_stack.length - 1"
                    title="下移"
                    @click="moveTech(i, 1)"
                  >
                    <ChevronDown class="size-3.5" />
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    size="icon"
                    class="size-8 text-muted-foreground hover:border-destructive/50 hover:text-destructive"
                    :disabled="form.tech_stack.length <= 1"
                    title="删除"
                    @click="removeTech(i)"
                  >
                    <Trash2 class="size-3.5" />
                  </Button>
                </div>
              </div>
              <div
                v-if="isCustomTechClass(item.class_name)"
                class="mt-3 flex flex-col gap-2 border-t border-border/50 pt-3 sm:flex-row sm:flex-wrap sm:items-center"
              >
                <Label :for="`tech_cls_${i}`" class="shrink-0 text-xs text-muted-foreground sm:w-28">
                  自定义 Tailwind 类
                </Label>
                <Input
                  :id="`tech_cls_${i}`"
                  v-model="item.class_name"
                  placeholder="bg-emerald-600"
                  class="h-8 min-w-0 flex-1 font-mono text-xs sm:min-w-[14rem]"
                />
                <p class="w-full text-[11px] text-muted-foreground">
                  在下拉中选预设可覆盖。
                </p>
              </div>
            </li>
          </ul>
        </div>
          </div>
        </div>

        <!-- 底部操作 -->
        <div class="sticky bottom-0 z-10 -mx-4 border-t border-border/80 bg-background/95 px-4 py-4 backdrop-blur-sm supports-[backdrop-filter]:bg-background/80 sm:-mx-6 sm:border sm:px-5 sm:shadow-sm lg:-mx-8">
          <div class="w-full space-y-3 text-left">
            <p
              v-if="saveError"
              class="text-sm text-destructive"
            >
              {{ saveError }}
            </p>
            <p
              v-if="saveSuccess"
              class="text-sm text-emerald-600 dark:text-emerald-400"
            >
              {{ saveSuccess }}
            </p>
            <div class="flex flex-wrap items-center gap-3">
              <Button type="submit" :disabled="saving">
                <LoaderCircle v-if="saving" class="mr-2 size-4 animate-spin" />
                保存配置
              </Button>
              <Button type="button" variant="outline" :disabled="loading || saving" @click="load">
                重新加载
              </Button>
            </div>
          </div>
        </div>
      </form>
    </div>
  </section>
</template>
