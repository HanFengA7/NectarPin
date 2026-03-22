<script setup lang="ts">
import { getLocalTimeZone, parseDate, today } from '@internationalized/date'
import { CalendarIcon, LoaderCircle } from 'lucide-vue-next'
import type { DateValue } from 'reka-ui'
import { computed, onMounted, reactive, ref, watch, type Ref } from 'vue'
import { RouterLink } from 'vue-router'

import {
  getPublicSiteHome,
  saveProtectedSiteHome,
  type SiteHomePayload,
  type SiteSocialLink,
  type SiteTechStackItem,
} from '@/api/site'
import { Button } from '@/components/ui/button'
import { Calendar } from '@/components/ui/calendar'
import { Input } from '@/components/ui/input'
import { Label } from '@/components/ui/label'
import { Popover, PopoverContent, PopoverTrigger } from '@/components/ui/popover'
import { RequestError } from '@/utils/req'
import { useSiteStore } from '@/stores/site'

const tz = getLocalTimeZone()

/** Reka Calendar 的 DateValue 与 @internationalized/date 在类型上不完全一致，运行时兼容。 */
function toRekaDateValue(v: ReturnType<typeof today> | ReturnType<typeof parseDate>): DateValue {
  return v as unknown as DateValue
}

function footerSinceToDateValue(s: string): DateValue | undefined {
  const t = s.trim()
  if (!/^\d{4}-\d{2}-\d{2}$/.test(t)) return undefined
  try {
    return toRekaDateValue(parseDate(t))
  } catch {
    return undefined
  }
}

function dateValueToYYYYMMDD(v: DateValue): string {
  return `${String(v.year).padStart(4, '0')}-${String(v.month).padStart(2, '0')}-${String(v.day).padStart(2, '0')}`
}

const loading = ref(true)
const saving = ref(false)
const loadError = ref('')
const saveError = ref('')
const saveSuccess = ref('')
const siteStore = useSiteStore()

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
  footer_icp_text: '',
  footer_icp_href: '',
  footer_since: '',
  footer_psb_text: '',
  footer_psb_href: '',
})

const footerSinceOpen = ref(false)
const footerSinceCalendar = ref(toRekaDateValue(today(tz))) as Ref<DateValue>

watch(
  () => form.footer_since,
  (s) => {
    const p = footerSinceToDateValue(s)
    footerSinceCalendar.value = p ?? toRekaDateValue(today(tz))
  },
  { immediate: true },
)

const footerSinceLabel = computed(() => {
  const s = form.footer_since.trim()
  if (!s) return '选择上线日期'
  const parsed = footerSinceToDateValue(s)
  if (!parsed) return s
  try {
    return parsed.toDate(tz).toLocaleDateString('zh-CN', {
      year: 'numeric',
      month: 'long',
      day: 'numeric',
    })
  } catch {
    return s
  }
})

function onFooterSinceCalendarUpdate(v: DateValue | undefined) {
  if (!v) return
  footerSinceCalendar.value = v
  form.footer_since = dateValueToYYYYMMDD(v)
  footerSinceOpen.value = false
}

function clearFooterSince() {
  form.footer_since = ''
  footerSinceCalendar.value = toRekaDateValue(today(tz))
  footerSinceOpen.value = false
}

const defaultCalendarPlaceholder = toRekaDateValue(today(tz)) as DateValue

function emptySocial(): SiteSocialLink {
  return { label: '', href: '' }
}

function emptyTech(): SiteTechStackItem {
  return { name: '', class_name: 'bg-zinc-600' }
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
  form.social_links = p.social_links?.length
    ? p.social_links.map((s) => ({ ...s }))
    : [emptySocial()]
  form.tech_stack = p.tech_stack?.length ? p.tech_stack.map((t) => ({ ...t })) : [emptyTech()]
  form.github_username = p.github_username ?? ''
  form.github_chart_hex = (p.github_chart_hex ?? '').trim().replace(/^#/, '').toLowerCase()
  form.footer_icp_text = p.footer_icp_text ?? ''
  form.footer_icp_href = p.footer_icp_href ?? ''
  form.footer_since = p.footer_since ?? ''
  form.footer_psb_text = p.footer_psb_text ?? ''
  form.footer_psb_href = p.footer_psb_href ?? ''
}

async function load() {
  loadError.value = ''
  loading.value = true
  try {
    const res = await getPublicSiteHome()
    applyPayload(res.data)
    siteStore.applyFromHomePayload(res.data)
  } catch (e) {
    loadError.value = e instanceof RequestError ? e.message : '无法加载站点配置'
  } finally {
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
    social_links: form.social_links.filter((s) => s.label.trim() || s.href.trim()),
    tech_stack: form.tech_stack.filter((t) => t.name.trim()),
    github_username: form.github_username.trim(),
    github_chart_hex: form.github_chart_hex.trim().replace(/^#/, '').toLowerCase(),
    footer_icp_text: form.footer_icp_text.trim(),
    footer_icp_href: form.footer_icp_href.trim(),
    footer_since: form.footer_since.trim(),
    footer_psb_text: form.footer_psb_text.trim(),
    footer_psb_href: form.footer_psb_href.trim(),
  }
  try {
    const res = await saveProtectedSiteHome(payload)
    applyPayload(res.data)
    siteStore.applyFromHomePayload(res.data)
    saveSuccess.value = '已保存'
  } catch (e) {
    saveError.value = e instanceof RequestError ? e.message : '保存失败，请稍后重试'
  } finally {
    saving.value = false
  }
}

onMounted(() => {
  void load()
})
</script>

<template>
  <section class="min-h-svh w-full bg-muted/25">
    <div class="w-full px-4 pt-6 pb-0 sm:px-6 sm:pt-8 lg:px-8 lg:pt-10">
      <header class="mb-6 text-left sm:mb-8">
        <p class="text-xs font-medium uppercase tracking-wider text-muted-foreground">
          全局 · 站点管理
        </p>
        <h1 class="mt-2 text-2xl font-semibold tracking-tight sm:text-3xl">页脚设置</h1>
        <p class="mt-2 text-sm leading-relaxed text-muted-foreground">
          配置前台页脚：运行天数、ICP 与公安备案。
        </p>
        <p class="mt-2 text-sm text-muted-foreground">
          社交链接仅在
          <RouterLink
            :to="{ name: 'global-siteSettings-index' }"
            class="font-medium text-blue-600 underline-offset-4 hover:underline dark:text-blue-400"
          >
            首页设置
          </RouterLink>
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

      <form v-else class="flex flex-col gap-5 sm:gap-6" @submit.prevent="handleSave">
        <div class="rounded-xl border bg-background p-5 shadow-sm sm:p-6">
          <div class="mb-4 border-b border-border/60 pb-3">
            <h2 class="text-base font-semibold tracking-tight">页脚（前台底部）</h2>
            <p class="mt-1 text-xs text-muted-foreground">
              左栏：© 站点名、固定「Powered by NectarPin」、运行天数。右栏：ICP 与公安备案。
            </p>
          </div>
          <div class="grid gap-6 lg:grid-cols-2">
            <div class="space-y-4 rounded-lg border border-border/60 bg-muted/10 p-4">
              <p class="text-xs font-medium uppercase tracking-wide text-muted-foreground">左栏</p>
              <div class="space-y-1.5">
                <Label for="footer_since_trigger">站点上线日期（可空）</Label>
                <Popover v-model:open="footerSinceOpen">
                  <PopoverTrigger as-child>
                    <Button
                      id="footer_since_trigger"
                      type="button"
                      variant="outline"
                      class="h-9 w-full justify-between font-normal sm:max-w-xs"
                    >
                      <span class="truncate">{{ footerSinceLabel }}</span>
                      <CalendarIcon class="size-4 shrink-0 opacity-60" />
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent class="w-auto overflow-hidden p-0" align="start">
                    <Calendar
                      :model-value="footerSinceCalendar"
                      layout="month-and-year"
                      class="rounded-md border-0 shadow-none"
                      :default-placeholder="defaultCalendarPlaceholder"
                      @update:model-value="onFooterSinceCalendarUpdate"
                    />
                    <div class="border-t border-border/60 p-2">
                      <Button
                        type="button"
                        variant="ghost"
                        size="sm"
                        class="w-full text-muted-foreground"
                        @click="clearFooterSince"
                      >
                        清除日期
                      </Button>
                    </div>
                  </PopoverContent>
                </Popover>
                <p class="text-xs text-muted-foreground">
                  用于「本站已运行 n 天」；保存为 YYYY-MM-DD。
                </p>
              </div>
            </div>
            <div class="space-y-4 rounded-lg border border-border/60 bg-muted/10 p-4">
              <p class="text-xs font-medium uppercase tracking-wide text-muted-foreground">
                右栏 · 备案
              </p>
              <div class="space-y-1.5">
                <Label for="footer_icp_text">ICP 备案号（可空）</Label>
                <Input
                  id="footer_icp_text"
                  v-model="form.footer_icp_text"
                  placeholder="例如 浙ICP备xxxxxxxx号-1"
                  autocomplete="off"
                />
              </div>
              <div class="space-y-1.5">
                <Label for="footer_icp_href">ICP 查询链接（可空）</Label>
                <Input
                  id="footer_icp_href"
                  v-model="form.footer_icp_href"
                  placeholder="https://beian.miit.gov.cn/…"
                  autocomplete="off"
                />
              </div>
              <div class="space-y-1.5">
                <Label for="footer_psb_text">公安备案号（可空）</Label>
                <Input
                  id="footer_psb_text"
                  v-model="form.footer_psb_text"
                  placeholder="例如 浙公网安备xxxxxxxxxxxxx号"
                  autocomplete="off"
                />
              </div>
              <div class="space-y-1.5">
                <Label for="footer_psb_href">公安备案查询链接（可空）</Label>
                <Input
                  id="footer_psb_href"
                  v-model="form.footer_psb_href"
                  placeholder="https://www.beian.gov.cn/…"
                  autocomplete="off"
                />
              </div>
            </div>
          </div>
        </div>

        <div
          class="sticky bottom-0 z-10 -mx-4 border-t border-border/80 bg-background/95 px-4 py-4 backdrop-blur-sm supports-[backdrop-filter]:bg-background/80 sm:-mx-6 sm:border sm:px-5 sm:shadow-sm lg:-mx-8"
        >
          <div class="w-full space-y-3 text-left">
            <p v-if="saveError" class="text-sm text-destructive">
              {{ saveError }}
            </p>
            <p v-if="saveSuccess" class="text-sm text-emerald-600 dark:text-emerald-400">
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
