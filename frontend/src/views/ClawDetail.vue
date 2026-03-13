<template>
  <div class="claw-detail">
    <div v-if="loading" class="loading">加载中...</div>
    <div v-else-if="error" class="error">{{ error }}</div>

    <div v-else-if="claw" class="content">
      <header class="header">
        <button @click="goBack" class="back-btn">← 返回</button>
        <div class="actions">
          <button @click="addToCompare" class="compare-btn">添加到对比</button>
        </div>
      </header>

      <section class="basic-info">
        <h1>{{ claw.name }}</h1>
        <p class="version">版本 {{ claw.version }}</p>
        <p class="description">{{ claw.description }}</p>
        <div class="overall-score">
          <span class="score-value">{{ claw.overallScore.toFixed(1) }}</span>
          <span class="score-label">综合评分</span>
        </div>
      </section>

      <section class="radar-section">
        <h2>多维度评分</h2>
        <div ref="radarChart" class="chart"></div>
      </section>

      <section class="metrics-table">
        <h2>详细指标</h2>
        <table>
          <thead>
            <tr>
              <th>指标</th>
              <th>分数</th>
              <th>等级</th>
            </tr>
          </thead>
          <tbody>
            <tr>
              <td>准确性</td>
              <td>{{ claw.metrics.accuracy.toFixed(2) }}</td>
              <td>{{ getGrade(claw.metrics.accuracy) }}</td>
            </tr>
            <tr>
              <td>速度</td>
              <td>{{ claw.metrics.speed.toFixed(2) }}</td>
              <td>{{ getGrade(claw.metrics.speed) }}</td>
            </tr>
            <tr>
              <td>稳定性</td>
              <td>{{ claw.metrics.stability.toFixed(2) }}</td>
              <td>{{ getGrade(claw.metrics.stability) }}</td>
            </tr>
            <tr>
              <td>资源使用</td>
              <td>{{ claw.metrics.resourceUsage.toFixed(2) }}</td>
              <td>{{ getGrade(claw.metrics.resourceUsage) }}</td>
            </tr>
            <tr>
              <td>可扩展性</td>
              <td>{{ claw.metrics.scalability.toFixed(2) }}</td>
              <td>{{ getGrade(claw.metrics.scalability) }}</td>
            </tr>
          </tbody>
        </table>
      </section>

      <section class="trend-section">
        <h2>历史趋势</h2>
        <p class="placeholder">趋势图功能开发中...</p>
      </section>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { clawApi } from '@/api/claw'
import type { Claw } from '@/types/claw'
import * as echarts from 'echarts'

const route = useRoute()
const router = useRouter()

const claw = ref<Claw | null>(null)
const loading = ref(true)
const error = ref<string | null>(null)
const radarChart = ref<HTMLElement>()
let chartInstance: echarts.ECharts | null = null

async function fetchDetail() {
  loading.value = true
  error.value = null
  try {
    const response = await clawApi.getDetail(route.params.id as string)
    claw.value = response.data
    await initRadarChart()
  } catch (e) {
    error.value = e instanceof Error ? e.message : '加载失败'
  } finally {
    loading.value = false
  }
}

async function initRadarChart() {
  if (!radarChart.value || !claw.value) return

  chartInstance = echarts.init(radarChart.value)

  const option = {
    radar: {
      indicator: [
        { name: '准确性', max: 100 },
        { name: '速度', max: 100 },
        { name: '稳定性', max: 100 },
        { name: '资源使用', max: 100 },
        { name: '可扩展性', max: 100 }
      ]
    },
    series: [{
      type: 'radar',
      data: [{
        value: [
          claw.value.metrics.accuracy,
          claw.value.metrics.speed,
          claw.value.metrics.stability,
          claw.value.metrics.resourceUsage,
          claw.value.metrics.scalability
        ],
        name: claw.value.name
      }]
    }]
  }

  chartInstance.setOption(option)
}

function getGrade(score: number): string {
  if (score >= 90) return 'A'
  if (score >= 80) return 'B'
  if (score >= 70) return 'C'
  if (score >= 60) return 'D'
  return 'F'
}

function goBack() {
  router.push('/')
}

function addToCompare() {
  if (claw.value) {
    const compareIds = JSON.parse(localStorage.getItem('compareIds') || '[]')
    if (!compareIds.includes(claw.value.id)) {
      compareIds.push(claw.value.id)
      localStorage.setItem('compareIds', JSON.stringify(compareIds))
      alert('已添加到对比列表')
    } else {
      alert('已在对比列表中')
    }
  }
}

onMounted(() => {
  fetchDetail()
  window.addEventListener('resize', () => {
    chartInstance?.resize()
  })
})

onUnmounted(() => {
  chartInstance?.dispose()
})
</script>

<style scoped>
.claw-detail {
  max-width: 1200px;
  margin: 0 auto;
  padding: 2rem;
}

.loading, .error {
  text-align: center;
  padding: 2rem;
  font-size: 1.1rem;
}

.error {
  color: #dc3545;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.back-btn, .compare-btn {
  padding: 0.5rem 1rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  background: white;
  cursor: pointer;
  font-size: 0.9rem;
}

.back-btn:hover, .compare-btn:hover {
  background: #f8f9fa;
}

.compare-btn {
  background: #007bff;
  color: white;
  border-color: #007bff;
}

.compare-btn:hover {
  background: #0056b3;
}

.basic-info {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 2rem;
  margin-bottom: 2rem;
}

.basic-info h1 {
  margin: 0 0 0.5rem 0;
  font-size: 2rem;
}

.version {
  color: #666;
  margin: 0 0 1rem 0;
}

.description {
  color: #333;
  line-height: 1.6;
  margin-bottom: 1.5rem;
}

.overall-score {
  text-align: center;
  padding: 1.5rem;
  background: #f8f9fa;
  border-radius: 6px;
}

.score-value {
  display: block;
  font-size: 3rem;
  font-weight: bold;
  color: #007bff;
}

.score-label {
  display: block;
  font-size: 1rem;
  color: #666;
  margin-top: 0.5rem;
}

.radar-section, .metrics-table, .trend-section {
  background: white;
  border: 1px solid #e0e0e0;
  border-radius: 8px;
  padding: 2rem;
  margin-bottom: 2rem;
}

h2 {
  margin: 0 0 1.5rem 0;
  font-size: 1.5rem;
}

.chart {
  width: 100%;
  height: 400px;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th, td {
  padding: 1rem;
  text-align: left;
  border-bottom: 1px solid #e0e0e0;
}

th {
  background: #f8f9fa;
  font-weight: 600;
}

tr:last-child td {
  border-bottom: none;
}

.placeholder {
  color: #999;
  text-align: center;
  padding: 2rem;
}
</style>
