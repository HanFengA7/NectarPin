<script setup lang="ts">
import { computed, shallowRef, ref, watch } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { ChevronLeft, ChevronRight, FilePlus2, LoaderCircle, PencilLine, Trash2 } from 'lucide-vue-next'

import { deleteArticle, listAdminArticles, type ArticleItem } from '@/api/article'
import { Button } from '@/components/ui/button'
import { RequestError } from '@/utils/req'

type StatusFilterValue = 'all' | '0' | '1' | '2'

const PAGE_SIZE = 10

const STATUS_OPTIONS = [
  { value: 'all', label: '全部状态' },
  { value: '0', label: '草稿' },
  { value: '1', label: '已发布' },
  { value: '2', label: '已归档' },
] as const

const STATUS_LABEL_MAP: Record<number, string> = {
  0: '草稿',
  1: '已发布',
  2: '已归档',
}

const STATUS_CLASS_MAP: Record<number, string> = {
  0: 'border-yellow-200 bg-yellow-50 text-yellow-700',
  1: 'border-emerald-200 bg-emerald-50 text-emerald-700',
  2: 'border-slate-200 bg-slate-100 text-slate-700',
}

const DATE_TIME_FORMATTER = new Intl.DateTimeFormat('zh-CN', {
  hour12: false,
})

const NUMBER_FORMATTER = new Intl.NumberFormat('zh-CN')

const route = useRoute()
const router = useRouter()

const articles = shallowRef<ArticleItem[]>([])
const isLoading = ref(false)
const deletingId = ref<number | null>(null)
const errorMessage = ref('')
const successMessage = ref('')
const statusFilter = ref<StatusFilterValue>('all')
const currentPage = ref(1)
const total = ref(0)
let listRequestId = 0

const totalPages = computed(() => {
  const pages = Math.ceil(total.value / PAGE_SIZE)
  return Math.max(1, pages)
})

function formatDateTime(value: string | null) {
  if (!value) {
    return '-'
  }

  const date = new Date(value)
  if (Number.isNaN(date.getTime())) {
    return '-'
  }

  return DATE_TIME_FORMATTER.format(date)
}

function formatNumber(value: number) {
  return NUMBER_FORMATTER.format(value)
}

function getStatusLabel(status: number) {
  return STATUS_LABEL_MAP[status] ?? '未知'
}

function getStatusClass(status: number) {
  return STATUS_CLASS_MAP[status] ?? 'border-slate-200 bg-slate-100 text-slate-700'
}

function consumeSavedMessage() {
  const saved = route.query.saved

  if (saved === 'created') {
    successMessage.value = '文章已创建。'
  } else if (saved === 'updated') {
    successMessage.value = '文章已更新。'
  }

  if (typeof saved === 'string') {
    const nextQuery = { ...route.query }
    delete nextQuery.saved
    void router.replace({ query: nextQuery })
  }
}

async function loadArticles() {
  const requestId = ++listRequestId
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await listAdminArticles({
      page: currentPage.value,
      page_size: PAGE_SIZE,
      ...(statusFilter.value === 'all' ? {} : { status: Number(statusFilter.value) }),
    })

    if (requestId !== listRequestId) {
      return
    }

    articles.value = response.data.items
    total.value = response.data.total

    if (currentPage.value > 1 && articles.value.length === 0) {
      currentPage.value = totalPages.value
    }
  } catch (error) {
    if (requestId !== listRequestId) {
      return
    }

    errorMessage.value =
      error instanceof RequestError ? error.message : '获取文章列表失败，请稍后重试。'
  } finally {
    if (requestId === listRequestId) {
      isLoading.value = false
    }
  }
}

async function handleDelete(article: ArticleItem) {
  const confirmed = window.confirm(`确定删除《${article.title}》吗？此操作无法撤销。`)

  if (!confirmed) {
    return
  }

  deletingId.value = article.id
  errorMessage.value = ''
  successMessage.value = ''

  try {
    await deleteArticle(article.id)
    successMessage.value = '文章已删除。'

    if (articles.value.length === 1 && currentPage.value > 1) {
      currentPage.value -= 1
      return
    }

    await loadArticles()
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '删除文章失败，请稍后重试。'
  } finally {
    deletingId.value = null
  }
}

watch(statusFilter, (value, previousValue) => {
  if (value !== previousValue && currentPage.value !== 1) {
    currentPage.value = 1
  }
})

watch([statusFilter, currentPage], () => {
  void loadArticles()
}, { immediate: true })

consumeSavedMessage()
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm lg:flex-row lg:items-center lg:justify-between">
        <div class="space-y-1">
          <p class="text-sm font-medium text-muted-foreground">内容管理</p>
          <h1 class="text-3xl font-semibold tracking-tight">文章列表</h1>
          <p class="text-sm text-muted-foreground">
            管理你的文章草稿、发布内容与归档内容。
          </p>
        </div>

        <Button as-child size="lg">
          <RouterLink :to="{ name: 'admin-article-create' }">
            <FilePlus2 class="size-4" />
            <span>新增文章</span>
          </RouterLink>
        </Button>
      </header>

      <div
        v-if="successMessage"
        class="rounded-xl border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-700"
      >
        {{ successMessage }}
      </div>

      <div
        v-if="errorMessage"
        class="rounded-xl border border-destructive/20 bg-destructive/8 px-4 py-3 text-sm text-destructive"
      >
        {{ errorMessage }}
      </div>

      <section class="rounded-2xl border bg-background shadow-sm">
        <div class="flex flex-col gap-4 border-b px-6 py-5 lg:flex-row lg:items-center lg:justify-between">
          <div>
            <h2 class="text-lg font-semibold">文章记录</h2>
            <p class="text-sm text-muted-foreground">
              共 {{ total }} 篇，当前第 {{ currentPage }} / {{ totalPages }} 页。
            </p>
          </div>

          <div class="flex flex-col gap-3 sm:flex-row sm:items-center">
            <label class="text-sm text-muted-foreground" for="status-filter">状态筛选</label>
            <select
              id="status-filter"
              v-model="statusFilter"
              class="h-9 min-w-40 rounded-md border border-input bg-background px-3 text-sm shadow-xs transition-[color,box-shadow] outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
            >
              <option
                v-for="option in STATUS_OPTIONS"
                :key="option.value"
                :value="option.value"
              >
                {{ option.label }}
              </option>
            </select>
          </div>
        </div>

        <div v-if="isLoading" class="flex min-h-64 items-center justify-center gap-3 px-6 py-10 text-sm text-muted-foreground">
          <LoaderCircle class="size-4 animate-spin" />
          <span>正在加载文章列表...</span>
        </div>

        <template v-else>
          <div v-if="articles.length === 0" class="flex min-h-64 flex-col items-center justify-center gap-3 px-6 py-10 text-center">
            <div class="rounded-full bg-muted p-3 text-muted-foreground">
              <FilePlus2 class="size-5" />
            </div>
            <div class="space-y-1">
              <h3 class="text-base font-medium">还没有文章</h3>
              <p class="text-sm text-muted-foreground">
                你可以先创建第一篇文章，或者切换筛选条件查看其他状态。
              </p>
            </div>
            <Button as-child>
              <RouterLink :to="{ name: 'admin-article-create' }">立即写文章</RouterLink>
            </Button>
          </div>

          <div v-else class="overflow-x-auto">
            <table class="min-w-full text-sm">
              <thead class="bg-muted/40 text-left text-muted-foreground">
                <tr>
                  <th class="px-6 py-3 font-medium">文章</th>
                  <th class="px-6 py-3 font-medium">状态</th>
                  <th class="px-6 py-3 font-medium">阅读量</th>
                  <th class="px-6 py-3 font-medium">发布时间</th>
                  <th class="px-6 py-3 font-medium">最后更新</th>
                  <th class="px-6 py-3 text-right font-medium">操作</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="article in articles"
                  :key="article.id"
                  class="border-t align-top transition-colors hover:bg-muted/20"
                >
                  <td class="px-6 py-4">
                    <div class="space-y-1">
                      <div class="font-medium text-foreground">{{ article.title }}</div>
                      <div class="text-xs text-muted-foreground">/{{ article.slug }}</div>
                      <p v-if="article.summary" class="line-clamp-2 max-w-xl text-sm text-muted-foreground">
                        {{ article.summary }}
                      </p>
                    </div>
                  </td>
                  <td class="px-6 py-4">
                    <span
                      class="inline-flex rounded-full border px-2.5 py-1 text-xs font-medium"
                      :class="getStatusClass(article.status)"
                    >
                      {{ getStatusLabel(article.status) }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">
                    {{ formatNumber(article.view_count) }}
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">
                    {{ formatDateTime(article.published_at) }}
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">
                    {{ formatDateTime(article.updated_at) }}
                  </td>
                  <td class="px-6 py-4">
                    <div class="flex justify-end gap-2">
                      <Button as-child variant="outline" size="sm">
                        <RouterLink :to="{ name: 'admin-article-edit', params: { id: article.id } }">
                          <PencilLine class="size-4" />
                          <span>编辑</span>
                        </RouterLink>
                      </Button>

                      <Button
                        variant="destructive"
                        size="sm"
                        :disabled="deletingId === article.id"
                        @click="handleDelete(article)"
                      >
                        <LoaderCircle v-if="deletingId === article.id" class="size-4 animate-spin" />
                        <Trash2 v-else class="size-4" />
                        <span>{{ deletingId === article.id ? '删除中...' : '删除' }}</span>
                      </Button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>

          <footer class="flex flex-col gap-3 border-t px-6 py-4 sm:flex-row sm:items-center sm:justify-between">
            <p class="text-sm text-muted-foreground">
              每页 {{ PAGE_SIZE }} 条，共 {{ total }} 条记录。
            </p>

            <div class="flex items-center gap-2">
              <Button
                variant="outline"
                size="sm"
                :disabled="currentPage <= 1 || isLoading"
                @click="currentPage -= 1"
              >
                <ChevronLeft class="size-4" />
                上一页
              </Button>

              <div class="min-w-20 text-center text-sm text-muted-foreground">
                第 {{ currentPage }} / {{ totalPages }} 页
              </div>

              <Button
                variant="outline"
                size="sm"
                :disabled="currentPage >= totalPages || isLoading"
                @click="currentPage += 1"
              >
                下一页
                <ChevronRight class="size-4" />
              </Button>
            </div>
          </footer>
        </template>
      </section>
    </div>
  </section>
</template>
