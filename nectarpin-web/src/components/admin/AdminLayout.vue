<script setup lang="ts">
import { onMounted } from 'vue'
import { RouterView } from 'vue-router'

import { useSiteStore } from '@/stores/site'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '../ui/sidebar'
import AdminSidebar from './AdminSidebar.vue'

const siteStore = useSiteStore()

onMounted(() => {
  void siteStore.hydrateFromApi()
})
</script>

<template>
  <SidebarProvider style="--sidebar-width: 13rem; --sidebar-mobile-width: 13rem">
    <AdminSidebar />
    <SidebarInset class="min-w-0">
      <header class="flex h-14 items-center border-b bg-background px-4 sm:px-6">
        <SidebarTrigger />
      </header>
      <div class="min-w-0 flex-1">
        <RouterView />
      </div>
    </SidebarInset>
  </SidebarProvider>
</template>
