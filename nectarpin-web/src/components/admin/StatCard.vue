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

const icons: Record<string, string> = {
  eye: `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M1 12s4-8 11-8 11 8 11 8-4 8-11 8-11-8-11-8z"/><circle cx="12" cy="12" r="3"/></svg>`,
  'file-text': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>`,
  'message-circle': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 11.5a8.38 8.38 0 0 1-.9 3.8 8.5 8.5 0 0 1-7.6 4.7 8.38 8.38 0 0 1-3.8-.9L3 21l1.9-5.7a8.38 8.38 0 0 1-.9-3.8 8.5 8.5 0 0 1 4.7-7.6 8.38 8.38 0 0 1 3.8-.9h.5a8.48 8.48 0 0 1 8 8v.5z"/></svg>`,
}

const getIcon = (name?: string): string => {
  if (!name) return ''
  return icons[name] || ''
}
</script>

<template>
  <div class="stat-card">
    <div class="stat-header">
      <div class="stat-icon" :style="{ backgroundColor: iconBgColor || 'var(--color-primary-soft)' }">
        <span class="icon" :style="{ color: iconColor || 'var(--color-primary)' }" v-html="getIcon(iconName)"></span>
      </div>
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
    <div class="stat-body">
      <span class="stat-value">{{ value }}</span>
      <span class="stat-title">{{ title }}</span>
    </div>
    <div class="stat-bg-pattern"></div>
  </div>
</template>

<style scoped>
.stat-card {
  position: relative;
  display: flex;
  flex-direction: column;
  gap: var(--space-4);
  padding: var(--space-6);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-sm);
  overflow: hidden;
  transition: all var(--transition-base);
  cursor: default;
}

.stat-card:hover {
  box-shadow: var(--shadow-lg);
  transform: translateY(-2px);
  border-color: var(--color-border-hover);
}

.stat-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
}

.stat-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 48px;
  height: 48px;
  border-radius: var(--radius-lg);
  transition: transform var(--transition-fast);
}

.stat-card:hover .stat-icon {
  transform: scale(1.05);
}

.icon {
  display: flex;
  align-items: center;
  justify-content: center;
}

.stat-delta {
  display: inline-flex;
  align-items: center;
  padding: var(--space-1) var(--space-3);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  border-radius: var(--radius-full);
  background: var(--color-success-soft);
  color: var(--color-success);
  transition: background var(--transition-fast);
}

.stat-delta.warning {
  background: var(--color-warning-soft);
  color: var(--color-warning);
}

.stat-delta.normal {
  background: var(--color-surface-alt);
  color: var(--color-text-muted);
}

.stat-body {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
  z-index: 1;
}

.stat-value {
  font-family: var(--font-heading);
  font-size: var(--font-size-4xl);
  font-weight: var(--font-weight-extrabold);
  letter-spacing: var(--letter-spacing-tight);
  color: var(--color-text-primary);
  line-height: 1;
  transition: color var(--transition-fast);
}

.stat-title {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-medium);
  color: var(--color-text-muted);
}

.stat-bg-pattern {
  position: absolute;
  bottom: -24px;
  right: -24px;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background: linear-gradient(135deg, var(--color-primary-soft) 0%, transparent 70%);
  opacity: 0.6;
  transition: all var(--transition-slow);
}

.stat-card:hover .stat-bg-pattern {
  transform: scale(1.1);
  opacity: 0.8;
}

@media (max-width: 768px) {
  .stat-card {
    padding: var(--space-5);
  }

  .stat-value {
    font-size: var(--font-size-3xl);
  }

  .stat-icon {
    width: 44px;
    height: 44px;
  }
}
</style>
