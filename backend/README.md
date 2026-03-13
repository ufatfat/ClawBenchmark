# ClawBenchmark Backend

基于 Gin 框架的 Golang 后端服务。

## 项目结构

```
backend/
├── cmd/
│   └── server/          # 主程序入口
├── internal/
│   ├── config/          # 配置管理
│   ├── middleware/      # 中间件
│   ├── database/        # 数据库连接
│   └── router/          # 路由配置
├── pkg/
│   └── response/        # 响应封装
├── go.mod
├── Makefile
├── .env.example
└── .gitignore
```

## 功能特性

- ✅ Gin HTTP 服务器
- ✅ Viper 配置管理（支持 .env）
- ✅ Zap 结构化日志
- ✅ PostgreSQL + GORM
- ✅ Redis 支持
- ✅ CORS 中间件
- ✅ 请求日志中间件
- ✅ 错误处理中间件
- ✅ 限流中间件
- ✅ 优雅关闭
- ✅ 单元测试

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod tidy
```

### 2. 配置环境变量

```bash
cp .env.example .env
# 编辑 .env 文件，配置数据库等信息
```

### 3. 运行服务

```bash
make run
```

### 4. 构建

```bash
make build
```

### 5. 测试

```bash
make test
```

## Makefile 命令

- `make run` - 运行服务器
- `make build` - 编译二进制文件
- `make test` - 运行测试
- `make lint` - 代码检查
- `make fmt` - 格式化代码
- `make tidy` - 整理依赖
- `make clean` - 清理构建文件
- `make help` - 显示帮助

## API 端点

- `GET /health` - 健康检查
- `GET /api/v1/ping` - 测试接口

## 技术栈

- **框架**: Gin
- **配置**: Viper
- **日志**: Zap
- **数据库**: PostgreSQL (GORM)
- **缓存**: Redis
- **限流**: ulule/limiter
- **测试**: testify
