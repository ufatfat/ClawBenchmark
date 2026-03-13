// API 通用响应结构
export interface ApiResponse<T = unknown> {
  code: number
  message: string
  data: T
}

export interface PageResult<T> {
  items: T[]
  total: number
  page: number
  pageSize: number
}

// 测评模型
export interface Benchmark {
  id: number
  name: string
  description: string
  category: string
  status: 'pending' | 'running' | 'completed' | 'failed'
  score: number | null
  createdAt: string
  updatedAt: string
}

// 测评结果
export interface BenchmarkResult {
  id: number
  benchmarkId: number
  modelName: string
  score: number
  metrics: Record<string, number>
  duration: number
  createdAt: string
}

// 模型信息
export interface Model {
  id: number
  name: string
  provider: string
  version: string
  description: string
}

// 排行榜条目
export interface LeaderboardEntry {
  rank: number
  model: Model
  avgScore: number
  benchmarkScores: Record<string, number>
}

// 用户
export interface User {
  id: number
  username: string
  email: string
  role: 'admin' | 'user'
  createdAt: string
}
