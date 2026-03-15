<script setup lang="ts">
defineProps<{
  title: string
  value: string
  delta?: string
  deltaType?: 'success' | 'warning' | 'normal'
  iconBgColor?: string
  iconColor?: string
  iconName?: string
}>()
</script>

<template>
  <div class="stat-card">
    <div class="stat-icon" :style="{ backgroundColor: iconBgColor || '#EEF4FF' }">
      <span
        class="icon"
        :style="{ color: iconColor || 'var(--color-primary)' }"
        v-html="getIcon(iconName)"
      ></span>
    </div>
    <span class="stat-title">{{ title }}</span>
    <span class="stat-value">{{ value }}</span>
    <span
      class="stat-delta"
      :class="{
        success: deltaType === 'success',
        warning: deltaType === 'warning',
      }"
      v-if="delta"
    >
      {{ delta }}
    </span>
  </div>
</template>

<script lang="ts">
const icons: Record<string, string> = {
  eye: `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>`,
  'file-text': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>`,
  'message-circle': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>`,
}

function getIcon(name?: string): string {
  if (!name) return ''
  return icons[name] || ''
}

export default {
  methods: { getIcon },
}
</script>

<style scoped>
.stat-card {
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 20px 24px;
  background-color: var(--color-surface);
  border-radius: var(--radius-card);
  border: 1px solid #e4edff;
  box-shadow: 0 4px 24px rgba(26, 63, 112, 0.0625);
}

.stat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: 10px;
}

.icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-title {
  font-size: 12px;
  font-weight: 600;
  letter-spacing: 0.5px;
  color: var(--color-text-muted);
}

.stat-value {
  font-family: var(--font-heading);
  font-size: 28px;
  font-weight: 800;
  letter-spacing: -1px;
  color: var(--color-text);
}

.stat-delta {
  font-size: 12px;
  font-weight: 600;
  color: var(--color-success);
}

.stat-delta.warning {
  color: var(--color-warning);
}

@media (max-width: 768px) {
  .stat-card {
    padding: 16px 20px;
  }

  .stat-value {
    font-size: 24px;
  }
}
</style>
