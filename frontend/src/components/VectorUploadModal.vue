<script setup lang="ts">
import { ref } from 'vue'
import { UploadFilled, Document, Delete } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import type { UploadFile, UploadFiles } from 'element-plus'
import { UploadVectorFiles } from '../Api/record'

const visible = defineModel<boolean>('visible', { required: true })

const fileList = ref<UploadFile[]>([])
const uploading = ref(false)

function handleFileChange(uploadFile: UploadFile, uploadFiles: UploadFiles) {
  const fileName = uploadFile.name.toLowerCase()
  const mimeType = uploadFile.raw?.type
  const isSupported =
    fileName.endsWith('.txt') ||
    fileName.endsWith('.csv') ||
    mimeType === 'text/plain' ||
    mimeType === 'text/csv'

  if (!isSupported) {
    ElMessage.warning('Only .txt and .csv files are supported')
    fileList.value = uploadFiles.filter(f => f.uid !== uploadFile.uid)
    return
  }

  fileList.value = uploadFiles
}

function handleRemove(uid: number) {
  fileList.value = fileList.value.filter(f => f.uid !== uid)
}

function closeDialog() {
  if (uploading.value) return
  visible.value = false
}

async function handleUploadToVectorDB() {
  const rawFiles = fileList.value
    .map(f => f.raw)
    .filter((f): f is NonNullable<typeof f> => Boolean(f))

  if (rawFiles.length === 0) {
    ElMessage.warning('Please select at least one .txt or .csv file')
    return
  }

  uploading.value = true
  try {
    const res = await UploadVectorFiles(rawFiles)
    ElMessage.success(
      `Uploaded ${res.files} file(s) (${res.chunks} chunks) to Vector Database`
    )
    fileList.value = []
    visible.value = false
  } catch (err: any) {
    ElMessage.error(err?.message || 'Failed to upload files to Vector Database')
  } finally {
    uploading.value = false
  }
}
</script>

<template>
  <el-dialog
    v-model="visible"
    title="Upload Files to Vector Database"
    width="480px"
    append-to-body
    :close-on-click-modal="!uploading"
    @close="closeDialog"
  >
    <el-upload
      drag
      multiple
      accept=".txt,.csv,text/plain,text/csv"
      :auto-upload="false"
      :show-file-list="false"
      :file-list="fileList"
      :on-change="handleFileChange"
    >
      <el-icon class="el-icon--upload"><UploadFilled /></el-icon>
      <div class="el-upload__text">
        Drop .txt or .csv files here, or <em>click to browse</em>
      </div>
    </el-upload>

    <el-space
      v-if="fileList.length > 0"
      direction="vertical"
      fill
      :size="6"
      class="selected-files"
    >
      <el-card
        v-for="file in fileList"
        :key="file.uid"
        shadow="never"
        body-style="padding: 8px 12px; display: flex; align-items: center; justify-content: space-between;"
      >
        <el-space :size="6">
          <el-icon><Document /></el-icon>
          <el-text truncated>{{ file.name }}</el-text>
        </el-space>
        <el-button
          type="danger"
          text
          size="small"
          :icon="Delete"
          :disabled="uploading"
          @click="handleRemove(file.uid)"
        />
      </el-card>
    </el-space>

    <template #footer>
      <el-button :disabled="uploading" @click="closeDialog">Cancel</el-button>
      <el-button
        type="primary"
        :loading="uploading"
        :disabled="fileList.length === 0"
        @click="handleUploadToVectorDB"
      >
        Upload
      </el-button>
    </template>
  </el-dialog>
</template>

<style scoped>
.selected-files {
  width: 100%;
  margin-top: 14px;
  max-height: 160px;
  overflow-y: auto;
}
</style>