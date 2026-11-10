<script setup lang="ts">
import { ref } from 'vue'
import { Promotion } from '@element-plus/icons-vue'

const props = defineProps<{
  placeholder?: string
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: 'send', text: string): void
}>()

const inputMessage = ref('')

function handleSend() {
  if (!inputMessage.value.trim() || props.disabled) return
  emit('send', inputMessage.value)
  inputMessage.value = ''
}
</script>

<template>
  <el-input
    v-model="inputMessage"
    :placeholder="placeholder || 'Ask anything...'"
    class="pill-input"
    :disabled="disabled"
    @keydown.enter="handleSend"
  >
    <template #suffix>
      <el-tooltip content="Send message" placement="top" :disabled="!inputMessage.trim()">
        <el-button
          circle
          type="primary"
          class="send-btn"
          :disabled="!inputMessage.trim() || disabled"
          @click="handleSend"
        >
          <el-icon><Promotion /></el-icon>
        </el-button>
      </el-tooltip>
    </template>
  </el-input>
</template>

<style scoped>
:deep(.pill-input .el-input__wrapper) {
  background: var(--card-bg);
  border: 1px solid var(--card-border);
  border-radius: 28px;
  box-shadow: 0 4px 16px var(--shadow) !important;
  padding: 6px 12px 6px 18px;
}

:deep(.pill-input .el-input__inner) {
  font-size: 15px;
  color: var(--text-primary);
}

.send-btn {
  border: none;
}
</style>