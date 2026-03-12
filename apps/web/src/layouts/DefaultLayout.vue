<script setup lang="ts">
import { NLayout, NLayoutHeader, NLayoutContent, NLayoutFooter, NMenu, NSpace, NButton } from 'naive-ui'
import { useRouter } from 'vue-router'
import { h } from 'vue'
import type { MenuOption } from 'naive-ui'

const router = useRouter()

const menuOptions: MenuOption[] = [
  { label: '首页', key: 'Home' },
  { label: 'Agent 目录', key: 'AgentList' },
  { label: '排行榜', key: 'Leaderboard' },
  { label: '对比', key: 'Compare' },
]

function handleMenuUpdate(key: string) {
  router.push({ name: key })
}
</script>

<template>
  <n-layout style="min-height: 100vh">
    <n-layout-header bordered style="padding: 0 24px; display: flex; align-items: center; justify-content: space-between; height: 64px;">
      <div style="display: flex; align-items: center; gap: 32px;">
        <router-link to="/" style="font-size: 20px; font-weight: 700; text-decoration: none; color: inherit;">
          ClawBenchmark
        </router-link>
        <n-menu
          mode="horizontal"
          :options="menuOptions"
          @update:value="handleMenuUpdate"
        />
      </div>
      <n-space>
        <n-button quaternary @click="router.push({ name: 'Login' })">登录</n-button>
      </n-space>
    </n-layout-header>

    <n-layout-content style="padding: 24px; max-width: 1200px; margin: 0 auto; width: 100%;">
      <router-view />
    </n-layout-content>

    <n-layout-footer bordered style="padding: 16px 24px; text-align: center; color: #999; font-size: 13px;">
      ClawBenchmark &copy; 2026 — AI Agent 测评平台
    </n-layout-footer>
  </n-layout>
</template>
