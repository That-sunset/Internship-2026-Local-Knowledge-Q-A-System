<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Getrecord, Getchat, Post, Deletebyid } from './api/record'
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

// 1. Load sidebar chat logs on startup via Getrecord
onMounted(async () => {
  const data = await Getrecord()
  if (Array.isArray(data)) {
    sessions.value = data.map((item: any) => ({
      id: item.id,
      title: item.content || item.text || `Chat ${item.id}`
    }))
  }
})

// 2. Load selected chat via Getchat when clicking a sidebar item
async function selectSession(id: number) {
  currentId.value = id
  const data = await Getchat(id)
  if (data) {
    messages.value = [
      {
        role: 'assistant',
        words: (data.content || data.text || '').split(' ')
      }
    ]
  }
}

function newSession() {
  currentId.value = null
  messages.value = []
}

// 3. Delete chat via Deletebyid when clicking the hover X in sidebar
async function deleteSession(id: number) {
  await Deletebyid(id)
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
    if (wordIndex < wordArray.length) {
      messages.value[assistantMsgIndex].words.push(wordArray[wordIndex])
      wordIndex++
    } else {
      clearInterval(interval)
      isGenerating.value = false
    }
  }, 40)
}

// 4. Send POST request when submitting through ChatInput / SubmitButton
async function handleSend(userText: string) {
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

  try {
    await Post({ text: userText })
  } catch (err) {
    console.error(err)
  }

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