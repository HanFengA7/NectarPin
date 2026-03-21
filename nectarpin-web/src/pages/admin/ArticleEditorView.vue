<script setup lang="ts">
import { computed, defineAsyncComponent, onMounted, reactive, ref } from 'vue'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import { ChevronLeft, LoaderCircle, Save } from 'lucide-vue-next'

import {
  createArticle,
  getAdminArticleById,
  updateArticle,
  type CreateArticlePayload,
  type UpdateArticlePayload,
} from '@/api/article'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import { RequestError } from '@/utils/req'

const MarkdownEditor = defineAsyncComponent(() => import('@/components/admin/VditorEditor.vue'))

type ArticleStatusValue = '0' | '1' | '2'

const STATUS_OPTIONS = [
  { value: '0', label: '草稿' },
  { value: '1', label: '已发布' },
  { value: '2', label: '已归档' },
] as const

const route = useRoute()
const router = useRouter()

const form = reactive({
  title: '',
  slug: '',
  summary: '',
  cover_image: '',
  status: '0' as ArticleStatusValue,
  content: '',
})

const isLoading = ref(false)
const isSubmitting = ref(false)
const errorMessage = ref('')
const loadErrorMessage = ref('')
let detailRequestId = 0

const articleId = computed(() => {
  const rawId = route.params.id

  if (typeof rawId !== 'string') {
    return null
  }

  const parsed = Number(rawId)

  if (!Number.isInteger(parsed) || parsed <= 0) {
    return null
  }

  return parsed
})

const isEditMode = computed(() => route.name === 'admin-article-edit')
const pageTitle = computed(() => (isEditMode.value ? '编辑文章' : '新增文章'))
const pageDescription = computed(() =>
  isEditMode.value
    ? '更新文章内容、状态与文章地址。'
    : '创建一篇新文章并保存为草稿或直接发布。',
)
const submitLabel = computed(() => {
  if (isSubmitting.value) {
    return isEditMode.value ? '保存中...' : '创建中...'
  }

  return isEditMode.value ? '保存修改' : '创建文章'
})

function buildPayload() {
  const payload = {
    title: form.title.trim(),
    content: form.content,
    status: Number(form.status),
    ...(form.slug.trim() ? { slug: form.slug.trim() } : {}),
    ...(form.summary.trim() ? { summary: form.summary.trim() } : {}),
    ...(form.cover_image.trim() ? { cover_image: form.cover_image.trim() } : {}),
  }

  return payload
}

async function loadArticle() {
  if (!articleId.value) {
    loadErrorMessage.value = '无效的文章 ID。'
    return
  }

  const requestId = ++detailRequestId
  isLoading.value = true
  loadErrorMessage.value = ''

  try {
    const response = await getAdminArticleById(articleId.value)

    if (requestId !== detailRequestId) {
      return
    }

    const article = response.data

    form.title = article.title
    form.slug = article.slug
    form.summary = article.summary
    form.cover_image = article.cover_image
    form.status = String(article.status) as ArticleStatusValue
    form.content = article.content
  } catch (error) {
    if (requestId !== detailRequestId) {
      return
    }

    loadErrorMessage.value =
      error instanceof RequestError ? error.message : '获取文章详情失败，请稍后重试。'
  } finally {
    if (requestId === detailRequestId) {
      isLoading.value = false
    }
  }
}

async function handleSubmit() {
  errorMessage.value = ''

  if (!form.title.trim()) {
    errorMessage.value = '请输入文章标题。'
    return
  }

  if (!form.content.trim()) {
    errorMessage.value = '请输入文章内容。'
    return
  }

  isSubmitting.value = true

  try {
    if (isEditMode.value) {
      if (!articleId.value) {
        throw new Error('无效的文章 ID')
      }

      await updateArticle(articleId.value, buildPayload() as UpdateArticlePayload)
    } else {
      await createArticle(buildPayload() as CreateArticlePayload)
    }

    await router.replace({
      name: 'admin-articles',
      query: {
        saved: isEditMode.value ? 'updated' : 'created',
      },
    })
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '保存文章失败，请稍后重试。'
  } finally {
    isSubmitting.value = false
  }
}

onMounted(() => {
  if (isEditMode.value) {
    void loadArticle()
  }
})
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm sm:flex-row sm:items-start sm:justify-between">
        <div class="space-y-1">
          <p class="text-sm font-medium text-muted-foreground">内容管理</p>
          <h1 class="text-3xl font-semibold tracking-tight">{{ pageTitle }}</h1>
          <p class="text-sm text-muted-foreground">{{ pageDescription }}</p>
        </div>

        <Button as-child variant="outline">
          <RouterLink :to="{ name: 'admin-articles' }">
            <ChevronLeft class="size-4" />
            <span>返回列表</span>
          </RouterLink>
        </Button>
      </header>

      <div
        v-if="loadErrorMessage"
        class="rounded-xl border border-destructive/20 bg-destructive/8 px-4 py-3 text-sm text-destructive"
      >
        {{ loadErrorMessage }}
      </div>

      <div
        v-if="errorMessage"
        class="rounded-xl border border-destructive/20 bg-destructive/8 px-4 py-3 text-sm text-destructive"
      >
        {{ errorMessage }}
      </div>

      <section v-if="isLoading" class="flex min-h-96 items-center justify-center gap-3 rounded-2xl border bg-background p-6 text-sm text-muted-foreground shadow-sm">
        <LoaderCircle class="size-4 animate-spin" />
        <span>正在加载文章内容...</span>
      </section>

      <section
        v-else-if="!loadErrorMessage"
        class="rounded-2xl border bg-background p-6 shadow-sm"
      >
        <form class="space-y-6" @submit.prevent="handleSubmit">
          <div class="grid gap-6 lg:grid-cols-2">
            <div class="space-y-2 lg:col-span-2">
              <label for="title" class="text-sm font-medium">文章标题</label>
              <Input
                id="title"
                v-model="form.title"
                type="text"
                placeholder="请输入文章标题"
              />
            </div>

            <div class="space-y-2">
              <label for="slug" class="text-sm font-medium">文章地址</label>
              <Input
                id="slug"
                v-model="form.slug"
                type="text"
                placeholder="留空时将根据标题自动生成"
              />
              <p class="text-xs text-muted-foreground">
                建议使用英文、数字与连字符，便于生成稳定的访问链接。
              </p>
            </div>

            <div class="space-y-2">
              <label for="status" class="text-sm font-medium">发布状态</label>
              <select
                id="status"
                v-model="form.status"
                class="h-9 w-full rounded-md border border-input bg-background px-3 text-sm shadow-xs transition-[color,box-shadow] outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
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

            <div class="space-y-2 lg:col-span-2">
              <label for="cover-image" class="text-sm font-medium">封面图片</label>
              <Input
                id="cover-image"
                v-model="form.cover_image"
                type="url"
                placeholder="请输入封面图片 URL（可选）"
              />
            </div>

            <div class="space-y-2 lg:col-span-2">
              <label for="summary" class="text-sm font-medium">文章摘要</label>
              <textarea
                id="summary"
                v-model="form.summary"
                rows="4"
                placeholder="请输入文章摘要（可选）"
                class="placeholder:text-muted-foreground selection:bg-primary selection:text-primary-foreground dark:bg-input/30 border-input min-h-28 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs transition-[color,box-shadow] outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50 disabled:pointer-events-none disabled:cursor-not-allowed disabled:opacity-50"
              />
            </div>

            <div class="space-y-2 lg:col-span-2">
              <label for="content" class="text-sm font-medium">正文内容</label>
              <MarkdownEditor
                id="content"
                v-model="form.content"
                placeholder="请输入文章正文（Markdown）"
                :min-height="560"
              />
            </div>
          </div>

          <div class="flex flex-col gap-3 border-t pt-6 sm:flex-row sm:items-center sm:justify-end">
            <Button as-child variant="outline">
              <RouterLink :to="{ name: 'admin-articles' }">取消</RouterLink>
            </Button>

            <Button type="submit" size="lg" :disabled="isSubmitting">
              <LoaderCircle v-if="isSubmitting" class="size-4 animate-spin" />
              <Save v-else class="size-4" />
              <span>{{ submitLabel }}</span>
            </Button>
          </div>
        </form>
      </section>
    </div>
  </section>
</template>
