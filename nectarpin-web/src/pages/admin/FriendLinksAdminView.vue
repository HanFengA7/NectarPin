<script setup lang="ts">
import { computed, ref, shallowRef } from 'vue'
import { ExternalLink, LoaderCircle, PencilLine, Plus, Trash2 } from 'lucide-vue-next'

import {
  createFriendLink,
  deleteFriendLink,
  listFriendLinkCategories,
  listFriendLinksAdmin,
  updateFriendLink,
  type FriendLinkAdminItem,
  type FriendLinkCategoryItem,
  type UpdateFriendLinkPayload,
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

const links = shallowRef<FriendLinkAdminItem[]>([])
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

const formCategoryId = ref<string>('')
const formTitle = ref('')
const formUrl = ref('')
const formDescription = ref('')
const formAvatarUrl = ref('')
const formSortOrder = ref(0)
const formEnabled = ref(true)

const categoryNameById = computed(() => {
  const m = new Map<number, string>()
  for (const c of categories.value) {
    m.set(c.id, c.name)
  }
  return m
})

function categoryLabel(id: number | null) {
  if (id == null) return '未分组'
  return categoryNameById.value.get(id) ?? `#${id}`
}

function openCreateDialog() {
  dialogMode.value = 'create'
  editingId.value = null
  formCategoryId.value = ''
  formTitle.value = ''
  formUrl.value = ''
  formDescription.value = ''
  formAvatarUrl.value = ''
  formSortOrder.value = 0
  formEnabled.value = true
  dialogError.value = ''
  dialogOpen.value = true
}

function openEditDialog(item: FriendLinkAdminItem) {
  dialogMode.value = 'edit'
  editingId.value = item.id
  formCategoryId.value =
    item.category_id != null && item.category_id > 0 ? String(item.category_id) : ''
  formTitle.value = item.title
  formUrl.value = item.url
  formDescription.value = item.description ?? ''
  formAvatarUrl.value = item.avatar_url ?? ''
  formSortOrder.value = item.sort_order
  formEnabled.value = item.is_enabled
  dialogError.value = ''
  dialogOpen.value = true
}

function parseCategoryIdForCreate(): number | null | undefined {
  const s = formCategoryId.value.trim()
  if (!s) return undefined
  const n = Number(s)
  if (!Number.isFinite(n) || n <= 0) return undefined
  return n
}

async function loadAll() {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const [linkRes, catRes] = await Promise.all([
      listFriendLinksAdmin(),
      listFriendLinkCategories(),
    ])
    links.value = linkRes.data.items ?? []
    categories.value = catRes.data.items ?? []
  } catch (error) {
    errorMessage.value = error instanceof RequestError ? error.message : '加载失败，请稍后重试。'
  } finally {
    isLoading.value = false
  }
}

async function handleDialogSubmit() {
  dialogError.value = ''

  if (!formTitle.value.trim()) {
    dialogError.value = '请输入标题。'
    return
  }
  if (!formUrl.value.trim()) {
    dialogError.value = '请输入链接 URL。'
    return
  }

  dialogSubmitting.value = true

  try {
    if (dialogMode.value === 'create') {
      const catId = parseCategoryIdForCreate()
      await createFriendLink({
        title: formTitle.value.trim(),
        url: formUrl.value.trim(),
        ...(formDescription.value.trim() ? { description: formDescription.value.trim() } : {}),
        ...(formAvatarUrl.value.trim() ? { avatar_url: formAvatarUrl.value.trim() } : {}),
        sort_order: formSortOrder.value,
        is_enabled: formEnabled.value,
        ...(catId != null ? { category_id: catId } : {}),
      })
      successMessage.value = '友链已创建。'
    } else {
      const catRaw = formCategoryId.value.trim()
      const payload: UpdateFriendLinkPayload = {
        title: formTitle.value.trim(),
        url: formUrl.value.trim(),
        description: formDescription.value.trim(),
        avatar_url: formAvatarUrl.value.trim(),
        sort_order: formSortOrder.value,
        is_enabled: formEnabled.value,
      }
      if (!catRaw) {
        payload.category_id = 0
      } else {
        const n = Number(catRaw)
        if (Number.isFinite(n) && n > 0) {
          payload.category_id = n
        }
      }
      await updateFriendLink(editingId.value!, payload)
      successMessage.value = '友链已更新。'
    }

    dialogOpen.value = false
    await loadAll()
  } catch (error) {
    dialogError.value = error instanceof RequestError ? error.message : '操作失败，请稍后重试。'
  } finally {
    dialogSubmitting.value = false
  }
}

async function handleDelete(item: FriendLinkAdminItem) {
  const confirmed = window.confirm(`确定删除友链「${item.title}」吗？`)
  if (!confirmed) return

  deletingId.value = item.id
  errorMessage.value = ''
  successMessage.value = ''

  try {
    await deleteFriendLink(item.id)
    successMessage.value = '友链已删除。'
    await loadAll()
  } catch (error) {
    errorMessage.value = error instanceof RequestError ? error.message : '删除失败，请稍后重试。'
  } finally {
    deletingId.value = null
  }
}

void loadAll()
</script>

<template>
  <section class="min-h-svh bg-muted/20">
    <div class="flex w-full flex-col gap-6 px-4 py-6 sm:px-6 lg:px-8">
      <header
        class="flex flex-col gap-4 rounded-2xl border bg-background p-6 shadow-sm lg:flex-row lg:items-center lg:justify-between"
      >
        <div class="space-y-1">
          <p class="text-sm font-medium text-muted-foreground">友链管理</p>
          <h1 class="text-3xl font-semibold tracking-tight">友链列表</h1>
          <p class="text-sm text-muted-foreground">
            维护前台「朋友」页面展示的链接，支持分组与启用状态。
          </p>
        </div>
        <Button size="lg" @click="openCreateDialog">
          <Plus class="size-4" />
          <span>新增友链</span>
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
        <div
          v-if="isLoading"
          class="flex min-h-64 items-center justify-center gap-3 px-6 py-10 text-sm text-muted-foreground"
        >
          <LoaderCircle class="size-4 animate-spin" />
          <span>正在加载...</span>
        </div>

        <template v-else>
          <div
            v-if="links.length === 0"
            class="flex min-h-64 flex-col items-center justify-center gap-3 px-6 py-10 text-center"
          >
            <div class="space-y-1">
              <h3 class="text-base font-medium">还没有友链</h3>
              <p class="text-sm text-muted-foreground">添加后将在前台「朋友」页展示（需启用）。</p>
            </div>
            <Button @click="openCreateDialog">新增友链</Button>
          </div>

          <div v-else class="overflow-x-auto">
            <table class="min-w-full text-sm">
              <thead class="bg-muted/40 text-left text-muted-foreground">
                <tr>
                  <th class="px-6 py-3 font-medium">标题</th>
                  <th class="px-6 py-3 font-medium">分组</th>
                  <th class="px-6 py-3 font-medium">链接</th>
                  <th class="px-6 py-3 font-medium">排序</th>
                  <th class="px-6 py-3 font-medium">启用</th>
                  <th class="px-6 py-3 text-right font-medium">操作</th>
                </tr>
              </thead>
              <tbody>
                <tr
                  v-for="item in links"
                  :key="item.id"
                  class="border-t transition-colors hover:bg-muted/20"
                >
                  <td class="px-6 py-4 font-medium text-foreground">
                    <span class="inline-flex items-center gap-2">
                      <img
                        v-if="item.avatar_url"
                        :src="item.avatar_url"
                        :alt="''"
                        class="size-8 rounded-full border border-border object-cover"
                        loading="lazy"
                      />
                      {{ item.title }}
                    </span>
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">
                    {{ categoryLabel(item.category_id) }}
                  </td>
                  <td class="max-w-[12rem] truncate px-6 py-4">
                    <a
                      :href="item.url"
                      class="inline-flex items-center gap-1 text-blue-600 underline-offset-4 hover:underline dark:text-blue-400"
                      target="_blank"
                      rel="noopener noreferrer"
                    >
                      <ExternalLink class="size-3.5 shrink-0 opacity-70" />
                      {{ item.url }}
                    </a>
                  </td>
                  <td class="px-6 py-4 text-muted-foreground">{{ item.sort_order }}</td>
                  <td class="px-6 py-4 text-muted-foreground">
                    {{ item.is_enabled ? '是' : '否' }}
                  </td>
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
          <DialogTitle>{{ dialogMode === 'create' ? '新增友链' : '编辑友链' }}</DialogTitle>
          <DialogDescription>链接须为 http(s) 完整地址。</DialogDescription>
        </DialogHeader>

        <form class="space-y-4" @submit.prevent="handleDialogSubmit">
          <div class="space-y-2">
            <label for="fl-cat" class="text-sm font-medium">分组</label>
            <select
              id="fl-cat"
              v-model="formCategoryId"
              class="border-input h-9 w-full rounded-md border bg-transparent px-3 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
            >
              <option value="">未分组</option>
              <option v-for="c in categories" :key="c.id" :value="String(c.id)">
                {{ c.name }}
              </option>
            </select>
          </div>
          <div class="space-y-2">
            <label for="fl-title" class="text-sm font-medium">标题</label>
            <Input id="fl-title" v-model="formTitle" placeholder="站点或博主名称" />
          </div>
          <div class="space-y-2">
            <label for="fl-url" class="text-sm font-medium">URL</label>
            <Input id="fl-url" v-model="formUrl" placeholder="https://example.com" />
          </div>
          <div class="space-y-2">
            <label for="fl-desc" class="text-sm font-medium">简介</label>
            <textarea
              id="fl-desc"
              v-model="formDescription"
              rows="2"
              placeholder="一行简介（可选）"
              class="placeholder:text-muted-foreground border-input min-h-16 w-full rounded-md border bg-transparent px-3 py-2 text-sm shadow-xs outline-none focus-visible:border-ring focus-visible:ring-[3px] focus-visible:ring-ring/50"
            />
          </div>
          <div class="space-y-2">
            <label for="fl-avatar" class="text-sm font-medium">头像图 URL</label>
            <Input id="fl-avatar" v-model="formAvatarUrl" placeholder="https://…（可选）" />
          </div>
          <div class="space-y-2">
            <label for="fl-sort" class="text-sm font-medium">排序</label>
            <Input id="fl-sort" v-model.number="formSortOrder" type="number" />
          </div>
          <div class="flex items-center gap-2">
            <input
              id="fl-en"
              v-model="formEnabled"
              type="checkbox"
              class="size-4 rounded border-input"
            />
            <label for="fl-en" class="text-sm font-medium">在前台展示</label>
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
