# ClawBenchmark

AI Agent 测评平台。当前阶段为 Frontend MVP，使用 mock 数据。

## 项目规范

### 架构
- Monorepo（npm workspaces）
- 前端：Vue 3 + TypeScript + Vite + Naive UI + Pinia + ECharts
- 数据：packages/mock-data 下的 JSON 文件
- 类型：packages/types 下的共享 TypeScript 类型

### 目录结构
- `apps/web/` — 前端 Web 应用（当前开发重点）
- `packages/types/` — 共享 TypeScript 类型定义
- `packages/mock-data/` — Mock 数据
- `docs/` — 项目文档（PRD.md 是核心需求文档）

### 前端编码规范
- Vue 3 Composition API + `<script setup lang="ts">` 语法
- 组件文件名 PascalCase（如 AgentCard.vue）
- composables 用 `use` 前缀（如 useAgentList.ts）
- 使用 Naive UI 组件库，不混用其他 UI 库
- 状态管理用 Pinia，store 文件以 `use` 前缀命名
- 所有组件和函数必须有 TypeScript 类型
- 样式使用 scoped CSS 或 CSS Modules

### 数据层规范
- 当前阶段所有数据来自 packages/mock-data
- API 层（apps/web/src/api/）封装数据访问，后续替换为真实 API 时只改这一层
- 类型定义统一放在 packages/types，前端通过 workspace 引用

### 测试要求
- 使用 Vitest 进行单元测试
- 测试文件与源文件同目录，命名 `*.test.ts` 或 `*.spec.ts`
- 每个 PR 必须包含对应的测试

### Git 规范
- 分支命名：`feat/<short-desc>`、`fix/<short-desc>`
- Commit message：`feat(scope): description` / `fix(scope): description`
- 每个 PR 只解决一个功能点

### 关键文档
- `docs/PRD.md` — 完整产品需求文档（实现功能前必读）
