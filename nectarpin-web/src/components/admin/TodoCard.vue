<script setup lang="ts">
defineProps<{
  title?: string
  items?: { text: string; dotColor: string }[]
}>()

const defaultItems = [
  { text: '审核 12 篇投稿并安排发布时间', dotColor: '#3B82F6' },
  { text: '处理 28 条待回复评论与举报内容', dotColor: '#F59E0B' },
]
</script>

<template>
  <div class="todo-card">
    <div class="todo-header">
      <h3 class="todo-title">{{ title || '今日待办' }}</h3>
      <span class="todo-count">{{ items?.length || defaultItems.length }}</span>
    </div>
    <div class="todo-items">
      <div class="todo-item" v-for="(item, index) in items || defaultItems" :key="index">
        <span class="todo-dot" :style="{ backgroundColor: item.dotColor }"></span>
        <span class="todo-text">{{ item.text }}</span>
        <svg class="todo-arrow" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
          <polyline points="9 18 15 12 9 6"/>
        </svg>
      </div>
    </div>
  </div>
</template>

<style scoped>
.todo-card {
  display: flex;
  flex-direction: column;
  gap: var(--space-5);
  padding: var(--space-6);
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-sm);
  height: 100%;
  transition: box-shadow var(--transition-fast), border-color var(--transition-fast);
}

.todo-card:hover {
  box-shadow: var(--shadow-md);
  border-color: var(--color-border-hover);
}

.todo-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.todo-title {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
}

.todo-count {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  min-width: 24px;
  height: 24px;
  padding: 0 var(--space-2);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  color: var(--color-primary);
  background: var(--color-primary-soft);
  border-radius: var(--radius-full);
  transition: background var(--transition-fast), color var(--transition-fast);
}

.todo-items {
  display: flex;
  flex-direction: column;
  gap: var(--space-2);
}

.todo-item {
  display: flex;
  align-items: center;
  gap: var(--space-3);
  padding: var(--space-3);
  background: var(--color-surface-alt);
  border-radius: var(--radius-lg);
  cursor: pointer;
  transition: all var(--transition-fast);
}

.todo-item:hover {
  background: var(--color-surface-hover);
  transform: translateX(4px);
}

.todo-item:hover .todo-arrow {
  opacity: 1;
  transform: translateX(2px);
}

.todo-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.todo-text {
  flex: 1;
  font-size: var(--font-size-sm);
  line-height: var(--line-height-normal);
  color: var(--color-text-primary);
}

.todo-arrow {
  color: var(--color-text-muted);
  opacity: 0;
  transition: opacity var(--transition-fast), transform var(--transition-fast);
}

@media (max-width: 768px) {
  .todo-card {
    padding: var(--space-5);
  }
}
</style>
