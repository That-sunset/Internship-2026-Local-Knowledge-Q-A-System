<script setup lang="ts">
import { Promotion } from '@element-plus/icons-vue'

defineProps<{
  disabled?: boolean
}>()

const emit = defineEmits<{
  (e: 'click'): void
}>()
</script>

<template>
  <el-button
    circle
    type="primary"
    class="send-btn"
    :disabled="disabled"
    @click="emit('click')"
  >
    <el-icon><Promotion /></el-icon>
  </el-button>
</template>

<style scoped>
.send-btn {
  position: absolute;
  right: 8px;
  bottom: 6px;
  z-index: 2;
  border: none;
}
</style>