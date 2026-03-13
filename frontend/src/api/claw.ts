import request from './request'
import type { Claw, ClawListResponse, ClawDetailResponse } from '@/types/claw'

export const clawApi = {
  getList(): Promise<ClawListResponse> {
    return request.get('/claws')
  },

  getDetail(id: string): Promise<ClawDetailResponse> {
    return request.get(`/claws/${id}`)
  },

  getMultiple(ids: string[]): Promise<{ data: Claw[] }> {
    return request.get('/claws/batch', { params: { ids: ids.join(',') } })
  }
}
