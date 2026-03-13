# ClawBenchmark 架构设计

## 系统架构

```
┌─────────────────────────────────────────┐
│         Frontend (Vue3 + TS)            │
│  ┌──────────┬──────────┬──────────┐    │
│  │  Claw    │  Skills  │  Common  │    │
│  │Benchmark │Benchmark │Components│    │
│  └──────────┴──────────┴──────────┘    │
└─────────────────┬───────────────────────┘
                  │ HTTP/REST API
┌─────────────────▼───────────────────────┐
│         Backend (Golang)                │
│  ┌──────────┬──────────┬──────────┐    │
│  │  Claw    │  Skills  │   Core   │    │
│  │ Service  │ Service  │  Service │    │
│  └──────────┴──────────┴──────────┘    │
│  ┌─────────────────────────────────┐   │
│  │      Repository Layer           │   │
│  └─────────────────────────────────┘   │
└─────────────────┬───────────────────────┘
                  │
    ┌─────────────┼─────────────┐
    │             │             │
┌───▼────┐   ┌───▼────┐   ┌───▼────┐
│PostgreSQL│   │ Redis  │   │MongoDB │
│  (主库)  │   │ (缓存) │   │(可选)  │
└─────────┘   └────────┘   └────────┘
```

## 模块划分

### 前端模块

#### 1. Common Module (公共模块)
**职责**：
- 布局组件（Header、Footer、Sidebar）
- 通用 UI 组件（Card、Table、Chart）
- 工具函数、类型定义
- API 请求封装

**技术要点**：
- Element-Plus 组件二次封装
- TypeScript 类型系统
- Axios 请求拦截

#### 2. Claw Benchmark Module
**职责**：
- 首页卡片列表展示
- Claw 详情页（图表 + 数据表）
- Claw 对比功能

**页面**：
- `/` - 首页（卡片列表）
- `/claw/:id` - Claw 详情页
- `/claw/compare` - 对比页面

#### 3. Skills Benchmark Module
**职责**：
- Skills 列表展示
- Skills 详情页
- 筛选和搜索功能

**页面**：
- `/skills` - Skills 列表
- `/skills/:id` - Skills 详情

### 后端模块

#### 1. Core Service (核心服务)
**职责**：
- HTTP 服务器初始化
- 路由注册
- 中间件（CORS、日志、错误处理、限流）
- 配置管理

**技术栈**：
- Gin 框架
- Viper 配置管理
- Zap 日志

#### 2. Claw Service
**职责**：
- Claw 测评数据 CRUD
- 评分计算和聚合
- 历史数据管理
- 自动化测试接口（预留）

**API 端点**：
- `GET /api/v1/claws` - 获取 Claw 列表
- `GET /api/v1/claws/:id` - 获取 Claw 详情
- `POST /api/v1/claws` - 创建测评数据
- `PUT /api/v1/claws/:id` - 更新测评数据
- `DELETE /api/v1/claws/:id` - 删除测评数据
- `POST /api/v1/claws/test` - 自动化测试接口（预留）

#### 3. Skills Service
**职责**：
- Skills 测评数据 CRUD
- 评分计算
- 筛选和搜索

**API 端点**：
- `GET /api/v1/skills` - 获取 Skills 列表
- `GET /api/v1/skills/:id` - 获取 Skills 详情
- `POST /api/v1/skills` - 创建测评数据
- `PUT /api/v1/skills/:id` - 更新测评数据
- `DELETE /api/v1/skills/:id` - 删除测评数据

#### 4. Repository Layer (数据层)
**职责**：
- 数据库连接管理
- CRUD 操作封装
- 缓存策略
- 数据模型定义

**技术栈**：
- GORM (PostgreSQL ORM)
- go-redis (Redis 客户端)

## 数据模型

### Claw 测评数据

```go
type ClawBenchmark struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"uniqueIndex;not null"`
    Version     string    `json:"version"`
    Description string    `json:"description"`
    
    // 性能指标
    ResponseTime    float64 `json:"response_time"`     // ms
    Concurrency     int     `json:"concurrency"`       // req/s
    MemoryUsage     float64 `json:"memory_usage"`      // MB
    CPUUsage        float64 `json:"cpu_usage"`         // %
    
    // 功能完整性
    ToolCount       int     `json:"tool_count"`
    SkillCompat     int     `json:"skill_compat"`
    APICoverage     float64 `json:"api_coverage"`      // %
    ExtensibilityScore int  `json:"extensibility_score"` // 1-5
    
    // 易用性
    InstallDifficulty int   `json:"install_difficulty"` // 1-5
    DocQuality        int   `json:"doc_quality"`        // 1-5
    CommunityActivity int   `json:"community_activity"` // 1-5
    LearningCurve     int   `json:"learning_curve"`     // 1-5
    
    // 稳定性
    ErrorRate         float64 `json:"error_rate"`       // %
    AvgUptime         float64 `json:"avg_uptime"`       // hours
    CrashFrequency    float64 `json:"crash_frequency"`  // per day
    UpdateFrequency   int     `json:"update_frequency"` // days
    
    // 综合评分
    OverallScore    float64   `json:"overall_score"`    // 0-100
    
    CreatedAt       time.Time `json:"created_at"`
    UpdatedAt       time.Time `json:"updated_at"`
}
```

### Skills 测评数据

```go
type SkillBenchmark struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"uniqueIndex;not null"`
    Version     string    `json:"version"`
    Description string    `json:"description"`
    Category    string    `json:"category"`
    
    // 功能性
    FunctionalityScore int `json:"functionality_score"` // 1-5
    UseCaseCoverage    int `json:"use_case_coverage"`   // 1-5
    
    // 性能
    ExecutionSpeed     float64 `json:"execution_speed"`  // ms
    ResourceConsumption float64 `json:"resource_consumption"` // MB
    
    // 兼容性
    SupportedClaws     string  `json:"supported_claws"`  // JSON array
    DependencyComplexity int   `json:"dependency_complexity"` // 1-5
    
    // 质量
    CodeQuality        int     `json:"code_quality"`     // 1-5
    DocQuality         int     `json:"doc_quality"`      // 1-5
    TestCoverage       float64 `json:"test_coverage"`    // %
    
    // 综合评分
    OverallScore       float64 `json:"overall_score"`    // 0-100
    
    CreatedAt          time.Time `json:"created_at"`
    UpdatedAt          time.Time `json:"updated_at"`
}
```

## 缓存策略

### Redis 缓存
- **Claw 列表**：缓存 5 分钟
- **Claw 详情**：缓存 10 分钟
- **Skills 列表**：缓存 5 分钟
- **Skills 详情**：缓存 10 分钟

### 缓存更新
- 数据更新时自动清除相关缓存
- 支持手动刷新缓存

## 安全性

### API 安全
- CORS 配置
- 请求限流（100 req/min per IP）
- SQL 注入防护（GORM 参数化查询）
- XSS 防护（前端输入验证）

### 数据安全
- 敏感数据加密存储
- 定期数据备份

## 扩展性设计

### 预留接口
1. **自动化测试接口**
   - `POST /api/v1/claws/test`
   - `POST /api/v1/skills/test`

2. **用户认证系统**
   - JWT Token 认证
   - OAuth 第三方登录

3. **论坛系统**
   - 人类论坛 API
   - AI Agent 论坛 API

### 微服务化准备
- 服务间通信接口预留
- 消息队列集成点预留
- 服务注册发现预留
