<template>
  <div class="claw-compare">
    <header class="header">
      <h1>Claw 对比</h1>
      <button @click="goBack" class="back-btn">← 返回</button>
    </header>

    <div class="selector">
      <h3>选择要对比的 Claw（最多 4 个）</h3>
      <div class="checkbox-group">
        <label v-for="claw in allClaws" :key="claw.id" class="checkbox-item">
          <input
            type="checkbox"
            :value="claw.id"
            v-model="selectedIds"
            :disabled="selectedIds.length >= 4 && !selectedIds.includes(claw.id)"
          />
          <span>{{ claw.name }} (v{{ claw.version }})</span>
        </label>
      </div>
    </div>

    <div v-if="selectedClaws.length >= 2" class="comparison">
      <section class="chart-section">
        <h2>对比图表</h2>
        <div ref="compareChart" class="chart"></div>
      </section>

      <section class="table-section">
        <h2>详细对比</h2>
        <div class="table-wrapper">
          <table>
            <thead>
              <tr>
                <th>指标</th>
                <th v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.name }}
                </th>
              </tr>
            </thead>
            <tbody>
              <tr>
                <td class="metric-name">综合评分</td>
                <td v-for="claw in selectedClaws" :key="claw.id" class="score-cell">
                  {{ claw.overallScore.toFixed(1) }}
                </td>
              </tr>
              <tr>
                <td class="metric-name">准确性</td>
                <td v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.metrics.accuracy.toFixed(1) }}
                </td>
              </tr>
              <tr>
                <td class="metric-name">速度</td>
                <td v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.metrics.speed.toFixed(1) }}
                </td>
              </tr>
              <tr>
                <td class="metric-name">稳定性</td>
                <td v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.metrics.stability.toFixed(1) }}
                </td>
              </tr>
              <tr>
                <td class="metric-name">资源使用</td>
                <td v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.metrics.resourceUsage.toFixed(1) }}
                </td>
              </tr>
              <tr>
                <td class="metric-name">可扩展性</td>
                <td v-for="claw in selectedClaws" :key="claw.id">
                  {{ claw.metrics.scalability.toFixed(1) }}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </section>
    </div>

    <div v-else class="empty-state">
      <p>请至少选择 2 个 Claw 进行对比</p>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed, watch, onMounted, onUnmounted } from 'vue'
import { useRouter } from 'vue-router'
import { useClawStore } from '@/stores/claw'
import type { Claw } from '@/types/claw'
import * as echarts from 'echarts'

const router = useRouter()
const clawStore = useClawStore()

const selectedIds = ref<string[]>([])
const compareChart = ref<HTMLElement>()
let chartInstance: echarts.ECharts | null = null

const allClaws = computed(() => clawStore.claws)
const selectedClaws = computed(() =>
  allClaws.value.filter(c => selectedIds.value.includes(c.id))
)

function goBack() {
  router.push('/')
}

function initChart() {
  if (!compareChart.value || selectedClaws.value.length < 2) return

  if (!chartInstance) {
    chartInstance = echarts.init(compareChart.value)
  }

  const indicators = [
    { name: '准确性', max: 100 },
    { name: '速度', max: 100 },
    { name: '稳定性', max: 100 },
    { name: '资源使用', max: 100 },
    { name: '可扩展性', max: 100 }
  ]

  const seriesData = selectedClaws.value.map(claw => ({
    value: [
      claw.metrics.accuracy,
      claw.metrics.speed,
      claw.metrics.stability,
      claw.metrics.resourceUsage,
      claw.metrics.scalability
    ],
    name: claw.name
  }))

  const option = {
    legend: {
      data: selectedClaws.value.map(c => c.name)
    },
    radar: {
      indicator: indicators
    },
    series: [{
      type: 'radar',
      data: seriesData
    }]
  }

  chartInstance.setOption(option)
}

watch(selectedClaws, () => {
  if (selectedClaws.value.length >= 2) {
    setTimeout(initChart, 100)
  }
}, { deep: true })

onMounted(async () => {
  if (allClaws.value.length === 0) {
    await clawStore.fetchClaws()
  }

  const savedIds = JSON.parse(localStorage.getItem('compareIds') || '[]')
  selectedIds.value = savedIds.slice(0, 4)

  window.addEventListener('resize', () => {
    chartInstance?.resize()
  })
})

onUnmounted(() => {
  chartInstance?.dispose()
})
</script>

<style scoped>
.claw-compare {
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

.back-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
}

.back-btn:hover {
  background: #f8f9fa;
}

.selector {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 1.5rem;
  margin-bottom: 2rem;
}

.selector h3 {
  margin: 0 0 1rem 0;
  font-size: 1.1rem;
}

.checkbox-group {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(250px, 1fr));
  gap: 0.75rem;
}

.checkbox-item {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  cursor: pointer;
}

.checkbox-item input[type="checkbox"] {
  cursor: pointer;
}

.checkbox-item input[type="checkbox"]:disabled {
  cursor: not-allowed;
}

.comparison {
  display: flex;
  flex-direction: column;
  gap: 2rem;
}

.chart-section, .table-section {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 1.5rem;
}

.chart-section h2, .table-section h2 {
  margin: 0 0 1rem 0;
  font-size: 1.3rem;
}

.chart {
  width: 100%;
  height: 400px;
}

.table-wrapper {
  overflow-x: auto;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 0.75rem;
  text-align: left;
  border-bottom: 1px solid #e0e0e0;
}

th {
  background: #f8f9fa;
  font-weight: 600;
}

.metric-name {
  font-weight: 600;
  color: #666;
}

.score-cell {
  font-weight: bold;
  color: #007bff;
}

.empty-state {
  text-align: center;
  padding: 3rem;
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
}

.empty-state p {
  color: #666;
  font-size: 1.1rem;
}
</style>
