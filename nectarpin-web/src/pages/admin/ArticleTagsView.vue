<script setup lang="ts">
import { ref, shallowRef } from 'vue'
import { LoaderCircle, PencilLine, Plus, Trash2 } from 'lucide-vue-next'

import {
  createTag,
  deleteTag,
  listTags,
  updateTag,
  type ArticleTagItem,
  type CreateTagPayload,
  type UpdateTagPayload,
} from '@/api/article'
import { Button } from '@/components/ui/button'
import { Input } from '@/components/ui/input'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
} from '@/components/ui/dialog'
import { RequestError } from '@/utils/req'

const tags = shallowRef<ArticleTagItem[]>([])
const isLoading = ref(false)
const errorMessage = ref('')
const successMessage = ref('')
const deletingId = ref<number | null>(null)

const dialogOpen = ref(false)
const dialogMode = ref<'create' | 'edit'>('create')
const dialogSubmitting = ref(false)
const dialogError = ref('')
const editingId = ref<number | null>(null)

const formName = ref('')
const formSlug = ref('')

function openCreateDialog() {
  dialogMode.value = 'create'
  editingId.value = null
  formName.value = ''
  formSlug.value = ''
  dialogError.value = ''
  dialogOpen.value = true
}

function openEditDialog(item: ArticleTagItem) {
  dialogMode.value = 'edit'
  editingId.value = item.id
  formName.value = item.name
  formSlug.value = item.slug
  dialogError.value = ''
  dialogOpen.value = true
}

async function loadTags() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await listTags()
    tags.value = response.data.items ?? []
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '获取标签列表失败，请稍后重试。'
  } finally {
    isLoading.value = false
  }
}

async function handleDialogSubmit() {
  dialogError.value = ''

  if (!formName.value.trim()) {
    dialogError.value = '请输入标签名称。'
    return
  }

  dialogSubmitting.value = true

  try {
    if (dialogMode.value === 'create') {
      const payload: CreateTagPayload = {
        name: formName.value.trim(),
        ...(formSlug.value.trim() ? { slug: formSlug.value.trim() } : {}),
      }
      await createTag(payload)
      successMessage.value = '标签已创建。'
    } else {
      const payload: UpdateTagPayload = {
        name: formName.value.trim(),
        slug: formSlug.value.trim(),
      }
      await updateTag(editingId.value!, payload)
      successMessage.value = '标签已更新。'
    }

    dialogOpen.value = false
    await loadTags()
  } catch (error) {
    dialogError.value =
      error instanceof RequestError ? error.message : '操作失败，请稍后重试。'
  } finally {
    dialogSubmitting.value = false
  }
}

async function handleDelete(item: ArticleTagItem) {
  const confirmed = window.confirm(
    `确定删除标签「${item.name}」吗？文章与此标签的关联将被移除。`,
  )
  if (!confirmed) return

  deletingId.value = item.id
  errorMessage.value = ''
  successMessage.value = ''

  try {
    await deleteTag(item.id)
    successMessage.value = '标签已删除。'
    await loadTags()
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '删除标签失败，请稍后重试。'
  } finally {
    deletingId.value = null
  }
}

void loadTags()
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm lg:flex-row lg:items-center lg:justify-between">
        <div class="space-y-1">
          <p class="text-sm font-medium text-muted-foreground">内容管理</p>
          <h1 class="text-3xl font-semibold tracking-tight">文章标签</h1>
          <p class="text-sm text-muted-foreground">管理文章标签，标签可灵活关联多篇文章。</p>
        </div>
        <Button size="lg" @click="openCreateDialog">
          <Plus class="size-4" />
          <span>新增标签</span>
        </Button>
      </header>

      <div v-if="successMessage" class="rounded-xl border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-700">
        {{ successMessage }}
      </div>
      <div v-if="errorMessage" class="rounded-xl border border-destructive/20 bg-destructive/8 px-4 py-3 text-sm text-destructive">
        {{ errorMessage }}
      </div>

      <section class="rounded-2xl border bg-background shadow-sm">
        <div v-if="isLoading" class="flex min-h-64 items-center justify-center gap-3 px-6 py-10 text-sm text-muted-foreground">
          <LoaderCircle class="size-4 animate-spin" />
          <span>正在加载标签列表...</span>
        </div>

        <template v-else>
          <div v-if="tags.length === 0" class="flex min-h-64 flex-col items-center justify-center gap-3 px-6 py-10 text-center">
            <div class="space-y-1">
              <h3 class="text-base font-medium">还没有标签</h3>
              <p class="text-sm text-muted-foreground">创建第一个标签来标记你的文章。</p>
            </div>
            <Button @click="openCreateDialog">新增标签</Button>
          </div>

          <div v-else class="overflow-x-auto">
            <table class="min-w-full text-sm">
              <thead class="bg-muted/40 text-left text-muted-foreground">
                <tr>
                  <th class="px-6 py-3 font-medium">标签名称</th>
                  <th class="px-6 py-3 font-medium">Slug</th>
                  <th class="px-6 py-3 font-medium">文章数</th>
                  <th class="px-6 py-3 text-right font-medium">操作</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in tags"
                  :key="item.id"
                  class="border-t transition-colors hover:bg-muted/20"
                >
                  <td class="px-6 py-4 font-medium text-foreground">{{ item.name }}</td>
                  <td class="px-6 py-4 text-muted-foreground">{{ item.slug }}</td>
                  <td class="px-6 py-4 text-muted-foreground">{{ item.article_count }}</td>
                  <td class="px-6 py-4">
                    <div class="flex justify-end gap-2">
                      <Button variant="outline" size="sm" @click="openEditDialog(item)">
                        <PencilLine class="size-4" />
                        <span>编辑</span>
                      </Button>
                      <Button
                        variant="destructive"
                        size="sm"
                        :disabled="deletingId === item.id"
                        @click="handleDelete(item)"
                      >
                        <LoaderCircle v-if="deletingId === item.id" class="size-4 animate-spin" />
                        <Trash2 v-else class="size-4" />
                        <span>{{ deletingId === item.id ? '删除中...' : '删除' }}</span>
                      </Button>
                    </div>
                  </td>
                </tr>
              </tbody>
            </table>
          </div>
        </template>
      </section>
    </div>

    <Dialog v-model:open="dialogOpen">
      <DialogContent class="sm:max-w-lg">
        <DialogHeader>
          <DialogTitle>{{ dialogMode === 'create' ? '新增标签' : '编辑标签' }}</DialogTitle>
          <DialogDescription>
            {{ dialogMode === 'create' ? '创建一个新的文章标签。' : '修改标签信息。' }}
          </DialogDescription>
        </DialogHeader>

        <form class="space-y-4" @submit.prevent="handleDialogSubmit">
          <div class="space-y-2">
            <label for="tag-name" class="text-sm font-medium">标签名称</label>
            <Input id="tag-name" v-model="formName" placeholder="请输入标签名称" />
          </div>
          <div class="space-y-2">
            <label for="tag-slug" class="text-sm font-medium">Slug</label>
            <Input id="tag-slug" v-model="formSlug" placeholder="留空自动生成" />
          </div>

          <p v-if="dialogError" class="rounded-md border border-destructive/20 bg-destructive/8 px-3 py-2 text-sm text-destructive">
            {{ dialogError }}
          </p>

          <DialogFooter>
            <Button type="button" variant="outline" @click="dialogOpen = false">取消</Button>
            <Button type="submit" :disabled="dialogSubmitting">
              <LoaderCircle v-if="dialogSubmitting" class="size-4 animate-spin" />
              <span>{{ dialogSubmitting ? '保存中...' : '保存' }}</span>
            </Button>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  </section>
</template>
