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
    <h3 class="todo-title">{{ title || '今日待办' }}</h3>
    <div class="todo-items">
      <div class="todo-item" v-for="(item, index) in items || defaultItems" :key="index">
        <span class="todo-dot" :style="{ backgroundColor: item.dotColor }"></span>
        <span class="todo-text">{{ item.text }}</span>
      </div>
    </div>
  </div>
</template>

<style scoped>
.todo-card {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 24px;
  background-color: var(--color-surface);
  border-radius: 16px;
  border: 1px solid #e4edff;
  box-shadow: 0 4px 24px rgba(26, 63, 112, 0.0625);
}

.todo-title {
  font-family: var(--font-heading);
  font-size: 16px;
  font-weight: 700;
  color: var(--color-text);
}

.todo-items {
  display: flex;
  flex-direction: column;
  gap: 10px;
}

.todo-item {
  display: flex;
  align-items: center;
  gap: 10px;
}

.todo-dot {
  width: 8px;
  height: 8px;
  border-radius: 50%;
  flex-shrink: 0;
}

.todo-text {
  font-size: 13px;
  line-height: 1.5;
  color: var(--color-text);
}

@media (max-width: 768px) {
  .todo-card {
    padding: 20px;
  }
}
</style>
