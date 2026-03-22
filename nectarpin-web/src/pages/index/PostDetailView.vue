<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'
import { useDark } from '@vueuse/core'
import { CalendarDays, Eye } from 'lucide-vue-next'
import { MdPreview } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'

import { getPublicArticleBySlug, type ArticleItem } from '@/api/article'
import { cn } from '@/lib/utils'
import { useSiteStore } from '@/stores/site'
import { RequestError } from '@/utils/req'

const route = useRoute()
const siteStore = useSiteStore()
const { siteName } = storeToRefs(siteStore)

const isDark = useDark()
const mdPreviewTheme = computed(() => (isDark.value ? 'dark' : 'light'))

const article = ref<ArticleItem | null>(null)
const loading = ref(true)
const err = ref('')

const publishedLabel = computed(() => {
  const raw = article.value?.published_at ?? article.value?.created_at
  if (!raw) return null
  const d = new Date(raw)
  if (Number.isNaN(d.getTime())) return null
  return d.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  })
})

const coverSrc = computed(() => article.value?.cover_image?.trim() ?? '')

async function load() {
  const slug = route.params.slug
  const s = typeof slug === 'string' ? slug : (slug?.[0] ?? '')
  if (!s) {
    loading.value = false
    err.value = '无效链接'
    article.value = null
    return
  }
  loading.value = true
  err.value = ''
  try {
    const res = await getPublicArticleBySlug(s)
    article.value = res.data
  } catch (e) {
    article.value = null
    err.value = e instanceof RequestError ? e.message : '加载失败'
  } finally {
    loading.value = false
  }
}

watch(() => route.params.slug, load, { immediate: true })

watch(
  [article, siteName],
  () => {
    const a = article.value
    const s = siteName.value
    if (a?.title) document.title = `${a.title} · ${s}`
    else document.title = s
  },
  { immediate: true },
)
</script>

<template>
  <div
    class="relative mx-auto w-full max-w-3xl px-4 pb-20 pt-8 sm:px-6 sm:pb-24 sm:pt-10 lg:px-8 lg:pt-12"
  >
    <div
      v-if="loading"
      class="flex min-h-[40vh] flex-col items-center justify-center gap-4 text-center"
      aria-busy="true"
    >
      <span
        class="size-9 animate-pulse rounded-full bg-gradient-to-br from-blue-200/80 to-violet-200/80 dark:from-blue-900/50 dark:to-violet-900/50"
        aria-hidden="true"
      />
      <p class="text-sm text-muted-foreground">加载中…</p>
    </div>

    <div
      v-else-if="err"
      class="rounded-2xl border border-destructive/25 bg-destructive/5 px-5 py-8 text-center sm:px-8"
    >
      <p class="text-sm font-medium text-destructive">{{ err }}</p>
    </div>

    <article v-else-if="article" class="post-detail-enter">
      <header class="space-y-5 sm:space-y-6">
        <h1
          class="text-[1.65rem] font-semibold leading-[1.2] tracking-tight text-foreground sm:text-4xl sm:leading-tight"
        >
          {{ article.title }}
        </h1>

        <p
          v-if="article.summary?.trim()"
          class="max-w-2xl text-base leading-relaxed text-muted-foreground sm:text-lg sm:leading-relaxed"
        >
          {{ article.summary.trim() }}
        </p>

        <div
          class="flex flex-wrap items-center gap-x-5 gap-y-2 border-b border-border/50 pb-6 text-sm text-muted-foreground sm:pb-8"
        >
          <span v-if="publishedLabel" class="inline-flex items-center gap-1.5">
            <CalendarDays class="size-4 shrink-0 opacity-70" aria-hidden="true" />
            <time :datetime="article.published_at ?? article.created_at">{{ publishedLabel }}</time>
          </span>
          <span class="inline-flex items-center gap-1.5">
            <Eye class="size-4 shrink-0 opacity-70" aria-hidden="true" />
            {{ article.view_count.toLocaleString('zh-CN') }} 次阅读
          </span>
        </div>
      </header>

      <figure v-if="coverSrc" class="mt-2 overflow-hidden rounded-2xl border border-border/50 bg-muted/20 shadow-sm">
        <img
          :src="coverSrc"
          :alt="article.title"
          class="aspect-[21/9] w-full object-cover sm:aspect-[2/1]"
          loading="eager"
          decoding="async"
        />
      </figure>

      <div
        :class="
          cn(
            'post-detail-md max-w-none',
            coverSrc ? 'mt-8 sm:mt-10' : 'mt-2 sm:mt-4',
            '[&_.md-editor-preview]:text-[0.9375rem] [&_.md-editor-preview]:leading-[1.75] [&_.md-editor-preview]:text-foreground',
            '[&_.md-editor-preview_a]:text-blue-600 [&_.md-editor-preview_a]:underline [&_.md-editor-preview_a]:underline-offset-4',
            'dark:[&_.md-editor-preview_a]:text-blue-400',
          )
        "
      >
        <MdPreview
          id="nectarpin-post-detail"
          :model-value="article.content"
          :theme="mdPreviewTheme"
          language="zh-CN"
          preview-theme="github"
          code-theme="atom"
          :show-code-row-number="false"
        />
      </div>
    </article>
  </div>
</template>

<style scoped>
.post-detail-md :deep(.md-editor-preview) {
  min-height: 0 !important;
}

.post-detail-md :deep(.md-editor) {
  border-radius: 0;
  box-shadow: none;
}

.post-detail-md :deep(.md-editor-preview-wrapper) {
  padding: 0;
}

@keyframes post-detail-enter {
  from {
    opacity: 0;
    transform: translateY(6px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.post-detail-enter {
  animation: post-detail-enter 0.35s ease-out both;
}
</style>
