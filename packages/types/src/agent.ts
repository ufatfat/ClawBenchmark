export type AgentCategory =
  | 'general'
  | 'task-automation'
  | 'code-generation'
  | 'data-analysis'
  | 'customer-service'
  | 'research'
  | 'devops'

export type AgentStatus = 'pending' | 'approved' | 'rejected' | 'archived'

export interface AgentScores {
  functionality: number
  performance: number
  security: number
}

export interface Agent {
  id: string
  name: string
  slug: string
  logo_url: string
  description: string
  full_description: string
  website_url: string
  repo_url: string
  license: string
  version: string
  category: AgentCategory
  supported_platforms: string[]
  supported_llms: string[]
  architecture_type: string
  is_open_source: boolean
  is_self_hosted: boolean
  tags: string[]
  submitted_by: string
  status: AgentStatus
  overall_score: number
  scores: AgentScores
  created_at: string
  updated_at: string
}
