<script setup lang="ts">
import { ref } from 'vue'
import AdminSidebar from '@/components/admin/AdminSidebar.vue'
import StatCard from '@/components/admin/StatCard.vue'
import TrendChart from '@/components/admin/TrendChart.vue'
import TodoCard from '@/components/admin/TodoCard.vue'
import HeatmapChart from '@/components/admin/HeatmapChart.vue'

const stats = ref([
  {
    title: '本月阅读量',
    value: '286K',
    delta: '+18.2% 较上月',
    deltaType: 'success' as const,
    iconBgColor: '#EEF4FF',
    iconColor: '#2F6BFF',
    iconName: 'eye',
  },
  {
    title: '已发布文章',
    value: '148',
    delta: '12 篇待审核',
    deltaType: 'warning' as const,
    iconBgColor: '#EDFAF4',
    iconColor: '#10B981',
    iconName: 'file-text',
  },
  {
    title: '评论互动',
    value: '3,842',
    delta: '本周优质评论 126',
    deltaType: 'success' as const,
    iconBgColor: '#FFF7ED',
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
  console.log('Create new article')
}
</script>

<template>
  <div class="dashboard-layout">
    <AdminSidebar />
    <main class="dashboard-main">
      <header class="dashboard-topbar">
        <div class="topbar-heading">
          <h1 class="topbar-title">仪表盘概览</h1>
          <p class="topbar-subtitle">跟踪站点运营、内容增长与待处理事项。</p>
        </div>
        <div class="topbar-actions">
          <button class="btn btn-secondary" @click="handleExport">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <path d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4" />
              <polyline points="7 10 12 15 17 10" />
              <line x1="12" y1="15" x2="12" y2="3" />
            </svg>
            <span>导出报告</span>
          </button>
          <button class="btn btn-primary" @click="handleCreateArticle">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              width="18"
              height="18"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="2"
              stroke-linecap="round"
              stroke-linejoin="round"
            >
              <line x1="12" y1="5" x2="12" y2="19" />
              <line x1="5" y1="12" x2="19" y2="12" />
            </svg>
            <span>新建文章</span>
          </button>
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
            subtitle="按月份查看评论、收藏和互动行为分布。"
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
          subtitle="过去 6 周文章访问量与互动变化。"
          badge="实时同步"
          :data="trendData"
        />
      </section>
    </main>
  </div>
</template>

<style scoped>
.dashboard-layout {
  display: flex;
  height: 100vh;
  background-color: var(--color-bg);
  overflow: hidden;
}

.dashboard-main {
  flex: 1;
  display: flex;
  flex-direction: column;
  gap: 28px;
  padding: 32px;
  min-width: 0;
  overflow-y: auto;
}

.dashboard-topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding-bottom: 20px;
  border-bottom: 1px solid var(--color-border);
}

.topbar-heading {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.topbar-title {
  font-family: var(--font-heading);
  font-size: 24px;
  font-weight: 800;
  letter-spacing: -0.5px;
  color: var(--color-text);
}

.topbar-subtitle {
  font-size: 13px;
  color: var(--color-text-muted);
}

.topbar-actions {
  display: flex;
  gap: 12px;
}

.btn {
  display: flex;
  align-items: center;
  gap: 8px;
  padding: 12px 18px;
  border-radius: var(--radius-pill);
  font-size: 14px;
  font-weight: 600;
  letter-spacing: 0.2px;
  cursor: pointer;
  transition: all 0.2s ease;
  border: none;
}

.btn-primary {
  background-color: var(--color-primary);
  color: #ffffff;
}

.btn-primary:hover {
  background-color: var(--color-primary-strong);
}

.btn-secondary {
  background-color: var(--color-surface);
  color: var(--color-text);
  border: 1px solid var(--color-border);
}

.btn-secondary:hover {
  background-color: var(--color-surface-alt);
}

.dashboard-stats {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.dashboard-overview {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 20px;
}

.overview-main {
  grid-column: span 2;
  min-width: 0;
}

.overview-side {
  width: 100%;
}

.dashboard-trend {
  flex: 1;
}

@media (max-width: 1200px) {
  .dashboard-stats {
    grid-template-columns: repeat(2, 1fr);
  }

  .dashboard-overview {
    grid-template-columns: repeat(2, 1fr);
  }

  .overview-main {
    grid-column: span 1;
  }
}

@media (max-width: 1024px) {
  .dashboard-main {
    padding: 24px;
  }

  .topbar-title {
    font-size: 20px;
  }

  .dashboard-stats {
    grid-template-columns: repeat(2, 1fr);
  }

  .dashboard-overview {
    grid-template-columns: 1fr;
  }

  .overview-main {
    grid-column: span 1;
  }
}

@media (max-width: 768px) {
  .dashboard-main {
    padding: 20px;
    gap: 20px;
  }

  .dashboard-topbar {
    flex-direction: column;
    align-items: flex-start;
    gap: 16px;
  }

  .topbar-actions {
    width: 100%;
  }

  .topbar-actions .btn {
    flex: 1;
    justify-content: center;
  }

  .dashboard-stats {
    grid-template-columns: 1fr;
  }

  .btn span {
    display: inline;
  }
}

@media (max-width: 480px) {
  .btn span {
    display: none;
  }

  .btn {
    padding: 12px;
  }
}
</style>
