# 开发指南

## 开发环境要求

### 前端
- Node.js >= 18
- npm >= 9

### 后端
- Go >= 1.21
- PostgreSQL >= 14
- Redis >= 7

## 本地开发

### 1. 数据库初始化

**PostgreSQL**
```bash
createdb clawbenchmark
psql clawbenchmark < backend/migrations/init.sql
```

**Redis**
```bash
redis-server
```

### 2. 后端开发

```bash
cd backend

# 安装依赖
go mod download

# 配置环境变量
cp .env.example .env
# 编辑 .env 文件

# 运行
make run

# 或者使用 air 热重载
air
```

### 3. 前端开发

```bash
cd frontend

# 安装依赖
npm install

# 开发模式
npm run dev

# 构建
npm run build
```

## 代码规范

### 前端

**ESLint + Prettier**
```bash
npm run lint
npm run format
```

**TypeScript 类型检查**
```bash
npm run type-check
```

### 后端

**golangci-lint**
```bash
make lint
```

**格式化**
```bash
make fmt
```

## 测试

### 前端测试
```bash
npm run test
npm run test:coverage
```

### 后端测试
```bash
make test
make test-coverage
```

## Git 工作流

### Commit 规范

使用 Conventional Commits：

```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type 类型：**
- `feat`: 新功能
- `fix`: 修复 bug
- `docs`: 文档更新
- `style`: 代码格式
- `refactor`: 重构
- `test`: 测试相关
- `chore`: 构建/工具/依赖
- `perf`: 性能优化

**示例：**
```
feat(claw): 添加 Claw 列表分页功能

- 实现后端分页逻辑
- 前端添加分页组件
- 更新 API 文档

Closes #12
```

### 分支策略

- `main` - 主分支，保持稳定
- `develop` - 开发分支
- `feature/*` - 功能分支
- `fix/*` - 修复分支

## CI/CD

### GitHub Actions

- **Lint**: 代码风格检查
- **Test**: 单元测试
- **Build**: 构建检查
- **Deploy**: 自动部署（main 分支）

## 部署

### Docker 部署

```bash
# 构建镜像
docker-compose build

# 启动服务
docker-compose up -d

# 查看日志
docker-compose logs -f
```

### 手动部署

```bash
# 前端
cd frontend
npm run build
# 将 dist/ 目录部署到 Web 服务器

# 后端
cd backend
make build
# 将二进制文件部署到服务器
```

## 常见问题

### 数据库连接失败
检查 PostgreSQL 是否启动，配置是否正确。

### Redis 连接失败
检查 Redis 是否启动，端口是否正确。

### 前端跨域问题
开发环境已配置代理，生产环境需要配置 CORS。

## 性能优化

### 前端
- 路由懒加载
- 组件按需引入
- 图片懒加载
- 打包优化

### 后端
- 数据库索引优化
- Redis 缓存策略
- 连接池配置
- 并发控制

## 监控和日志

### 日志
- 前端：Console + Sentry
- 后端：Zap 日志库

### 监控
- 性能监控
- 错误追踪
- 用户行为分析
