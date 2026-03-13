<template>
  <el-card class="benchmark-card" shadow="hover" @click="emit('click', benchmark)">
    <div class="card-header">
      <el-tag :type="statusType" size="small">{{ statusLabel }}</el-tag>
      <span class="category">{{ benchmark.category }}</span>
    </div>

    <h3 class="card-title">{{ benchmark.name }}</h3>
    <p class="card-desc">{{ benchmark.description }}</p>

    <div class="card-footer">
      <div class="score" v-if="benchmark.score !== null">
        <span class="score-value">{{ benchmark.score.toFixed(1) }}</span>
        <span class="score-label">分</span>
      </div>
      <span class="date">{{ formatDate(benchmark.createdAt) }}</span>
    </div>
  </el-card>
</template>

<script setup lang="ts">
  import { computed } from 'vue'
  import type { Benchmark } from '@/types'

  const props = defineProps<{ benchmark: Benchmark }>()
  const emit = defineEmits<{ click: [benchmark: Benchmark] }>()

  const statusMap = {
    pending: { type: 'info', label: '待运行' },
    running: { type: 'warning', label: '运行中' },
    completed: { type: 'success', label: '已完成' },
    failed: { type: 'danger', label: '失败' },
  } as const

  const statusType = computed(() => statusMap[props.benchmark.status].type)
  const statusLabel = computed(() => statusMap[props.benchmark.status].label)

  function formatDate(dateStr: string) {
    return new Date(dateStr).toLocaleDateString('zh-CN')
  }
</script>

<style scoped>
  .benchmark-card {
    cursor: pointer;
    transition: transform 0.2s;
  }

  .benchmark-card:hover {
    transform: translateY(-2px);
  }

  .card-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 12px;
  }

  .category {
    font-size: 12px;
    color: #909399;
  }

  .card-title {
    font-size: 16px;
    font-weight: 600;
    margin: 0 0 8px;
    color: #303133;
  }

  .card-desc {
    font-size: 13px;
    color: #606266;
    margin: 0 0 16px;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .card-footer {
    display: flex;
    justify-content: space-between;
    align-items: flex-end;
  }

  .score-value {
    font-size: 24px;
    font-weight: 700;
    color: #409eff;
  }

  .score-label {
    font-size: 13px;
    color: #909399;
    margin-left: 2px;
  }

  .date {
    font-size: 12px;
    color: #c0c4cc;
  }
</style>
