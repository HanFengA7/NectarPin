<script setup lang="ts">
import { computed } from 'vue'
import { useRoute, useRouter } from 'vue-router'

interface NavItem {
  name: string
  label: string
  icon: string
  path: string
}

const navItems: NavItem[] = [
  { name: 'admin-dashboard', label: 'Dashboard', icon: 'layout-dashboard', path: '/admin' },
  { name: 'admin-articles', label: 'Articles', icon: 'file-text', path: '/admin/articles' },
  { name: 'admin-comments', label: 'Comments', icon: 'message-square', path: '/admin/comments' },
  { name: 'admin-settings', label: 'Settings', icon: 'settings', path: '/admin/settings' },
]

const route = useRoute()
const router = useRouter()

const currentRoute = computed(() => route.name as string)

const navigateTo = (path: string) => {
  router.push(path)
}

const handleLogout = () => {
  localStorage.removeItem('token')
  localStorage.removeItem('refresh_token')
  localStorage.removeItem('user')
  sessionStorage.removeItem('token')
  sessionStorage.removeItem('refresh_token')
  sessionStorage.removeItem('user')
  router.push('/admin/login')
}
</script>

<template>
  <aside class="sidebar">
    <div class="sidebar-brand">
      <div class="brand-logo">
        <svg width="36" height="36" viewBox="0 0 36 36" fill="none">
          <defs>
            <linearGradient id="logoGradient" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#2F6BFF" />
              <stop offset="100%" stop-color="#1A4FCC" />
            </linearGradient>
          </defs>
          <circle cx="18" cy="18" r="18" fill="url(#logoGradient)" />
        </svg>
      </div>
      <div class="brand-info">
        <span class="brand-name">NectarPin Admin</span>
        <span class="brand-tagline">Content control center</span>
      </div>
    </div>

    <nav class="sidebar-nav">
      <button
        v-for="item in navItems"
        :key="item.name"
        class="nav-item"
        :class="{ active: currentRoute === item.name }"
        @click="navigateTo(item.path)"
      >
        <span class="nav-icon" v-html="getIcon(item.icon)"></span>
        <span class="nav-label">{{ item.label }}</span>
      </button>
    </nav>

    <div class="sidebar-user">
      <div class="user-avatar">
        <svg width="36" height="36" viewBox="0 0 36 36" fill="none">
          <defs>
            <linearGradient id="avatarGradient" x1="0%" y1="0%" x2="100%" y2="100%">
              <stop offset="0%" stop-color="#2F6BFF" />
              <stop offset="100%" stop-color="#1A4FCC" />
            </linearGradient>
          </defs>
          <circle cx="18" cy="18" r="18" fill="url(#avatarGradient)" />
        </svg>
      </div>
      <div class="user-info">
        <span class="user-name">Admin</span>
        <span class="user-role">超级管理员</span>
      </div>
      <button class="logout-btn" @click="handleLogout" title="退出登录">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          stroke="currentColor"
          stroke-width="2"
          stroke-linecap="round"
          stroke-linejoin="round"
        >
          <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
          <polyline points="16 17 21 12 16 7" />
          <line x1="21" y1="12" x2="9" y2="12" />
        </svg>
      </button>
    </div>
  </aside>
</template>

<script lang="ts">
const icons: Record<string, string> = {
  'layout-dashboard': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect x="3" y="3" width="7" height="9"/><rect x="14" y="3" width="7" height="5"/><rect x="14" y="12" width="7" height="9"/><rect x="3" y="16" width="7" height="5"/></svg>`,
  'file-text': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M14 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8z"/><polyline points="14 2 14 8 20 8"/><line x1="16" y1="13" x2="8" y2="13"/><line x1="16" y1="17" x2="8" y2="17"/><polyline points="10 9 9 9 8 9"/></svg>`,
  'message-square': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><path d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/></svg>`,
  'settings': `<svg xmlns="http://www.w3.org/2000/svg" width="18" height="18" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="3"/><path d="M19.4 15a1.65 1.65 0 0 0 .33 1.82l.06.06a2 2 0 0 1 0 2.83 2 2 0 0 1-2.83 0l-.06-.06a1.65 1.65 0 0 0-1.82-.33 1.65 1.65 0 0 0-1 1.51V21a2 2 0 0 1-2 2 2 2 0 0 1-2-2v-.09A1.65 1.65 0 0 0 9 19.4a1.65 1.65 0 0 0-1.82.33l-.06.06a2 2 0 0 1-2.83 0 2 2 0 0 1 0-2.83l.06-.06a1.65 1.65 0 0 0 .33-1.82 1.65 1.65 0 0 0-1.51-1H3a2 2 0 0 1-2-2 2 2 0 0 1 2-2h.09A1.65 1.65 0 0 0 4.6 9a1.65 1.65 0 0 0-.33-1.82l-.06-.06a2 2 0 0 1 0-2.83 2 2 0 0 1 2.83 0l.06.06a1.65 1.65 0 0 0 1.82.33H9a1.65 1.65 0 0 0 1-1.51V3a2 2 0 0 1 2-2 2 2 0 0 1 2 2v.09a1.65 1.65 0 0 0 1 1.51 1.65 1.65 0 0 0 1.82-.33l.06-.06a2 2 0 0 1 2.83 0 2 2 0 0 1 0 2.83l-.06.06a1.65 1.65 0 0 0-.33 1.82V9a1.65 1.65 0 0 0 1.51 1H21a2 2 0 0 1 2 2 2 2 0 0 1-2 2h-.09a1.65 1.65 0 0 0-1.51 1z"/></svg>`,
}

function getIcon(name: string): string {
  return icons[name] || ''
}

export default {
  methods: { getIcon },
}
</script>

<style scoped>
.sidebar {
  display: flex;
  flex-direction: column;
  width: var(--sidebar-width);
  min-width: var(--sidebar-width);
  height: 100vh;
  background-color: var(--color-surface);
  border-right: 1px solid #E4EDFF;
  padding: 28px 20px;
  gap: 20px;
}

.sidebar-brand {
  display: flex;
  align-items: center;
  gap: 12px;
  padding-bottom: 16px;
}

.brand-logo {
  width: 36px;
  height: 36px;
  flex-shrink: 0;
}

.brand-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
}

.brand-name {
  font-family: var(--font-heading);
  font-size: 16px;
  font-weight: 700;
  color: var(--color-text);
}

.brand-tagline {
  font-size: 11px;
  color: var(--color-text-muted);
}

.sidebar-nav {
  display: flex;
  flex-direction: column;
  gap: 4px;
  flex: 1;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border: none;
  background: transparent;
  border-radius: 12px;
  cursor: pointer;
  transition: all 0.2s ease;
  width: 100%;
  text-align: left;
}

.nav-item:hover {
  background-color: var(--color-primary-soft);
}

.nav-item:hover .nav-icon,
.nav-item:hover .nav-label {
  color: var(--color-primary-strong);
}

.nav-item.active {
  background-color: var(--color-primary-soft);
}

.nav-item.active .nav-icon,
.nav-item.active .nav-label {
  color: var(--color-primary-strong);
}

.nav-icon {
  display: flex;
  align-items: center;
  justify-content: center;
  color: var(--color-text-muted);
  flex-shrink: 0;
}

.nav-label {
  font-size: 14px;
  font-weight: 600;
  color: var(--color-text-muted);
}

.sidebar-user {
  display: flex;
  align-items: flex-end;
  gap: 10px;
  padding-top: 16px;
}

.user-avatar {
  width: 36px;
  height: 36px;
  border-radius: 50%;
  flex-shrink: 0;
}

.user-info {
  display: flex;
  flex-direction: column;
  gap: 2px;
  flex: 1;
  min-width: 0;
}

.user-name {
  font-size: 13px;
  font-weight: 600;
  color: var(--color-text);
}

.user-role {
  font-size: 11px;
  color: var(--color-text-muted);
}

.logout-btn {
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 8px;
  border: none;
  background: transparent;
  color: var(--color-text-muted);
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.2s ease;
}

.logout-btn:hover {
  background-color: var(--color-primary-soft);
  color: var(--color-primary);
}

@media (max-width: 1024px) {
  .sidebar {
    width: 72px;
    min-width: 72px;
    padding: 20px 12px;
    align-items: center;
  }

  .brand-info,
  .nav-label,
  .user-info {
    display: none;
  }

  .sidebar-brand {
    justify-content: center;
    padding-bottom: 16px;
  }

  .nav-item {
    justify-content: center;
    padding: 12px;
  }

  .sidebar-user {
    flex-direction: column;
    align-items: center;
    gap: 8px;
  }

  .logout-btn {
    padding: 10px;
  }
}

@media (max-width: 768px) {
  .sidebar {
    display: none;
  }
}
</style>
