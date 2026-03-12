export type UserRole = 'user' | 'developer' | 'admin'

export interface User {
  id: string
  email: string
  username: string
  display_name: string
  avatar_url: string
  role: UserRole
  github_id: string
  bio: string
  company: string
  is_verified: boolean
  created_at: string
  updated_at: string
}

export type CommentStatus = 'active' | 'hidden' | 'deleted'

export interface Comment {
  id: string
  agent_id: string
  user_id: string
  parent_id: string | null
  content: string
  upvotes: number
  status: CommentStatus
  created_at: string
  updated_at: string
}
