<script setup lang="ts">
import { ref } from 'vue'
import { FolderAdd } from '@element-plus/icons-vue'
import VectorUploadModal from './VectorUploadModal.vue'

const dialogVisible = ref(false)
</script>

<template>
  <el-button
    type="primary"
    plain
    size="small"
    :icon="FolderAdd"
    @click="dialogVisible = true"
  >
    Vector Database
  </el-button>

  <VectorUploadModal v-model:visible="dialogVisible" />
</template>