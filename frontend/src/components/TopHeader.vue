<script setup lang="ts">
import { Expand } from '@element-plus/icons-vue'
import BackendButton from './BackendButton.vue'

defineProps<{ isSidebarOpen: boolean }>()
const emit = defineEmits<{ (e: 'toggle-sidebar'): void }>()
</script>

<template>
  <el-header class="top-header-banner">
    <el-row align="middle" justify="space-between" class="header-row">
      <el-col :span="16">
        <el-space :size="10" alignment="center">
          <el-tooltip v-if="!isSidebarOpen" content="Open sidebar" placement="bottom">
            <el-button circle text class="expand-btn" @click="emit('toggle-sidebar')">
              <el-icon><Expand /></el-icon>
            </el-button>
          </el-tooltip>

          <el-tag type="success" effect="dark" round class="status-badge" />
          <el-text tag="h1" size="large" class="header-main-title">
            AI Q&A System
          </el-text>
        </el-space>
      </el-col>

      <el-col :span="8" class="header-right-col">
        <BackendButton />
      </el-col>
    </el-row>
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

.header-row {
  width: 100%;
}

.header-right-col {
  display: flex;
  justify-content: flex-end;
  align-items: center;
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
  margin: 0;
  color: var(--el-text-color-primary);
}

.expand-btn {
  color: var(--text-secondary);
  font-size: 18px;
}
</style>