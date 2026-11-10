<script setup lang="ts">
import ChatInput from './ChatInput.vue'

defineProps<{
  isGenerating: boolean
}>()

const emit = defineEmits<{
  (e: 'send', text: string): void
}>()
</script>

<template>
  <div class="welcome-screen">
    <div class="input-wrapper">
      <ChatInput
        placeholder="Ask anything..."
        :disabled="isGenerating"
        @send="text => emit('send', text)"
      />
    </div>
  </div>
</template>

<style scoped>
.welcome-screen {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding-bottom: 120px;
}

.input-wrapper {
  width: 90%;
  max-width: 680px;
}
</style>