package service

import (
	"context"
	"fmt"

	"github.com/ufatfat/clawbenchmark/internal/models"
	"github.com/ufatfat/clawbenchmark/internal/repository"
	"go.uber.org/zap"
)

type ClawService interface {
	CreateBenchmark(ctx context.Context, req *CreateBenchmarkRequest) (*models.ClawBenchmark, error)
	GetBenchmark(ctx context.Context, id uint) (*models.ClawBenchmark, error)
	GetBenchmarkByName(ctx context.Context, name string) (*models.ClawBenchmark, error)
	UpdateBenchmark(ctx context.Context, id uint, req *UpdateBenchmarkRequest) (*models.ClawBenchmark, error)
	DeleteBenchmark(ctx context.Context, id uint) error
	ListBenchmarks(ctx context.Context, req *ListBenchmarksRequest) (*ListBenchmarksResponse, error)
	GetTopBenchmarks(ctx context.Context, n int) ([]*models.ClawBenchmark, error)
}

type clawService struct {
	repo repository.ClawRepository
	log  *zap.Logger
}

func NewClawService(repo repository.ClawRepository, log *zap.Logger) ClawService {
	return &clawService{
		repo: repo,
		log:  log,
	}
}

// CreateBenchmarkRequest 创建测评请求
type CreateBenchmarkRequest struct {
	Name            string  `json:"name" binding:"required,min=1,max=100"`
	Version         string  `json:"version" binding:"required,max=50"`
	Description     string  `json:"description"`
	Performance     float64 `json:"performance" binding:"required,min=0,max=100"`
	Functionality   float64 `json:"functionality" binding:"required,min=0,max=100"`
	Usability       float64 `json:"usability" binding:"required,min=0,max=100"`
	Stability       float64 `json:"stability" binding:"required,min=0,max=100"`
	ResponseTime    int     `json:"response_time" binding:"required,min=0"`
	Throughput      int     `json:"throughput" binding:"required,min=0"`
	MemoryUsage     int     `json:"memory_usage" binding:"required,min=0"`
	CPUUsage        float64 `json:"cpu_usage" binding:"required,min=0,max=100"`
	Features        string  `json:"features"`
	SupportedSkills int     `json:"supported_skills" binding:"min=0"`
	TestEnv         string  `json:"test_env"`
}

// UpdateBenchmarkRequest 更新测评请求
type UpdateBenchmarkRequest struct {
	Version         *string  `json:"version"`
	Description     *string  `json:"description"`
	Performance     *float64 `json:"performance" binding:"omitempty,min=0,max=100"`
	Functionality   *float64 `json:"functionality" binding:"omitempty,min=0,max=100"`
	Usability       *float64 `json:"usability" binding:"omitempty,min=0,max=100"`
	Stability       *float64 `json:"stability" binding:"omitempty,min=0,max=100"`
	ResponseTime    *int     `json:"response_time" binding:"omitempty,min=0"`
	Throughput      *int     `json:"throughput" binding:"omitempty,min=0"`
	MemoryUsage     *int     `json:"memory_usage" binding:"omitempty,min=0"`
	CPUUsage        *float64 `json:"cpu_usage" binding:"omitempty,min=0,max=100"`
	Features        *string  `json:"features"`
	SupportedSkills *int     `json:"supported_skills" binding:"omitempty,min=0"`
	TestEnv         *string  `json:"test_env"`
	Status          *string  `json:"status" binding:"omitempty,oneof=active archived"`
}

// ListBenchmarksRequest 列表查询请求
type ListBenchmarksRequest struct {
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	SortBy   string `form:"sort_by"`
	Order    string `form:"order" binding:"omitempty,oneof=asc desc"`
}

// ListBenchmarksResponse 列表查询响应
type ListBenchmarksResponse struct {
	Items      []*models.ClawBenchmark `json:"items"`
	Total      int64                   `json:"total"`
	Page       int                     `json:"page"`
	PageSize   int                     `json:"page_size"`
	TotalPages int                     `json:"total_pages"`
}

// CreateBenchmark 创建测评记录
func (s *clawService) CreateBenchmark(ctx context.Context, req *CreateBenchmarkRequest) (*models.ClawBenchmark, error) {
	// 检查名称是否已存在
	existing, _ := s.repo.GetByName(ctx, req.Name)
	if existing != nil {
		return nil, fmt.Errorf("benchmark with name '%s' already exists", req.Name)
	}

	benchmark := &models.ClawBenchmark{
		Name:            req.Name,
		Version:         req.Version,
		Description:     req.Description,
		Performance:     req.Performance,
		Functionality:   req.Functionality,
		Usability:       req.Usability,
		Stability:       req.Stability,
		ResponseTime:    req.ResponseTime,
		Throughput:      req.Throughput,
		MemoryUsage:     req.MemoryUsage,
		CPUUsage:        req.CPUUsage,
		Features:        req.Features,
		SupportedSkills: req.SupportedSkills,
		TestEnv:         req.TestEnv,
		Status:          "active",
	}

	// 计算综合评分
	benchmark.CalculateOverallScore()

	if err := s.repo.Create(ctx, benchmark); err != nil {
		s.log.Error("failed to create benchmark", zap.Error(err))
		return nil, fmt.Errorf("failed to create benchmark: %w", err)
	}

	s.log.Info("benchmark created", zap.String("name", benchmark.Name), zap.Uint("id", benchmark.ID))
	return benchmark, nil
}

// GetBenchmark 获取测评记录
func (s *clawService) GetBenchmark(ctx context.Context, id uint) (*models.ClawBenchmark, error) {
	benchmark, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return benchmark, nil
}

// GetBenchmarkByName 根据名称获取测评记录
func (s *clawService) GetBenchmarkByName(ctx context.Context, name string) (*models.ClawBenchmark, error) {
	benchmark, err := s.repo.GetByName(ctx, name)
	if err != nil {
		return nil, err
	}
	return benchmark, nil
}

// UpdateBenchmark 更新测评记录
func (s *clawService) UpdateBenchmark(ctx context.Context, id uint, req *UpdateBenchmarkRequest) (*models.ClawBenchmark, error) {
	benchmark, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// 更新字段
	if req.Version != nil {
		benchmark.Version = *req.Version
	}
	if req.Description != nil {
		benchmark.Description = *req.Description
	}
	if req.Performance != nil {
		benchmark.Performance = *req.Performance
	}
	if req.Functionality != nil {
		benchmark.Functionality = *req.Functionality
	}
	if req.Usability != nil {
		benchmark.Usability = *req.Usability
	}
	if req.Stability != nil {
		benchmark.Stability = *req.Stability
	}
	if req.ResponseTime != nil {
		benchmark.ResponseTime = *req.ResponseTime
	}
	if req.Throughput != nil {
		benchmark.Throughput = *req.Throughput
	}
	if req.MemoryUsage != nil {
		benchmark.MemoryUsage = *req.MemoryUsage
	}
	if req.CPUUsage != nil {
		benchmark.CPUUsage = *req.CPUUsage
	}
	if req.Features != nil {
		benchmark.Features = *req.Features
	}
	if req.SupportedSkills != nil {
		benchmark.SupportedSkills = *req.SupportedSkills
	}
	if req.TestEnv != nil {
		benchmark.TestEnv = *req.TestEnv
	}
	if req.Status != nil {
		benchmark.Status = *req.Status
	}

	// 重新计算综合评分
	benchmark.CalculateOverallScore()

	if err := s.repo.Update(ctx, benchmark); err != nil {
		s.log.Error("failed to update benchmark", zap.Error(err), zap.Uint("id", id))
		return nil, fmt.Errorf("failed to update benchmark: %w", err)
	}

	s.log.Info("benchmark updated", zap.Uint("id", id))
	return benchmark, nil
}

// DeleteBenchmark 删除测评记录
func (s *clawService) DeleteBenchmark(ctx context.Context, id uint) error {
	if err := s.repo.Delete(ctx, id); err != nil {
		s.log.Error("failed to delete benchmark", zap.Error(err), zap.Uint("id", id))
		return fmt.Errorf("failed to delete benchmark: %w", err)
	}

	s.log.Info("benchmark deleted", zap.Uint("id", id))
	return nil
}

// ListBenchmarks 列表查询
func (s *clawService) ListBenchmarks(ctx context.Context, req *ListBenchmarksRequest) (*ListBenchmarksResponse, error) {
	if req.Page < 1 {
		req.Page = 1
	}
	if req.PageSize < 1 {
		req.PageSize = 10
	}
	if req.SortBy == "" {
		req.SortBy = "overall_score"
	}
	if req.Order == "" {
		req.Order = "desc"
	}

	items, total, err := s.repo.List(ctx, req.Page, req.PageSize, req.SortBy, req.Order)
	if err != nil {
		s.log.Error("failed to list benchmarks", zap.Error(err))
		return nil, fmt.Errorf("failed to list benchmarks: %w", err)
	}

	totalPages := int(total) / req.PageSize
	if int(total)%req.PageSize > 0 {
		totalPages++
	}

	return &ListBenchmarksResponse{
		Items:      items,
		Total:      total,
		Page:       req.Page,
		PageSize:   req.PageSize,
		TotalPages: totalPages,
	}, nil
}

// GetTopBenchmarks 获取评分最高的测评记录
func (s *clawService) GetTopBenchmarks(ctx context.Context, n int) ([]*models.ClawBenchmark, error) {
	if n < 1 {
		n = 10
	}

	benchmarks, err := s.repo.GetTopN(ctx, n)
	if err != nil {
		s.log.Error("failed to get top benchmarks", zap.Error(err))
		return nil, fmt.Errorf("failed to get top benchmarks: %w", err)
	}

	return benchmarks, nil
}
