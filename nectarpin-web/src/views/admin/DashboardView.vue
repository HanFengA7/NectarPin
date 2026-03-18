<script setup lang="ts">
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import StatCard from '@/components/admin/StatCard.vue'
import TrendChart from '@/components/admin/TrendChart.vue'
import TodoCard from '@/components/admin/TodoCard.vue'
import HeatmapChart from '@/components/admin/HeatmapChart.vue'

const router = useRouter()

const stats = ref([
  {
    title: '本月阅读量',
    value: '286K',
    delta: '+18.2%',
    deltaType: 'success' as const,
    iconBgColor: '#EFF6FF',
    iconColor: '#3B82F6',
    iconName: 'eye',
  },
  {
    title: '已发布文章',
    value: '148',
    delta: '12 待审核',
    deltaType: 'warning' as const,
    iconBgColor: '#ECFDF5',
    iconColor: '#10B981',
    iconName: 'file-text',
  },
  {
    title: '评论互动',
    value: '3,842',
    delta: '+126 本周',
    deltaType: 'success' as const,
    iconBgColor: '#FFFBEB',
    iconColor: '#F59E0B',
    iconName: 'message-circle',
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
])

const handleExport = () => {
  console.log('Export report')
}

const handleCreateArticle = () => {
  router.push('/admin/articles/add')
}
</script>

<template>
  <div class="dashboard-view">
    <main class="dashboard-main">
      <header class="dashboard-header">
        <div class="header-content">
          <div class="header-text">
            <h1 class="header-title">仪表盘</h1>
            <p class="header-subtitle">跟踪站点运营、内容增长与待处理事项</p>
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
          <HeatmapChart
            title="社区活跃热力图"
            subtitle="按月份查看评论、收藏和互动行为分布"
            badge="最近 12 个月"
          />
        </div>
        <div class="overview-side">
          <TodoCard title="今日待办" :items="todoItems" />
        </div>
      </section>

      <section class="dashboard-trend">
        <TrendChart
          title="内容流量趋势"
          subtitle="过去 6 周文章访问量与互动变化"
          badge="实时同步"
          :data="trendData"
        />
      </section>
    </main>
  </div>
</template>

<style scoped>
.dashboard-view {
  padding: var(--space-page);
}

.dashboard-main {
  width: 100%;
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
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: var(--space-5);
  margin-bottom: var(--space-6);
}

.dashboard-overview {
  display: grid;
  grid-template-columns: 1fr 360px;
  gap: var(--space-5);
  margin-bottom: var(--space-6);
}

.overview-main {
  min-width: 0;
}

.overview-side {
  display: flex;
}

.dashboard-trend {
  width: 100%;
}

@media (max-width: 1400px) {
  .dashboard-overview {
    grid-template-columns: 1fr 320px;
  }
}

@media (max-width: 1200px) {
  .dashboard-main {
    margin-left: calc(var(--sidebar-collapsed) + 32px);
    padding: var(--space-6);
  }

  .dashboard-stats {
    grid-template-columns: repeat(2, 1fr);
  }

  .dashboard-overview {
    grid-template-columns: 1fr;
  }

  .overview-side {
    max-width: 400px;
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
    grid-template-columns: 1fr;
  }

  .header-title {
    font-size: var(--font-size-3xl);
  }
}
</style>
