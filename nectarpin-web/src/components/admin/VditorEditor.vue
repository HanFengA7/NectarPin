<script setup lang="ts">
import 'md-editor-v3/lib/style.css'

import { computed } from 'vue'

import { MdEditor } from 'md-editor-v3'

const props = withDefaults(
  defineProps<{
    modelValue: string
    placeholder?: string
    minHeight?: number
  }>(),
  {
    placeholder: '请输入 Markdown 内容',
    minHeight: 480,
  },
)

const emit = defineEmits<{
  (e: 'update:modelValue', payload: string): void
}>()

const value = computed({
  get() {
    return props.modelValue
  },
  set(nextValue: string) {
    emit('update:modelValue', nextValue)
  },
})

const editorStyle = computed(() => ({
  height: `${props.minHeight}px`,
}))
</script>

<template>
  <div class="overflow-hidden border bg-background">
    <MdEditor
      v-model="value"
      class="[--md-bk-color:var(--background)] [--md-border-color:var(--border)] [--md-radius:0px]"
      language="zh-CN"
      preview-theme="github"
      code-theme="atom"
      :style="editorStyle"
      :placeholder="placeholder"
      :toolbars-exclude="['github']"
      show-code-row-number
    />
  </div>
</template>
