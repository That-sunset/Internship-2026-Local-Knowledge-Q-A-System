<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from './components/Sidebar.vue'
import TopHeader from './components/TopHeader.vue'
import WelcomeScreen from './components/WelcomeScreen.vue'
import ChatWorkspace from './components/ChatWorkspace.vue'
import BackendButton from './components/BackendButton.vue'

interface Session {
  id: number
  title: string
}

interface Message {
  role: 'user' | 'assistant'
  words: string[]
  rawText?: string
}

const isSidebarOpen = ref(true)
const isGenerating = ref(false)

const sessions = ref<Session[]>([])
const currentId = ref<number | null>(null)
const messages = ref<Message[]>([])

function selectSession(id: number) {
  currentId.value = id
}

function newSession() {
  currentId.value = null
  messages.value = []
}

function streamWordsEffect(fullText: string) {
  const assistantMsgIndex = messages.value.length
  messages.value.push({ role: 'assistant', words: [] })

  const wordArray = fullText.split(' ')
  let wordIndex = 0

  const interval = setInterval(() => {
    if (wordIndex < wordArray.length) {
      messages.value[assistantMsgIndex].words.push(wordArray[wordIndex])
      wordIndex++
    } else {
      clearInterval(interval)
      isGenerating.value = false
    }
  }, 40)
}

function handleSend(userText: string) {
  if (currentId.value === null) {
    const id = Date.now()
    const title = userText.slice(0, 30)
    sessions.value.unshift({ id, title })
    currentId.value = id
  }

  messages.value.push({
    role: 'user',
    words: [],
    rawText: userText
  })

  isGenerating.value = true

  setTimeout(() => {
    const mockResponse = `Thank you for reaching out! Based on your query regarding "${userText}", here is the relevant information retrieved from our knowledge base. Please let me know if you need further assistance.`
    streamWordsEffect(mockResponse)
  }, 300)
}
</script>

<template>
  <div class="app-layout">
    <Sidebar
      v-model:is-open="isSidebarOpen"
      :sessions="sessions"
      :current-id="currentId"
      @select-session="selectSession"
      @new-session="newSession"
    />

    <main class="main-area">
      <!-- Placed inside main for proper visibility -->
      <BackendButton />

      <TopHeader
        :is-sidebar-open="isSidebarOpen"
        @toggle-sidebar="isSidebarOpen = !isSidebarOpen"
      />

      <WelcomeScreen
        v-if="messages.length === 0"
        :is-generating="isGenerating"
        @send="handleSend"
      />

      <ChatWorkspace
        v-else
        :messages="messages"
        :is-generating="isGenerating"
        @send="handleSend"
      />
    </main>
  </div>
</template>

<style scoped>
.app-layout {
  display: flex;
  min-height: 100vh;
}

.main-area {
  position: relative;
  flex: 1;
}
</style>