<script setup lang="ts">
import 'md-editor-v3/lib/style.css'

import { computed, ref } from 'vue'

import { MdEditor } from 'md-editor-v3'

const props = withDefaults(
  defineProps<{
    modelValue: string
    placeholder?: string
    minHeight?: number
    /** 是否默认展开右侧预览；文章编辑等保持 true，仅个别页需 false */
    defaultPreviewOpen?: boolean
  }>(),
  {
    placeholder: '请输入 Markdown 内容',
    minHeight: 480,
    defaultPreviewOpen: true,
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

/** 预览分栏初始是否展开；可通过工具栏「预览」切换 */
const preview = ref(props.defaultPreviewOpen)
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
      v-model:preview="preview"
      show-code-row-number
    />
  </div>
</template>
