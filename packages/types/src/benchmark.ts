export type BenchmarkDimension = 'functionality' | 'performance' | 'security'

export type BenchmarkTaskStatus =
  | 'pending'
  | 'running'
  | 'completed'
  | 'failed'
  | 'cancelled'

export interface BenchmarkEnvironment {
  os: string
  cpu: string
  memory: string
  llm_provider: string
  llm_model: string
}

export interface MetricResult {
  metric_name: string
  value: number
  unit: string
  score: number
  details?: Record<string, unknown>
}

export interface BenchmarkResult {
  id: string
  task_id: string
  agent_id: string
  dimension: BenchmarkDimension
  overall_score: number
  metrics: MetricResult[]
  raw_data?: Record<string, unknown>
  summary: string
  created_at: string
}

export interface BenchmarkTask {
  id: string
  agent_id: string
  agent_version: string
  test_suite_id: string
  dimension: BenchmarkDimension
  status: BenchmarkTaskStatus
  progress: number
  started_at: string
  completed_at: string
  created_by: string
  environment: BenchmarkEnvironment
  results?: BenchmarkResult
  created_at: string
}
