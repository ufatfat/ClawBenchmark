import http from '@/utils/http'
import type { Benchmark, BenchmarkResult, PageResult } from '@/types'

export const benchmarkApi = {
  list: (params?: { page?: number; pageSize?: number; category?: string }) =>
    http.get<PageResult<Benchmark>>('/benchmarks', { params }),

  get: (id: number) => http.get<Benchmark>(`/benchmarks/${id}`),

  create: (data: Partial<Benchmark>) => http.post<Benchmark>('/benchmarks', data),

  update: (id: number, data: Partial<Benchmark>) =>
    http.put<Benchmark>(`/benchmarks/${id}`, data),

  delete: (id: number) => http.delete(`/benchmarks/${id}`),

  run: (id: number) => http.post(`/benchmarks/${id}/run`),

  results: (id: number) => http.get<BenchmarkResult[]>(`/benchmarks/${id}/results`),
}
