<script setup lang="ts">
import { computed } from 'vue'
import { storeToRefs } from 'pinia'

import { useSiteStore } from '@/stores/site'

/** 开源仓库（页脚 Powered by 固定指向此处） */
const NECTARPIN_REPO_URL = 'https://github.com/HanFengA7/NectarPin'

const {
  siteName,
  footerIcpText,
  footerIcpHref,
  footerSince,
  footerPsbText,
  footerPsbHref,
} = storeToRefs(useSiteStore())

const showFilingColumn = computed(
  () => Boolean(footerIcpText.value.trim() || footerPsbText.value.trim()),
)

const year = computed(() => new Date().getFullYear())

const runningDays = computed(() => {
  const raw = footerSince.value.trim()
  if (!/^\d{4}-\d{2}-\d{2}$/.test(raw)) return null
  const start = Date.parse(`${raw}T00:00:00`)
  if (!Number.isFinite(start)) return null
  const diff = Date.now() - start
  if (diff < 0) return 0
  return Math.floor(diff / 86_400_000)
})
</script>

<template>
  <footer
    class="mt-auto border-t border-border/40 bg-background dark:bg-zinc-950/30"
    role="contentinfo"
  >
    <div class="mx-auto w-full max-w-6xl px-4 py-10 sm:px-6 lg:px-8">
      <div
        class="flex flex-col gap-8 sm:flex-row sm:items-start sm:justify-between sm:gap-12"
      >
        <!-- 左侧：版权 / Powered by / 运行天数 -->
        <div class="min-w-0 flex-1 space-y-1 text-sm leading-relaxed text-muted-foreground">
          <p>
            ©{{ year }}
            <span class="text-foreground/90">{{ siteName }}</span>
          </p>
          <p>
            <a
              :href="NECTARPIN_REPO_URL"
              class="underline-offset-4 transition-colors hover:text-foreground hover:underline"
              target="_blank"
              rel="noopener noreferrer"
            >
              Powered by NectarPin
            </a>
          </p>
          <p v-if="runningDays !== null">本站已运行 {{ runningDays }} 天</p>
        </div>

        <!-- 右侧：ICP 备案 / 公安备案 -->
        <div
          v-if="showFilingColumn"
          class="space-y-1.5 text-sm leading-relaxed text-muted-foreground sm:max-w-md sm:text-right"
        >
          <p v-if="footerIcpText">
            <template v-if="footerIcpHref">
              <a
                :href="footerIcpHref"
                class="underline-offset-4 transition-colors hover:text-foreground hover:underline"
                target="_blank"
                rel="noopener noreferrer"
              >
                {{ footerIcpText }}
              </a>
            </template>
            <template v-else>{{ footerIcpText }}</template>
          </p>
          <p v-if="footerPsbText">
            <template v-if="footerPsbHref">
              <a
                :href="footerPsbHref"
                class="underline-offset-4 transition-colors hover:text-foreground hover:underline"
                target="_blank"
                rel="noopener noreferrer"
              >
                {{ footerPsbText }}
              </a>
            </template>
            <template v-else>{{ footerPsbText }}</template>
          </p>
        </div>
      </div>
    </div>
  </footer>
</template>
