<script setup lang="ts">
import { ref, computed } from 'vue'
import { Search, Plus, Fold, ChatDotRound } from '@element-plus/icons-vue'

interface Session {
  id: number
  title: string
}

const props = defineProps<{
  isOpen: boolean
  sessions: Session[]
  currentId: number | null
}>()

const emit = defineEmits<{
  (e: 'update:isOpen', value: boolean): void
  (e: 'select-session', id: number): void
  (e: 'new-session'): void
}>()

const searchQuery = ref('')
const isSearching = ref(false)

const filteredSessions = computed(() => {
  if (!searchQuery.value.trim()) return props.sessions
  return props.sessions.filter(s =>
    s.title.toLowerCase().includes(searchQuery.value.toLowerCase())
  )
})

function toggleSearch() {
  isSearching.value = !isSearching.value
  if (!isSearching.value) {
    searchQuery.value = ''
  }
}
</script>

<template>
  <aside class="sidebar-aside" :class="{ closed: !isOpen }">
    <!-- Top Action Controls -->
    <div class="sidebar-top">
      <div class="expanding-search" :class="{ expanded: isSearching }">
        <el-tooltip content="Search conversations" placement="bottom" :disabled="isSearching">
          <el-button
            circle
            text
            class="icon-btn"
            @click="toggleSearch"
          >
            <el-icon><Search /></el-icon>
          </el-button>
        </el-tooltip>
        <el-input
          v-model="searchQuery"
          placeholder="Search history..."
          class="search-input"
          size="small"
          clearable
        />
      </div>

      <el-tooltip content="Close sidebar" placement="bottom">
        <el-button
          circle
          text
          class="icon-btn"
          @click="emit('update:isOpen', false)"
        >
          <el-icon><Fold /></el-icon>
        </el-button>
      </el-tooltip>
    </div>

    <!-- New Chat Action -->
    <div class="new-chat-wrapper">
      <el-button
        class="new-chat-btn"
        round
        type="primary"
        plain
        @click="emit('new-session')"
      >
        <el-icon><Plus /></el-icon>
        <span>New chat</span>
      </el-button>
    </div>

    <!-- Scrollable Conversation History -->
    <div class="sidebar-history-container">
      <div class="history-label">Recent Conversations</div>
      <el-scrollbar class="sidebar-scrollbar">
        <div
          v-for="session in filteredSessions"
          :key="session.id"
          class="history-item"
          :class="{ active: session.id === currentId }"
          @click="emit('select-session', session.id)"
        >
          <el-icon class="item-icon"><ChatDotRound /></el-icon>
          <span class="chat-title">{{ session.title }}</span>
        </div>
      </el-scrollbar>
    </div>
  </aside>
</template>

<style scoped>
.sidebar-aside {
  width: 260px;
  background-color: var(--sidebar-bg);
  display: flex;
  flex-direction: column;
  padding: 12px;
  border-right: 1px solid var(--sidebar-border);
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.2s ease, padding 0.3s ease;
  overflow: hidden;
  white-space: nowrap;
  flex-shrink: 0;
}

.sidebar-aside.closed {
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

.expanding-search {
  display: flex;
  align-items: center;
  width: 36px;
  height: 36px;
  border-radius: 18px;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), background-color 0.2s ease;
  overflow: hidden;
}

.expanding-search:hover {
  background-color: var(--item-hover);
}

.expanding-search.expanded {
  width: 180px;
  background-color: var(--search-bg);
}

.search-input {
  width: 0;
  opacity: 0;
  transition: width 0.3s cubic-bezier(0.4, 0, 0.2, 1), opacity 0.2s ease;
}

.expanding-search.expanded .search-input {
  width: 130px;
  opacity: 1;
}

:deep(.search-input .el-input__wrapper) {
  background: transparent;
  box-shadow: none !important;
}

.new-chat-wrapper {
  margin-bottom: 16px;
}

.new-chat-btn {
  width: 100%;
  justify-content: flex-start;
  padding: 10px 16px;
  background-color: var(--card-bg);
  border: 1px solid var(--card-border);
  color: var(--text-primary);
  box-shadow: 0 1px 3px var(--shadow);
}

.icon-btn {
  color: var(--text-secondary);
  font-size: 16px;
}

.sidebar-history-container {
  flex: 1;
  display: flex;
  flex-direction: column;
  overflow: hidden;
}

.history-label {
  font-size: 11px;
  font-weight: 600;
  color: var(--text-secondary);
  padding: 6px 8px;
  text-transform: uppercase;
}

.sidebar-scrollbar {
  flex: 1;
}

.history-item {
  display: flex;
  align-items: center;
  gap: 8px;
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

.item-icon {
  font-size: 14px;
}
</style>