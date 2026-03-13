import http from '@/utils/http'
import type { User } from '@/types'

export const authApi = {
  login: (data: { username: string; password: string }) =>
    http.post<{ token: string; user: User }>('/auth/login', data),

  logout: () => http.post('/auth/logout'),

  profile: () => http.get<User>('/auth/profile'),
}
