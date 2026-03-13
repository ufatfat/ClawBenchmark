import http from '@/utils/http'
import type { LeaderboardEntry, Model } from '@/types'

export const modelApi = {
  list: () => http.get<Model[]>('/models'),
  get: (id: number) => http.get<Model>(`/models/${id}`),
  leaderboard: () => http.get<LeaderboardEntry[]>('/leaderboard'),
}
