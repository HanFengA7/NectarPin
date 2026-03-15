<script setup lang="ts">
import { ref, computed } from 'vue'

const props = defineProps<{
  title?: string
  subtitle?: string
  badge?: string
}>()

const generateHeatmapData = () => {
  const data: { level: number }[][] = []
  for (let week = 0; week < 53; week++) {
    const weekData: { level: number }[] = []
    for (let day = 0; day < 7; day++) {
      weekData.push({ level: Math.floor(Math.random() * 5) })
    }
    data.push(weekData)
  }
  return data
}

const heatmapData = ref(generateHeatmapData())

const months = ['Mar', 'May', 'Jul', 'Sep', 'Nov', 'Jan', 'Mar']
const dayLabels = ['Mon', 'Wed', 'Fri']

const getLevelColor = (level: number): string => {
  const colors: string[] = ['#EEF4FF', '#C5D8FF', '#8BB8FF', '#4F8FFF', '#2F6BFF']
  return colors[level] ?? colors[0]!
}
</script>

<template>
  <div class="heatmap-card">
    <div class="heatmap-header">
      <div class="heatmap-text">
        <h3 class="heatmap-title">{{ title || '社区活跃热力图' }}</h3>
        <p class="heatmap-subtitle">{{ subtitle || '按月份查看评论、收藏和互动行为分布。' }}</p>
      </div>
      <span class="heatmap-badge" v-if="badge">{{ badge }}</span>
    </div>
    <div class="heatmap-body">
      <div class="day-labels">
        <span class="day-label" v-for="day in dayLabels" :key="day">{{ day }}</span>
      </div>
      <div class="heatmap-content">
        <div class="month-labels">
          <span class="month-label" v-for="month in months" :key="month">{{ month }}</span>
        </div>
        <div class="heatmap-grid">
          <div class="heatmap-week" v-for="(week, weekIndex) in heatmapData" :key="weekIndex">
            <div
              class="heatmap-cell"
              v-for="(day, dayIndex) in week"
              :key="dayIndex"
              :style="{ backgroundColor: getLevelColor(day.level) }"
            ></div>
          </div>
        </div>
        <div class="heatmap-legend">
          <span class="legend-text">基于评论、收藏与互动行为统计</span>
          <div class="legend-scale">
            <span class="legend-label">Less</span>
            <div class="legend-cells">
              <span class="legend-cell" v-for="i in 5" :key="i" :style="{ backgroundColor: getLevelColor(i - 1) }"></span>
            </div>
            <span class="legend-label">More</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.heatmap-card {
  display: flex;
  flex-direction: column;
  gap: 18px;
  padding: 24px;
  background-color: var(--color-surface);
  border-radius: 16px;
  border: 1px solid #E4EDFF;
  box-shadow: 0 4px 24px rgba(26, 63, 112, 0.0625);
}

.heatmap-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.heatmap-text {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.heatmap-title {
  font-family: var(--font-heading);
  font-size: 16px;
  font-weight: 700;
  color: var(--color-text);
}

.heatmap-subtitle {
  font-size: 13px;
  color: var(--color-text-muted);
}

.heatmap-badge {
  padding: 5px 12px;
  font-size: 11px;
  font-weight: 700;
  letter-spacing: 0.4px;
  color: var(--color-primary-strong);
  background-color: var(--color-primary-soft);
  border-radius: var(--radius-pill);
}

.heatmap-body {
  display: flex;
  gap: 14px;
}

.day-labels {
  display: flex;
  flex-direction: column;
  gap: 18px;
  padding-top: 26px;
}

.day-label {
  font-size: 12px;
  color: var(--color-text-muted);
}

.heatmap-content {
  display: flex;
  flex-direction: column;
  gap: 12px;
  flex: 1;
}

.month-labels {
  display: flex;
  justify-content: space-between;
}

.month-label {
  font-size: 12px;
  color: var(--color-text-muted);
}

.heatmap-grid {
  display: flex;
  gap: 4px;
}

.heatmap-week {
  display: flex;
  flex-direction: column;
  gap: 4px;
}

.heatmap-cell {
  width: 12px;
  height: 12px;
  border-radius: 2px;
  transition: all 0.2s ease;
}

.heatmap-cell:hover {
  transform: scale(1.2);
}

.heatmap-legend {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.legend-text {
  font-size: 12px;
  color: var(--color-text-muted);
}

.legend-scale {
  display: flex;
  align-items: center;
  gap: 6px;
}

.legend-label {
  font-size: 12px;
  color: var(--color-text-muted);
}

.legend-cells {
  display: flex;
  gap: 3px;
}

.legend-cell {
  width: 12px;
  height: 12px;
  border-radius: 2px;
}

@media (max-width: 1024px) {
  .heatmap-grid {
    overflow-x: auto;
    padding-bottom: 8px;
  }

  .heatmap-cell {
    width: 10px;
    height: 10px;
  }
}

@media (max-width: 768px) {
  .heatmap-card {
    padding: 20px;
  }

  .heatmap-body {
    flex-direction: column;
  }

  .day-labels {
    flex-direction: row;
    padding-top: 0;
    gap: 24px;
  }

  .heatmap-cell {
    width: 8px;
    height: 8px;
  }

  .legend-text {
    display: none;
  }
}
</style>
