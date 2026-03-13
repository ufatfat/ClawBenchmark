<template>
  <div ref="chartRef" :style="{ width, height }" />
</template>

<script setup lang="ts">
  import { ref, onMounted, onUnmounted, watch } from 'vue'
  import * as echarts from 'echarts'

  const props = withDefaults(
    defineProps<{
      data: { name: string; value: number }[]
      type?: 'bar' | 'radar' | 'line'
      title?: string
      width?: string
      height?: string
    }>(),
    { type: 'bar', width: '100%', height: '320px' }
  )

  const chartRef = ref<HTMLElement>()
  let chart: echarts.ECharts | null = null

  function buildOption() {
    if (props.type === 'radar') {
      return {
        title: props.title ? { text: props.title, left: 'center' } : undefined,
        radar: {
          indicator: props.data.map((d) => ({ name: d.name, max: 100 })),
        },
        series: [{ type: 'radar', data: [{ value: props.data.map((d) => d.value), name: '得分' }] }],
      }
    }

    return {
      title: props.title ? { text: props.title, left: 'center' } : undefined,
      tooltip: { trigger: 'axis' },
      xAxis: { type: 'category', data: props.data.map((d) => d.name), axisLabel: { rotate: 30 } },
      yAxis: { type: 'value', min: 0, max: 100 },
      series: [
        {
          type: props.type,
          data: props.data.map((d) => d.value),
          itemStyle: { color: '#409eff' },
          label: { show: true, position: 'top', formatter: '{c}' },
        },
      ],
    }
  }

  function initChart() {
    if (!chartRef.value) return
    chart = echarts.init(chartRef.value)
    chart.setOption(buildOption())
  }

  watch(() => props.data, () => chart?.setOption(buildOption()), { deep: true })

  const resizeObserver = new ResizeObserver(() => chart?.resize())

  onMounted(() => {
    initChart()
    if (chartRef.value) resizeObserver.observe(chartRef.value)
  })

  onUnmounted(() => {
    resizeObserver.disconnect()
    chart?.dispose()
  })
</script>
