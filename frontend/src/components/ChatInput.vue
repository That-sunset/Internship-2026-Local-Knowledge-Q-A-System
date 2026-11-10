<script setup lang="ts">
import { ref } from 'vue'
import SubmitButton from './SubmitButton.vue'

const props = defineProps<{
  placeholder?: string
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: 'send', text: string): void
}>()

const inputMessage = ref('')

function handleSend(event?: KeyboardEvent) {
  if (event && event.shiftKey) return
  if (event) event.preventDefault()

  if (!inputMessage.value.trim() || props.disabled) return
  emit('send', inputMessage.value)
  inputMessage.value = ''
}
</script>

<template>
  <div class="chat-input-wrapper">
    <el-input
      v-model="inputMessage"
      style="width: 100%"
      :autosize="{ minRows: 1, maxRows: 4 }"
      type="textarea"
      :placeholder="placeholder || 'Ask anything...'"
      :disabled="disabled"
      resize="none"
      class="chat-textarea"
      @keydown.enter="handleSend"
    />
    <SubmitButton
      :disabled="!inputMessage.trim() || disabled"
      @click="handleSend"
    />
  </div>
</template>

<style scoped>
.chat-input-wrapper {
  position: relative;
  width: 100%;
  display: flex;
  align-items: flex-end;
}

:deep(.chat-textarea .el-textarea__inner) {
  border-radius: 20px;
  padding: 10px 48px 10px 16px;
  font-size: 15px;
  line-height: 1.5;
  background-color: var(--card-bg);
  border-color: var(--card-border);
  color: var(--text-primary);
  box-shadow: 0 4px 16px var(--shadow);
}
</style>