<script setup lang="ts">
import { ref } from 'vue'
import Sidebar from './components/Sidebar.vue'
import TopHeader from './components/TopHeader.vue'
import WelcomeScreen from './components/WelcomeScreen.vue'
import ChatWorkspace from './components/ChatWorkspace.vue'
import BackendButton from './components/BackendButton.vue'
import { Postrecord, Postchat } from './Api/record'
import type { ChatRecord } from './Api/record'

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

function selectSession(payload: { id: number; chats: ChatRecord[] }) {
  currentId.value = payload.id
  messages.value = payload.chats.map(chat =>
    chat.source === 'user'
      ? { role: 'user', words: [], rawText: chat.content }
      : { role: 'assistant', words: chat.content.split(' ') }
  )
}

function newSession() {
  currentId.value = null
  messages.value = []
}

function deleteSession(id: number) {
  sessions.value = sessions.value.filter(s => s.id !== id)
  if (currentId.value === id) {
    newSession()
  }
}

function streamWordsEffect(fullText: string) {
  const assistantMsgIndex = messages.value.length
  messages.value.push({ role: 'assistant', words: [] })

  const wordArray = fullText.split(' ')
  let wordIndex = 0

  const interval = setInterval(() => {
    const currentMsg = messages.value[assistantMsgIndex]
    const currentWord = wordArray[wordIndex]

    if (currentMsg && currentWord !== undefined) {
      currentMsg.words.push(currentWord)
      wordIndex++
    } else {
      clearInterval(interval)
      isGenerating.value = false
    }
  }, 40)
}

async function handleSend(userText: string) {
  try {
    // 1. Create a real SQLite session if this is a brand-new chat
    if (currentId.value === null) {
      const res = await Postrecord()
      const newId = res.data.id
      const title = userText.slice(0, 30)
      sessions.value.unshift({ id: newId, title })
      currentId.value = newId
    }

    const sessionId = currentId.value!

    // 2. Display the user message immediately
    messages.value.push({
      role: 'user',
      words: [],
      rawText: userText
    })

    // 3. Send { prompt } -> backend inserts user + model rows -> returns { answer }
    isGenerating.value = true
    const res = await Postchat(sessionId, { prompt: userText })
    const aiAnswer = res.answer || res.data?.content || 'No response from model.'

    // 4. Stream the real AI answer in the UI
    streamWordsEffect(aiAnswer)
  } catch (err) {
    console.error('Failed to send message:', err)
    isGenerating.value = false
  }
}
</script>

<template>
  <div class="app-layout">
    <Sidebar
      v-model:is-open="isSidebarOpen"
      :sessions="sessions"
      :current-id="currentId"
      @loaded-sessions="sessions = $event"
      @select-session="selectSession"
      @new-session="newSession"
      @delete-session="deleteSession"
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