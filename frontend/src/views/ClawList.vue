<template>
  <div class="claw-list">
    <header class="header">
      <h1>Claw 测评</h1>
      <div class="sort-controls">
        <label>排序：</label>
        <select v-model="sortField" @change="handleSort">
          <option value="overallScore">综合评分</option>
          <option value="accuracy">准确性</option>
          <option value="speed">速度</option>
          <option value="stability">稳定性</option>
          <option value="resourceUsage">资源使用</option>
          <option value="scalability">可扩展性</option>
        </select>
        <button @click="toggleSortOrder" class="sort-btn">
          {{ sortOrder === 'desc' ? '↓' : '↑' }}
        </button>
      </div>
    </header>

    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else class="card-grid">
      <div
        v-for="claw in sortedClaws"
        :key="claw.id"
        class="card"
        @click="goToDetail(claw.id)"
      >
        <h3>{{ claw.name }}</h3>
        <p class="version">v{{ claw.version }}</p>
        <div class="score">
          <span class="score-value">{{ claw.overallScore.toFixed(1) }}</span>
          <span class="score-label">综合评分</span>
        </div>
        <div class="metrics">
          <div class="metric">
            <span class="metric-label">准确性</span>
            <span class="metric-value">{{ claw.metrics.accuracy.toFixed(1) }}</span>
          </div>
          <div class="metric">
            <span class="metric-label">速度</span>
            <span class="metric-value">{{ claw.metrics.speed.toFixed(1) }}</span>
          </div>
          <div class="metric">
            <span class="metric-label">稳定性</span>
            <span class="metric-value">{{ claw.metrics.stability.toFixed(1) }}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useClawStore } from '@/stores/claw'
import type { SortField, SortOrder } from '@/types/claw'

const router = useRouter()
const clawStore = useClawStore()

const sortField = ref<SortField>('overallScore')
const sortOrder = ref<SortOrder>('desc')

const loading = computed(() => clawStore.loading)
const error = computed(() => clawStore.error)
const sortedClaws = computed(() => clawStore.sortedClaws)

function handleSort() {
  clawStore.sortClaws(sortField.value, sortOrder.value)
}

function toggleSortOrder() {
  sortOrder.value = sortOrder.value === 'desc' ? 'asc' : 'desc'
  handleSort()
}

function goToDetail(id: string) {
  router.push(`/claw/${id}`)
}

onMounted(async () => {
  await clawStore.fetchClaws()
  handleSort()
})
</script>

<style scoped>
.claw-list {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

h1 {
  font-size: 2rem;
  margin: 0;
}

.sort-controls {
  display: flex;
  align-items: center;
  gap: 0.5rem;
}

select {
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 0.9rem;
}

.sort-btn {
  padding: 0.5rem 1rem;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1.2rem;
}

.sort-btn:hover {
  background: #0056b3;
}

.loading, .error {
  text-align: center;
  padding: 2rem;
  font-size: 1.1rem;
}

.error {
  color: #dc3545;
}

.card-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
}

.card {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 1.5rem;
  cursor: pointer;
  transition: all 0.2s;
}

.card:hover {
  box-shadow: 0 4px 12px rgba(0,0,0,0.1);
  transform: translateY(-2px);
}

.card h3 {
  margin: 0 0 0.5rem 0;
  font-size: 1.3rem;
}

.version {
  color: #666;
  font-size: 0.9rem;
  margin: 0 0 1rem 0;
}

.score {
  text-align: center;
  padding: 1rem;
  background: #f8f9fa;
  border-radius: 6px;
  margin-bottom: 1rem;
}

.score-value {
  display: block;
  font-size: 2.5rem;
  font-weight: bold;
  color: #007bff;
}

.score-label {
  display: block;
  font-size: 0.9rem;
  color: #666;
  margin-top: 0.25rem;
}

.metrics {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.metric {
  display: flex;
  justify-content: space-between;
  padding: 0.5rem;
  background: #f8f9fa;
  border-radius: 4px;
}

.metric-label {
  color: #666;
  font-size: 0.9rem;
}

.metric-value {
  font-weight: 600;
  color: #333;
}
</style>
