<script setup lang="ts">
import { ref, computed } from 'vue'
import { RouterView } from 'vue-router'
import { NConfigProvider, NMessageProvider, darkTheme } from 'naive-ui'
import { lightThemeOverrides, darkThemeOverrides } from './config/theme'

// 主题切换（可以根据需要改为响应式，比如从 localStorage 读取）
const isDark = ref(false)
const theme = computed(() => isDark.value ? darkTheme : null)
const themeOverrides = computed(() => isDark.value ? darkThemeOverrides : lightThemeOverrides)
</script>

<template>
  <NConfigProvider :theme="theme" :theme-overrides="themeOverrides">
    <NMessageProvider>
      <RouterView />
    </NMessageProvider>
  </NConfigProvider>
</template>

<style scoped>
:deep(> *) {
  min-height: 100vh;
}
</style>
