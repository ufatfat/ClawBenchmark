<template>
  <el-container class="main-layout">
    <AppHeader />
    <el-main class="main-content">
      <div class="content-wrapper">
        <router-view />
      </div>
    </el-main>
    <AppFooter />
  </el-container>
</template>

<script setup lang="ts">
  import AppHeader from './AppHeader.vue'
  import AppFooter from './AppFooter.vue'
</script>

<style scoped>
  .main-layout {
    min-height: 100vh;
    display: flex;
    flex-direction: column;
  }

  .main-content {
    flex: 1;
    background: #f5f7fa;
    padding: 24px;
  }

  .content-wrapper {
    max-width: 1280px;
    margin: 0 auto;
  }
</style>
