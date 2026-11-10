<script setup lang="ts">
import { ref, computed, nextTick } from 'vue'

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
const isDarkMode = ref(false)
const searchQuery = ref('')
const isSearching = ref(false)
const searchInputRef = ref<HTMLInputElement | null>(null)
const inputMessage = ref('')
const isGenerating = ref(false)

const sessions = ref<Session[]>([])
const currentId = ref<number | null>(null)
const messages = ref<Message[]>([])

const filteredSessions = computed(() => {
  if (!searchQuery.value.trim()) return sessions.value
  return sessions.value.filter(s =>
    s.title.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

function toggleSidebar() {
  isSidebarOpen.value = !isSidebarOpen.value
}

function toggleSearch() {
  isSearching.value = !isSearching.value
  if (isSearching.value) {
    nextTick(() => {
      searchInputRef.value?.focus()
    })
  } else {
    searchQuery.value = ''
  }
}

function selectSession(id: number) {
  currentId.value = id
}

function newSession() {
  currentId.value = null
  messages.value = []
  inputMessage.value = ''
}

// Gemini-style word-by-word fade-in streaming effect
function streamWordsEffect(fullText: string) {
  const assistantMsgIndex = messages.value.length
  messages.value.push({ role: 'assistant', words: [] })

  const wordArray = fullText.split(' ')
  let wordIndex = 0

  const interval = setInterval(() => {
    if (wordIndex < wordArray.length) {
      messages.value[assistantMsgIndex].words.push(wordArray[wordIndex])
      wordIndex++
      
      nextTick(() => {
        const viewport = document.querySelector('.chat-viewport')
        if (viewport) viewport.scrollTop = viewport.scrollHeight
      })
    } else {
      clearInterval(interval)
      isGenerating.value = false
    }
  }, 40)
}

function handleSend() {
  if (!inputMessage.value.trim() || isGenerating.value) return
  
  if (currentId.value === null) {
    const id = Date.now()
    const title = inputMessage.value.slice(0, 30)
    sessions.value.unshift({ id, title })
    currentId.value = id
  }

  messages.value.push({ 
    role: 'user', 
    words: [], 
    rawText: inputMessage.value 
  })
  
  const userText = inputMessage.value
  inputMessage.value = ''
  isGenerating.value = true

  setTimeout(() => {
    const mockResponse = `Thank you for reaching out! Based on your query regarding "${userText}", here is the relevant information retrieved from our knowledge base. Please let me know if you need further clarification or help.`
    streamWordsEffect(mockResponse)
  }, 300)
}
</script>

<template>
  <div class="app-layout" :class="{ 'dark-theme': isDarkMode }">
    <!-- Collapsible Sidebar -->
    <aside class="sidebar" :class="{ closed: !isSidebarOpen }">
      <!-- Sidebar Top Controls -->
      <div class="sidebar-top">
        <!-- Expanding Search Pill Component -->
        <div class="expanding-search-container" :class="{ expanded: isSearching }">
          <button class="icon-btn search-toggle-btn" title="Search conversations" @click="toggleSearch">
            <svg class="svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="11" cy="11" r="7" />
              <path d="M16 16l4.5 4.5" stroke-linecap="round" />
            </svg>
          </button>
          <input
            ref="searchInputRef"
            v-model="searchQuery"
            type="text"
            placeholder="Search history..."
            class="expanding-search-input"
          />
        </div>

        <button class="sidebar-toggle-btn" title="Close sidebar" @click="toggleSidebar">
          <svg class="panel-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <rect x="3" y="3" width="18" height="18" rx="2" />
            <path d="M9 3v18" />
          </svg>
          <span class="hover-arrow">«</span>
        </button>
      </div>

      <!-- New Chat Action -->
      <div class="new-chat-wrapper">
        <button class="new-chat-btn" @click="newSession">
          <svg class="svg-icon add-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
            <line x1="12" y1="5" x2="12" y2="19" stroke-linecap="round" />
            <line x1="5" y1="12" x2="19" y2="12" stroke-linecap="round" />
          </svg>
          <span class="btn-text">New chat</span>
        </button>
      </div>

      <!-- Conversation History -->
      <div class="sidebar-history">
        <div class="history-label">Recent Conversations</div>
        <div
          v-for="s in filteredSessions"
          :key="s.id"
          class="history-item"
          :class="{ active: s.id === currentId }"
          @click="selectSession(s.id)"
        >
          <span class="chat-title">{{ s.title }}</span>
        </div>
      </div>

      <!-- Dark / Light Mode Switch -->
      <div class="sidebar-footer">
        <label class="switch-container">
          <div class="switch-mode-icon">
            <svg v-if="!isDarkMode" class="theme-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <circle cx="12" cy="12" r="4" />
              <path d="M12 2v2m0 16v2M4.93 4.93l1.41 1.41m11.32 11.32l1.41 1.41M2 12h2m16 0h2M4.93 19.07l1.41-1.41m11.32-11.32l1.41-1.41" stroke-linecap="round" />
            </svg>
            <svg v-else class="theme-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <path d="M21 12.79A9 9 0 1 1 11.21 3 7 7 0 0 0 21 12.79z" stroke-linecap="round" stroke-linejoin="round" />
            </svg>
            <span class="switch-label">{{ isDarkMode ? 'Dark Mode' : 'Light Mode' }}</span>
          </div>
          <input type="checkbox" v-model="isDarkMode" />
          <span class="slider">
            <span class="knob"></span>
          </span>
        </label>
      </div>
    </aside>

    <!-- Main Workspace -->
    <main class="main-area">
      <header class="topbar">
        <div class="topbar-left">
          <button
            v-if="!isSidebarOpen"
            class="sidebar-toggle-btn expand-btn"
            title="Open sidebar"
            @click="toggleSidebar"
          >
            <svg class="panel-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
              <rect x="3" y="3" width="18" height="18" rx="2" />
              <path d="M9 3v18" />
            </svg>
            <span class="hover-arrow">»</span>
          </button>
        </div>
      </header>

      <!-- Welcome Landing Screen -->
      <div v-if="messages.length === 0" class="welcome-screen">
        <h1 class="welcome-title">How can we help you today?</h1>

        <div class="floating-pill-input">
          <input
            v-model="inputMessage"
            type="text"
            placeholder="Ask anything..."
            class="pill-input"
            @keydown.enter="handleSend"
          />
          <div class="pill-actions">
            <button class="pill-btn send-icon-btn" :disabled="!inputMessage.trim() || isGenerating" @click="handleSend">
              <svg class="send-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="5" y1="12" x2="19" y2="12" stroke-linecap="round" />
                <polyline points="12 5 19 12 12 19" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </button>
          </div>
        </div>
      </div>

      <!-- Chat Feed View with Gemini Word Fade-in -->
      <div v-else class="chat-workspace">
        <div class="chat-viewport">
          <div class="chat-container">
            <div
              v-for="(msg, index) in messages"
              :key="index"
              class="message-row"
              :class="msg.role"
            >
              <div class="message-body">
                <div class="author-name">
                  {{ msg.role === 'user' ? 'You' : 'Support Assistant' }}
                </div>
                <div class="bubble">
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
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Sticky Bottom Chat Input -->
        <footer class="bottom-input-area">
          <div class="floating-pill-input chat-bottom-input">
            <input
              v-model="inputMessage"
              type="text"
              placeholder="Ask a follow-up question..."
              class="pill-input"
              :disabled="isGenerating"
              @keydown.enter="handleSend"
            />
            <button class="pill-btn send-icon-btn" :disabled="!inputMessage.trim() || isGenerating" @click="handleSend">
              <svg class="send-svg-icon" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2">
                <line x1="5" y1="12" x2="19" y2="12" stroke-linecap="round" />
                <polyline points="12 5 19 12 12 19" stroke-linecap="round" stroke-linejoin="round" />
              </svg>
            </button>
          </div>
        </footer>
      </div>
    </main>
  </div>
</template>

<style scoped>
/* Base Theme Variables */
.app-layout {
  --bg-main: radial-gradient(circle at 50% 50%, #f0f6ff 0%, #e8f1fd 40%, #f8fafc 100%);
  --sidebar-bg: #f0f4f9;
  --sidebar-border: #e1e3e1;
  --text-primary: #1f1f1f;
  --text-secondary: #444746;
  --btn-bg: #dde3ea;
  --btn-hover: #d3dbdf;
  --search-bg: #e4ebf3;
  --item-hover: #e1e5ea;
  --item-active-bg: #c2e7ff;
  --item-active-text: #001d35;
  --card-bg: #ffffff;
  --card-border: #e1e3e1;
  --user-bubble-bg: #c2e7ff;
  --user-bubble-text: #001d35;
  --assistant-bubble-bg: #ffffff;
  --assistant-bubble-text: #1f1f1f;
  --shadow: rgba(0, 0, 0, 0.04);
  --switch-track: #dde3ea;
  --switch-knob: #ffffff;

  height: 100vh;
  width: 100vw;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, sans-serif;
  color: var(--text-primary);
  background: var(--bg-main);
  display: flex;
  overflow: hidden;
  transition: background 0.3s ease, color 0.3s ease;
}

/* Dark Theme Variables */
.app-layout.dark-theme {
  --bg-main: #1e1e24;
  --sidebar-bg: #18181c;
  --sidebar-border: #2e2e34;
  --text-primary: #ececf1;
  --text-secondary: #9e9ea8;
  --btn-bg: #2d2d35;
  --btn-hover: #3a3a45;
  --search-bg: #282830;
  --item-hover: #2a2a32;
  --item-active-bg: #343541;
  --item-active-text: #ffffff;
  --card-bg: #2a2a32;
  --card-border: #383842;
  --user-bubble-bg: #2b5278;
  --user-bubble-text: #ffffff;
  --assistant-bubble-bg: #2a2a32;
  --assistant-bubble-text: #d1d5db;
  --shadow: rgba(0, 0, 0, 0.2);
  --switch-track: #3a3a45;
  --switch-knob: #10a37f;
}

/* Sidebar Layout */
.sidebar {
  width: 260px;
  background-color: var(--sidebar-bg);
  display: flex;
  flex-direction: column;
  padding: 12px;
  border-right: 1px solid var(--sidebar-border);
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), padding 0.3s ease, opacity 0.2s ease;
  overflow: hidden;
  white-space: nowrap;
  flex-shrink: 0;
}

.sidebar.closed {
  width: 0;
  padding: 12px 0;
  border-right: none;
  opacity: 0;
}

.sidebar-top {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 12px;
  height: 36px;
}

/* Expanding Circular Search Button to Pill */
.expanding-search-container {
  display: flex;
  align-items: center;
  width: 36px;
  height: 36px;
  border-radius: 18px;
  background-color: transparent;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), background-color 0.2s ease;
  overflow: hidden;
}

.expanding-search-container:hover {
  background-color: var(--item-hover);
}

.expanding-search-container.expanded {
  width: 180px;
  background-color: var(--search-bg);
}

.expanding-search-input {
  width: 0;
  opacity: 0;
  border: none;
  outline: none;
  background: transparent;
  font-size: 13px;
  color: var(--text-primary);
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.2s ease;
  padding-right: 10px;
}

.expanding-search-container.expanded .expanding-search-input {
  width: 130px;
  opacity: 1;
}

.new-chat-wrapper {
  margin-bottom: 16px;
}

.new-chat-btn {
  width: 100%;
  padding: 10px 14px;
  background-color: var(--card-bg);
  border: 1px solid var(--card-border);
  border-radius: 20px;
  color: var(--text-primary);
  font-size: 14px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 10px;
  cursor: pointer;
  box-shadow: 0 1px 3px var(--shadow);
  transition: background-color 0.2s ease, border-color 0.2s ease;
}

.new-chat-btn:hover {
  background-color: var(--btn-hover);
}

.add-icon {
  width: 16px;
  height: 16px;
}

.icon-btn {
  width: 36px;
  height: 36px;
  background: transparent;
  border: none;
  border-radius: 50%;
  color: var(--text-secondary);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
}

.svg-icon {
  width: 18px;
  height: 18px;
}

.sidebar-toggle-btn {
  width: 36px;
  height: 36px;
  background: transparent;
  border: none;
  border-radius: 50%;
  color: var(--text-secondary);
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  position: relative;
  flex-shrink: 0;
}

.sidebar-toggle-btn:hover {
  background-color: var(--item-hover);
  color: var(--text-primary);
}

.panel-icon {
  width: 20px;
  height: 20px;
  transition: opacity 0.2s ease;
}

.hover-arrow {
  position: absolute;
  font-size: 18px;
  font-weight: bold;
  opacity: 0;
  transition: opacity 0.2s ease;
}

.sidebar-toggle-btn:hover .panel-icon {
  opacity: 0;
}

.sidebar-toggle-btn:hover .hover-arrow {
  opacity: 1;
}

.sidebar-history {
  flex: 1;
  overflow-y: auto;
}

.history-label {
  font-size: 11px;
  font-weight: 600;
  color: var(--text-secondary);
  padding: 6px 8px;
  text-transform: uppercase;
}

.history-item {
  display: flex;
  align-items: center;
  padding: 8px 12px;
  border-radius: 18px;
  font-size: 14px;
  color: var(--text-secondary);
  cursor: pointer;
}

.history-item:hover {
  background-color: var(--item-hover);
  color: var(--text-primary);
}

.history-item.active {
  background-color: var(--item-active-bg);
  color: var(--item-active-text);
  font-weight: 500;
}

.sidebar-footer {
  padding-top: 12px;
  border-top: 1px solid var(--sidebar-border);
  display: flex;
  align-items: center;
}

.switch-container {
  display: flex;
  align-items: center;
  justify-content: space-between;
  width: 100%;
  padding: 4px 6px;
  cursor: pointer;
  user-select: none;
}

.switch-mode-icon {
  display: flex;
  align-items: center;
  gap: 8px;
}

.theme-svg-icon {
  width: 18px;
  height: 18px;
  color: var(--text-secondary);
}

.switch-label {
  font-size: 13px;
  color: var(--text-secondary);
  font-weight: 500;
}

.switch-container input {
  display: none;
}

.slider {
  width: 44px;
  height: 24px;
  background-color: var(--switch-track);
  border-radius: 12px;
  position: relative;
  transition: background-color 0.3s ease;
  display: flex;
  align-items: center;
  padding: 2px;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
}

.knob {
  width: 20px;
  height: 20px;
  background-color: var(--switch-knob);
  border-radius: 50%;
  transition: transform 0.3s cubic-bezier(0.4, 0, 0.2, 1);
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.2);
}

.switch-container input:checked + .slider .knob {
  transform: translateX(20px);
}

/* Main Workspace */
.main-area {
  flex: 1;
  display: flex;
  flex-direction: column;
  height: 100vh;
  position: relative;
}

.topbar {
  height: 56px;
  display: flex;
  align-items: center;
  padding: 0 20px;
}

.welcome-screen {
  flex: 1;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  padding-bottom: 120px;
}

.welcome-title {
  font-size: 32px;
  font-weight: 400;
  color: var(--text-primary);
  margin-bottom: 32px;
}

.floating-pill-input {
  width: 90%;
  max-width: 680px;
  background: var(--card-bg);
  border: 1px solid var(--card-border);
  border-radius: 28px;
  box-shadow: 0 4px 16px var(--shadow);
  display: flex;
  align-items: center;
  padding: 6px 16px;
  gap: 10px;
}

.pill-input {
  flex: 1;
  border: none;
  outline: none;
  font-size: 16px;
  background: transparent;
  color: var(--text-primary);
}

.pill-btn {
  background: transparent;
  border: none;
  font-size: 18px;
  color: var(--text-secondary);
  cursor: pointer;
  padding: 6px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
}

.pill-btn:hover {
  background-color: var(--item-hover);
}

.send-svg-icon {
  width: 18px;
  height: 18px;
}

.send-icon-btn:disabled {
  opacity: 0.3;
  cursor: not-allowed;
}

.chat-workspace {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.chat-viewport {
  flex: 1;
  overflow-y: auto;
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
  margin-bottom: 24px;
  width: 100%;
}

.message-row.user {
  justify-content: flex-end;
}

.message-row.assistant {
  justify-content: flex-start;
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

.bubble {
  font-size: 15px;
  line-height: 1.6;
  padding: 12px 18px;
  border-radius: 20px;
  box-shadow: 0 2px 6px var(--shadow);
  word-break: break-word;
}

.user .bubble {
  background: var(--user-bubble-bg);
  color: var(--user-bubble-text);
  border-bottom-right-radius: 4px;
}

.assistant .bubble {
  background: var(--assistant-bubble-bg);
  color: var(--assistant-bubble-text);
  border-bottom-left-radius: 4px;
}

/* Gemini-style Word Fade-in Keyframe Animation */
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
  max-width: 768px;
}
</style>