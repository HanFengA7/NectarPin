<script setup lang="ts">
import AdminSidebar from './AdminSidebar.vue'
import { RouterView } from 'vue-router'
import { SidebarProvider, SidebarTrigger } from '../ui/sidebar'
</script>

<template>
    <SidebarProvider
        style="
            --sidebar-width: 13rem;
            --sidebar-mobile-width: 13rem;
        ">
        <AdminSidebar />
        <main>
            <SidebarTrigger />
            <RouterView />
        </main>
    </SidebarProvider>
</template>