import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { clawApi } from '@/api/claw'
import type { Claw, SortField, SortOrder } from '@/types/claw'

export const useClawStore = defineStore('claw', () => {
  const claws = ref<Claw[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  const sortedClaws = computed(() => {
    return [...claws.value]
  })

  async function fetchClaws() {
    loading.value = true
    error.value = null
    try {
      const response = await clawApi.getList()
      claws.value = response.data
    } catch (e) {
      error.value = e instanceof Error ? e.message : 'Failed to fetch claws'
      throw e
    } finally {
      loading.value = false
    }
  }

  function sortClaws(field: SortField, order: SortOrder) {
    claws.value.sort((a, b) => {
      let aVal: number, bVal: number

      if (field === 'overallScore') {
        aVal = a.overallScore
        bVal = b.overallScore
      } else {
        aVal = a.metrics[field]
        bVal = b.metrics[field]
      }

      return order === 'asc' ? aVal - bVal : bVal - aVal
    })
  }

  function getClawById(id: string) {
    return claws.value.find(c => c.id === id)
  }

  return {
    claws,
    sortedClaws,
    loading,
    error,
    fetchClaws,
    sortClaws,
    getClawById
  }
})
