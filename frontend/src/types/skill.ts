export interface Skill {
  id: string
  name: string
  version: string
  category: string
  description: string
  overallScore: number
  scores: {
    functionality: number
    performance: number
    reliability: number
    usability: number
    maintainability: number
  }
  supportedClaws: string[]
  dependencies: string[]
  createdAt: string
  updatedAt: string
}

export interface SkillListParams {
  category?: string
  search?: string
  sortBy?: 'name' | 'overallScore' | 'createdAt'
  sortOrder?: 'asc' | 'desc'
  page?: number
  pageSize?: number
}

export interface SkillListResponse {
  data: Skill[]
  total: number
  page: number
  pageSize: number
}
