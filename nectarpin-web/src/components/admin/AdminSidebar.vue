<script setup lang="ts">
import type { Component } from 'vue'
import { markRaw, ref } from 'vue'
import { storeToRefs } from 'pinia'
import { RouterLink, useRoute, useRouter } from 'vue-router'
import {
  ChevronRight,
  FileText,
  FolderTree,
  GalleryVerticalEnd,
  LayoutDashboard,
  LogOut,
  Tags,
  User,
  Link,
  Settings,
  List,
} from 'lucide-vue-next'
import { CollapsibleContent, CollapsibleRoot, CollapsibleTrigger } from 'reka-ui'
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
  SidebarMenuSub,
  SidebarMenuSubButton,
  SidebarMenuSubItem,
  SidebarRail,
  SidebarSeparator,
} from '@/components/ui/sidebar'
import { getAdminProfile, getRefreshToken, logoutAdmin } from '@/lib/admin-auth'
import { useSiteStore } from '@/stores/site'

const router = useRouter()
const route = useRoute()
const profile = getAdminProfile()
const siteStore = useSiteStore()
const { siteName } = storeToRefs(siteStore)

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

interface NavItem {
  name: string
  label: string
  icon?: Component
  url: string
}

interface NavItemWithChildren {
  name: string
  label: string
  icon: Component
  children: NavItem[]
}

type NavEntry = NavItem | NavItemWithChildren

function hasChildren(entry: NavEntry): entry is NavItemWithChildren {
  return 'children' in entry
}

interface AdminNavGroup {
  group: string
  menuItems: NavEntry[]
}

const sidebarGroups: AdminNavGroup[] = [
  {
    group: '概览',
    menuItems: [
      {
        name: 'dashboard',
        label: '仪表盘',
        icon: markRaw(LayoutDashboard),
        url: '/admin/dashboard',
      },
    ],
  },
  {
    group: '内容',
    menuItems: [
      {
        name: 'article',
        label: '文章管理',
        icon: markRaw(FileText),
        children: [
          {
            name: 'articles',
            label: '文章列表',
            icon: markRaw(FileText),
            url: '/admin/articles',
          },
          {
            name: 'articleCategories',
            label: '文章分类',
            icon: markRaw(FolderTree),
            url: '/admin/articleCategories',
          },
          {
            name: 'articleTags',
            label: '文章标签',
            icon: markRaw(Tags),
            url: '/admin/articleTags',
          },
        ],
      },
      {
        name: 'links',
        label: '友链管理',
        icon: markRaw(Link),
        children: [
          {
            name: 'links',
            label: '友链列表',
            icon: markRaw(List),
            url: '/admin/links',
          },
          {
            name: 'linkCategories',
            label: '页面设置',
            icon: markRaw(Settings),
            url: '/admin/linkCategories',
          },
        ],
      },
    ],
  },
  {
    group: '全局',
    menuItems: [
      {
        name: 'global-siteSettings',
        label: '站点管理',
        icon: markRaw(FileText),
        children: [
          {
            name: 'global-siteSettings-index',
            label: '首页设置',
            url: '/admin/global/siteSettings/index',
          },
          {
            name: 'global-siteSettings-footer',
            label: '页脚设置',
            url: '/admin/global/siteSettings/footer',
          },
        ],
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
              class="flex aspect-square size-8 items-center justify-center rounded-lg bg-sidebar-primary text-sidebar-primary-foreground"
            >
              <GalleryVerticalEnd class="size-4" />
            </div>
            <div class="grid flex-1 text-left text-sm leading-tight">
              <span class="truncate font-semibold">{{ siteName }}</span>
              <span
                class="flex items-center gap-1.5 truncate text-xs text-muted-foreground font-normal"
                >0.0.3
                <span
                  class="inline-flex items-center rounded-full border border-transparent bg-primary px-1.5 py-px text-[8px] font-medium leading-none text-primary-foreground"
                  >Alpha</span
                ></span
              >
            </div>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarHeader>
    <SidebarContent>
      <SidebarGroup v-for="group in sidebarGroups" :key="group.group">
        <SidebarGroupLabel>{{ group.group }}</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            <template v-for="item in group.menuItems" :key="item.name">
              <SidebarMenuItem v-if="!hasChildren(item)">
                <SidebarMenuButton as-child :is-active="route.path === (item as NavItem).url">
                  <RouterLink :to="(item as NavItem).url">
                    <component :is="item.icon" />
                    <span>{{ item.label }}</span>
                  </RouterLink>
                </SidebarMenuButton>
              </SidebarMenuItem>

              <CollapsibleRoot v-else as-child :default-open="true">
                <SidebarMenuItem>
                  <CollapsibleTrigger as-child>
                    <SidebarMenuButton>
                      <component :is="item.icon" />
                      <span>{{ item.label }}</span>
                      <ChevronRight
                        class="ml-auto size-4 transition-transform duration-200 group-data-[state=open]/collapsible:rotate-90"
                      />
                    </SidebarMenuButton>
                  </CollapsibleTrigger>
                  <CollapsibleContent>
                    <SidebarMenuSub>
                      <SidebarMenuSubItem
                        v-for="child in (item as NavItemWithChildren).children"
                        :key="child.name"
                      >
                        <SidebarMenuSubButton as-child :is-active="route.path === child.url">
                          <RouterLink :to="child.url">
                            <component v-if="child.icon" :is="child.icon" />
                            <span>{{ child.label }}</span>
                          </RouterLink>
                        </SidebarMenuSubButton>
                      </SidebarMenuSubItem>
                    </SidebarMenuSub>
                  </CollapsibleContent>
                </SidebarMenuItem>
              </CollapsibleRoot>
            </template>
          </SidebarMenu>
        </SidebarGroupContent>
      </SidebarGroup>
    </SidebarContent>
    <SidebarFooter>
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton
            size="lg"
            tooltip="账户"
            as-child
            :is-active="route.name === 'admin-profile'"
          >
            <RouterLink :to="{ name: 'admin-profile' }">
              <div
                class="flex aspect-square size-8 shrink-0 items-center justify-center rounded-md border border-sidebar-border bg-sidebar-accent/40"
              >
                <User class="size-4 text-muted-foreground" />
              </div>
              <div class="grid min-w-0 flex-1 text-left text-sm leading-tight">
                <span class="truncate font-medium">{{
                  profile?.nickname || profile?.username || '管理员'
                }}</span>
                <span class="truncate text-xs font-normal text-muted-foreground">
                  {{ profile?.email || '个人资料' }}
                </span>
              </div>
            </RouterLink>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
      <SidebarSeparator class="mx-2" />
      <SidebarMenu>
        <SidebarMenuItem>
          <SidebarMenuButton
            tooltip="退出登录"
            class="text-muted-foreground hover:bg-sidebar-accent hover:text-destructive"
            @click="handleLogout"
          >
            <LogOut class="size-4" />
            <span>退出登录</span>
          </SidebarMenuButton>
        </SidebarMenuItem>
      </SidebarMenu>
    </SidebarFooter>
    <SidebarRail />
  </Sidebar>
</template>
