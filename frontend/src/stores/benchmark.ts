import { defineStore } from 'pinia'
import { ref } from 'vue'
import { benchmarkApi } from '@/api/benchmark'
import type { Benchmark, PageResult } from '@/types'

export const useBenchmarkStore = defineStore('benchmark', () => {
  const list = ref<Benchmark[]>([])
  const total = ref(0)
  const loading = ref(false)
  const current = ref<Benchmark | null>(null)

  async function fetchList(params?: { page?: number; pageSize?: number; category?: string }) {
    loading.value = true
    try {
      const res = await benchmarkApi.list(params) as unknown as PageResult<Benchmark>
      list.value = res.items
      total.value = res.total
    } finally {
      loading.value = false
    }
  }

  async function fetchOne(id: number) {
    current.value = await benchmarkApi.get(id) as unknown as Benchmark
  }

  return { list, total, loading, current, fetchList, fetchOne }
})
