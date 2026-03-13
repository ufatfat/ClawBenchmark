package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/ufatfat/clawbenchmark/internal/models"
	"gorm.io/gorm"
)

type ClawRepository interface {
	Create(ctx context.Context, benchmark *models.ClawBenchmark) error
	GetByID(ctx context.Context, id uint) (*models.ClawBenchmark, error)
	GetByName(ctx context.Context, name string) (*models.ClawBenchmark, error)
	Update(ctx context.Context, benchmark *models.ClawBenchmark) error
	Delete(ctx context.Context, id uint) error
	List(ctx context.Context, page, pageSize int, sortBy, order string) ([]*models.ClawBenchmark, int64, error)
	GetTopN(ctx context.Context, n int) ([]*models.ClawBenchmark, error)
}

type clawRepository struct {
	db    *gorm.DB
	redis *redis.Client
}

func NewClawRepository(db *gorm.DB, rdb *redis.Client) ClawRepository {
	return &clawRepository{
		db:    db,
		redis: rdb,
	}
}

const (
	cacheKeyPrefix = "claw:benchmark:"
	cacheTTL       = 10 * time.Minute
	listCacheKey   = "claw:benchmarks:list:"
)

// Create 创建测评记录
func (r *clawRepository) Create(ctx context.Context, benchmark *models.ClawBenchmark) error {
	if err := r.db.WithContext(ctx).Create(benchmark).Error; err != nil {
		return fmt.Errorf("failed to create benchmark: %w", err)
	}

	// 清除列表缓存
	r.clearListCache(ctx)
	return nil
}

// GetByID 根据 ID 获取测评记录（带缓存）
func (r *clawRepository) GetByID(ctx context.Context, id uint) (*models.ClawBenchmark, error) {
	cacheKey := fmt.Sprintf("%s%d", cacheKeyPrefix, id)

	// 尝试从缓存获取
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var benchmark models.ClawBenchmark
		if err := json.Unmarshal([]byte(cached), &benchmark); err == nil {
			return &benchmark, nil
		}
	}

	// 从数据库获取
	var benchmark models.ClawBenchmark
	if err := r.db.WithContext(ctx).First(&benchmark, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("benchmark not found")
		}
		return nil, fmt.Errorf("failed to get benchmark: %w", err)
	}

	// 写入缓存
	r.setCache(ctx, cacheKey, &benchmark)
	return &benchmark, nil
}

// GetByName 根据名称获取测评记录
func (r *clawRepository) GetByName(ctx context.Context, name string) (*models.ClawBenchmark, error) {
	var benchmark models.ClawBenchmark
	if err := r.db.WithContext(ctx).Where("name = ?", name).First(&benchmark).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("benchmark not found")
		}
		return nil, fmt.Errorf("failed to get benchmark: %w", err)
	}
	return &benchmark, nil
}

// Update 更新测评记录
func (r *clawRepository) Update(ctx context.Context, benchmark *models.ClawBenchmark) error {
	if err := r.db.WithContext(ctx).Save(benchmark).Error; err != nil {
		return fmt.Errorf("failed to update benchmark: %w", err)
	}

	// 清除缓存
	cacheKey := fmt.Sprintf("%s%d", cacheKeyPrefix, benchmark.ID)
	r.redis.Del(ctx, cacheKey)
	r.clearListCache(ctx)
	return nil
}

// Delete 删除测评记录
func (r *clawRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.WithContext(ctx).Delete(&models.ClawBenchmark{}, id).Error; err != nil {
		return fmt.Errorf("failed to delete benchmark: %w", err)
	}

	// 清除缓存
	cacheKey := fmt.Sprintf("%s%d", cacheKeyPrefix, id)
	r.redis.Del(ctx, cacheKey)
	r.clearListCache(ctx)
	return nil
}

// List 分页查询测评记录
func (r *clawRepository) List(ctx context.Context, page, pageSize int, sortBy, order string) ([]*models.ClawBenchmark, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	// 验证排序字段
	validSortFields := map[string]bool{
		"id": true, "name": true, "overall_score": true,
		"performance": true, "functionality": true,
		"usability": true, "stability": true, "created_at": true,
	}
	if !validSortFields[sortBy] {
		sortBy = "overall_score"
	}
	if order != "asc" && order != "desc" {
		order = "desc"
	}

	// 尝试从缓存获取
	cacheKey := fmt.Sprintf("%s%d:%d:%s:%s", listCacheKey, page, pageSize, sortBy, order)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var result struct {
			Benchmarks []*models.ClawBenchmark
			Total      int64
		}
		if err := json.Unmarshal([]byte(cached), &result); err == nil {
			return result.Benchmarks, result.Total, nil
		}
	}

	// 从数据库查询
	var benchmarks []*models.ClawBenchmark
	var total int64

	db := r.db.WithContext(ctx).Model(&models.ClawBenchmark{})
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to count benchmarks: %w", err)
	}

	offset := (page - 1) * pageSize
	orderClause := fmt.Sprintf("%s %s", sortBy, order)
	if err := db.Order(orderClause).Offset(offset).Limit(pageSize).Find(&benchmarks).Error; err != nil {
		return nil, 0, fmt.Errorf("failed to list benchmarks: %w", err)
	}

	// 写入缓存
	result := struct {
		Benchmarks []*models.ClawBenchmark
		Total      int64
	}{benchmarks, total}
	if data, err := json.Marshal(result); err == nil {
		r.redis.Set(ctx, cacheKey, data, cacheTTL)
	}

	return benchmarks, total, nil
}

// GetTopN 获取评分最高的 N 个测评记录
func (r *clawRepository) GetTopN(ctx context.Context, n int) ([]*models.ClawBenchmark, error) {
	if n < 1 || n > 100 {
		n = 10
	}

	cacheKey := fmt.Sprintf("claw:benchmarks:top:%d", n)
	cached, err := r.redis.Get(ctx, cacheKey).Result()
	if err == nil {
		var benchmarks []*models.ClawBenchmark
		if err := json.Unmarshal([]byte(cached), &benchmarks); err == nil {
			return benchmarks, nil
		}
	}

	var benchmarks []*models.ClawBenchmark
	if err := r.db.WithContext(ctx).
		Where("status = ?", "active").
		Order("overall_score desc").
		Limit(n).
		Find(&benchmarks).Error; err != nil {
		return nil, fmt.Errorf("failed to get top benchmarks: %w", err)
	}

	// 写入缓存
	if data, err := json.Marshal(benchmarks); err == nil {
		r.redis.Set(ctx, cacheKey, data, cacheTTL)
	}

	return benchmarks, nil
}

// setCache 设置缓存
func (r *clawRepository) setCache(ctx context.Context, key string, value interface{}) {
	if data, err := json.Marshal(value); err == nil {
		r.redis.Set(ctx, key, data, cacheTTL)
	}
}

// clearListCache 清除列表缓存
func (r *clawRepository) clearListCache(ctx context.Context) {
	pattern := listCacheKey + "*"
	iter := r.redis.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		r.redis.Del(ctx, iter.Val())
	}
}
