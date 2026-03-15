<script setup lang="ts">
defineProps<{
  title?: string
  subtitle?: string
  badge?: string
  data?: { value: number; highlight?: boolean }[]
}>()

const defaultData = [
  { value: 120 },
  { value: 156 },
  { value: 210, highlight: true },
  { value: 172 },
  { value: 228, highlight: true },
  { value: 194 },
]
</script>

<template>
  <div class="trend-card">
    <div class="trend-header">
      <div class="trend-text">
        <h3 class="trend-title">{{ title || '内容流量趋势' }}</h3>
        <p class="trend-subtitle">{{ subtitle || '过去 6 周文章访问量与互动变化。' }}</p>
      </div>
      <span class="trend-badge" v-if="badge">{{ badge }}</span>
    </div>
    <div class="trend-chart">
      <div
        class="chart-bar"
        :class="{ highlight: bar.highlight }"
        v-for="(bar, index) in data || defaultData"
        :key="index"
        :style="{ height: `${bar.value}px` }"
      ></div>
    </div>
  </div>
</template>

<style scoped>
.trend-card {
  display: flex;
  flex-direction: column;
  gap: 20px;
  padding: 24px;
  background-color: var(--color-surface);
  border-radius: 16px;
  border: 1px solid #E4EDFF;
  box-shadow: 0 4px 24px rgba(26, 63, 112, 0.0625);
}

.trend-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.trend-text {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.trend-title {
  font-family: var(--font-heading);
  font-size: 16px;
  font-weight: 700;
  color: var(--color-text);
}

.trend-subtitle {
  font-size: 13px;
  color: var(--color-text-muted);
}

.trend-badge {
  padding: 5px 12px;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.4px;
  color: var(--color-primary-strong);
  background-color: var(--color-primary-soft);
  border-radius: var(--radius-pill);
}

.trend-chart {
  display: flex;
  align-items: flex-end;
  gap: 12px;
  height: 260px;
  padding-top: 16px;
}

.chart-bar {
  flex: 1;
  min-width: 40px;
  border-radius: 10px 10px 4px 4px;
  background: linear-gradient(180deg, #DDEAFF 0%, #C5D8FF 100%);
  transition: all 0.3s ease;
}

.chart-bar.highlight {
  background: linear-gradient(180deg, #3B7BFF 0%, #1A4FCC 100%);
}

.chart-bar:hover {
  opacity: 0.8;
  transform: scaleY(1.02);
  transform-origin: bottom;
}

@media (max-width: 768px) {
  .trend-card {
    padding: 20px;
  }

  .trend-chart {
    height: 180px;
    gap: 8px;
  }

  .chart-bar {
    min-width: 24px;
  }
}
</style>
