<script setup lang="ts">
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import { getPublicArticleBySlug, type ArticleItem } from '@/api/article'
import { useSiteStore } from '@/stores/site'
import { RequestError } from '@/utils/req'

const route = useRoute()
const siteStore = useSiteStore()
const { siteName } = storeToRefs(siteStore)

const article = ref<ArticleItem | null>(null)
const loading = ref(true)
const err = ref('')

async function load() {
  const slug = route.params.slug
  const s = typeof slug === 'string' ? slug : slug?.[0] ?? ''
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
  }
  catch (e) {
    article.value = null
    err.value = e instanceof RequestError ? e.message : '加载失败'
  }
  finally {
    loading.value = false
  }
}

watch(() => route.params.slug, load, { immediate: true })

watch(
  [article, siteName],
  () => {
    const a = article.value
    const s = siteName.value
    if (a?.title)
      document.title = `${a.title} · ${s}`
    else
      document.title = s
  },
  { immediate: true },
)
</script>

<template>
  <div class="mx-auto w-full max-w-3xl px-4 py-12 sm:px-6 lg:px-10">
    <div v-if="loading" class="text-sm text-muted-foreground">
      加载中…
    </div>
    <div v-else-if="err" class="text-sm text-destructive">
      {{ err }}
    </div>
    <article v-else-if="article" class="space-y-6">
      <header class="space-y-2">
        <h1 class="text-2xl font-semibold tracking-tight sm:text-3xl">
          {{ article.title }}
        </h1>
        <p v-if="article.summary" class="text-muted-foreground">
          {{ article.summary }}
        </p>
      </header>
      <div
        class="max-w-none whitespace-pre-wrap text-sm leading-relaxed text-foreground"
      >
        {{ article.content }}
      </div>
    </article>
  </div>
</template>
