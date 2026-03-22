<script setup lang="ts">
import { computed, onMounted, ref, shallowRef } from 'vue'
import { ExternalLink, LoaderCircle } from 'lucide-vue-next'

import { getPublicFriendPage, type PublicFriendSection } from '@/api/link'
import { RequestError } from '@/utils/req'

/** 与侧栏「朋友」一致 */
const PAGE_HEADING = '朋友'

const pageLoading = ref(true)
const pageError = ref('')
const introHtml = ref('')
const sections = shallowRef<PublicFriendSection[]>([])

const showIntro = computed(() => introHtml.value.trim().length > 0)

function friendLinkHost(url: string) {
  try {
    return new URL(url).host
  } catch {
    return url
  }
}

async function loadPage() {
  pageLoading.value = true
  pageError.value = ''
  try {
    const res = await getPublicFriendPage()
    if (res.code !== 200 || !res.data) {
      pageError.value = '暂时无法加载友链。'
      introHtml.value = ''
      sections.value = []
      return
    }
    introHtml.value = res.data.intro_html ?? ''
    sections.value = res.data.sections ?? []
  } catch (e) {
    introHtml.value = ''
    sections.value = []
    pageError.value =
      e instanceof RequestError ? e.message : '暂时无法加载友链，请稍后重试。'
  } finally {
    pageLoading.value = false
  }
}

onMounted(() => {
  void loadPage()
})
</script>

<template>
  <div
    class="friends-page w-full px-4 pb-24 pt-12 sm:px-6 sm:pt-14 lg:px-10 lg:pt-16"
  >
    <div class="mx-auto max-w-5xl xl:max-w-6xl">
      <header class="border-b border-border/80 pb-8 dark:border-border/50">
        <h1 class="mt-2 text-3xl font-semibold tracking-tight text-foreground sm:text-4xl">
          {{ PAGE_HEADING }}
        </h1>
        <div
          v-if="showIntro"
          class="friend-intro-html mt-6 max-w-3xl text-base leading-relaxed text-muted-foreground sm:text-[1.05rem] sm:leading-8 [&_a]:text-blue-600 [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-blue-700 dark:[&_a]:text-blue-400 dark:hover:[&_a]:text-blue-300 [&_p]:mb-3 [&_p:last-child]:mb-0 [&_ul]:my-3 [&_ol]:my-3 [&_li]:my-1 [&_strong]:font-semibold"
          v-html="introHtml"
        />
      </header>

      <div class="mt-10">
        <div
          v-if="pageLoading"
          class="flex min-h-[12rem] flex-col items-center justify-center gap-3 rounded-2xl border border-border/60 bg-muted/15 py-16 text-sm text-muted-foreground"
        >
          <LoaderCircle
            class="size-8 animate-spin text-blue-600/70 dark:text-blue-400/70"
            aria-hidden="true"
          />
          <span>加载中…</span>
        </div>
        <div
          v-else-if="pageError"
          class="rounded-2xl border border-destructive/25 bg-destructive/5 px-6 py-12 text-center text-sm text-destructive"
        >
          {{ pageError }}
        </div>
        <template v-else>
          <div
            v-if="sections.length === 0"
            class="rounded-2xl border border-dashed border-border/80 bg-muted/10 py-20 text-center"
          >
            <p class="text-muted-foreground">暂无友链，稍后再来看看。</p>
          </div>
          <div v-else class="space-y-16 sm:space-y-20">
            <section
              v-for="(sec, si) in sections"
              :key="`${sec.category_id ?? 'x'}-${si}`"
              class="scroll-mt-20 space-y-6"
            >
              <div class="border-l-[3px] border-blue-500 pl-4 dark:border-blue-400">
                <h2 class="text-xl font-semibold tracking-tight text-foreground sm:text-2xl">
                  {{ sec.category_name }}
                </h2>
                <p
                  v-if="sec.category_description?.trim()"
                  class="mt-2 max-w-3xl text-sm leading-relaxed text-muted-foreground sm:text-[0.9375rem]"
                >
                  {{ sec.category_description }}
                </p>
              </div>
              <ul
                class="grid grid-cols-1 gap-5 sm:grid-cols-2 sm:gap-6 xl:grid-cols-3"
                role="list"
              >
                <li v-for="(fl, fi) in sec.links" :key="`${si}-${fi}-${fl.url}`">
                  <a
                    :href="fl.url"
                    class="group relative flex h-full min-h-[5.5rem] gap-4 rounded-2xl border border-border/80 bg-card/45 p-5 shadow-sm outline-none ring-offset-background transition-all duration-200 hover:-translate-y-0.5 hover:border-blue-500/40 hover:shadow-md focus-visible:ring-2 focus-visible:ring-ring dark:bg-card/25 dark:hover:border-blue-400/35"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <div
                      class="relative size-14 shrink-0 overflow-hidden rounded-2xl border border-border/70 bg-gradient-to-br from-muted to-muted/50 shadow-inner"
                    >
                      <img
                        v-if="fl.avatar_url?.trim()"
                        :src="fl.avatar_url"
                        :alt="fl.title"
                        class="size-full object-cover"
                        loading="lazy"
                        decoding="async"
                      />
                      <div
                        v-else
                        class="flex size-full items-center justify-center text-lg font-semibold tabular-nums text-muted-foreground"
                      >
                        {{ fl.title.slice(0, 1) }}
                      </div>
                    </div>
                    <div class="flex min-w-0 flex-1 flex-col py-0.5">
                      <div class="flex items-start justify-between gap-3">
                        <p
                          class="font-semibold leading-snug text-foreground transition-colors group-hover:text-blue-600 dark:group-hover:text-blue-400"
                        >
                          {{ fl.title }}
                        </p>
                        <ExternalLink
                          class="size-4 shrink-0 text-muted-foreground/45 transition-colors group-hover:text-blue-600 dark:group-hover:text-blue-400"
                          aria-hidden="true"
                          :stroke-width="2"
                        />
                      </div>
                      <p
                        v-if="fl.description?.trim()"
                        class="mt-2 line-clamp-2 text-sm leading-relaxed text-muted-foreground"
                      >
                        {{ fl.description }}
                      </p>
                    </div>
                  </a>
                </li>
              </ul>
            </section>
          </div>
        </template>
      </div>
    </div>
  </div>
</template>
