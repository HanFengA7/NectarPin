<script setup lang="ts">
import { computed, ref, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { RouterLink, useRoute } from 'vue-router'
import { Moon, Search, Sun } from 'lucide-vue-next'
import { useDark, useToggle, useWindowScroll } from '@vueuse/core'
import { Input } from '@/components/ui/input'
import { cn } from '@/lib/utils'
import { useSiteStore } from '@/stores/site'

interface IndexNavItem {
  label: string
  name: string
  path: string
}

const INDEX_NAV_ITEMS: IndexNavItem[] = [
  { label: '首页', name: 'index', path: '' },
  { label: '文章', name: 'index-articles', path: 'articles' },
  { label: '项目', name: 'index-projects', path: 'projects' },
  { label: '朋友', name: 'index-friends', path: 'friends' },
  { label: '关于', name: 'index-about', path: 'about' },
]

const route = useRoute()
const { siteName, avatarUrl, avatarInitials, heroName } = storeToRefs(useSiteStore())
const avatarLoadFailed = ref(false)

watch(avatarUrl, () => {
  avatarLoadFailed.value = false
})

const avatarFallbackText = computed(() => {
  const cfg = avatarInitials.value.trim()
  if (cfg) return Array.from(cfg).slice(0, 4).join('')
  const n = heroName.value.trim()
  const chars = Array.from(n)
  if (chars.length >= 2) return chars.slice(0, 2).join('')
  if (chars.length === 1) return chars[0] ?? ''
  return 'NP'
})
const { y: scrollY } = useWindowScroll()
/** 页面在顶部时顶栏透明，滚动后为毛玻璃 */
const isAtTop = computed(() => scrollY.value < 8)
const isDark = useDark()
const toggleDark = useToggle(isDark)

const searchQuery = defineModel<string>('searchQuery', { default: '' })

const navItems = INDEX_NAV_ITEMS

function linkActive(item: IndexNavItem) {
  if (route.name === item.name) return true
  if (item.path && route.path === `/${item.path}`) return true
  return false
}
</script>

<template>
  <header
    :class="
      cn(
        'sticky top-0 z-50 border-b transition-[background-color,backdrop-filter,border-color,box-shadow] duration-300 ease-out',
        isAtTop
          ? 'border-transparent bg-transparent shadow-none backdrop-blur-none backdrop-saturate-100 dark:bg-transparent'
          : 'border-border/40 bg-white/55 shadow-sm backdrop-blur-xl backdrop-saturate-150 dark:border-white/10 dark:bg-zinc-950/45',
      )
    "
  >
    <div class="flex h-14 w-full items-center gap-3 px-4 sm:gap-4 sm:px-6 lg:px-8">
      <div class="flex min-w-0 flex-1 justify-start">
        <RouterLink
          :to="{ name: 'index' }"
          class="flex min-w-0 shrink-0 items-center gap-2 outline-none focus-visible:ring-2 focus-visible:ring-blue-500/40 focus-visible:ring-offset-2 focus-visible:ring-offset-zinc-50 dark:focus-visible:ring-offset-zinc-950"
          aria-label="首页"
        >
          <span
            class="relative flex size-8 shrink-0 overflow-hidden rounded-full border border-border/60 bg-gradient-to-br from-blue-100 to-violet-200 ring-1 ring-border/40 dark:from-blue-950 dark:to-violet-950 dark:ring-border/50"
            aria-hidden="true"
          >
            <img
              v-if="avatarUrl && !avatarLoadFailed"
              :src="avatarUrl"
              alt=""
              class="size-full object-cover"
              @error="avatarLoadFailed = true"
            />
            <span
              v-else
              class="flex size-full items-center justify-center text-[10px] font-bold leading-none tracking-tight text-blue-700/90 dark:text-blue-200/90"
            >
              {{ avatarFallbackText }}
            </span>
          </span>
          <span
            class="hidden max-w-[10rem] truncate text-sm font-semibold tracking-tight text-foreground/90 sm:inline"
            >{{ siteName }}</span
          >
        </RouterLink>
      </div>

      <div class="flex min-w-0 flex-1 items-center justify-end gap-2 sm:gap-3 md:gap-2 lg:gap-3">
        <nav
          class="hidden shrink-0 items-center justify-end gap-0.5 md:flex md:gap-0.5 lg:gap-1"
          aria-label="主导航"
        >
          <RouterLink
            v-for="item in navItems"
            :key="item.name"
            :to="{ name: item.name }"
            :class="
              cn(
                'relative px-2.5 py-2 text-sm font-medium text-foreground/85 transition-colors hover:text-foreground lg:px-3.5',
                linkActive(item) && 'text-blue-600 dark:text-blue-400',
              )
            "
          >
            <span class="relative inline-block">{{ item.label }}</span>
            <span
              v-if="linkActive(item)"
              class="absolute bottom-0 left-2 right-2 h-0.5 rounded-full bg-blue-600 dark:bg-blue-500 lg:left-3 lg:right-3"
              aria-hidden="true"
            />
          </RouterLink>
        </nav>

        <div class="relative hidden w-[min(100%,220px)] sm:block md:w-[min(100%,260px)]">
          <Input
            v-model="searchQuery"
            type="search"
            placeholder=""
            class="h-9 rounded-full border-border/80 bg-background pr-9 shadow-sm dark:border-border"
            aria-label="搜索"
          />
          <Search
            class="pointer-events-none absolute right-3 top-1/2 size-4 -translate-y-1/2 text-muted-foreground"
            aria-hidden="true"
          />
        </div>

        <button
          type="button"
          class="flex h-9 items-center gap-0 rounded-full border border-border/80 bg-background px-1 shadow-sm transition-colors hover:bg-muted/60 dark:border-border"
          :aria-pressed="isDark"
          aria-label="切换浅色 / 深色模式"
          @click="toggleDark()"
        >
          <span
            :class="
              cn(
                'flex size-7 items-center justify-center rounded-full transition-colors',
                isDark ? 'bg-blue-600 text-white shadow-sm' : 'text-muted-foreground',
              )
            "
          >
            <Moon class="size-4" aria-hidden="true" />
          </span>
          <span
            :class="
              cn(
                'flex size-7 items-center justify-center rounded-full transition-colors',
                isDark ? 'text-muted-foreground' : 'bg-blue-600 text-white shadow-sm',
              )
            "
          >
            <Sun class="size-4" aria-hidden="true" />
          </span>
        </button>
      </div>
    </div>

    <nav
      :class="
        cn(
          'flex justify-end gap-1 overflow-x-auto border-t px-4 pb-2 pt-1 transition-[border-color] duration-300 md:hidden',
          isAtTop ? 'border-transparent' : 'border-border/40',
        )
      "
      aria-label="主导航"
    >
      <RouterLink
        v-for="item in navItems"
        :key="`m-${item.name}`"
        :to="{ name: item.name }"
        :class="
          cn(
            'relative shrink-0 whitespace-nowrap rounded-full px-3 py-1.5 text-sm font-medium',
            linkActive(item)
              ? 'bg-blue-600 text-white dark:bg-blue-500'
              : 'bg-muted/60 text-foreground/80',
          )
        "
      >
        {{ item.label }}
      </RouterLink>
    </nav>
  </header>
</template>
