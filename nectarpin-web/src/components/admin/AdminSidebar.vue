<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'

interface NavItem {
  name: string
  label: string
  icon: string
  path: string
}

const navItems: NavItem[] = [
  { name: 'admin-dashboard', label: '仪表盘', icon: 'layout-dashboard', path: '/admin' },
  { name: 'admin-articles', label: '文章管理', icon: 'file-text', path: '/admin/articles' },
  { name: 'admin-comments', label: '评论管理', icon: 'message-square', path: '/admin/comments' },
  { name: 'admin-settings', label: '系统设置', icon: 'settings', path: '/admin/settings' },
]

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const currentRoute = computed(() => route.name as string)

const userName = computed(() => {
  const user = userStore.user
  if (!user) return '管理员'
  if (user.nickname && user.nickname.trim() !== '' && user.nickname !== user.username) {
    return user.nickname
  }
  return user.username
})

const userRole = computed(() => (userStore.isAdmin ? '超级管理员' : '普通用户'))

const navigateTo = (path: string) => {
  router.push(path)
}

const handleLogout = async () => {
  await userStore.logout()
  router.push('/admin/login')
}

const icons: Record<string, string> = {
  'layout-dashboard': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="9"/><rect x="14" y="3" width="7" height="5"/><rect x="14" y="12" width="7" height="9"/><rect x="3" y="16" width="7" height="5"/></svg>`,
  'file-text': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>`,
  'message-square': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>`,
  'settings': `<svg xmlns="http://www.w3.org/2000/svg" width="20" height="20" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>`,
}

const getIcon = (name: string): string => {
  return icons[name] || ''
}
</script>

<template>
  <aside class="sidebar">
    <div class="sidebar-header">
      <div class="brand">
        <div class="brand-logo">
          <svg width="40" height="40" viewBox="0 0 40 40" fill="none">
            <defs>
              <linearGradient id="sidebarLogoGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                <stop offset="0%" stop-color="#3B82F6" />
                <stop offset="100%" stop-color="#1D4ED8" />
              </linearGradient>
            </defs>
            <rect width="40" height="40" rx="12" fill="url(#sidebarLogoGradient)" />
            <path d="M12 20L17 25L28 14" stroke="white" stroke-width="3" stroke-linecap="round" stroke-linejoin="round" />
          </svg>
        </div>
        <div class="brand-text">
          <span class="brand-name">NectarPin</span>
          <span class="brand-role">内容管理</span>
        </div>
      </div>
    </div>

    <nav class="sidebar-nav">
      <div class="nav-section">
        <span class="nav-section-label">主菜单</span>
        <button
          v-for="item in navItems"
          :key="item.name"
          class="nav-item"
          :class="{ active: currentRoute === item.name }"
          @click="navigateTo(item.path)"
        >
          <span class="nav-icon" v-html="getIcon(item.icon)"></span>
          <span class="nav-label">{{ item.label }}</span>
          <span v-if="currentRoute === item.name" class="nav-indicator"></span>
        </button>
      </div>
    </nav>

    <div class="sidebar-footer">
      <div class="user-card">
        <div class="user-avatar">
          <svg width="36" height="36" viewBox="0 0 36 36" fill="none">
            <defs>
              <linearGradient id="userAvatarGradient" x1="0%" y1="0%" x2="100%" y2="100%">
                <stop offset="0%" stop-color="#3B82F6" />
                <stop offset="100%" stop-color="#1D4ED8" />
              </linearGradient>
            </defs>
            <circle cx="18" cy="18" r="18" fill="url(#userAvatarGradient)" />
            <circle cx="18" cy="14" r="5" fill="white" fill-opacity="0.9" />
            <path d="M8 28C8 23.5817 12.0294 20 18 20C23.9706 20 28 23.5817 28 28" stroke="white" stroke-opacity="0.9" stroke-width="2" stroke-linecap="round" />
          </svg>
        </div>
        <div class="user-info">
          <span class="user-name">{{ userName }}</span>
          <span class="user-role-badge">{{ userRole }}</span>
        </div>
        <button class="logout-btn" @click="handleLogout" title="退出登录" aria-label="退出登录">
          <svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
            <polyline points="16 17 21 12 16 7" />
            <line x1="21" y1="12" x2="9" y2="12" />
          </svg>
        </button>
      </div>
    </div>
  </aside>
</template>

<style scoped>
.sidebar {
  position: fixed;
  top: 16px;
  left: 16px;
  bottom: 16px;
  width: var(--sidebar-width);
  display: flex;
  flex-direction: column;
  background: var(--color-surface);
  border-radius: var(--radius-xl);
  border: 1px solid var(--color-border);
  box-shadow: var(--shadow-lg);
  overflow: hidden;
  z-index: var(--z-fixed);
  transition: box-shadow var(--transition-base);
}

.sidebar:hover {
  box-shadow: var(--shadow-xl);
}

.sidebar-header {
  padding: var(--space-5);
  border-bottom: 1px solid var(--color-border);
}

.brand {
  display: flex;
  align-items: center;
  gap: var(--space-3);
}

.brand-logo {
  flex-shrink: 0;
}

.brand-text {
  display: flex;
  flex-direction: column;
}

.brand-name {
  font-family: var(--font-heading);
  font-size: var(--font-size-lg);
  font-weight: var(--font-weight-bold);
  color: var(--color-text-primary);
  letter-spacing: var(--letter-spacing-tight);
}

.brand-role {
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-medium);
  color: var(--color-text-muted);
  text-transform: uppercase;
  letter-spacing: var(--letter-spacing-wider);
}

.sidebar-nav {
  flex: 1;
  padding: var(--space-4) var(--space-3);
  overflow-y: auto;
}

.nav-section {
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.nav-section-label {
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-muted);
  text-transform: uppercase;
  letter-spacing: var(--letter-spacing-wider);
  padding: var(--space-2) var(--space-3);
  margin-bottom: var(--space-1);
}

.nav-item {
  position: relative;
  display: flex;
  align-items: center;
  gap: var(--space-3);
  padding: var(--space-3) var(--space-3);
  border-radius: var(--radius-lg);
  border: none;
  background: transparent;
  color: var(--color-text-secondary);
  cursor: pointer;
  transition: all var(--transition-fast);
  width: 100%;
  text-align: left;
}

.nav-item:hover {
  background: var(--color-surface-alt);
  color: var(--color-text-primary);
}

.nav-item.active {
  background: var(--color-primary-soft);
  color: var(--color-primary);
}

.nav-item.active:hover {
  background: var(--color-primary-soft);
}

.nav-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  width: 20px;
  height: 20px;
}

.nav-label {
  font-size: var(--font-size-base);
  font-weight: var(--font-weight-medium);
  flex: 1;
}

.nav-indicator {
  position: absolute;
  right: var(--space-3);
  width: 6px;
  height: 6px;
  background: var(--color-primary);
  border-radius: var(--radius-full);
}

.sidebar-footer {
  padding: var(--space-4);
  border-top: 1px solid var(--color-border);
}

.user-card {
  display: flex;
  align-items: center;
  gap: var(--space-3);
  padding: var(--space-3);
  background: var(--color-surface-alt);
  border-radius: var(--radius-lg);
  transition: background var(--transition-fast);
}

.user-card:hover {
  background: var(--color-surface-hover);
}

.user-avatar {
  flex-shrink: 0;
}

.user-info {
  flex: 1;
  min-width: 0;
  display: flex;
  flex-direction: column;
  gap: var(--space-1);
}

.user-name {
  font-size: var(--font-size-sm);
  font-weight: var(--font-weight-semibold);
  color: var(--color-text-primary);
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
}

.user-role-badge {
  display: inline-flex;
  align-items: center;
  padding: 2px var(--space-2);
  font-size: var(--font-size-xs);
  font-weight: var(--font-weight-medium);
  color: var(--color-primary);
  background: var(--color-primary-soft);
  border-radius: var(--radius-full);
  width: fit-content;
}

.logout-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 36px;
  height: 36px;
  border-radius: var(--radius-md);
  color: var(--color-text-muted);
  transition: all var(--transition-fast);
  flex-shrink: 0;
}

.logout-btn:hover {
  background: var(--color-danger-soft);
  color: var(--color-danger);
}

@media (max-width: 1200px) {
  .sidebar {
    width: var(--sidebar-collapsed);
  }

  .brand-text,
  .nav-label,
  .nav-section-label,
  .user-info,
  .nav-indicator {
    display: none;
  }

  .brand {
    justify-content: center;
  }

  .nav-item {
    justify-content: center;
    padding: var(--space-3);
  }

  .user-card {
    justify-content: center;
    padding: var(--space-2);
  }
}

@media (max-width: 768px) {
  .sidebar {
    display: none;
  }
}
</style>
