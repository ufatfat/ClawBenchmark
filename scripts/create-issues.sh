#!/bin/bash
# 批量创建 ClawBenchmark MVP 阶段 Issue
# 使用方法：
#   brew install gh
#   gh auth login
#   chmod +x scripts/create-issues.sh
#   ./scripts/create-issues.sh

REPO="ufatfat/ClawBenchmark"

# 先创建标签
echo "=== 创建标签 ==="
gh label create "ai-task" --color "0E8A16" --description "AI 自动实现的任务" --repo $REPO 2>/dev/null
gh label create "frontend" --color "1D76DB" --description "前端相关" --repo $REPO 2>/dev/null
gh label create "infra" --color "D93F0B" --description "基础设施" --repo $REPO 2>/dev/null
gh label create "P0" --color "B60205" --description "最高优先级" --repo $REPO 2>/dev/null
gh label create "P1" --color "FBCA04" --description "中优先级" --repo $REPO 2>/dev/null
gh label create "needs-human" --color "D876E3" --description "需要人工介入" --repo $REPO 2>/dev/null

# 创建里程碑
echo "=== 创建里程碑 ==="
gh api repos/$REPO/milestones -f title="M1 - 首页与布局" -f description="完成首页、导航布局、全局样式" 2>/dev/null
gh api repos/$REPO/milestones -f title="M2 - Agent 目录" -f description="Agent 列表页、详情页、搜索筛选" 2>/dev/null
gh api repos/$REPO/milestones -f title="M3 - 排行榜与对比" -f description="排行榜页面、Agent 对比页面、图表" 2>/dev/null
gh api repos/$REPO/milestones -f title="M4 - 用户系统" -f description="登录页、个人中心、评论系统" 2>/dev/null
gh api repos/$REPO/milestones -f title="M5 - 联调优化" -f description="整体联调、响应式适配、性能优化" 2>/dev/null

echo "=== 创建 Issue ==="

# ============================================================
# M1 - 首页与布局
# ============================================================

gh issue create --repo $REPO \
  --title "feat(layout): 实现全局导航栏和页面布局" \
  --label "ai-task,frontend,P0" \
  --milestone "M1 - 首页与布局" \
  --body '## 任务描述
完善 `apps/web/src/layouts/DefaultLayout.vue`，实现完整的全局导航栏和页面布局。

## 验收标准
- [ ] 顶部导航栏包含 Logo、菜单项（首页/Agent目录/排行榜/对比）、登录按钮
- [ ] 导航栏高亮当前路由对应的菜单项
- [ ] 页面底部 Footer 包含版权信息和链接
- [ ] 响应式布局，移动端导航折叠为汉堡菜单
- [ ] 使用 Naive UI 组件（NLayout, NMenu, NButton 等）

## 上下文
- 当前骨架代码：`apps/web/src/layouts/DefaultLayout.vue`
- 路由配置：`apps/web/src/router/index.ts`
- PRD 参考：docs/PRD.md 第 3 章'

gh issue create --repo $REPO \
  --title "feat(home): 实现首页" \
  --label "ai-task,frontend,P0" \
  --milestone "M1 - 首页与布局" \
  --body '## 任务描述
实现 `apps/web/src/pages/HomePage.vue`，作为网站首页展示平台核心价值。

## 验收标准
- [ ] Hero 区域：标题、副标题、CTA 按钮（浏览 Agent / 查看排行榜）
- [ ] 统计数据展示区：Agent 数量、测评报告数、注册用户数（从 mock 数据计算）
- [ ] 热门 Agent 卡片展示（取 overall_score 前 6 名）
- [ ] 最新测评动态列表（取最近 5 条 benchmark）
- [ ] 使用 Naive UI 组件，视觉美观

## 上下文
- Mock 数据：`packages/mock-data`
- 类型定义：`packages/types`
- PRD 参考：docs/PRD.md 第 3 章'

gh issue create --repo $REPO \
  --title "feat(store): 创建 Pinia stores 和 mock 数据服务层" \
  --label "ai-task,frontend,P0" \
  --milestone "M1 - 首页与布局" \
  --body '## 任务描述
创建 Pinia stores 和 API 服务层，封装对 mock 数据的访问。后续替换为真实 API 时只需改 api 层。

## 验收标准
- [ ] `apps/web/src/api/agents.ts` — Agent 数据访问（列表、详情、搜索、筛选）
- [ ] `apps/web/src/api/benchmarks.ts` — Benchmark 数据访问（排行榜、对比）
- [ ] `apps/web/src/api/users.ts` — User 数据访问
- [ ] `apps/web/src/api/comments.ts` — Comment 数据访问
- [ ] `apps/web/src/stores/useAgentStore.ts` — Agent 状态管理
- [ ] `apps/web/src/stores/useBenchmarkStore.ts` — Benchmark 状态管理
- [ ] `apps/web/src/stores/useUserStore.ts` — 用户状态管理（mock 登录态）
- [ ] 所有 API 函数返回 Promise，模拟异步行为
- [ ] 包含单元测试

## 上下文
- Mock 数据：`packages/mock-data`
- 类型定义：`packages/types`
- 数据层规范见 CLAUDE.md'

# ============================================================
# M2 - Agent 目录
# ============================================================

gh issue create --repo $REPO \
  --title "feat(agent): 实现 AgentCard 组件" \
  --label "ai-task,frontend,P0" \
  --milestone "M2 - Agent 目录" \
  --body '## 任务描述
创建可复用的 AgentCard 组件，用于在列表页和首页展示 Agent 摘要信息。

## 验收标准
- [ ] `apps/web/src/components/AgentCard.vue`
- [ ] 展示：Logo（占位图）、名称、分类标签、简介、综合评分、三维评分条
- [ ] 点击跳转到 Agent 详情页
- [ ] 支持 props：agent 对象
- [ ] 视觉美观，使用 Naive UI 的 NCard 组件
- [ ] 包含单元测试

## 上下文
- 类型定义：`packages/types/src/agent.ts`
- 依赖 store：Issue #5 (Pinia stores)'

gh issue create --repo $REPO \
  --title "feat(agent): 实现 Agent 列表页" \
  --label "ai-task,frontend,P0" \
  --milestone "M2 - Agent 目录" \
  --body '## 任务描述
实现 `apps/web/src/pages/AgentListPage.vue`，展示所有 Agent 的列表。

## 验收标准
- [ ] Agent 卡片网格布局（每行 3 个，响应式）
- [ ] 搜索框：按名称/描述模糊搜索
- [ ] 筛选：按分类（category）筛选
- [ ] 排序：按综合评分、按名称、按最新更新
- [ ] 分页或无限滚动
- [ ] 空状态提示
- [ ] 使用 AgentCard 组件

## 上下文
- 依赖组件：AgentCard（Issue #6）
- 依赖 store：Issue #5
- PRD 参考：docs/PRD.md 第 3.1 章'

gh issue create --repo $REPO \
  --title "feat(agent): 实现 Agent 详情页" \
  --label "ai-task,frontend,P0" \
  --milestone "M2 - Agent 目录" \
  --body '## 任务描述
实现 `apps/web/src/pages/AgentDetailPage.vue`，展示单个 Agent 的完整信息。

## 验收标准
- [ ] 顶部：Logo、名称、分类、标签、简介、外部链接（官网/仓库）
- [ ] 评分区域：综合评分 + 三维评分雷达图（ECharts）
- [ ] 详细描述区域：渲染 full_description 的 Markdown 内容
- [ ] 基本信息：版本、许可证、支持平台、支持 LLM、架构类型
- [ ] 测评历史：该 Agent 的所有 benchmark 结果列表
- [ ] 评论区域：展示该 Agent 的评论（嵌套回复）
- [ ] 根据路由参数 slug 获取 Agent 数据

## 上下文
- 路由：`/agents/:slug`
- 依赖 store：Issue #5
- 需要 ECharts 雷达图
- PRD 参考：docs/PRD.md 第 3.1 章'

# ============================================================
# M3 - 排行榜与对比
# ============================================================

gh issue create --repo $REPO \
  --title "feat(leaderboard): 实现排行榜页面" \
  --label "ai-task,frontend,P0" \
  --milestone "M3 - 排行榜与对比" \
  --body '## 任务描述
实现 `apps/web/src/pages/LeaderboardPage.vue`，展示 Agent 综合排名。

## 验收标准
- [ ] 排行榜表格：排名、Logo、名称、综合评分、功能性/性能/安全性分项评分
- [ ] 支持按维度切换排序（综合/功能性/性能/安全性）
- [ ] 支持按分类筛选
- [ ] 评分用颜色区分等级（绿色 > 80，黄色 60-80，红色 < 60）
- [ ] 点击 Agent 名称跳转详情页
- [ ] 支持勾选 Agent 跳转到对比页
- [ ] 使用 Naive UI NDataTable 组件

## 上下文
- 依赖 store：Issue #5
- PRD 参考：docs/PRD.md 第 3.2 章'

gh issue create --repo $REPO \
  --title "feat(compare): 实现 Agent 对比页面" \
  --label "ai-task,frontend,P0" \
  --milestone "M3 - 排行榜与对比" \
  --body '## 任务描述
实现 `apps/web/src/pages/ComparePage.vue`，支持 2-4 个 Agent 的并排对比。

## 验收标准
- [ ] Agent 选择器：下拉搜索选择 2-4 个 Agent
- [ ] 支持从 URL query 参数读取预选 Agent（如 ?agents=openclaw,autogpt）
- [ ] 并排对比表格：基本信息、各维度评分
- [ ] 雷达图对比（ECharts，多个 Agent 叠加在同一雷达图）
- [ ] 柱状图对比（各维度分项得分）
- [ ] 支持平台、LLM 等特性对比矩阵

## 上下文
- 依赖 store：Issue #5
- 需要 ECharts 雷达图 + 柱状图
- PRD 参考：docs/PRD.md 第 3.2 章'

# ============================================================
# M4 - 用户系统
# ============================================================

gh issue create --repo $REPO \
  --title "feat(auth): 实现登录页面（Mock）" \
  --label "ai-task,frontend,P1" \
  --milestone "M4 - 用户系统" \
  --body '## 任务描述
实现 `apps/web/src/pages/LoginPage.vue`，提供 Mock 登录功能。

## 验收标准
- [ ] 登录表单：邮箱 + 密码
- [ ] GitHub 登录按钮（Mock，点击直接以 mock 用户登录）
- [ ] 注册表单：用户名 + 邮箱 + 密码
- [ ] 登录/注册切换 Tab
- [ ] Mock 登录逻辑：匹配 mock-data 中的用户，存入 Pinia store
- [ ] 登录成功后跳转首页，导航栏显示用户头像和名称
- [ ] 表单验证（Naive UI NForm）

## 上下文
- 依赖 store：Issue #5 (useUserStore)
- Mock 用户数据：`packages/mock-data/users.json`
- PRD 参考：docs/PRD.md 第 3.3 章'

gh issue create --repo $REPO \
  --title "feat(profile): 实现个人中心页面" \
  --label "ai-task,frontend,P1" \
  --milestone "M4 - 用户系统" \
  --body '## 任务描述
实现 `apps/web/src/pages/ProfilePage.vue`，展示用户个人信息。

## 验收标准
- [ ] 用户信息展示：头像、用户名、邮箱、角色、简介、公司
- [ ] 编辑个人信息（Mock，修改 store 中的数据）
- [ ] 我的评论列表
- [ ] 未登录时重定向到登录页
- [ ] 退出登录按钮

## 上下文
- 依赖 store：Issue #5 (useUserStore)
- PRD 参考：docs/PRD.md 第 3.3 章'

gh issue create --repo $REPO \
  --title "feat(comment): 实现评论组件" \
  --label "ai-task,frontend,P1" \
  --milestone "M4 - 用户系统" \
  --body '## 任务描述
创建可复用的评论组件，用于 Agent 详情页的评论区。

## 验收标准
- [ ] `apps/web/src/components/CommentList.vue` — 评论列表
- [ ] `apps/web/src/components/CommentItem.vue` — 单条评论（支持嵌套回复缩进）
- [ ] `apps/web/src/components/CommentForm.vue` — 发表评论表单
- [ ] 展示：用户头像、用户名、时间、内容、点赞数
- [ ] 点赞功能（Mock，修改 store 数据）
- [ ] 回复功能（Mock，添加到 store）
- [ ] 嵌套回复最多 3 层
- [ ] 未登录时提示登录

## 上下文
- 依赖 store：Issue #5
- Mock 评论数据：`packages/mock-data/comments.json`
- 类型定义：`packages/types/src/user.ts`'

# ============================================================
# M5 - 联调优化
# ============================================================

gh issue create --repo $REPO \
  --title "feat(responsive): 全站响应式适配" \
  --label "ai-task,frontend,P1" \
  --milestone "M5 - 联调优化" \
  --body '## 任务描述
确保所有页面在移动端、平板、桌面端都有良好的展示效果。

## 验收标准
- [ ] 导航栏移动端折叠
- [ ] Agent 卡片网格响应式（桌面3列 → 平板2列 → 手机1列）
- [ ] 排行榜表格移动端横向滚动
- [ ] 对比页面移动端纵向堆叠
- [ ] 图表自适应容器宽度
- [ ] 断点：640px / 768px / 1024px / 1280px

## 上下文
- 所有页面组件
- 使用 Naive UI 的 NGrid 或 CSS media queries'

gh issue create --repo $REPO \
  --title "feat(dark-mode): 实现暗色模式切换" \
  --label "ai-task,frontend,P1" \
  --milestone "M5 - 联调优化" \
  --body '## 任务描述
支持亮色/暗色模式切换。

## 验收标准
- [ ] 导航栏添加主题切换按钮（太阳/月亮图标）
- [ ] 使用 Naive UI 的 NConfigProvider 切换主题
- [ ] ECharts 图表跟随主题变化
- [ ] 主题偏好存储到 localStorage
- [ ] 默认跟随系统偏好

## 上下文
- 布局组件：`apps/web/src/layouts/DefaultLayout.vue`
- Naive UI 主题配置文档'

echo ""
echo "=== 全部 Issue 创建完成 ==="
echo "访问 https://github.com/$REPO/issues 查看"
