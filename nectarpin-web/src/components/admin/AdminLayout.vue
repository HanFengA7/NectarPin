<script setup lang="ts">
import AdminSidebar from './AdminSidebar.vue'
import { RouterView } from 'vue-router'
import { SidebarInset, SidebarProvider, SidebarTrigger } from '../ui/sidebar'
</script>

<template>
  <SidebarProvider
    style="
      --sidebar-width: 13rem;
      --sidebar-mobile-width: 13rem;
    "
  >
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