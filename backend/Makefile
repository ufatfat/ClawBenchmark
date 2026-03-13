.PHONY: run build test lint fmt clean help

# 变量定义
BINARY_NAME=server
MAIN_PATH=./cmd/server
BUILD_DIR=./bin

# 默认目标
.DEFAULT_GOAL := help

## run: 运行服务器
run:
	@echo "Running server..."
	@go run $(MAIN_PATH)/main.go

## build: 编译二进制文件
build:
	@echo "Building..."
	@mkdir -p $(BUILD_DIR)
	@go build -o $(BUILD_DIR)/$(BINARY_NAME) $(MAIN_PATH)/main.go
	@echo "Build complete: $(BUILD_DIR)/$(BINARY_NAME)"

## test: 运行测试
test:
	@echo "Running tests..."
	@go test -v -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report: coverage.html"

## lint: 运行代码检查
lint:
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install: https://golangci-lint.run/usage/install/"; \
	fi

## fmt: 格式化代码
fmt:
	@echo "Formatting code..."
	@go fmt ./...
	@gofmt -s -w .

## tidy: 整理依赖
tidy:
	@echo "Tidying dependencies..."
	@go mod tidy

## clean: 清理构建文件
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

## help: 显示帮助信息
help:
	@echo "Usage: make [target]"
	@echo ""
	@echo "Targets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'
