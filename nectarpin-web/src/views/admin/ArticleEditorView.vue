<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted, nextTick } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { MdEditor, type ExposeParam } from 'md-editor-v3'
import 'md-editor-v3/lib/style.css'
import { articleApi, type Article } from '@/api/article'

const route = useRoute()
const router = useRouter()

const editorRef = ref<ExposeParam | null>(null)
const articleId = computed(() => route.query.id ? Number(route.query.id) : null)
const isEditMode = computed(() => !!articleId.value)

const title = ref('')
const content = ref('')
const summary = ref('')
const coverImage = ref('')
const status = ref(0)
const tags = ref<string[]>([])
const tagInput = ref('')

const isPreview = ref(false)
const isSaving = ref(false)
const isLoading = ref(false)
const lastSavedContent = ref('')

const toastMessage = ref('')
const toastType = ref<'success' | 'error' | 'info'>('info')
const showToast = ref(false)
const editorModeRef = ref<HTMLElement | null>(null)
const editorIndicatorStyle = ref({ left: '0px', width: '0px' })

const handleUploadImg = async (files: File[]) => {
  const file = files[0]
  if (!file) return ''
  return new URL(URL.createObjectURL(file), window.location.origin).href
}

const hasUnsavedChanges = ref(false)
const autoSaveTimer = ref<ReturnType<typeof setTimeout> | null>(null)
const AUTO_SAVE_DELAY = 30000

const statusOptions = [
  { value: 0, label: '草稿', icon: '📝', color: '#F59E0B' },
  { value: 1, label: '已发布', icon: '✅', color: '#10B981' },
  { value: 2, label: '已下架', icon: '📦', color: '#6B7280' },
]

const titleMaxLength = 200
const isTitleValid = computed(() => {
  return title.value.trim().length > 0 && title.value.length <= titleMaxLength
})
const isContentValid = computed(() => content.value.trim().length > 0)
const canSubmit = computed(() => isTitleValid.value && isContentValid.value && !isSaving.value)

watch([title, content, summary, coverImage, status, tags], () => {
  hasUnsavedChanges.value = title.value !== lastSavedContent.value || content.value !== lastSavedContent.value
}, { deep: true })

watch(content, () => {
  scheduleAutoSave()
})

watch(isPreview, async () => {
  await nextTick()
  updateEditorIndicator()
})

watch(isLoading, async (loading) => {
  if (!loading) {
    await nextTick()
    updateEditorIndicator()
  }
})

const updateEditorIndicator = () => {
  if (!editorModeRef.value) {
    setTimeout(updateEditorIndicator, 10)
    return
  }
  const activeBtn = editorModeRef.value.querySelector('.mode-btn.active') as HTMLElement
  if (!activeBtn) return
  const containerRect = editorModeRef.value.getBoundingClientRect()
  const btnRect = activeBtn.getBoundingClientRect()
  editorIndicatorStyle.value = {
    left: `${btnRect.left - containerRect.left}px`,
    width: `${btnRect.width}px`,
  }
}

const showToastMessage = (msg: string, type: 'success' | 'error' | 'info' = 'info') => {
  toastMessage.value = msg
  toastType.value = type
  showToast.value = true
  setTimeout(() => {
    showToast.value = false
  }, 3000)
}

const scheduleAutoSave = () => {
  if (autoSaveTimer.value) {
    clearTimeout(autoSaveTimer.value)
  }
  if (isEditMode.value && content.value.trim()) {
    autoSaveTimer.value = setTimeout(() => {
      handleAutoSave()
    }, AUTO_SAVE_DELAY)
  }
}

const handleAutoSave = async () => {
  if (!isEditMode.value || !articleId.value || !hasUnsavedChanges.value) return
  try {
    await articleApi.update(articleId.value, {
      title: title.value,
      content: content.value,
      summary: summary.value,
      cover_image: coverImage.value,
      status: status.value,
    })
    lastSavedContent.value = content.value
    hasUnsavedChanges.value = false
    showToastMessage('已自动保存', 'success')
  } catch (error) {
    console.error('Auto-save failed:', error)
  }
}

const loadArticle = async (id: number) => {
  isLoading.value = true
  try {
    const response = await articleApi.getById(id)
    if (response.data) {
      const article = response.data
      title.value = article.title || ''
      content.value = article.content || ''
      summary.value = article.summary || ''
      coverImage.value = article.cover_image || ''
      status.value = article.status ?? 0
      lastSavedContent.value = content.value
    }
  } catch (error: any) {
    showToastMessage(error?.response?.data?.message || '加载文章失败', 'error')
  } finally {
    isLoading.value = false
  }
}

const handleSubmit = async (publishStatus?: number) => {
  if (!isTitleValid.value) {
    showToastMessage('请输入有效的文章标题（1-200字符）', 'error')
    return
  }
  if (!isContentValid.value) {
    showToastMessage('文章内容不能为空', 'error')
    return
  }

  isSaving.value = true
  const articleData: Partial<Article> = {
    title: title.value.trim(),
    content: content.value,
    summary: summary.value.trim() || undefined,
    cover_image: coverImage.value.trim() || undefined,
    status: publishStatus !== undefined ? publishStatus : status.value,
  }

  try {
    if (isEditMode.value && articleId.value) {
      const response = await articleApi.update(articleId.value, articleData)
      showToastMessage('文章更新成功', 'success')
      lastSavedContent.value = content.value
      hasUnsavedChanges.value = false
      if (response.data?.id) {
        router.replace({ query: { id: response.data.id.toString() } })
      }
    } else {
      const response = await articleApi.create(articleData)
      showToastMessage('文章创建成功', 'success')
      if (response.data?.id) {
        router.replace({ query: { id: response.data.id.toString() } })
      }
      setTimeout(() => {
        router.push('/admin/articles')
      }, 1500)
    }
  } catch (error: any) {
    showToastMessage(error?.response?.data?.message || (isEditMode.value ? '更新失败' : '创建失败'), 'error')
  } finally {
    isSaving.value = false
  }
}

const handleSaveDraft = () => {
  status.value = 0
  handleSubmit(0)
}

const handlePublish = () => {
  handleSubmit()
}

const addTag = () => {
  const tag = tagInput.value.trim()
  if (tag && !tags.value.includes(tag)) {
    tags.value.push(tag)
  }
  tagInput.value = ''
}

const removeTag = (index: number) => {
  tags.value.splice(index, 1)
}

const handleTagKeydown = (e: KeyboardEvent) => {
  if (e.key === 'Enter' || e.key === ',') {
    e.preventDefault()
    addTag()
  }
}

const handleCancel = () => {
  if (hasUnsavedChanges.value) {
    if (confirm('您有未保存的更改，确定要离开吗？')) {
      router.push('/admin/articles')
    }
  } else {
    router.push('/admin/articles')
  }
}

onMounted(() => {
  if (articleId.value) {
    loadArticle(articleId.value)
  }
  setTimeout(() => {
    updateEditorIndicator()
  }, 0)
})

onUnmounted(() => {
  if (autoSaveTimer.value) {
    clearTimeout(autoSaveTimer.value)
  }
})
</script>

<template>
  <div class="editor-view">
    <main class="editor-main">
      <header class="editor-header">
        <div class="header-left">
          <button class="back-btn" @click="handleCancel">
            <svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 12H5M12 19l-7-7 7-7"/>
            </svg>
          </button>
          <div class="header-title-group">
            <h1 class="header-title">{{ isEditMode ? '编辑文章' : '创建文章' }}</h1>
            <span v-if="hasUnsavedChanges" class="unsaved-indicator">
              <span class="dot"></span>
              未保存
            </span>
          </div>
        </div>
        <div class="header-actions">
          <button
            class="action-btn draft"
            :disabled="isSaving"
            @click="handleSaveDraft"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/>
              <polyline points="17 21 17 13 7 13 7 21"/><polyline points="7 3 7 8 15 8"/>
            </svg>
            保存草稿
          </button>
          <button
            class="action-btn publish"
            :disabled="!canSubmit"
            @click="handlePublish"
          >
            <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <line x1="22" y1="2" x2="11" y2="13"/><polygon points="22 2 15 22 11 13 2 9 22 2"/>
            </svg>
            {{ isSaving ? '处理中...' : (isEditMode ? '更新文章' : '发布文章') }}
          </button>
        </div>
      </header>

      <div v-if="isLoading" class="loading-state">
        <div class="loading-spinner">
          <div class="spinner-ring"></div>
          <div class="spinner-ring"></div>
          <div class="spinner-ring"></div>
        </div>
        <span class="loading-text">加载文章中...</span>
      </div>

      <div v-else class="editor-content">
        <div class="editor-main-area">
          <div class="title-section">
            <input
              v-model="title"
              type="text"
              class="title-input"
              placeholder="在这里输入文章标题..."
              maxlength="200"
            />
            <div class="title-footer">
              <span class="char-counter" :class="{ warning: title.length > 180 }">
                {{ title.length }} / {{ titleMaxLength }}
              </span>
              <div class="editor-mode-toggle" ref="editorModeRef">
                <div class="mode-indicator" :style="editorIndicatorStyle"></div>
                <button
                  class="mode-btn"
                  :class="{ active: !isPreview }"
                  @click="isPreview = false"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                    <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
                  </svg>
                  编辑
                </button>
                <button
                  class="mode-btn"
                  :class="{ active: isPreview }"
                  @click="isPreview = true"
                >
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                    <circle cx="12" cy="12" r="3"/>
                  </svg>
                  预览
                </button>
              </div>
            </div>
          </div>

          <div class="editor-container">
            <MdEditor
              v-model="content"
              ref="editorRef"
              :editorId="`article-editor-${articleId || 'new'}`"
              :preview="isPreview"
              language="zh-CN"
              :style="{ minHeight: '480px' }"
              placeholder="开始撰写你的文章内容，支持 Markdown 语法..."
              @upload-img="handleUploadImg"
            />
          </div>
        </div>

        <aside class="settings-panel">
          <div class="panel-card">
            <div class="panel-header">
              <h3 class="panel-title">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/>
                </svg>
                文章设置
              </h3>
            </div>
            <div class="panel-body">
              <div class="setting-group">
                <label class="setting-label">发布状态</label>
                <div class="status-selector">
                  <button
                    v-for="opt in statusOptions"
                    :key="opt.value"
                    class="status-option"
                    :class="{ active: status === opt.value }"
                    @click="status = opt.value"
                  >
                    <span class="status-indicator" :style="{ background: opt.color }"></span>
                    <span class="status-text">{{ opt.label }}</span>
                  </button>
                </div>
              </div>

              <div class="setting-group">
                <label class="setting-label">
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                    <circle cx="8.5" cy="8.5" r="1.5"/>
                    <polyline points="21 15 16 10 5 21"/>
                  </svg>
                  封面图片
                </label>
                <input
                  v-model="coverImage"
                  type="url"
                  class="setting-input"
                  placeholder="输入图片 URL..."
                />
                <div v-if="coverImage" class="cover-preview">
                  <img :src="coverImage" alt="封面预览" />
                </div>
              </div>

              <div class="setting-group">
                <label class="setting-label">
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <line x1="17" y1="10" x2="3" y2="10"/><line x1="21" y1="6" x2="3" y2="6"/><line x1="21" y1="14" x2="3" y2="14"/><line x1="17" y1="18" x2="3" y2="18"/>
                  </svg>
                  文章摘要
                </label>
                <textarea
                  v-model="summary"
                  class="setting-textarea"
                  placeholder="简短描述文章内容..."
                  rows="3"
                ></textarea>
              </div>

              <div class="setting-group">
                <label class="setting-label">
                  <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                    <path d="M20.59 13.41l-7.17 7.17a2 2 0 0 1-2.83 0L2 12V2h10l8.59 8.59a2 2 0 0 1 0 2.82z"/>
                    <line x1="7" y1="7" x2="7.01" y2="7"/>
                  </svg>
                  标签
                </label>
                <div class="tags-input-wrapper">
                  <div v-if="tags.length > 0" class="tags-container">
                    <span v-for="(tag, index) in tags" :key="index" class="tag-badge">
                      {{ tag }}
                      <button class="tag-remove-btn" @click="removeTag(index)">
                        <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                          <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
                        </svg>
                      </button>
                    </span>
                  </div>
                  <input
                    v-model="tagInput"
                    type="text"
                    class="tag-input-field"
                    placeholder="输入标签后按回车..."
                    @keydown="handleTagKeydown"
                    @blur="addTag"
                  />
                </div>
              </div>
            </div>
          </div>

          <div class="panel-card tips-card">
            <div class="panel-header">
              <h3 class="panel-title">
                <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/>
                </svg>
                写作提示
              </h3>
            </div>
            <div class="panel-body">
              <ul class="tips-list">
                <li>使用 <code>#</code> 标记标题层级</li>
                <li>使用 <code>**文字**</code> 加粗文本</li>
                <li>使用 <code>`code`</code> 标记代码</li>
                <li>拖拽或粘贴图片即可上传</li>
              </ul>
            </div>
          </div>
        </aside>
      </div>
    </main>

    <Transition name="toast">
      <div v-if="showToast" class="toast-notification" :class="toastType">
        <svg v-if="toastType === 'success'" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/>
        </svg>
        <svg v-else-if="toastType === 'error'" xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/><line x1="15" y1="9" x2="9" y2="15"/><line x1="9" y1="9" x2="15" y2="15"/>
        </svg>
        <svg v-else xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/><line x1="12" y1="16" x2="12" y2="12"/><line x1="12" y1="8" x2="12.01" y2="8"/>
        </svg>
        <span>{{ toastMessage }}</span>
      </div>
    </Transition>
  </div>
</template>

<style scoped>
.editor-view {
  padding: var(--space-page);
  min-height: 100vh;
  background: var(--color-bg);
}

.editor-main {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: var(--space-section);
}

.editor-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: var(--space-section);
  padding-bottom: var(--space-gap);
  border-bottom: 1px solid var(--color-border);
}

.header-left {
  display: flex;
  align-items: center;
  gap: 16px;
}

.back-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  background: var(--color-surface);
  color: var(--color-text);
  cursor: pointer;
  transition: all 0.2s ease;
}

.back-btn:hover {
  background: var(--color-surface-alt);
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.header-title-group {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title {
  font-family: var(--font-heading);
  font-size: 1.5rem;
  font-weight: 700;
  color: var(--color-text);
  letter-spacing: -0.02em;
}

.unsaved-indicator {
  display: flex;
  align-items: center;
  gap: 6px;
  padding: 4px 10px;
  background: #FEF3C7;
  color: #92400E;
  border-radius: var(--radius-full);
  font-size: 0.75rem;
  font-weight: 500;
}

.unsaved-indicator .dot {
  width: 6px;
  height: 6px;
  background: #F59E0B;
  border-radius: 50%;
  animation: pulse 1.5s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.4; }
}

.header-actions {
  display: flex;
  gap: 10px;
}

.action-btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 18px;
  border-radius: var(--radius-lg);
  font-size: 0.875rem;
  font-weight: 600;
  border: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-btn.draft {
  background: var(--color-surface);
  color: var(--color-text);
  border: 1px solid var(--color-border);
}

.action-btn.draft:hover:not(:disabled) {
  background: var(--color-surface-alt);
  border-color: var(--color-primary);
  color: var(--color-primary);
}

.action-btn.publish {
  background: var(--color-primary);
  color: white;
  box-shadow: 0 2px 8px rgba(47, 107, 255, 0.25);
}

.action-btn.publish:hover:not(:disabled) {
  background: var(--color-primary-strong);
  box-shadow: 0 4px 12px rgba(47, 107, 255, 0.35);
  transform: translateY(-1px);
}

.action-btn:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none !important;
}

.loading-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 100px 20px;
  gap: 24px;
}

.loading-spinner {
  position: relative;
  width: 48px;
  height: 48px;
}

.spinner-ring {
  position: absolute;
  width: 100%;
  height: 100%;
  border: 3px solid transparent;
  border-top-color: var(--color-primary);
  border-radius: 50%;
  animation: spin 1.2s linear infinite;
}

.spinner-ring:nth-child(1) { animation-delay: 0s; }
.spinner-ring:nth-child(2) { animation-delay: 0.2s; width: 80%; height: 80%; top: 10%; left: 10%; }
.spinner-ring:nth-child(3) { animation-delay: 0.4s; width: 60%; height: 60%; top: 20%; left: 20%; }

@keyframes spin {
  to { transform: rotate(360deg); }
}

.loading-text {
  font-size: 0.9375rem;
  color: var(--color-text-muted);
}

.editor-content {
  display: grid;
  grid-template-columns: 1fr 300px;
  gap: var(--space-page);
  align-items: start;
}

.editor-main-area {
  display: flex;
  flex-direction: column;
  gap: var(--space-gap);
}

.title-section {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  padding: 16px 20px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  box-shadow: var(--shadow-lg);
}

.title-input {
  width: 100%;
  font-size: 1.5rem;
  font-weight: 700;
  font-family: var(--font-heading);
  border: none;
  background: transparent;
  color: var(--color-text);
  letter-spacing: -0.01em;
}

.title-input:focus {
  outline: none;
}

.title-input::placeholder {
  color: var(--color-text-muted);
  font-weight: 400;
}

.title-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.char-counter {
  font-size: 0.75rem;
  color: var(--color-text-muted);
  transition: color 0.2s;
}

.char-counter.warning {
  color: var(--color-warning);
}

.editor-mode-toggle {
  display: flex;
  gap: 4px;
  padding: 4px;
  background: var(--color-surface-alt);
  border-radius: 10px;
  position: relative;
}

.mode-indicator {
  position: absolute;
  top: 4px;
  bottom: 4px;
  background: var(--color-primary);
  border-radius: 8px;
  transition: left var(--transition-base), width var(--transition-base);
  z-index: 0;
  pointer-events: none;
}

.mode-btn {
  position: relative;
  z-index: 1;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 6px 14px;
  border: none;
  background: transparent;
  color: var(--color-text-muted);
  font-size: 0.8125rem;
  font-weight: 500;
  border-radius: 8px;
  cursor: pointer;
  transition: color var(--transition-fast);
}

.mode-btn:hover {
  color: var(--color-text);
}

.mode-btn.active {
  color: white;
}

.editor-container {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  overflow: hidden;
  box-shadow: var(--shadow-lg);
}

.settings-panel {
  display: flex;
  flex-direction: column;
  gap: var(--space-gap);
  position: sticky;
  top: var(--space-page);
}

.panel-card {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  overflow: hidden;
  box-shadow: var(--shadow-lg);
}

.panel-header {
  padding: 14px 16px;
  border-bottom: 1px solid var(--color-border);
  background: var(--color-surface-alt);
}

.panel-title {
  display: flex;
  align-items: center;
  gap: 8px;
  font-size: 0.875rem;
  font-weight: 600;
  color: var(--color-text);
}

.panel-body {
  padding: 16px;
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.setting-group {
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.setting-label {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 0.8125rem;
  font-weight: 500;
  color: var(--color-text-muted);
}

.status-selector {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.status-option {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 10px 12px;
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: transparent;
  cursor: pointer;
  transition: all 0.2s ease;
}

.status-option:hover {
  background: var(--color-surface-alt);
  border-color: var(--color-primary);
}

.status-option.active {
  background: var(--color-primary-soft);
  border-color: var(--color-primary);
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.status-text {
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-text);
}

.setting-input {
  width: 100%;
  padding: 10px 12px;
  font-size: 0.875rem;
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: var(--color-surface);
  color: var(--color-text);
  transition: all 0.2s ease;
}

.setting-input:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px var(--color-primary-soft);
}

.cover-preview {
  margin-top: 8px;
  border-radius: 10px;
  overflow: hidden;
  border: 1px solid var(--color-border);
}

.cover-preview img {
  width: 100%;
  height: 120px;
  object-fit: cover;
}

.setting-textarea {
  width: 100%;
  padding: 10px 12px;
  font-size: 0.875rem;
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: var(--color-surface);
  color: var(--color-text);
  resize: vertical;
  min-height: 80px;
  transition: all 0.2s ease;
}

.setting-textarea:focus {
  outline: none;
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px var(--color-primary-soft);
}

.tags-input-wrapper {
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: var(--color-surface);
  padding: 8px;
  transition: all 0.2s ease;
}

.tags-input-wrapper:focus-within {
  border-color: var(--color-primary);
  box-shadow: 0 0 0 3px var(--color-primary-soft);
}

.tags-container {
  display: flex;
  flex-wrap: wrap;
  gap: 6px;
  margin-bottom: 6px;
}

.tag-badge {
  display: inline-flex;
  align-items: center;
  gap: 4px;
  padding: 4px 8px 4px 10px;
  background: var(--color-primary-soft);
  color: var(--color-primary-strong);
  border-radius: var(--radius-full);
  font-size: 0.75rem;
  font-weight: 500;
}

.tag-remove-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 16px;
  height: 16px;
  border: none;
  background: transparent;
  color: var(--color-primary-strong);
  cursor: pointer;
  border-radius: 50%;
  transition: all 0.15s ease;
}

.tag-remove-btn:hover {
  background: var(--color-primary);
  color: white;
}

.tag-input-field {
  width: 100%;
  padding: 4px 6px;
  border: none;
  background: transparent;
  font-size: 0.875rem;
  color: var(--color-text);
}

.tag-input-field:focus {
  outline: none;
}

.tips-card .panel-body {
  padding: 12px 16px;
}

.tips-list {
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 8px;
}

.tips-list li {
  font-size: 0.8125rem;
  color: var(--color-text-muted);
  line-height: 1.5;
}

.tips-list code {
  padding: 2px 6px;
  background: var(--color-surface-alt);
  border-radius: 4px;
  font-family: monospace;
  font-size: 0.75rem;
  color: var(--color-primary-strong);
}

.toast-notification {
  position: fixed;
  bottom: 24px;
  right: 24px;
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 20px;
  border-radius: var(--radius-lg);
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-xl);
  font-size: 0.875rem;
  font-weight: 500;
  color: var(--color-text);
  z-index: 1000;
}

.toast-notification.success {
  border-color: #10B981;
  color: #065F46;
  background: #F0FDF4;
}

.toast-notification.error {
  border-color: #EF4444;
  color: #991B1B;
  background: #FEF2F2;
}

.toast-notification.info {
  border-color: var(--color-primary);
  color: var(--color-primary-strong);
  background: var(--color-primary-soft);
}

.toast-enter-active,
.toast-leave-active {
  transition: opacity var(--transition-slow), transform var(--transition-slow);
}

.toast-enter-from,
.toast-leave-to {
  opacity: 0;
  transform: translateY(24px) scale(0.95);
}

@media (max-width: 1024px) {
  .editor-content {
    grid-template-columns: 1fr;
  }

  .settings-panel {
    position: static;
    order: -1;
  }

  .editor-header {
    flex-direction: column;
    align-items: flex-start;
    gap: var(--space-gap);
  }

  .header-actions {
    width: 100%;
    flex-wrap: wrap;
  }

  .action-btn {
    flex: 1;
    justify-content: center;
  }
}

@media (max-width: 768px) {
  .editor-view {
    padding: var(--space-gap);
  }

  .editor-header {
    padding-bottom: var(--space-gap);
  }

  .header-title {
    font-size: 1.25rem;
  }

  .title-input {
    font-size: 1.25rem;
  }

  .title-footer {
    flex-direction: column;
    align-items: flex-start;
    gap: 8px;
  }

  .mode-btn {
    padding: 6px 10px;
    font-size: 0.75rem;
  }

  .mode-btn svg {
    display: none;
  }
}

@media (max-width: 640px) {
  .editor-view {
    padding: var(--space-4);
  }
}
</style>
