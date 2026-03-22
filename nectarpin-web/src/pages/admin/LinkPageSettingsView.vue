<script setup lang="ts">
import { ref, shallowRef } from 'vue'
import { LoaderCircle, PencilLine, Plus, Trash2 } from 'lucide-vue-next'

import {
  createFriendLinkCategory,
  deleteFriendLinkCategory,
  getFriendLinkPageSettings,
  listFriendLinkCategories,
  saveFriendLinkPageSettings,
  updateFriendLinkCategory,
  type FriendLinkCategoryItem,
  type CreateFriendLinkCategoryPayload,
  type UpdateFriendLinkCategoryPayload,
} from '@/api/link'
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

const introHtml = ref('')
const pageLoading = ref(false)
const pageSaving = ref(false)
const pageError = ref('')

const categories = shallowRef<FriendLinkCategoryItem[]>([])
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
const formDescription = ref('')
const formSortOrder = ref(0)

async function loadPageSettings() {
  pageLoading.value = true
  pageError.value = ''
  try {
    const res = await getFriendLinkPageSettings()
    if (res.code === 200 && res.data) {
      introHtml.value = res.data.intro_html ?? ''
    }
  } catch (e) {
    pageError.value = e instanceof RequestError ? e.message : '加载友链页设置失败。'
  } finally {
    pageLoading.value = false
  }
}

async function handleSaveIntro() {
  pageSaving.value = true
  pageError.value = ''
  successMessage.value = ''
  try {
    const res = await saveFriendLinkPageSettings({ intro_html: introHtml.value })
    if (res.code === 200) {
      successMessage.value = '友链页简介已保存。'
    }
  } catch (e) {
    pageError.value = e instanceof RequestError ? e.message : '保存失败。'
  } finally {
    pageSaving.value = false
  }
}

function openCreateDialog() {
  dialogMode.value = 'create'
  editingId.value = null
  formName.value = ''
  formSlug.value = ''
  formDescription.value = ''
  formSortOrder.value = 0
  dialogError.value = ''
  dialogOpen.value = true
}

function openEditDialog(item: FriendLinkCategoryItem) {
  dialogMode.value = 'edit'
  editingId.value = item.id
  formName.value = item.name
  formSlug.value = item.slug
  formDescription.value = item.description
  formSortOrder.value = item.sort_order
  dialogError.value = ''
  dialogOpen.value = true
}

async function loadCategories() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await listFriendLinkCategories()
    categories.value = response.data.items ?? []
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '获取分组列表失败，请稍后重试。'
  } finally {
    isLoading.value = false
  }
}

async function handleDialogSubmit() {
  dialogError.value = ''

  if (!formName.value.trim()) {
    dialogError.value = '请输入分组名称。'
    return
  }

  dialogSubmitting.value = true

  try {
    if (dialogMode.value === 'create') {
      const payload: CreateFriendLinkCategoryPayload = {
        name: formName.value.trim(),
        ...(formSlug.value.trim() ? { slug: formSlug.value.trim() } : {}),
        ...(formDescription.value.trim() ? { description: formDescription.value.trim() } : {}),
        sort_order: formSortOrder.value,
      }
      await createFriendLinkCategory(payload)
      successMessage.value = '分组已创建。'
    } else {
      const payload: UpdateFriendLinkCategoryPayload = {
        name: formName.value.trim(),
        slug: formSlug.value.trim(),
        description: formDescription.value.trim(),
        sort_order: formSortOrder.value,
      }
      await updateFriendLinkCategory(editingId.value!, payload)
      successMessage.value = '分组已更新。'
    }

    dialogOpen.value = false
    await loadCategories()
  } catch (error) {
    dialogError.value = error instanceof RequestError ? error.message : '操作失败，请稍后重试。'
  } finally {
    dialogSubmitting.value = false
  }
}

async function handleDelete(item: FriendLinkCategoryItem) {
  const confirmed = window.confirm(`确定删除分组「${item.name}」吗？该分组下的友链将变为未分组。`)
  if (!confirmed) return

  deletingId.value = item.id
  errorMessage.value = ''
  successMessage.value = ''

  try {
    await deleteFriendLinkCategory(item.id)
    successMessage.value = '分组已删除。'
    await loadCategories()
  } catch (error) {
    errorMessage.value =
      error instanceof RequestError ? error.message : '删除分组失败，请稍后重试。'
  } finally {
    deletingId.value = null
  }
}

void loadPageSettings()
void loadCategories()
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header
        class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm lg:flex-row lg:items-center lg:justify-between"
      >
        <div class="space-y-1">
          <p class="text-sm font-medium text-muted-foreground">友链管理</p>
          <h1 class="text-3xl font-semibold tracking-tight">页面设置</h1>
          <p class="text-sm text-muted-foreground">
            编辑友链页顶部简介（HTML 片段，与首页简介相同策略），并管理友链分组。
          </p>
        </div>
      </header>

      <div
        v-if="successMessage"
        class="rounded-xl border border-emerald-200 bg-emerald-50 px-4 py-3 text-sm text-emerald-700"
      >
        {{ successMessage }}
      </div>
      <div
        v-if="errorMessage || pageError"
        class="rounded-xl border border-destructive/20 bg-destructive/8 px-4 py-3 text-sm text-destructive"
      >
        {{ errorMessage || pageError }}
      </div>

      <section class="rounded-2xl border bg-background p-6 shadow-sm">
        <h2 class="text-lg font-semibold tracking-tight">友链页简介</h2>
        <p class="mt-1 text-sm text-muted-foreground">
          显示在「朋友」页面标题下方；可留空。仅管理员可编辑，请自行确保 HTML 安全。
        </p>
        <div v-if="pageLoading" class="mt-4 flex items-center gap-2 text-sm text-muted-foreground">
          <LoaderCircle class="size-4 animate-spin" />
          加载中…
        </div>
        <template v-else>
          <textarea
            v-model="introHtml"
            rows="6"
            placeholder="例如简短说明换链规则或致谢…"
            class="placeholder:text-muted-foreground border-input mt-4 min-h-32 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
          />
          <div class="mt-4">
            <Button :disabled="pageSaving" @click="handleSaveIntro">
              <LoaderCircle v-if="pageSaving" class="size-4 animate-spin" />
              <span>{{ pageSaving ? '保存中…' : '保存简介' }}</span>
            </Button>
          </div>
        </template>
      </section>

      <header
        class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm lg:flex-row lg:items-center lg:justify-between"
      >
        <div class="space-y-1">
          <h2 class="text-xl font-semibold tracking-tight">友链分组</h2>
          <p class="text-sm text-muted-foreground">用于前台分区展示；删除分组不会删除友链。</p>
        </div>
        <Button size="lg" @click="openCreateDialog">
          <Plus class="size-4" />
          <span>新增分组</span>
        </Button>
      </header>

      <section class="rounded-2xl border bg-background shadow-sm">
        <div
          v-if="isLoading"
          class="flex min-h-64 items-center justify-center gap-3 px-6 py-10 text-sm text-muted-foreground"
        >
          <LoaderCircle class="size-4 animate-spin" />
          <span>正在加载分组列表...</span>
        </div>

        <template v-else>
          <div
            v-if="categories.length === 0"
            class="flex min-h-64 flex-col items-center justify-center gap-3 px-6 py-10 text-center"
          >
            <div class="space-y-1">
              <h3 class="text-base font-medium">还没有分组</h3>
              <p class="text-sm text-muted-foreground">
                创建分组以便在前台归类友链；也可仅用未分组。
              </p>
            </div>
            <Button @click="openCreateDialog">新增分组</Button>
          </div>

          <div v-else class="overflow-x-auto">
            <table class="min-w-full text-sm">
              <thead class="bg-muted/40 text-left text-muted-foreground">
                <tr>
                  <th class="px-6 py-3 font-medium">分组名称</th>
                  <th class="px-6 py-3 font-medium">Slug</th>
                  <th class="px-6 py-3 font-medium">描述</th>
                  <th class="px-6 py-3 font-medium">排序</th>
                  <th class="px-6 py-3 text-right font-medium">操作</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in categories"
                  :key="item.id"
                  class="border-t transition-colors hover:bg-muted/20"
                >
                  <td class="px-6 py-4 font-medium text-foreground">{{ item.name }}</td>
                  <td class="px-6 py-4 text-muted-foreground">{{ item.slug }}</td>
                  <td class="max-w-xs truncate px-6 py-4 text-muted-foreground">
                    {{ item.description || '-' }}
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">{{ item.sort_order }}</td>
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
          <DialogTitle>{{ dialogMode === 'create' ? '新增分组' : '编辑分组' }}</DialogTitle>
          <DialogDescription>
            {{ dialogMode === 'create' ? '创建一个新的友链分组。' : '修改分组信息。' }}
          </DialogDescription>
        </DialogHeader>

        <form class="space-y-4" @submit.prevent="handleDialogSubmit">
          <div class="space-y-2">
            <label for="fl-cat-name" class="text-sm font-medium">分组名称</label>
            <Input id="fl-cat-name" v-model="formName" placeholder="请输入分组名称" />
          </div>
          <div class="space-y-2">
            <label for="fl-cat-slug" class="text-sm font-medium">Slug</label>
            <Input id="fl-cat-slug" v-model="formSlug" placeholder="留空自动生成" />
          </div>
          <div class="space-y-2">
            <label for="fl-cat-desc" class="text-sm font-medium">描述</label>
            <textarea
              id="fl-cat-desc"
              v-model="formDescription"
              rows="3"
              placeholder="该分组下的说明文字（可选，可展示在前台）"
              class="placeholder:text-muted-foreground border-input min-h-20 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
            />
          </div>
          <div class="space-y-2">
            <label for="fl-cat-sort" class="text-sm font-medium">排序权重</label>
            <Input id="fl-cat-sort" v-model.number="formSortOrder" type="number" placeholder="0" />
            <p class="text-xs text-muted-foreground">数值越小越靠前。</p>
          </div>

          <p
            v-if="dialogError"
            class="rounded-md border border-destructive/20 bg-destructive/8 px-3 py-2 text-sm text-destructive"
          >
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
