# ClawBenchmark

针对 OpenClaw、ZeroClaw、NanoClaw 等 Claw-like 项目的测评网站。

## 技术栈

### 前端
- TypeScript
- Vue 3
- Element-Plus

### 后端
- Golang
- PostgreSQL
- Redis

## 功能模块

### MVP 阶段
1. **Claw 测评** - 展示各 Claw 项目的性能、功能、易用性等多维度评测
2. **Skills 测评** - 评测各类 Skills 的功能性、性能、兼容性

### 后续规划
3. **使用场景案例** - 实际应用案例展示
4. **社区论坛** - 人类论坛 + AI Agent 论坛

## 项目结构

```
ClawBenchmark/
├── frontend/          # Vue3 前端
├── backend/           # Golang 后端
├── docs/              # 项目文档
├── scripts/           # 部署脚本
└── docker/            # Docker 配置
```

## 快速开始

### 前端
```bash
cd frontend
npm install
npm run dev
```

### 后端
```bash
cd backend
make install
make run
```

## 开发规范

- Commit 格式：Conventional Commits
- 代码风格：统一使用 lint 工具
- 测试：单元测试覆盖核心逻辑
