export interface ClawMetrics {
  accuracy: number
  speed: number
  stability: number
  resourceUsage: number
  scalability: number
}

export interface Claw {
  id: string
  name: string
  version: string
  description: string
  overallScore: number
  metrics: ClawMetrics
  createdAt: string
  updatedAt: string
}

export interface ClawListResponse {
  data: Claw[]
  total: number
}

export interface ClawDetailResponse {
  data: Claw
}

export type SortField = 'overallScore' | 'accuracy' | 'speed' | 'stability' | 'resourceUsage' | 'scalability'
export type SortOrder = 'asc' | 'desc'
