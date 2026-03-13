import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { skillApi } from '../api/skill'
import type { Skill, SkillListParams } from '../types/skill'

export const useSkillStore = defineStore('skill', () => {
  const skills = ref<Skill[]>([])
  const currentSkill = ref<Skill | null>(null)
  const categories = ref<string[]>([])
  const total = ref(0)
  const loading = ref(false)
  const error = ref<string | null>(null)

  const filters = ref<SkillListParams>({
    category: undefined,
    search: '',
    sortBy: 'overallScore',
    sortOrder: 'desc',
    page: 1,
    pageSize: 20
  })

  const filteredSkills = computed(() => skills.value)

  async function fetchSkills() {
    loading.value = true
    error.value = null
    try {
      const response = await skillApi.getList(filters.value)
      skills.value = response.data
      total.value = response.total
    } catch (e) {
      error.value = e instanceof Error ? e.message : '获取列表失败'
      throw e
    } finally {
      loading.value = false
    }
  }

  async function fetchSkillDetail(id: string) {
    loading.value = true
    error.value = null
    try {
      currentSkill.value = await skillApi.getDetail(id)
    } catch (e) {
      error.value = e instanceof Error ? e.message : '获取详情失败'
      throw e
    } finally {
      loading.value = false
    }
  }

  async function fetchCategories() {
    try {
      categories.value = await skillApi.getCategories()
    } catch (e) {
      console.error('获取分类失败:', e)
    }
  }

  function setFilter(key: keyof SkillListParams, value: any) {
    filters.value[key] = value
    filters.value.page = 1
  }

  function resetFilters() {
    filters.value = {
      category: undefined,
      search: '',
      sortBy: 'overallScore',
      sortOrder: 'desc',
      page: 1,
      pageSize: 20
    }
  }

  return {
    skills,
    currentSkill,
    categories,
    total,
    loading,
    error,
    filters,
    filteredSkills,
    fetchSkills,
    fetchSkillDetail,
    fetchCategories,
    setFilter,
    resetFilters
  }
})
