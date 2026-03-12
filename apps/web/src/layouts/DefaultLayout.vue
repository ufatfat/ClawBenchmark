<script setup lang="ts">
import {
  NLayout, NLayoutHeader, NLayoutContent, NLayoutFooter,
  NMenu, NButton, NSpace, NDrawer, NDrawerContent
} from 'naive-ui'
import { useRouter, useRoute } from 'vue-router'
import { computed, ref } from 'vue'
import type { MenuOption } from 'naive-ui'

const router = useRouter()
const route = useRoute()

const menuOptions: MenuOption[] = [
  { label: '首页', key: 'Home' },
  { label: 'Agent 目录', key: 'AgentList' },
  { label: '排行榜', key: 'Leaderboard' },
  { label: '对比', key: 'Compare' },
]

// Map route names to menu keys (AgentDetail highlights AgentList)
const activeKey = computed(() => {
  const name = route.name as string
  if (name === 'AgentDetail') return 'AgentList'
  return name
})

function handleMenuUpdate(key: string) {
  drawerOpen.value = false
  router.push({ name: key })
}

// Mobile drawer
const drawerOpen = ref(false)
</script>

<template>
  <n-layout style="min-height: 100vh">
    <n-layout-header bordered style="padding: 0 24px; height: 64px; display: flex; align-items: center; justify-content: space-between; position: sticky; top: 0; z-index: 100;">
      <!-- Logo + desktop nav -->
      <div style="display: flex; align-items: center; gap: 32px;">
        <router-link to="/" class="logo">ClawBenchmark</router-link>

        <!-- Desktop menu -->
        <n-menu
          class="desktop-menu"
          mode="horizontal"
          :options="menuOptions"
          :value="activeKey"
          @update:value="handleMenuUpdate"
        />
      </div>

      <!-- Right side -->
      <n-space align="center">
        <n-button
          class="desktop-menu"
          quaternary
          @click="router.push({ name: 'Login' })"
        >
          登录
        </n-button>

        <!-- Hamburger for mobile -->
        <n-button
          class="mobile-menu-btn"
          quaternary
          @click="drawerOpen = true"
          aria-label="打开菜单"
        >
          ☰
        </n-button>
      </n-space>
    </n-layout-header>

    <!-- Mobile drawer -->
    <n-drawer v-model:show="drawerOpen" placement="right" :width="240">
      <n-drawer-content title="菜单" closable>
        <n-menu
          :options="menuOptions"
          :value="activeKey"
          @update:value="handleMenuUpdate"
        />
        <div style="padding: 16px 0;">
          <n-button block @click="() => { drawerOpen = false; router.push({ name: 'Login' }) }">
            登录
          </n-button>
        </div>
      </n-drawer-content>
    </n-drawer>

    <n-layout-content style="padding: 24px; max-width: 1200px; margin: 0 auto; width: 100%; box-sizing: border-box;">
      <router-view />
    </n-layout-content>

    <n-layout-footer bordered style="padding: 24px; text-align: center; color: #999; font-size: 13px;">
      <div>
        ClawBenchmark &copy; 2026 &mdash; AI Agent 测评平台
      </div>
      <div style="margin-top: 8px; display: flex; justify-content: center; gap: 16px;">
        <router-link to="/agents" class="footer-link">Agent 目录</router-link>
        <router-link to="/leaderboard" class="footer-link">排行榜</router-link>
        <router-link to="/compare" class="footer-link">对比</router-link>
      </div>
    </n-layout-footer>
  </n-layout>
</template>

<style scoped>
.logo {
  font-size: 20px;
  font-weight: 700;
  text-decoration: none;
  color: inherit;
  white-space: nowrap;
}

.footer-link {
  color: #999;
  text-decoration: none;
}

.footer-link:hover {
  color: #18a058;
}

.mobile-menu-btn {
  display: none;
}

@media (max-width: 768px) {
  .desktop-menu {
    display: none;
  }

  .mobile-menu-btn {
    display: inline-flex;
  }
}
</style>
