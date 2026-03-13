# API 文档

## 基础信息

- **Base URL**: `http://localhost:8080/api/v1`
- **Content-Type**: `application/json`
- **字符编码**: UTF-8

## 通用响应格式

### 成功响应
```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 错误响应
```json
{
  "code": 1001,
  "message": "error message",
  "data": null
}
```

### 错误码
- `0` - 成功
- `1001` - 参数错误
- `1002` - 资源不存在
- `1003` - 数据库错误
- `1004` - 缓存错误
- `5000` - 服务器内部错误

---

## Claw 测评 API

### 1. 获取 Claw 列表

**请求**
```
GET /claws
```

**查询参数**
- `page` (int, optional): 页码，默认 1
- `page_size` (int, optional): 每页数量，默认 10
- `sort_by` (string, optional): 排序字段，默认 `overall_score`
- `order` (string, optional): 排序方式，`asc` 或 `desc`，默认 `desc`

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 10,
    "page": 1,
    "page_size": 10,
    "items": [
      {
        "id": 1,
        "name": "OpenClaw",
        "version": "1.0.0",
        "description": "开源 AI 助手框架",
        "response_time": 120.5,
        "concurrency": 100,
        "memory_usage": 256.0,
        "cpu_usage": 15.5,
        "tool_count": 50,
        "skill_compat": 45,
        "api_coverage": 95.0,
        "extensibility_score": 5,
        "install_difficulty": 2,
        "doc_quality": 5,
        "community_activity": 4,
        "learning_curve": 3,
        "error_rate": 0.5,
        "avg_uptime": 720.0,
        "crash_frequency": 0.1,
        "update_frequency": 7,
        "overall_score": 92.5,
        "created_at": "2026-03-13T15:00:00Z",
        "updated_at": "2026-03-13T15:00:00Z"
      }
    ]
  }
}
```

### 2. 获取 Claw 详情

**请求**
```
GET /claws/:id
```

**路径参数**
- `id` (int, required): Claw ID

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "name": "OpenClaw",
    "version": "1.0.0",
    "description": "开源 AI 助手框架",
    "response_time": 120.5,
    "concurrency": 100,
    "memory_usage": 256.0,
    "cpu_usage": 15.5,
    "tool_count": 50,
    "skill_compat": 45,
    "api_coverage": 95.0,
    "extensibility_score": 5,
    "install_difficulty": 2,
    "doc_quality": 5,
    "community_activity": 4,
    "learning_curve": 3,
    "error_rate": 0.5,
    "avg_uptime": 720.0,
    "crash_frequency": 0.1,
    "update_frequency": 7,
    "overall_score": 92.5,
    "created_at": "2026-03-13T15:00:00Z",
    "updated_at": "2026-03-13T15:00:00Z"
  }
}
```

### 3. 创建 Claw 测评数据

**请求**
```
POST /claws
```

**请求体**
```json
{
  "name": "OpenClaw",
  "version": "1.0.0",
  "description": "开源 AI 助手框架",
  "response_time": 120.5,
  "concurrency": 100,
  "memory_usage": 256.0,
  "cpu_usage": 15.5,
  "tool_count": 50,
  "skill_compat": 45,
  "api_coverage": 95.0,
  "extensibility_score": 5,
  "install_difficulty": 2,
  "doc_quality": 5,
  "community_activity": 4,
  "learning_curve": 3,
  "error_rate": 0.5,
  "avg_uptime": 720.0,
  "crash_frequency": 0.1,
  "update_frequency": 7
}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "overall_score": 92.5,
    "created_at": "2026-03-13T15:00:00Z"
  }
}
```

### 4. 更新 Claw 测评数据

**请求**
```
PUT /claws/:id
```

**路径参数**
- `id` (int, required): Claw ID

**请求体**
```json
{
  "version": "1.0.1",
  "response_time": 115.0,
  "memory_usage": 240.0
}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "overall_score": 93.0,
    "updated_at": "2026-03-13T16:00:00Z"
  }
}
```

### 5. 删除 Claw 测评数据

**请求**
```
DELETE /claws/:id
```

**路径参数**
- `id` (int, required): Claw ID

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## Skills 测评 API

### 1. 获取 Skills 列表

**请求**
```
GET /skills
```

**查询参数**
- `page` (int, optional): 页码，默认 1
- `page_size` (int, optional): 每页数量，默认 10
- `category` (string, optional): 分类筛选
- `sort_by` (string, optional): 排序字段，默认 `overall_score`
- `order` (string, optional): 排序方式，`asc` 或 `desc`，默认 `desc`

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "total": 20,
    "page": 1,
    "page_size": 10,
    "items": [
      {
        "id": 1,
        "name": "weather",
        "version": "1.0.0",
        "description": "天气查询 Skill",
        "category": "工具",
        "functionality_score": 5,
        "use_case_coverage": 4,
        "execution_speed": 50.0,
        "resource_consumption": 10.0,
        "supported_claws": "[\"OpenClaw\", \"ZeroClaw\"]",
        "dependency_complexity": 2,
        "code_quality": 5,
        "doc_quality": 5,
        "test_coverage": 85.0,
        "overall_score": 88.0,
        "created_at": "2026-03-13T15:00:00Z",
        "updated_at": "2026-03-13T15:00:00Z"
      }
    ]
  }
}
```

### 2. 获取 Skill 详情

**请求**
```
GET /skills/:id
```

**路径参数**
- `id` (int, required): Skill ID

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "name": "weather",
    "version": "1.0.0",
    "description": "天气查询 Skill",
    "category": "工具",
    "functionality_score": 5,
    "use_case_coverage": 4,
    "execution_speed": 50.0,
    "resource_consumption": 10.0,
    "supported_claws": "[\"OpenClaw\", \"ZeroClaw\"]",
    "dependency_complexity": 2,
    "code_quality": 5,
    "doc_quality": 5,
    "test_coverage": 85.0,
    "overall_score": 88.0,
    "created_at": "2026-03-13T15:00:00Z",
    "updated_at": "2026-03-13T15:00:00Z"
  }
}
```

### 3. 创建 Skill 测评数据

**请求**
```
POST /skills
```

**请求体**
```json
{
  "name": "weather",
  "version": "1.0.0",
  "description": "天气查询 Skill",
  "category": "工具",
  "functionality_score": 5,
  "use_case_coverage": 4,
  "execution_speed": 50.0,
  "resource_consumption": 10.0,
  "supported_claws": "[\"OpenClaw\", \"ZeroClaw\"]",
  "dependency_complexity": 2,
  "code_quality": 5,
  "doc_quality": 5,
  "test_coverage": 85.0
}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "overall_score": 88.0,
    "created_at": "2026-03-13T15:00:00Z"
  }
}
```

### 4. 更新 Skill 测评数据

**请求**
```
PUT /skills/:id
```

**路径参数**
- `id` (int, required): Skill ID

**请求体**
```json
{
  "version": "1.0.1",
  "execution_speed": 45.0,
  "test_coverage": 90.0
}
```

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "overall_score": 89.5,
    "updated_at": "2026-03-13T16:00:00Z"
  }
}
```

### 5. 删除 Skill 测评数据

**请求**
```
DELETE /skills/:id
```

**路径参数**
- `id` (int, required): Skill ID

**响应**
```json
{
  "code": 0,
  "message": "success",
  "data": null
}
```

---

## 预留接口

### 自动化测试接口

**Claw 测试**
```
POST /claws/test
```

**Skills 测试**
```
POST /skills/test
```

这些接口暂时返回 501 Not Implemented，待后续实现。
