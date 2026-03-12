<script setup lang="ts">
import { computed } from 'vue'
import { useRouter } from 'vue-router'
import {
  NButton,
  NCard,
  NGrid,
  NGridItem,
  NTag,
  NThing,
  NText,
} from 'naive-ui'
import { agents, benchmarks, users } from '@clawbenchmark/mock-data'

const router = useRouter()

const heroActions = {
  browseAgents: () => router.push({ name: 'AgentList' }),
  viewLeaderboard: () => router.push({ name: 'Leaderboard' }),
}

const stats = computed(() => {
  const benchmarkAgentCount = new Set(benchmarks.map((item) => item.agent_id)).size

  return [
    { label: 'Agent 数量', value: agents.length, suffix: '个' },
    { label: '测评报告数', value: benchmarks.length, suffix: '份' },
    { label: '注册用户数', value: users.length, suffix: '人' },
    { label: '参与测评 Agent', value: benchmarkAgentCount, suffix: '个' },
  ]
})

const topAgents = computed(() => {
  return [...agents]
    .sort((a, b) => b.overall_score - a.overall_score)
    .slice(0, 6)
})

const latestBenchmarks = computed(() => {
  const agentMap = new Map(agents.map((agent) => [agent.id, agent]))

  return [...benchmarks]
    .sort(
      (a, b) =>
        new Date(b.created_at).getTime() - new Date(a.created_at).getTime(),
    )
    .slice(0, 5)
    .map((benchmark) => ({
      ...benchmark,
      agentName: agentMap.get(benchmark.agent_id)?.name ?? 'Unknown Agent',
    }))
})
</script>

<template>
  <div class="home-page">
    <section class="hero">
      <n-tag round type="success" size="small">AI Agent Benchmarking</n-tag>
      <h1>ClawBenchmark：客观、透明、可复现的 Agent 测评平台</h1>
      <p>
        聚合主流 AI Agent 的评测数据、能力画像与趋势变化，帮助团队快速做出技术选型。
      </p>
      <div class="hero-actions">
        <n-button type="primary" size="large" @click="heroActions.browseAgents">
          浏览 Agent
        </n-button>
        <n-button size="large" @click="heroActions.viewLeaderboard">
          查看排行榜
        </n-button>
      </div>
    </section>

    <section>
      <h2 class="section-title">平台数据总览</h2>
      <n-grid cols="1 s:2 m:4" responsive="screen" :x-gap="16" :y-gap="16">
        <n-grid-item v-for="stat in stats" :key="stat.label">
          <n-card size="small" class="stat-card">
            <n-text depth="3">{{ stat.label }}</n-text>
            <div class="stat-value">{{ stat.value }}<span>{{ stat.suffix }}</span></div>
          </n-card>
        </n-grid-item>
      </n-grid>
    </section>

    <section>
      <h2 class="section-title">热门 Agent（Top 6）</h2>
      <n-grid cols="1 s:2 m:3" responsive="screen" :x-gap="16" :y-gap="16">
        <n-grid-item v-for="agent in topAgents" :key="agent.id">
          <n-card hoverable class="agent-card" @click="router.push({ name: 'AgentDetail', params: { slug: agent.slug } })">
            <n-thing>
              <template #header>{{ agent.name }}</template>
              <template #description>
                <n-tag size="small" type="info">{{ agent.category }}</n-tag>
              </template>
              <p class="agent-description">{{ agent.description }}</p>
              <div class="agent-score">综合评分：{{ agent.overall_score }}</div>
            </n-thing>
          </n-card>
        </n-grid-item>
      </n-grid>
    </section>

    <section>
      <h2 class="section-title">最新测评动态</h2>
      <n-card>
        <ul class="benchmark-list">
          <li v-for="bench in latestBenchmarks" :key="bench.id" class="benchmark-item">
            <div>
              <strong>{{ bench.agentName }}</strong>
              <n-tag size="small" type="warning" class="dimension-tag">{{ bench.dimension }}</n-tag>
              <span class="benchmark-summary">{{ bench.summary }}</span>
            </div>
            <div class="benchmark-meta">
              <span>评分：{{ bench.overall_score }}</span>
              <span>{{ new Date(bench.created_at).toLocaleDateString() }}</span>
            </div>
          </li>
        </ul>
      </n-card>
    </section>
  </div>
</template>

<style scoped>
.home-page {
  display: flex;
  flex-direction: column;
  gap: 28px;
}

.hero {
  background: linear-gradient(120deg, #ecfdf5 0%, #f5f3ff 100%);
  border: 1px solid #e5e7eb;
  border-radius: 16px;
  padding: 24px;
}

.hero h1 {
  margin: 12px 0;
  font-size: 30px;
  line-height: 1.25;
}

.hero p {
  margin: 0;
  color: #4b5563;
  max-width: 760px;
}

.hero-actions {
  margin-top: 20px;
  display: flex;
  flex-wrap: wrap;
  gap: 12px;
}

.section-title {
  margin: 0 0 12px;
}

.stat-card {
  min-height: 110px;
}

.stat-value {
  margin-top: 8px;
  font-size: 32px;
  font-weight: 700;
  color: #111827;
}

.stat-value span {
  margin-left: 4px;
  font-size: 15px;
  color: #6b7280;
}

.agent-card {
  cursor: pointer;
  min-height: 190px;
}

.agent-description {
  color: #4b5563;
  margin: 10px 0;
}

.agent-score {
  font-weight: 600;
  color: #059669;
}

.benchmark-list {
  margin: 0;
  padding: 0;
  list-style: none;
  display: flex;
  flex-direction: column;
  gap: 12px;
}

.benchmark-item {
  display: flex;
  gap: 8px;
  justify-content: space-between;
  align-items: flex-start;
  border-bottom: 1px solid #f3f4f6;
  padding-bottom: 10px;
}

.benchmark-item:last-child {
  border-bottom: none;
  padding-bottom: 0;
}

.dimension-tag {
  margin-left: 8px;
}

.benchmark-summary {
  margin-left: 8px;
  color: #4b5563;
}

.benchmark-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 4px;
  color: #6b7280;
  white-space: nowrap;
  font-size: 13px;
}

@media (max-width: 768px) {
  .hero h1 {
    font-size: 24px;
  }

  .benchmark-item {
    flex-direction: column;
  }

  .benchmark-meta {
    align-items: flex-start;
  }
}
</style>
