<template>
  <el-footer class="app-footer">
    <div class="footer-inner">
      <span>© 2024 ClawBenchmark. All rights reserved.</span>
      <span class="divider">|</span>
      <span>AI 模型评测平台</span>
    </div>
  </el-footer>
</template>

<script setup lang="ts"></script>

<style scoped>
  .app-footer {
    background: #f5f7fa;
    border-top: 1px solid #e4e7ed;
    height: 48px;
    display: flex;
    align-items: center;
    justify-content: center;
  }

  .footer-inner {
    display: flex;
    align-items: center;
    gap: 12px;
    color: #909399;
    font-size: 13px;
  }

  .divider {
    color: #dcdfe6;
  }
</style>
