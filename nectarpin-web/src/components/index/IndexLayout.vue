<script setup lang="ts">
import { onMounted, watch } from 'vue'
import { RouterView, useRoute } from 'vue-router'
import { storeToRefs } from 'pinia'

import IndexFooter from '@/components/index/IndexFooter.vue'
import IndexSidebar from '@/components/index/IndexSidebar.vue'
import { useSiteStore } from '@/stores/site'

const route = useRoute()
const siteStore = useSiteStore()
const { siteName } = storeToRefs(siteStore)

const INDEX_TITLE_SUFFIX: Record<string, string> = {
  index: '',
  'index-articles': ' · 文章',
  'index-projects': ' · 项目',
  'index-friends': ' · 朋友',
  'index-about': ' · 关于',
}

function syncPublicDocumentTitle() {
  if (route.name === 'index-post-detail') return
  const base = siteName.value
  const suf = INDEX_TITLE_SUFFIX[route.name as string] ?? ''
  document.title = `${base}${suf}`
}

onMounted(() => {
  void siteStore.hydrateFromApi()
})

watch(
  [siteName, () => route.name],
  () => {
    syncPublicDocumentTitle()
  },
  { immediate: true },
)
</script>

<template>
  <div class="flex min-h-dvh flex-col bg-background text-foreground">
    <IndexSidebar />
    <main class="flex-1">
      <RouterView />
    </main>
    <IndexFooter />
  </div>
</template>
