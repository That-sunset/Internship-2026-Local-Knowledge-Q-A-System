<script setup lang="ts">
import { ref } from 'vue'

const connected = ref(false)
const errorMsg = ref('')

async function connect() {
  errorMsg.value = ''
  console.log('Button clicked! Fetching backend...')

  try {
    const res = await fetch('http://127.0.0.1:8080/api/info')
    console.log('Response received, status:', res.status)

    if (res.ok) {
      connected.value = true
    } else {
      errorMsg.value = `Server error: ${res.status}`
    }
  } catch (err: any) {
    console.error('Fetch failed:', err)
    errorMsg.value = 'Cannot reach backend'
    connected.value = false
  }
}
</script>

<template>
  <div style="position: absolute; top: 12px; right: 12px; z-index: 9999; display: flex; align-items: center; gap: 8px; background: rgba(255,255,255,0.9); padding: 6px; border-radius: 4px;">
    <button @click="connect">Communicate with Backend</button>
    <span v-if="connected" style="color: green; font-weight: bold;">backend connected</span>
    <span v-if="errorMsg" style="color: red; font-weight: bold;">{{ errorMsg }}</span>
  </div>
</template>