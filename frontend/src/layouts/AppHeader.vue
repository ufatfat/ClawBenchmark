<template>
  <el-header class="app-header">
    <div class="header-inner">
      <router-link to="/" class="logo">
        <span class="logo-icon">⚡</span>
        <span class="logo-text">ClawBenchmark</span>
      </router-link>

      <el-menu mode="horizontal" :router="true" :default-active="route.path" class="nav-menu">
        <el-menu-item index="/benchmarks">测评列表</el-menu-item>
        <el-menu-item index="/leaderboard">排行榜</el-menu-item>
        <el-menu-item index="/models">模型库</el-menu-item>
        <el-menu-item index="/reports">报告</el-menu-item>
      </el-menu>

      <div class="header-actions">
        <template v-if="auth.isLoggedIn()">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" icon="UserFilled" />
              <span>{{ auth.user?.username }}</span>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="logout">退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </template>
        <el-button v-else type="primary" @click="router.push('/login')">登录</el-button>
      </div>
    </div>
  </el-header>
</template>

<script setup lang="ts">
  import { useRoute, useRouter } from 'vue-router'
  import { useAuthStore } from '@/stores/auth'

  const route = useRoute()
  const router = useRouter()
  const auth = useAuthStore()

  function handleCommand(cmd: string) {
    if (cmd === 'logout') {
      auth.logout()
      router.push('/login')
    }
  }
</script>

<style scoped>
  .app-header {
    background: #fff;
    border-bottom: 1px solid #e4e7ed;
    padding: 0;
    height: 60px;
    position: sticky;
    top: 0;
    z-index: 100;
  }

  .header-inner {
    max-width: 1280px;
    margin: 0 auto;
    height: 100%;
    display: flex;
    align-items: center;
    padding: 0 24px;
    gap: 32px;
  }

  .logo {
    display: flex;
    align-items: center;
    gap: 8px;
    text-decoration: none;
    font-size: 18px;
    font-weight: 700;
    color: #303133;
    white-space: nowrap;
  }

  .logo-icon {
    font-size: 22px;
  }

  .nav-menu {
    flex: 1;
    border-bottom: none;
  }

  .header-actions {
    margin-left: auto;
  }

  .user-info {
    display: flex;
    align-items: center;
    gap: 8px;
    cursor: pointer;
    color: #303133;
  }
</style>
