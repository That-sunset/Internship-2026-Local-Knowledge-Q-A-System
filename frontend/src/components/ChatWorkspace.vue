<script setup lang="ts">
import { ref, watch, nextTick } from 'vue'
import { UserFilled, Service } from '@element-plus/icons-vue'
import ChatInput from './ChatInput.vue'

interface Message {
  role: 'user' | 'assistant'
  words: string[]
  rawText?: string
}

const props = defineProps<{
  messages: Message[]
  isGenerating: boolean
}>()

const emit = defineEmits<{
  (e: 'send', text: string): void
}>()

const scrollbarRef = ref<any>(null)

watch(
  () => props.messages,
  () => {
    nextTick(() => {
      if (scrollbarRef.value) {
        scrollbarRef.value.setScrollTop(100000)
      }
    })
  },
  { deep: true }
)
</script>

<template>
  <div class="chat-workspace">
    <el-scrollbar ref="scrollbarRef" class="chat-viewport">
      <div class="chat-container">
        <div
          v-for="(msg, index) in messages"
          :key="index"
          class="message-row"
          :class="msg.role"
        >
          <!-- Role Avatar -->
          <el-avatar
            :size="32"
            :icon="msg.role === 'user' ? UserFilled : Service"
            class="message-avatar"
            :class="msg.role"
          />

          <div class="message-body">
            <div class="author-name">
              {{ msg.role === 'user' ? 'You' : 'AI Assistant' }}
            </div>
            
            <!-- Element Plus Card Bubble -->
            <el-card shadow="never" class="bubble-card" :class="msg.role">
              <template v-if="msg.role === 'user'">
                {{ msg.rawText }}
              </template>

              <template v-else>
                <span
                  v-for="(word, wIdx) in msg.words"
                  :key="wIdx"
                  class="word-fade"
                >
                  {{ word }}&nbsp;
                </span>
              </template>
            </el-card>
          </div>
        </div>
      </div>
    </el-scrollbar>

    <footer class="bottom-input-area">
      <div class="chat-bottom-input">
        <ChatInput
          placeholder="Ask a follow-up question..."
          :disabled="isGenerating"
          @send="text => emit('send', text)"
        />
      </div>
    </footer>
  </div>
</template>

<style scoped>
.chat-workspace {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-viewport {
  flex: 1;
  padding: 20px;
}

.chat-container {
  max-width: 768px;
  margin: 0 auto;
  display: flex;
  flex-direction: column;
}

.message-row {
  display: flex;
  gap: 12px;
  margin-bottom: 24px;
  width: 100%;
}

.message-row.user {
  flex-direction: row-reverse;
}

.message-avatar.user {
  background-color: #409eff;
}

.message-avatar.assistant {
  background-color: #67c23a;
}

.message-body {
  max-width: 75%;
  display: flex;
  flex-direction: column;
}

.message-row.user .message-body {
  align-items: flex-end;
}

.message-row.assistant .message-body {
  align-items: flex-start;
}

.author-name {
  font-size: 12px;
  font-weight: 600;
  color: var(--text-secondary);
  margin-bottom: 4px;
}

:deep(.bubble-card) {
  border-radius: 18px;
  border: none;
}

:deep(.bubble-card .el-card__body) {
  padding: 12px 18px;
  font-size: 15px;
  line-height: 1.6;
  word-break: break-word;
}

:deep(.bubble-card.user) {
  background: var(--user-bubble-bg);
  color: var(--user-bubble-text);
  border-bottom-right-radius: 4px;
}

:deep(.bubble-card.assistant) {
  background: var(--assistant-bubble-bg);
  color: var(--assistant-bubble-text);
  border: 1px solid var(--card-border);
  border-bottom-left-radius: 4px;
}

.word-fade {
  display: inline-block;
  animation: wordFadeIn 0.15s cubic-bezier(0.1, 0.9, 0.2, 1) forwards;
}

@keyframes wordFadeIn {
  0% {
    opacity: 0;
    transform: translateY(4px) scale(0.96);
    filter: blur(1px);
  }
  100% {
    opacity: 1;
    transform: translateY(0) scale(1);
    filter: blur(0);
  }
}

.bottom-input-area {
  padding: 16px;
  display: flex;
  justify-content: center;
}

.chat-bottom-input {
  width: 100%;
  max-width: 768px;
}
</style>