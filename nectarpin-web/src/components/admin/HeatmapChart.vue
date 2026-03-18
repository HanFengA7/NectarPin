<script setup lang="ts">
import { ref, computed, onMounted, onUnmounted, watch } from 'vue'

interface HeatmapValue {
  date: string
  count: number
}

const props = withDefaults(
  defineProps<{
    title?: string
    subtitle?: string
    badge?: string
    values?: HeatmapValue[]
    endDate?: Date
    rangeColor?: string[]
    max?: number
    locale?: string
    tooltip?: boolean
    vertical?: boolean
  }>(),
  {
    title: '社区活跃热力图',
    subtitle: '按月份查看评论、收藏和互动行为分布。',
    badge: '',
    values: () => [],
    endDate: () => new Date(),
    rangeColor: () => ['#ebedf0', '#c6e48b', '#7bc96f', '#239a3b', '#196127'],
    max: 0,
    locale: 'zh-CN',
    tooltip: true,
    vertical: false,
  },
)

const emit = defineEmits<{
  (e: 'day-click', value: HeatmapValue): void
}>()

const containerRef = ref<HTMLElement | null>(null)
const containerWidth = ref(800)

const SQUARE_SIZE = 12
const SQUARE_GAP = 4
const SQUARE_BORDER_RADIUS = 2
const LEFT_PANEL_WIDTH = 40
const TOP_PANEL_HEIGHT = 20

const MONTH_LABELS = ['一月', '二月', '三月', '四月', '五月', '六月', '七月', '八月', '九月', '十月', '十一月', '十二月']
const DAY_LABELS = ['周一', '', '周三', '', '周五', '', '周日']

const updateWidth = () => {
  if (containerRef.value) {
    containerWidth.value = containerRef.value.offsetWidth
  }
}

onMounted(() => {
  updateWidth()
  window.addEventListener('resize', updateWidth)
})

onUnmounted(() => {
  window.removeEventListener('resize', updateWidth)
})

const getStartDate = () => {
  const end = props.endDate
  const start = new Date(end)
  start.setFullYear(start.getFullYear() - 1)
  start.setDate(start.getDate() - start.getDay() + 1)
  return start
}

const getDays = () => {
  const start = getStartDate()
  const end = props.endDate
  const days: Date[] = []
  const current = new Date(start)
  while (current <= end) {
    days.push(new Date(current))
    current.setDate(current.getDate() + 1)
  }
  return days
}

const getWeekCount = computed(() => {
  const days = getDays()
  return Math.ceil(days.length / 7)
})

const getCellSize = computed(() => {
  const availableWidth = containerWidth.value - LEFT_PANEL_WIDTH - 20
  const weeks = getWeekCount.value
  const cellWidth = (availableWidth - (weeks + 1) * SQUARE_GAP) / weeks
  return Math.max(SQUARE_SIZE, Math.min(20, cellWidth))
})

const svgWidth = computed(() => {
  return LEFT_PANEL_WIDTH + getWeekCount.value * (getCellSize.value + SQUARE_GAP) + SQUARE_GAP
})

const svgHeight = computed(() => {
  return TOP_PANEL_HEIGHT + 7 * (getCellSize.value + SQUARE_GAP) + SQUARE_GAP
})

const formatDate = (date: Date): string => {
  const year = date.getFullYear()
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  return `${year}-${month}-${day}`
}

const getCount = (date: Date): number => {
  const dateStr = formatDate(date)
  const value = props.values.find((v) => v.date === dateStr)
  return value?.count ?? 0
}

const getMaxCount = computed(() => {
  if (props.max > 0) return props.max
  const counts = props.values.map((v) => v.count)
  return Math.max(...counts, 1)
})

const getColorIndex = (count: number): number => {
  if (count === 0) return 0
  const ratio = count / getMaxCount.value
  return Math.min(4, Math.ceil(ratio * 4))
}

const getColor = (count: number): string => {
  const index = getColorIndex(count)
  return props.rangeColor[index] ?? props.rangeColor[0]!
}

const getCoordinates = (dayIndex: number) => {
  const week = Math.floor(dayIndex / 7)
  const dayOfWeek = dayIndex % 7
  return {
    x: LEFT_PANEL_WIDTH + week * (getCellSize.value + SQUARE_GAP),
    y: TOP_PANEL_HEIGHT + dayOfWeek * (getCellSize.value + SQUARE_GAP),
  }
}

const getMonthLabels = computed(() => {
  const days = getDays()
  const labels: { month: number; x: number }[] = []
  let lastMonth = -1
  days.forEach((day, index) => {
    const month = day.getMonth()
    if (month !== lastMonth) {
      const coords = getCoordinates(index)
      labels.push({ month, x: coords.x })
      lastMonth = month
    }
  })
  return labels
})

const handleDayClick = (date: Date) => {
  const dateStr = formatDate(date)
  const value = props.values.find((v) => v.date === dateStr) || { date: dateStr, count: 0 }
  emit('day-click', value)
}

const generateRandomData = (): HeatmapValue[] => {
  const data: HeatmapValue[] = []
  const days = getDays()
  days.forEach((day) => {
    data.push({ date: formatDate(day), count: Math.floor(Math.random() * 50) })
  })
  return data
}

const internalValues = ref<HeatmapValue[]>(props.values.length > 0 ? props.values : generateRandomData())

watch(() => props.values, (newValues) => {
  if (newValues.length > 0) {
    internalValues.value = newValues
  }
})

const getCountFromDate = (date: Date): number => {
  const dateStr = formatDate(date)
  const value = internalValues.value.find((v) => v.date === dateStr)
  return value?.count ?? 0
}

const tooltipContent = ref<string>('')
const tooltipPosition = ref({ x: 0, y: 0 })
const showTooltip = ref(false)

const handleMouseEnter = (event: MouseEvent, date: Date) => {
  if (!props.tooltip) return
  const count = getCountFromDate(date)
  const dateStr = formatDate(date)
  tooltipContent.value = `${dateStr}: ${count} 次活动`
  tooltipPosition.value = { x: event.clientX, y: event.clientY }
  showTooltip.value = true
}

const handleMouseLeave = () => {
  showTooltip.value = false
}
</script>

<template>
  <div class="heatmap-card">
    <div class="heatmap-header">
      <div class="heatmap-text">
        <h3 class="heatmap-title">{{ title }}</h3>
        <p class="heatmap-subtitle">{{ subtitle }}</p>
      </div>
      <span class="heatmap-badge" v-if="badge">{{ badge }}</span>
    </div>
    <div class="heatmap-body">
      <div class="heatmap-scroll" ref="containerRef">
        <svg
          class="heatmap-svg"
          :width="svgWidth"
          :height="svgHeight"
          :viewBox="`0 0 ${svgWidth} ${svgHeight}`"
          preserveAspectRatio="xMinYMid meet"
        >
          <g class="month-labels">
            <text v-for="(label, index) in getMonthLabels" :key="index" :x="label.x" :y="12" class="month-label">
              {{ MONTH_LABELS[label.month] }}
            </text>
          </g>
          <g class="day-labels">
            <text
              v-for="(label, index) in DAY_LABELS"
              :key="index"
              :x="0"
              :y="TOP_PANEL_HEIGHT + index * (getCellSize + SQUARE_GAP) + getCellSize - 2"
              class="day-label"
            >
              {{ label }}
            </text>
          </g>
          <g class="heatmap-cells">
            <rect
              v-for="(day, index) in getDays()"
              :key="index"
              :x="getCoordinates(index).x"
              :y="getCoordinates(index).y"
              :width="getCellSize"
              :height="getCellSize"
              :rx="SQUARE_BORDER_RADIUS"
              :ry="SQUARE_BORDER_RADIUS"
              :fill="getColor(getCountFromDate(day))"
              class="heatmap-cell"
              @click="handleDayClick(day)"
              @mouseenter="handleMouseEnter($event, day)"
              @mouseleave="handleMouseLeave"
            />
          </g>
        </svg>
      </div>
      <div class="heatmap-legend">
        <span class="legend-text">少</span>
        <div class="legend-scale">
          <span
            v-for="(color, index) in rangeColor"
            :key="index"
            class="legend-cell"
            :style="{ backgroundColor: color, width: `${getCellSize}px`, height: `${getCellSize}px` }"
          ></span>
        </div>
        <span class="legend-text">多</span>
      </div>
    </div>
    <div v-if="showTooltip && tooltip" class="heatmap-tooltip" :style="{ left: `${tooltipPosition.x + 10}px`, top: `${tooltipPosition.y - 30}px` }">
      {{ tooltipContent }}
    </div>
  </div>
</template>

<style scoped>
.heatmap-card {
  display: flex;
  flex-direction: column;
  gap: var(--space-5);
  padding: var(--space-6);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-sm);
  width: 100%;
  box-sizing: border-box;
  position: relative;
  transition: box-shadow var(--transition-fast), border-color var(--transition-fast);
}

.heatmap-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--color-border-hover);
}

.heatmap-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.heatmap-text {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.heatmap-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
}

.heatmap-subtitle {
  font-size: var(--font-size-sm);
  color: var(--color-text-muted);
}

.heatmap-badge {
  display: inline-flex;
  align-items: center;
  padding: var(--space-1) var(--space-3);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  color: var(--color-primary);
  background: var(--color-primary-soft);
  border-radius: var(--radius-full);
}

.heatmap-body {
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
}

.heatmap-scroll {
  overflow-x: auto;
  padding-bottom: var(--space-2);
  width: 100%;
}

.heatmap-svg {
  display: block;
  width: 100%;
  height: auto;
  min-width: 600px;
}

.month-label {
  font-size: var(--font-size-xs);
  fill: var(--color-text-muted);
  font-family: var(--font-sans);
}

.day-label {
  font-size: var(--font-size-xs);
  fill: var(--color-text-muted);
  font-family: var(--font-sans);
}

.heatmap-cell {
  cursor: pointer;
  transition: opacity var(--transition-fast);
}

.heatmap-cell:hover {
  opacity: 0.8;
  stroke: var(--color-text-primary);
  stroke-width: 1px;
}

.heatmap-legend {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: var(--space-2);
  padding-left: LEFT_PANEL_WIDTH;
}

.legend-text {
  font-size: var(--font-size-xs);
  color: var(--color-text-muted);
}

.legend-scale {
  display: flex;
  gap: 3px;
}

.legend-cell {
  border-radius: var(--radius-xs);
}

.heatmap-tooltip {
  position: fixed;
  padding: var(--space-2) var(--space-3);
  background: var(--color-text-primary);
  color: var(--color-bg);
  font-size: var(--font-size-xs);
  border-radius: var(--radius-sm);
  pointer-events: none;
  z-index: var(--z-tooltip);
  white-space: nowrap;
  box-shadow: var(--shadow-lg);
}
</style>
