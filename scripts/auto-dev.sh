#!/bin/bash
# 本地 AI 自动开发脚本
# 定时检查带 ai-task 标签的 Issue，自动编码并创建 PR

set -e

REPO="ufatfat/ClawBenchmark"
REPO_DIR="/Users/ufatfat/Projects/ClawBenchmark"

cd "$REPO_DIR"

echo "=== 检查待处理的 Issue ==="

# 获取所有带 ai-task 标签且未关闭的 Issue
ISSUES=$(gh issue list --repo "$REPO" --label "ai-task" --state open --json number,title,body --limit 50)

# 解析 JSON，逐个处理
echo "$ISSUES" | jq -c '.[]' | while read -r issue; do
  ISSUE_NUMBER=$(echo "$issue" | jq -r '.number')
  ISSUE_TITLE=$(echo "$issue" | jq -r '.title')
  ISSUE_BODY=$(echo "$issue" | jq -r '.body')

  BRANCH_NAME="ai/issue-${ISSUE_NUMBER}"

  # 检查是否已经有对应的分支或 PR
  if git show-ref --verify --quiet "refs/heads/$BRANCH_NAME" 2>/dev/null; then
    echo "Issue #${ISSUE_NUMBER} 已有分支 $BRANCH_NAME，跳过"
    continue
  fi

  if gh pr list --repo "$REPO" --head "$BRANCH_NAME" --state all --json number | jq -e '. | length > 0' >/dev/null; then
    echo "Issue #${ISSUE_NUMBER} 已有 PR，跳过"
    continue
  fi

  echo ""
  echo "=========================================="
  echo "开始处理 Issue #${ISSUE_NUMBER}"
  echo "标题: ${ISSUE_TITLE}"
  echo "=========================================="

  # 确保在 main 分支
  git checkout main
  git pull origin main

  # 创建新分支
  git checkout -b "$BRANCH_NAME"

  # 运行 Claude Code（清除嵌套会话检查）
  unset CLAUDECODE
  claude --allowedTools Bash,Read,Write,Edit,Glob,Grep -p "你是 ClawBenchmark 项目的全栈工程师。请根据以下 Issue 实现代码。

## Issue #${ISSUE_NUMBER}
标题: ${ISSUE_TITLE}

内容:
${ISSUE_BODY}

## 要求
1. 先阅读 CLAUDE.md 了解项目规范
2. 阅读 docs/PRD.md 中与本 Issue 相关的章节
3. 检查现有代码，理解当前架构和模式
4. 实现功能代码
5. 确保代码没有语法错误
6. 不要运行 dev server 或任何长时间运行的命令
7. 完成后简要说明你做了什么"

  # 检查是否有变更
  if git diff --quiet && git diff --cached --quiet; then
    echo "⚠️  没有代码变更，跳过"
    git checkout main
    git branch -D "$BRANCH_NAME"
    gh issue comment "$ISSUE_NUMBER" --repo "$REPO" --body "⚠️ AI 未产生任何代码变更。可能需要人工介入。"
    continue
  fi

  echo "=== 检测到代码变更，创建 PR ==="

  # 提交变更
  git add -A
  git commit -m "feat: ${ISSUE_TITLE}

Closes #${ISSUE_NUMBER}

Co-Authored-By: Claude Code <noreply@anthropic.com>"

  # 推送到远程
  git push origin "$BRANCH_NAME"

  # 创建 PR
  gh pr create \
    --repo "$REPO" \
    --title "feat: ${ISSUE_TITLE}" \
    --body "## 自动实现 Issue #${ISSUE_NUMBER}

${ISSUE_TITLE}

由 Claude Code 本地自动生成。

Closes #${ISSUE_NUMBER}

🤖 Generated with Claude Code" \
    --base main \
    --head "$BRANCH_NAME"

  echo "✅ Issue #${ISSUE_NUMBER} 处理完成"

  # 回到 main 分支
  git checkout main

  # 每个 Issue 之间暂停 10 秒
  sleep 10
done

echo ""
echo "=== 所有 Issue 处理完毕 ==="
