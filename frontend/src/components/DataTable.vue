<template>
  <div class="data-table">
    <div class="table-toolbar" v-if="$slots.toolbar || searchable">
      <slot name="toolbar" />
      <el-input
        v-if="searchable"
        v-model="searchText"
        placeholder="搜索..."
        clearable
        style="width: 240px"
        @input="handleSearch"
      >
        <template #prefix><el-icon><Search /></el-icon></template>
      </el-input>
    </div>

    <el-table
      :data="filteredData"
      v-loading="loading"
      stripe
      border
      style="width: 100%"
      @sort-change="handleSortChange"
    >
      <slot />
    </el-table>

    <el-pagination
      v-if="total > 0"
      v-model:current-page="currentPage"
      v-model:page-size="pageSize"
      :total="total"
      :page-sizes="[10, 20, 50]"
      layout="total, sizes, prev, pager, next"
      class="pagination"
      @change="emit('page-change', { page: currentPage, pageSize })"
    />
  </div>
</template>

<script setup lang="ts">
  import { ref, computed } from 'vue'
  import { Search } from '@element-plus/icons-vue'

  const props = withDefaults(
    defineProps<{
      data: Record<string, unknown>[]
      total?: number
      loading?: boolean
      searchable?: boolean
      searchFields?: string[]
    }>(),
    { total: 0, loading: false, searchable: false, searchFields: () => [] }
  )

  const emit = defineEmits<{
    'page-change': [{ page: number; pageSize: number }]
    'sort-change': [{ prop: string; order: string }]
  }>()

  const searchText = ref('')
  const currentPage = ref(1)
  const pageSize = ref(10)

  const filteredData = computed(() => {
    if (!searchText.value || !props.searchFields.length) return props.data
    const q = searchText.value.toLowerCase()
    return props.data.filter((row) =>
      props.searchFields.some((field) => String(row[field] ?? '').toLowerCase().includes(q))
    )
  })

  function handleSearch() {
    currentPage.value = 1
  }

  function handleSortChange({ prop, order }: { prop: string; order: string }) {
    emit('sort-change', { prop, order })
  }
</script>

<style scoped>
  .table-toolbar {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 16px;
  }

  .pagination {
    margin-top: 16px;
    justify-content: flex-end;
  }
</style>
