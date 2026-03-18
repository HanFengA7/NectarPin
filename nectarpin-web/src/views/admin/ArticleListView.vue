<script setup lang="ts">
import { ref, computed, onMounted, nextTick } from 'vue'
import { useRouter } from 'vue-router'
import { articleApi, type Article } from '@/api/article'

const router = useRouter()

const articles = ref<Article[]>([])
const isLoading = ref(false)
const errorMessage = ref('')
const currentPage = ref(1)
const pageSize = ref(12)
const total = ref(0)

const statusFilter = ref<number | undefined>(undefined)
const filterTabsRef = ref<HTMLElement | null>(null)
const indicatorStyle = ref({ left: '0px', width: '0px' })

const statusOptions: { value: number | undefined; label: string }[] = [
  { value: undefined, label: '全部文章' },
  { value: 0, label: '草稿箱' },
  { value: 1, label: '已发布' },
  { value: 2, label: '已下架' },
]

const statusMap: Record<number, { label: string; class: string; dot: string }> = {
  0: { label: '草稿', class: 'status-draft', dot: '#F59E0B' },
  1: { label: '已发布', class: 'status-published', dot: '#10B981' },
  2: { label: '已下架', class: 'status-archived', dot: '#6B7280' },
}

const fetchArticles = async () => {
  isLoading.value = true
  errorMessage.value = ''
  try {
    const response = await articleApi.listForAdmin({
      page: currentPage.value,
      page_size: pageSize.value,
      status: statusFilter.value,
    })
    articles.value = response.data.items || []
    total.value = response.data.total || 0
  } catch (error: any) {
    errorMessage.value = error?.response?.data?.message || '加载文章列表失败'
  } finally {
    isLoading.value = false
  }
}

const handlePageChange = (page: number) => {
  currentPage.value = page
  fetchArticles()
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const handleFilterChange = () => {
  currentPage.value = 1
  fetchArticles()
}

const handleCreateArticle = () => {
  router.push('/admin/articles/add')
}

const handleEditArticle = (id: number) => {
  router.push({ path: '/admin/articles/edit', query: { id: id.toString() } })
}

const handleDeleteArticle = async (id: number) => {
  if (!confirm('确定要删除这篇文章吗？此操作不可恢复。')) return
  try {
    await articleApi.delete(id)
    fetchArticles()
  } catch (error: any) {
    alert(error?.response?.data?.message || '删除失败')
  }
}

const formatDate = (dateStr?: string) => {
  if (!dateStr) return '-'
  const date = new Date(dateStr)
  const now = new Date()
  const diff = now.getTime() - date.getTime()
  const days = Math.floor(diff / (1000 * 60 * 60 * 24))
  if (days === 0) return '今天'
  if (days === 1) return '昨天'
  if (days < 7) return `${days}天前`
  return date.toLocaleDateString('zh-CN', { month: 'short', day: 'numeric' })
}

const totalPages = computed(() => Math.ceil(total.value / pageSize.value))

const visiblePages = computed(() => {
  const pages: (number | string)[] = []
  const total = totalPages.value
  const current = currentPage.value
  if (total <= 5) {
    for (let i = 1; i <= total; i++) pages.push(i)
  } else {
    if (current <= 3) {
      pages.push(1, 2, 3, 4, total)
    } else if (current >= total - 2) {
      pages.push(1, total - 3, total - 2, total - 1, total)
    } else {
      pages.push(1, current - 1, current, current + 1, total)
    }
  }
  return pages
})

const handleTabClick = async (value: number | undefined) => {
  statusFilter.value = value
  await nextTick()
  updateIndicator()
  handleFilterChange()
}

const updateIndicator = () => {
  if (!filterTabsRef.value) return

  const activeTab = filterTabsRef.value.querySelector('.filter-tab.active') as HTMLElement
  if (!activeTab) return

  const tabsContainer = filterTabsRef.value
  const containerRect = tabsContainer.getBoundingClientRect()
  const tabRect = activeTab.getBoundingClientRect()

  indicatorStyle.value = {
    left: `${tabRect.left - containerRect.left}px`,
    width: `${tabRect.width}px`,
  }
}

onMounted(() => {
  fetchArticles()
  nextTick(() => {
    updateIndicator()
  })
})

</script>

<template>
  <div class="list-view">
    <main class="list-main">
      <header class="list-header">
        <div class="header-content">
          <div class="header-text">
            <h1 class="header-title">文章管理</h1>
            <p class="header-subtitle">共 {{ total }} 篇文章</p>
          </div>
          <button class="btn btn-primary" @click="handleCreateArticle">
            <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M12 5v14M5 12h14"/>
            </svg>
            写文章
          </button>
        </div>

        <div class="filter-tabs" ref="filterTabsRef">
          <div class="filter-indicator" :style="indicatorStyle"></div>
          <button
            v-for="opt in statusOptions"
            :key="opt.value"
            class="filter-tab"
            :class="{ active: statusFilter === opt.value }"
            @click="handleTabClick(opt.value)"
          >
            {{ opt.label }}
          </button>
        </div>
      </header>

      <div v-if="errorMessage" class="alert alert-error">
        <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
          <circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/>
        </svg>
        {{ errorMessage }}
      </div>

      <div v-if="isLoading" class="loading-state">
        <div class="skeleton-grid">
          <div v-for="i in 6" :key="i" class="skeleton-card">
            <div class="skeleton skeleton-cover"></div>
            <div class="skeleton skeleton-title"></div>
            <div class="skeleton skeleton-text"></div>
            <div class="skeleton-footer">
              <div class="skeleton skeleton-badge"></div>
              <div class="skeleton skeleton-date"></div>
            </div>
          </div>
        </div>
      </div>

      <div v-else-if="articles.length === 0" class="empty-state">
        <div class="empty-illustration">
          <svg xmlns="http://www.w3.org/2000/svg" width="120" height="120" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="0.75">
            <path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z" stroke="#D7E4FF"/>
            <polyline points="14 2 14 8 20 8" stroke="#D7E4FF"/>
            <line x1="16" y1="13" x2="8" y2="13" stroke="#B8CCFF"/>
            <line x1="16" y1="17" x2="8" y2="17" stroke="#B8CCFF"/>
            <polyline points="10 9 9 9 8 9" stroke="#B8CCFF"/>
          </svg>
        </div>
        <h3 class="empty-title">还没有文章</h3>
        <p class="empty-desc">开始创作你的第一篇文章吧</p>
        <button class="btn btn-primary" @click="handleCreateArticle">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <path d="M12 5v14M5 12h14"/>
          </svg>
          写文章
        </button>
      </div>

      <div v-else class="article-grid-wrapper">
        <TransitionGroup name="list" tag="div" class="article-grid">
        <article
          v-for="article in articles"
          :key="article.id"
          class="article-card"
        >
          <div class="card-cover" :style="article.cover_image ? { backgroundImage: `url(${article.cover_image})` } : {}">
            <div v-if="!article.cover_image" class="cover-placeholder">
              <svg xmlns="http://www.w3.org/2000/svg" width="32" height="32" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5">
                <rect x="3" y="3" width="18" height="18" rx="2" ry="2"/>
                <circle cx="8.5" cy="8.5" r="1.5"/>
                <polyline points="21 15 16 10 5 21"/>
              </svg>
            </div>
            <span class="card-status" :class="statusMap[article.status || 0]?.class">
              <span class="status-dot" :style="{ background: statusMap[article.status || 0]?.dot }"></span>
              {{ statusMap[article.status || 0]?.label }}
            </span>
          </div>
          <div class="card-body">
            <h3 class="card-title">{{ article.title }}</h3>
            <p v-if="article.summary" class="card-summary">{{ article.summary }}</p>
            <div class="card-meta">
              <span class="meta-date">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <rect x="3" y="4" width="18" height="18" rx="2" ry="2"/>
                  <line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/>
                  <line x1="3" y1="10" x2="21" y2="10"/>
                </svg>
                {{ formatDate(article.updated_at) }}
              </span>
              <span v-if="article.view_count" class="meta-views">
                <svg xmlns="http://www.w3.org/2000/svg" width="14" height="14" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                  <path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/>
                  <circle cx="12" cy="12" r="3"/>
                </svg>
                {{ article.view_count }}
              </span>
            </div>
          </div>
          <div class="card-actions">
            <button class="action-btn edit" @click="handleEditArticle(article.id!)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/>
                <path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1 1-4 9.5-9.5z"/>
              </svg>
              编辑
            </button>
            <button class="action-btn delete" @click="handleDeleteArticle(article.id!)">
              <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <polyline points="3 6 5 6 21 6"/>
                <path d="M19 6v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6m3 0V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/>
              </svg>
              删除
            </button>
          </div>
        </article>
        </TransitionGroup>
      </div>

      <div v-if="totalPages > 1" class="pagination">
        <button
          class="page-btn nav"
          :disabled="currentPage === 1"
          @click="handlePageChange(currentPage - 1)"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="15 18 9 12 15 6"/>
          </svg>
        </button>
        <template v-for="page in visiblePages" :key="page">
          <span v-if="page === '...'" class="page-ellipsis">···</span>
          <button
            v-else
            class="page-btn"
            :class="{ active: page === currentPage }"
            @click="handlePageChange(page as number)"
          >
            {{ page }}
          </button>
        </template>
        <button
          class="page-btn nav"
          :disabled="currentPage === totalPages"
          @click="handlePageChange(currentPage + 1)"
        >
          <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <polyline points="9 18 15 12 9 6"/>
          </svg>
        </button>
      </div>
    </main>
  </div>
</template>

<style scoped>
.list-view {
  padding: var(--space-page);
}

.list-main {
  width: 100%;
}

.list-header {
  margin-bottom: var(--space-section);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--space-gap);
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.header-title {
  font-family: var(--font-heading);
  font-size: 1.75rem;
  font-weight: 700;
  color: var(--color-text);
  letter-spacing: -0.02em;
}

.header-subtitle {
  font-size: 0.9375rem;
  color: var(--color-text-muted);
}

.btn {
  display: inline-flex;
  align-items: center;
  gap: 8px;
  padding: 10px 20px;
  border-radius: var(--radius-lg);
  font-size: 0.9375rem;
  font-weight: 500;
  border: none;
  transition: all 0.2s ease;
  cursor: pointer;
}

.btn-primary {
  background: var(--color-primary);
  color: white;
  box-shadow: 0 2px 8px rgba(47, 107, 255, 0.25);
}

.btn-primary:hover {
  background: var(--color-primary-strong);
  box-shadow: 0 4px 12px rgba(47, 107, 255, 0.35);
  transform: translateY(-1px);
}

.btn-primary:active {
  transform: translateY(0);
}

.filter-tabs {
  display: flex;
  gap: 4px;
  padding: 4px;
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  width: fit-content;
  position: relative;
}

.filter-indicator {
  position: absolute;
  top: 4px;
  bottom: 4px;
  background: var(--color-primary);
  border-radius: calc(var(--radius-lg) - 2px);
  transition: left var(--transition-base), width var(--transition-base);
  z-index: 0;
  pointer-events: none;
}

.filter-tab {
  position: relative;
  z-index: 1;
  padding: 8px 16px;
  border: none;
  background: transparent;
  color: var(--color-text-muted);
  font-size: 0.875rem;
  font-weight: 500;
  border-radius: calc(var(--radius-lg) - 2px);
  cursor: pointer;
  transition: color var(--transition-fast);
}

.filter-tab:hover {
  color: var(--color-text);
}

.filter-tab.active {
  color: white;
}

.alert {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 14px 18px;
  border-radius: var(--radius-lg);
  margin-bottom: var(--space-gap);
  font-size: 0.875rem;
}

.alert-error {
  background: #FEF2F2;
  color: var(--color-danger);
  border: 1px solid #FECACA;
}

.loading-state {
  padding: var(--space-gap) 0;
}

.skeleton-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: var(--space-gap);
}

.skeleton-card {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  overflow: hidden;
  padding: 0;
}

.skeleton {
  background: linear-gradient(90deg, var(--color-surface-alt) 25%, var(--color-border) 50%, var(--color-surface-alt) 75%);
  background-size: 200% 100%;
  animation: shimmer 1.2s infinite;
  border-radius: 4px;
}

.skeleton-cover {
  height: 140px;
  border-radius: 0;
}

.skeleton-title {
  height: 20px;
  width: 70%;
  margin: 16px;
}

.skeleton-text {
  height: 14px;
  width: 50%;
  margin: 0 16px 16px;
}

.skeleton-footer {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  border-top: 1px solid var(--color-border);
}

.skeleton-badge {
  height: 24px;
  width: 60px;
  border-radius: var(--radius-full);
}

.skeleton-date {
  height: 14px;
  width: 50px;
}

@keyframes shimmer {
  0% { background-position: 200% 0; }
  100% { background-position: -200% 0; }
}

.empty-state {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding: 80px 20px;
  text-align: center;
}

.empty-illustration {
  margin-bottom: 24px;
  opacity: 0.8;
}

.empty-title {
  font-family: var(--font-heading);
  font-size: 1.25rem;
  font-weight: 600;
  color: var(--color-text);
  margin-bottom: 8px;
}

.empty-desc {
  font-size: 0.9375rem;
  color: var(--color-text-muted);
  margin-bottom: 24px;
}

.article-grid-wrapper {
  position: relative;
}

.article-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(280px, 1fr));
  gap: var(--space-gap);
}

.article-card {
  background: var(--color-surface);
  border-radius: var(--radius-lg);
  border: 1px solid var(--color-border);
  overflow: hidden;
  transition: all 0.25s ease;
  display: flex;
  flex-direction: column;
}

.article-card:hover {
  border-color: var(--color-primary);
  box-shadow: var(--shadow-lg);
  transform: translateY(-2px);
}

.card-cover {
  position: relative;
  height: 140px;
  background: var(--color-surface-alt);
  background-size: cover;
  background-position: center;
  display: flex;
  align-items: center;
  justify-content: center;
}

.cover-placeholder {
  color: var(--color-text-muted);
  opacity: 0.4;
}

.card-status {
  position: absolute;
  top: 12px;
  right: 12px;
  display: inline-flex;
  align-items: center;
  gap: 6px;
  padding: 4px 10px;
  border-radius: var(--radius-full);
  font-size: 0.75rem;
  font-weight: 500;
  background: rgba(255, 255, 255, 0.9);
  backdrop-filter: blur(8px);
  color: var(--color-text);
}

.status-dot {
  width: 6px;
  height: 6px;
  border-radius: 50%;
}

.card-body {
  padding: 16px;
  flex: 1;
  display: flex;
  flex-direction: column;
}

.card-title {
  font-family: var(--font-heading);
  font-size: 1rem;
  font-weight: 600;
  color: var(--color-text);
  margin-bottom: 8px;
  line-height: 1.4;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-summary {
  font-size: 0.8125rem;
  color: var(--color-text-muted);
  line-height: 1.5;
  margin-bottom: 12px;
  flex: 1;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.card-meta {
  display: flex;
  align-items: center;
  gap: 16px;
  font-size: 0.75rem;
  color: var(--color-text-muted);
}

.meta-date,
.meta-views {
  display: flex;
  align-items: center;
  gap: 4px;
}

.card-actions {
  display: flex;
  gap: 8px;
  padding: 12px 16px;
  border-top: 1px solid var(--color-border);
  background: var(--color-surface-alt);
}

.action-btn {
  flex: 1;
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: 6px;
  padding: 8px 12px;
  border: none;
  border-radius: 10px;
  font-size: 0.8125rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.action-btn.edit {
  background: var(--color-primary-soft);
  color: var(--color-primary-strong);
}

.action-btn.edit:hover {
  background: var(--color-primary);
  color: white;
}

.action-btn.delete {
  background: #FEF2F2;
  color: var(--color-danger);
}

.action-btn.delete:hover {
  background: var(--color-danger);
  color: white;
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 4px;
  margin-top: var(--space-page);
}

.page-btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 36px;
  height: 36px;
  padding: 0 12px;
  border: 1px solid var(--color-border);
  border-radius: 10px;
  background: var(--color-surface);
  color: var(--color-text);
  font-size: 0.875rem;
  font-weight: 500;
  cursor: pointer;
  transition: all 0.2s ease;
}

.page-btn:hover:not(:disabled) {
  border-color: var(--color-primary);
  color: var(--color-primary);
  background: var(--color-primary-soft);
}

.page-btn.active {
  background: var(--color-primary);
  border-color: var(--color-primary);
  color: white;
}

.page-btn:disabled {
  opacity: 0.4;
  cursor: not-allowed;
}

.page-btn.nav {
  padding: 0 8px;
}

.page-ellipsis {
  color: var(--color-text-muted);
  padding: 0 4px;
}

@media (max-width: 1024px) {
  .list-main {
    margin-left: 0;
    max-width: 100vw;
    padding: var(--space-gap);
  }

  .header-content {
    flex-direction: column;
    gap: var(--space-gap);
  }

  .header-actions {
    width: 100%;
  }

  .filter-tabs {
    width: 100%;
    overflow-x: auto;
  }
}

@media (max-width: 640px) {
  .header-title {
    font-size: 1.5rem;
  }

  .article-grid {
    grid-template-columns: 1fr;
  }

  .filter-tabs {
    flex-wrap: nowrap;
    overflow-x: auto;
    -webkit-overflow-scrolling: touch;
  }

  .filter-tab {
    white-space: nowrap;
  }
}
</style>
