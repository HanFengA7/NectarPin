<script setup lang="ts">
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'
import StatCard from '@/components/admin/StatCard.vue'
import TrendChart from '@/components/admin/TrendChart.vue'
import TodoCard from '@/components/admin/TodoCard.vue'
import HeatmapChart from '@/components/admin/HeatmapChart.vue'

const router = useRouter()

const currentTime = computed(() => {
  const now = new Date()
  const hour = now.getHours()
  if (hour < 6) return '凌晨好'
  if (hour < 12) return '早上好'
  if (hour < 14) return '中午好'
  if (hour < 18) return '下午好'
  return '晚上好'
})

const currentDate = computed(() => {
  const now = new Date()
  const options: Intl.DateTimeFormatOptions = {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    weekday: 'long'
  }
  return now.toLocaleDateString('zh-CN', options)
})

const stats = ref([
  {
    title: '用户总数',
    value: '12,847',
    delta: '+12.5%',
    deltaType: 'success' as const,
    iconBgColor: 'rgba(59, 130, 246, 0.1)',
    iconColor: '#3B82F6',
    iconName: 'users',
  },
  {
    title: '总收入',
    value: '$84,230',
    delta: '+8.2%',
    deltaType: 'success' as const,
    iconBgColor: 'rgba(16, 185, 129, 0.1)',
    iconColor: '#10B981',
    iconName: 'dollar-sign',
  },
  {
    title: '订单数',
    value: '3,421',
    delta: '+2.6%',
    deltaType: 'success' as const,
    iconBgColor: 'rgba(245, 158, 11, 0.1)',
    iconColor: '#F59E0B',
    iconName: 'shopping-cart',
  },
  {
    title: '转化率',
    value: '4.28%',
    delta: '+0.8%',
    deltaType: 'success' as const,
    iconBgColor: 'rgba(139, 92, 246, 0.1)',
    iconColor: '#8B5CF6',
    iconName: 'trending-up',
  },
])

const trendData = ref([
  { value: 120 },
  { value: 156 },
  { value: 210, highlight: true },
  { value: 172 },
  { value: 228, highlight: true },
  { value: 194 },
])

const todoItems = ref([
  { text: '审核 12 篇投稿并安排发布时间', dotColor: '#3B82F6' },
  { text: '处理 28 条待回复评论与举报内容', dotColor: '#F59E0B' },
  { text: '更新系统安全补丁', dotColor: '#10B981' },
  { text: '分析上周流量数据', dotColor: '#8B5CF6' },
])

const quickLinks = ref([
  {
    title: '新增订单',
    icon: 'plus-circle',
    color: '#3B82F6',
    path: '/admin/orders/add',
  },
  {
    title: '添加用户',
    icon: 'user-plus',
    color: '#10B981',
    path: '/admin/users/add',
  },
  {
    title: '生成报告',
    icon: 'file-text',
    color: '#F59E0B',
    path: '/admin/reports',
  },
  {
    title: '系统设置',
    icon: 'settings',
    color: '#8B5CF6',
    path: '/admin/settings',
  },
])

const systemStatus = ref([
  {
    name: 'API 服务',
    status: 'online',
    color: '#10B981',
  },
  {
    name: '数据库',
    status: 'online',
    color: '#10B981',
  },
  {
    name: '缓存服务',
    status: 'online',
    color: '#10B981',
  },
  {
    name: '存储服务',
    status: 'online',
    color: '#10B981',
  },
])

const handleExport = () => {
  console.log('Export report')
}

const handleCreateArticle = () => {
  router.push('/admin/articles/add')
}

const navigateTo = (path: string) => {
  router.push(path)
}

const icons: Record<string, string> = {
  'users': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87"/><path d="M16 3.13a4 4 0 0 1 0 7.75"/></svg>`,
  'dollar-sign': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><line x1="12" y1="2" x2="12" y2="22"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/></svg>`,
  'shopping-cart': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="21" r="1"/><circle cx="20" cy="21" r="1"/><path d="M1 1h4l2.68 13.39a2 2 0 0 0 2 1.61h9.72a2 2 0 0 0 2-1.61L23 6H6"/></svg>`,
  'trending-up': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><polyline points="22 7 13.5 15.5 8.5 10.5 2 17"/><polyline points="16 7 22 7 22 13"/></svg>`,
  'plus-circle': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="16"/><line x1="8" y1="12" x2="16" y2="12"/></svg>`,
  'user-plus': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><line x1="19" y1="8" x2="19" y2="14"/><line x1="22" y1="11" x2="16" y2="11"/></svg>`,
  'file-text': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>`,
  'settings': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>`,
}

const getIcon = (name: string): string => {
  return icons[name] || ''
}
</script>

<template>
  <div class="dashboard-view">
    <main class="dashboard-main">
      <header class="dashboard-header">
        <div class="header-content">
          <div class="header-text">
            <h1 class="header-title">{{ currentTime }}，管理员</h1>
            <p class="header-subtitle">{{ currentDate }}</p>
          </div>
          <div class="header-actions">
            <button class="btn btn-secondary" @click="handleExport">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4"/>
                <polyline points="7 10 12 15 17 10"/>
                <line x1="12" y1="15" x2="12" y2="3"/>
              </svg>
              导出报告
            </button>
            <button class="btn btn-primary" @click="handleCreateArticle">
              <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                <line x1="12" y1="5" x2="12" y2="19"/>
                <line x1="5" y1="12" x2="19" y2="12"/>
              </svg>
              新建文章
            </button>
          </div>
        </div>
      </header>

      <section class="dashboard-stats">
        <StatCard
          v-for="stat in stats"
          :key="stat.title"
          :title="stat.title"
          :value="stat.value"
          :delta="stat.delta"
          :delta-type="stat.deltaType"
          :icon-bg-color="stat.iconBgColor"
          :icon-color="stat.iconColor"
          :icon-name="stat.iconName"
        />
      </section>

      <section class="dashboard-overview">
        <div class="overview-main">
          <div class="revenue-card">
            <div class="revenue-header">
              <div>
                <h3 class="revenue-title">本月收入</h3>
                <p class="revenue-subtitle">实时监控您的业务表现</p>
              </div>
              <span class="revenue-badge">+12.5%</span>
            </div>
            <div class="revenue-content">
              <div class="revenue-amount">$84,230</div>
              <div class="revenue-stats">
                <div class="revenue-stat">
                  <span class="stat-label">本月收入</span>
                  <span class="stat-value">$118.5k</span>
                </div>
                <div class="revenue-stat">
                  <span class="stat-label">本月订单</span>
                  <span class="stat-value">856</span>
                </div>
                <div class="revenue-stat">
                  <span class="stat-label">新用户</span>
                  <span class="stat-value">128</span>
                </div>
              </div>
              <button class="revenue-btn">查看详情 →</button>
            </div>
          </div>

          <div class="quick-access">
            <h3 class="section-title">快速入口</h3>
            <div class="quick-links">
              <button
                v-for="link in quickLinks"
                :key="link.title"
                class="quick-link-item"
                :style="{ borderColor: link.color }"
                @click="navigateTo(link.path)"
              >
                <span class="quick-link-icon" :style="{ color: link.color }" v-html="getIcon(link.icon)"></span>
                <span class="quick-link-title">{{ link.title }}</span>
              </button>
            </div>
          </div>

          <div class="dashboard-trend">
            <TrendChart
              title="内容流量趋势"
              subtitle="过去 6 周文章访问量与互动变化"
              badge="实时同步"
              :data="trendData"
            />
          </div>
        </div>

        <div class="overview-side">
          <div class="system-status">
            <h3 class="section-title">系统状态</h3>
            <div class="status-items">
              <div v-for="item in systemStatus" :key="item.name" class="status-item">
                <span class="status-name">{{ item.name }}</span>
                <div class="status-indicator" :style="{ background: item.color }"></div>
                <span class="status-text" :style="{ color: item.color }">正常运行</span>
              </div>
            </div>
          </div>

          <TodoCard title="今日待办" :items="todoItems" />

          <div class="recent-activity">
            <h3 class="section-title">最近活动</h3>
            <div class="activity-items">
              <div class="activity-item">
                <div class="activity-icon" style="background: rgba(59, 130, 246, 0.1); color: #3B82F6;">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="10"/><line x1="12" y1="8" x2="12" y2="12"/><line x1="12" y1="16" x2="12.01" y2="16"/></svg>
                </div>
                <div class="activity-content">
                  <p class="activity-text">新用户注册</p>
                  <p class="activity-time">2分钟前</p>
                </div>
              </div>
              <div class="activity-item">
                <div class="activity-icon" style="background: rgba(16, 185, 129, 0.1); color: #10B981;">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><polyline points="22 4 12 14.01 9 11.01"/></svg>
                </div>
                <div class="activity-content">
                  <p class="activity-text">订单完成</p>
                  <p class="activity-time">15分钟前</p>
                </div>
              </div>
              <div class="activity-item">
                <div class="activity-icon" style="background: rgba(245, 158, 11, 0.1); color: #F59E0B;">
                  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>
                </div>
                <div class="activity-content">
                  <p class="activity-text">文章发布</p>
                  <p class="activity-time">1小时前</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </section>
    </main>
  </div>
</template>

<style scoped>
.dashboard-view {
  padding: var(--space-page);
  background: var(--color-bg);
  min-height: 100vh;
}

.dashboard-main {
  width: 100%;
  max-width: var(--content-max-width);
  margin: 0 auto;
}

.dashboard-header {
  margin-bottom: var(--space-section);
}

.header-content {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  gap: var(--space-6);
}

.header-text {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.header-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-4xl);
  font-weight: var(--font-weight-extrabold);
  letter-spacing: var(--letter-spacing-tight);
  color: var(--color-text-primary);
  line-height: var(--line-height-tight);
}

.header-subtitle {
  font-size: var(--font-size-base);
  color: var(--color-text-muted);
  margin-top: var(--space-1);
}

.header-actions {
  display: flex;
  gap: var(--space-3);
  flex-shrink: 0;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  gap: var(--space-2);
  padding: var(--space-3) var(--space-5);
  font-size: var(--font-size-base);
  font-weight: var(--font-weight-semibold);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--transition-fast);
  border: none;
  white-space: nowrap;
}

.btn-primary {
  background: var(--color-primary);
  color: #ffffff;
  box-shadow: var(--shadow-sm);
}

.btn-primary:hover {
  background: var(--color-primary-hover);
  box-shadow: var(--shadow-md);
  transform: translateY(-1px);
}

.btn-primary:active {
  transform: translateY(0);
}

.btn-secondary {
  background: var(--color-surface);
  color: var(--color-text-primary);
  border: 1px solid var(--color-border);
}

.btn-secondary:hover {
  background: var(--color-surface-alt);
  border-color: var(--color-border-hover);
  color: var(--color-text-primary);
}

.dashboard-stats {
  display: flex;
  gap: var(--space-5);
  margin-bottom: var(--space-6);
}

.dashboard-stats > * {
  flex: 1;
  min-width: 0;
}

.dashboard-overview {
  display: flex;
  gap: var(--space-5);
  margin-bottom: var(--space-6);
}

.overview-main {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: var(--space-5);
}

.overview-side {
  width: 360px;
  flex-shrink: 0;
  display: flex;
  flex-direction: column;
  gap: var(--space-5);
}

.revenue-card {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  padding: var(--space-6);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--color-border);
}

.revenue-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: var(--space-4);
}

.revenue-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-xl);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
  margin-bottom: var(--space-1);
}

.revenue-subtitle {
  font-size: var(--font-size-base);
  color: var(--color-text-muted);
}

.revenue-badge {
  padding: var(--space-1) var(--space-3);
  background: rgba(16, 185, 129, 0.1);
  color: #10B981;
  border-radius: var(--radius-full);
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
}

.revenue-content {
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
}

.revenue-amount {
  font-family: var(--font-heading);
  font-size: var(--font-size-4xl);
  font-weight: var(--font-weight-extrabold);
  color: var(--color-text-primary);
  line-height: var(--line-height-tight);
}

.revenue-stats {
  display: flex;
  gap: var(--space-6);
  margin: var(--space-4) 0;
}

.revenue-stat {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.stat-label {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
}

.stat-value {
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-primary);
}

.revenue-btn {
  align-self: flex-start;
  padding: var(--space-2) var(--space-4);
  background: transparent;
  color: var(--color-primary);
  border: 1px solid var(--color-primary);
  border-radius: var(--radius-lg);
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.revenue-btn:hover {
  background: var(--color-primary-soft);
  transform: translateY(-1px);
}

.quick-access {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  padding: var(--space-6);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--color-border);
}

.section-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
  margin-bottom: var(--space-4);
}

.quick-links {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: var(--space-3);
}

.quick-link-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: var(--space-2);
  padding: var(--space-4);
  background: var(--color-surface-alt);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--transition-fast);
  text-align: center;
}

.quick-link-item:hover {
  transform: translateY(-2px);
  box-shadow: var(--shadow-md);
}

.quick-link-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40px;
  height: 40px;
  background: var(--color-surface);
  border-radius: var(--radius-full);
}

.quick-link-title {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-primary);
}

.dashboard-trend {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  padding: var(--space-6);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--color-border);
}

.system-status {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  padding: var(--space-6);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--color-border);
}

.status-items {
  display: flex;
  flex-direction: column;
  gap: var(--space-3);
}

.status-item {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: var(--space-3);
  background: var(--color-surface-alt);
  border-radius: var(--radius-lg);
}

.status-name {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-medium);
  color: var(--color-text-primary);
}

.status-indicator {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.status-text {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-medium);
}

.recent-activity {
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  padding: var(--space-6);
  box-shadow: var(--shadow-lg);
  border: 1px solid var(--color-border);
}

.activity-items {
  display: flex;
  flex-direction: column;
  gap: var(--space-3);
}

.activity-item {
  display: flex;
  align-items: flex-start;
  gap: var(--space-3);
  padding: var(--space-3);
  background: var(--color-surface-alt);
  border-radius: var(--radius-lg);
}

.activity-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: var(--radius-full);
  flex-shrink: 0;
}

.activity-content {
  flex: 1;
  min-width: 0;
}

.activity-text {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-medium);
  color: var(--color-text-primary);
  margin-bottom: var(--space-1);
}

.activity-time {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
}

@media (max-width: 1400px) {
  .dashboard-overview {
    gap: var(--space-4);
  }

  .overview-side {
    width: 320px;
  }

  .revenue-stats {
    gap: var(--space-4);
  }
}

@media (max-width: 1200px) {
  .dashboard-main {
    margin-left: calc(var(--sidebar-collapsed) + 32px);
    padding: var(--space-6);
  }

  .dashboard-stats {
    flex-wrap: wrap;
  }

  .dashboard-stats > * {
    flex: 1 1 calc(50% - var(--space-5) / 2);
  }

  .dashboard-overview {
    flex-direction: column;
  }

  .overview-side {
    width: 100%;
    max-width: 100%;
    flex-direction: row;
    flex-wrap: wrap;
  }

  .overview-side > * {
    flex: 1 1 calc(50% - var(--space-5) / 2);
    min-width: 300px;
  }

  .quick-links {
    grid-template-columns: repeat(4, 1fr);
  }
}

@media (max-width: 768px) {
  .dashboard-main {
    margin-left: 0;
    padding: var(--space-5);
  }

  .header-content {
    flex-direction: column;
    align-items: stretch;
    gap: var(--space-4);
  }

  .header-actions {
    justify-content: flex-start;
  }

  .btn {
    flex: 1;
    justify-content: center;
  }

  .dashboard-stats {
    flex-direction: column;
  }

  .header-title {
    font-size: var(--font-size-3xl);
  }

  .revenue-amount {
    font-size: var(--font-size-3xl);
  }

  .revenue-stats {
    flex-direction: column;
    gap: var(--space-3);
  }

  .quick-links {
    grid-template-columns: repeat(2, 1fr);
  }

  .overview-side {
    flex-direction: column;
  }

  .overview-side > * {
    flex: 1 1 100%;
  }
}
</style>