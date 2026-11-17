<script setup lang="ts">
import { ref } from 'vue'
import { Connection, Check, Close } from '@element-plus/icons-vue'
import { api } from '../Api/api'

const connected = ref(false)
const loading = ref(false)
const errorMsg = ref('')

async function connect() {
  errorMsg.value = ''
  loading.value = true
  console.log('Button clicked! Fetching backend...')

  try {
    const data = await api.get('/api/info')
    console.log('Response received:', data)
    connected.value = true
  } catch (err: any) {
    console.error('Fetch failed:', err)
    errorMsg.value = err.message?.startsWith('Server error:')
      ? err.message
      : 'Cannot reach backend'
    connected.value = false
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="backend-status-widget">
    <el-button 
      type="primary" 
      :loading="loading" 
      :icon="Connection"
      @click="connect"
    >
      Communicate with Backend
    </el-button>

    <el-tag v-if="connected" type="success" effect="light" class="status-tag">
      <el-icon><Check /></el-icon>
      Backend Connected
    </el-tag>

    <el-tag v-if="errorMsg" type="danger" effect="light" class="status-tag">
      <el-icon><Close /></el-icon>
      {{ errorMsg }}
    </el-tag>
  </div>
</template>

<style scoped>
.backend-status-widget {
  position: absolute;
  top: 12px;
  right: 12px;
  z-index: 9999;
  display: flex;
  align-items: center;
  gap: 8px;
  background: rgba(255, 255, 255, 0.95);
  padding: 8px 12px;
  border-radius: 6px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  backdrop-filter: blur(4px);
}

.status-tag {
  display: flex;
  align-items: center;
  gap: 4px;
}
</style>