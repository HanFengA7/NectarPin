<script setup lang="ts">
import AdminSidebar from '@/components/admin/AdminSidebar.vue'
</script>

<template>
  <div class="admin-layout">
    <AdminSidebar />
    <div class="admin-content">
      <router-view v-slot="{ Component, route }">
        <Transition name="page" mode="out-in">
          <component :is="Component" :key="route.path" />
        </Transition>
      </router-view>
    </div>
  </div>
</template>

<style scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
  background: var(--color-bg);
}

.admin-content {
  flex: 1;
  margin-left: calc(var(--sidebar-width) + 32px);
  min-height: 100vh;
  overflow: hidden;
}

@media (max-width: 1200px) {
  .admin-content {
    margin-left: calc(var(--sidebar-collapsed) + 32px);
  }
}

@media (max-width: 768px) {
  .admin-content {
    margin-left: 0;
  }
}
</style>
