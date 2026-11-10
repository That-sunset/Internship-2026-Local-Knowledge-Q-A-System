<script setup lang="ts">
import { Expand } from '@element-plus/icons-vue'

defineProps<{
  isSidebarOpen: boolean
}>()

const emit = defineEmits<{
  (e: 'toggle-sidebar'): void
}>()
</script>

<template>
  <el-header class="top-header-banner">
    <div class="header-left-content">
      <el-tooltip v-if="!isSidebarOpen" content="Open sidebar" placement="bottom">
        <el-button
          circle
          text
          class="expand-btn"
          @click="emit('toggle-sidebar')"
        >
          <el-icon><Expand /></el-icon>
        </el-button>
      </el-tooltip>

      <div class="header-title-box">
        <el-tag type="success" effect="dark" round class="status-badge" />
        <h1 class="header-main-title">AI Q&A System</h1>
      </div>
    </div>
  </el-header>
</template>

<style scoped>
.top-header-banner {
  height: 56px !important;
  background-color: var(--header-bg);
  border-bottom: 1px solid var(--header-border);
  display: flex;
  align-items: center;
  padding: 0 20px;
  box-shadow: 0 1px 4px var(--shadow);
}

.header-left-content {
  display: flex;
  align-items: center;
  gap: 12px;
}

.header-title-box {
  display: flex;
  align-items: center;
  gap: 8px;
}

.status-badge {
  width: 8px;
  height: 8px;
  padding: 0;
  min-width: 8px;
}

.header-main-title {
  font-size: 16px;
  font-weight: 600;
  color: var(--text-primary);
  margin: 0;
  letter-spacing: -0.2px;
}

.expand-btn {
  color: var(--text-secondary);
  font-size: 18px;
}
</style>