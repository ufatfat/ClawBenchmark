import axios from 'axios'
import type { Skill, SkillListParams, SkillListResponse } from '../types/skill'

const API_BASE = import.meta.env.VITE_API_BASE || '/api'

export const skillApi = {
  async getList(params: SkillListParams = {}): Promise<SkillListResponse> {
    const { data } = await axios.get(`${API_BASE}/skills`, { params })
    return data
  },

  async getDetail(id: string): Promise<Skill> {
    const { data } = await axios.get(`${API_BASE}/skills/${id}`)
    return data
  },

  async getCategories(): Promise<string[]> {
    const { data } = await axios.get(`${API_BASE}/skills/categories`)
    return data
  }
}
