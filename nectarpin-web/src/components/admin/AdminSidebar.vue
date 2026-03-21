<script setup lang="ts">
import type { Component } from 'vue'
import { markRaw } from 'vue'
import { RouterLink, useRouter } from 'vue-router'
import { FileText, GalleryVerticalEnd, LayoutDashboard, LogOut, User } from 'lucide-vue-next'
import { logoutUser, logoutUserByRefreshToken } from '@/api/user'
import {
    Sidebar,
    SidebarContent,
    SidebarFooter,
    SidebarGroup,
    SidebarGroupContent,
    SidebarGroupLabel,
    SidebarHeader,
    SidebarMenu,
    SidebarMenuButton,
    SidebarMenuItem,
    SidebarRail,
    SidebarSeparator,
} from '@/components/ui/sidebar'
import { getAdminProfile, getRefreshToken, logoutAdmin } from '@/lib/admin-auth'

const router = useRouter()
const profile = getAdminProfile()

async function handleLogout() {
    const refreshToken = getRefreshToken()

    try {
        if (refreshToken) {
            await logoutUserByRefreshToken({ refresh_token: refreshToken })
        } else {
            await logoutUser()
        }
    } catch {
        // 登出失败时仍清理本地会话，避免前端停留在错误登录态。
    } finally {
        logoutAdmin()
        await router.replace({ name: 'admin-login' })
    }
}

interface AdminNavGroup {
  group: string
  menuItems: {
    name: string
    label: string
    icon: Component
    url: string
  }[]
}

const sidebarGroups: AdminNavGroup[] = [
  {
    group: '常用功能',
    menuItems: [
      {
        name: 'dashboard',
        label: '概览',
        icon: markRaw(LayoutDashboard),
        url: '/admin/dashboard',
      },
      {
        name: 'articles',
        label: '文章',
        icon: markRaw(FileText),
        url: '/admin/articles',
      },
    ],
  },
]
</script>

<template>
    <Sidebar>
        <SidebarHeader>
            <SidebarMenu>
                <SidebarMenuItem>
                    <SidebarMenuButton size="lg">
                        <div
                            class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground">
                            <GalleryVerticalEnd class="size-4" />
                        </div>
                        <div class="grid flex-1 text-left text-sm leading-tight">
                            <span class="truncate font-semibold">NectarPin</span>
                            <span class="truncate text-xs text-muted-foreground font-normal">V1.0.0 Alpha</span>
                        </div>
                    </SidebarMenuButton>
                </SidebarMenuItem>
            </SidebarMenu>
        </SidebarHeader>
        <SidebarContent>
            <SidebarGroup v-for="group in sidebarGroups" :key="group.group">
                    <SidebarGroupLabel>{{ group.group }}</SidebarGroupLabel>
                <SidebarGroupContent>
                    <SidebarMenuButton as-child v-for="item in group.menuItems" :key="item.name">
                        <RouterLink :to="item.url">
                            <component :is="item.icon" />
                            <span>{{ item.label }}</span>
                        </RouterLink>
                    </SidebarMenuButton>
                </SidebarGroupContent>
            </SidebarGroup>
        </SidebarContent>
        <SidebarFooter>
            <SidebarMenu>
                <SidebarMenuItem>
                    <SidebarMenuButton size="lg" tooltip="账户">
                        <div
                            class="flex aspect-square size-8 shrink-0 items-center justify-center rounded-md border border-sidebar-border bg-sidebar-accent/40">
                            <User class="size-4 text-muted-foreground" />
                        </div>
                        <div class="grid min-w-0 flex-1 text-left text-sm leading-tight">
                            <span class="truncate font-medium">{{ profile?.nickname || profile?.username || '管理员' }}</span>
                            <span class="truncate text-xs font-normal text-muted-foreground">
                                {{ profile?.email || '账户与偏好' }}
                            </span>
                        </div>
                    </SidebarMenuButton>
                </SidebarMenuItem>
            </SidebarMenu>
            <SidebarSeparator class="mx-2" />
            <SidebarMenu>
                <SidebarMenuItem>
                    <SidebarMenuButton
                        tooltip="退出登录"
                        class="text-muted-foreground hover:bg-sidebar-accent hover:text-destructive"
                        @click="handleLogout">
                        <LogOut class="size-4" />
                        <span>退出登录</span>
                    </SidebarMenuButton>
                </SidebarMenuItem>
            </SidebarMenu>
        </SidebarFooter>
        <SidebarRail />
    </Sidebar>
</template>