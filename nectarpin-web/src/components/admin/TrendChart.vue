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

const getBarHeight = (value: number, max: number): number => {
  return Math.max(20, (value / max) * 100)
}
</script>

<template>
  <div class="trend-card">
    <div class="trend-header">
      <div class="trend-text">
        <h3 class="trend-title">{{ title || '内容流量趋势' }}</h3>
        <p class="trend-subtitle">{{ subtitle || '过去 6 周文章访问量与互动变化。' }}</p>
      </div>
      <span class="trend-badge" v-if="badge">
        <svg xmlns="http://www.w3.org/2000/svg" width="12" height="12" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="23 6 13.5 15.5 8.5 10.5 1 18"/>
          <polyline points="17 6 23 6 23 12"/>
        </svg>
        {{ badge }}
      </span>
    </div>
    <div class="trend-chart">
      <div
        class="chart-bar-wrapper"
        v-for="(bar, index) in data || defaultData"
        :key="index"
      >
        <div
          class="chart-bar"
          :class="{ highlight: bar.highlight }"
          :style="{ height: `${getBarHeight(bar.value, Math.max(...(data || defaultData).map(d => d.value)))}%` }"
        >
          <span class="bar-tooltip">{{ bar.value }}</span>
        </div>
        <span class="bar-label">周{{ index + 1 }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.trend-card {
  display: flex;
  flex-direction: column;
  gap: var(--space-6);
  padding: var(--space-6);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-sm);
  transition: box-shadow var(--transition-fast), border-color var(--transition-fast);
}

.trend-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--color-border-hover);
}

.trend-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.trend-text {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.trend-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
}

.trend-subtitle {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
}

.trend-badge {
  display: inline-flex;
  align-items: center;
  gap: var(--space-2);
  padding: var(--space-2) var(--space-3);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  color: var(--color-primary);
  background: var(--color-primary-soft);
  border-radius: var(--radius-full);
  transition: background var(--transition-fast);
}

.trend-chart {
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  gap: var(--space-4);
  height: 200px;
  padding-top: var(--space-5);
}

.chart-bar-wrapper {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  gap: var(--space-3);
  height: 100%;
}

.chart-bar {
  position: relative;
  width: 100%;
  max-width: 60px;
  border-radius: var(--radius-md) var(--radius-md) var(--radius-xs) var(--radius-xs);
  background: linear-gradient(180deg, var(--color-surface-alt) 0%, var(--color-border) 100%);
  transition: all var(--transition-base);
  cursor: pointer;
}

.chart-bar:hover {
  transform: scaleY(1.02);
  transform-origin: bottom;
  box-shadow: var(--shadow-md);
}

.chart-bar:hover .bar-tooltip {
  opacity: 1;
  transform: translateX(-50%) translateY(-8px);
}

.chart-bar.highlight {
  background: linear-gradient(180deg, var(--color-primary) 0%, var(--color-primary-active) 100%);
}

.chart-bar.highlight:hover {
  background: linear-gradient(180deg, var(--color-primary-hover) 0%, var(--color-primary) 100%);
}

.bar-tooltip {
  position: absolute;
  top: -28px;
  left: 50%;
  transform: translateX(-50%) translateY(0);
  padding: var(--space-1) var(--space-2);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-primary);
  background: var(--color-surface);
  border-radius: var(--radius-sm);
  box-shadow: var(--shadow-sm);
  opacity: 0;
  transition: all var(--transition-fast);
  white-space: nowrap;
}

.bar-label {
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-medium);
  color: var(--color-text-muted);
}

@media (max-width: 768px) {
  .trend-card {
    padding: var(--space-5);
  }

  .trend-chart {
    height: 160px;
    gap: var(--space-3);
  }

  .chart-bar {
    max-width: 40px;
  }
}
</style>
