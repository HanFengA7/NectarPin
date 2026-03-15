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

const MONTH_LABELS = [
  '一月',
  '二月',
  '三月',
  '四月',
  '五月',
  '六月',
  '七月',
  '八月',
  '九月',
  '十月',
  '十一月',
  '十二月',
]
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
    data.push({
      date: formatDate(day),
      count: Math.floor(Math.random() * 50),
    })
  })
  return data
}

const internalValues = ref<HeatmapValue[]>(
  props.values.length > 0 ? props.values : generateRandomData(),
)

watch(
  () => props.values,
  (newValues) => {
    if (newValues.length > 0) {
      internalValues.value = newValues
    }
  },
)

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
            <text
              v-for="(label, index) in getMonthLabels"
              :key="index"
              :x="label.x"
              :y="12"
              class="month-label"
            >
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
            :style="{
              backgroundColor: color,
              width: `${getCellSize}px`,
              height: `${getCellSize}px`,
            }"
          ></span>
        </div>
        <span class="legend-text">多</span>
      </div>
    </div>

    <div
      v-if="showTooltip && tooltip"
      class="heatmap-tooltip"
      :style="{ left: `${tooltipPosition.x + 10}px`, top: `${tooltipPosition.y - 30}px` }"
    >
      {{ tooltipContent }}
    </div>
  </div>
</template>

<style scoped>
.heatmap-card {
  display: flex;
  flex-direction: column;
  gap: 18px;
  padding: 20px 24px;
  background-color: var(--color-surface);
  border-radius: var(--radius-card);
  border: 1px solid #e4edff;
  box-shadow: 0 4px 24px rgba(26, 63, 112, 0.0625);
  width: 100%;
  box-sizing: border-box;
  position: relative;
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
  flex-direction: column;
  gap: 12px;
}

.heatmap-scroll {
  overflow-x: auto;
  padding-bottom: 8px;
  width: 100%;
}

.heatmap-svg {
  display: block;
  width: 100%;
  height: auto;
  min-width: 600px;
}

.month-label {
  font-size: 12px;
  fill: var(--color-text-muted);
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans', Helvetica, Arial, sans-serif;
}

.day-label {
  font-size: 12px;
  fill: var(--color-text-muted);
  font-family:
    -apple-system, BlinkMacSystemFont, 'Segoe UI', 'Noto Sans', Helvetica, Arial, sans-serif;
}

.heatmap-cell {
  cursor: pointer;
  transition: opacity 0.2s ease;
}

.heatmap-cell:hover {
  opacity: 0.8;
  stroke: #1b1f23;
  stroke-width: 1px;
}

.heatmap-legend {
  display: flex;
  justify-content: flex-end;
  align-items: center;
  gap: 6px;
  padding-left: 40px;
}

.legend-text {
  font-size: 12px;
  color: var(--color-text-muted);
}

.legend-scale {
  display: flex;
  gap: 3px;
}

.legend-cell {
  border-radius: 2px;
}

.heatmap-tooltip {
  position: fixed;
  padding: 6px 10px;
  background-color: rgba(0, 0, 0, 0.85);
  color: #fff;
  font-size: 12px;
  border-radius: 4px;
  pointer-events: none;
  z-index: 1000;
  white-space: nowrap;
}

@media (max-width: 768px) {
  .heatmap-card {
    padding: 20px;
  }

  .legend-text:first-child {
    display: none;
  }
}
</style>
