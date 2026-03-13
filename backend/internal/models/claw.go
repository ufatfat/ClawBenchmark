package models

import (
	"time"
)

// ClawBenchmark Claw 测评数据模型
type ClawBenchmark struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Name        string    `gorm:"type:varchar(100);not null;uniqueIndex" json:"name"`
	Version     string    `gorm:"type:varchar(50);not null" json:"version"`
	Description string    `gorm:"type:text" json:"description"`

	// 性能指标 (0-100)
	Performance float64 `gorm:"type:decimal(5,2);not null" json:"performance"`
	// 功能完整性 (0-100)
	Functionality float64 `gorm:"type:decimal(5,2);not null" json:"functionality"`
	// 易用性 (0-100)
	Usability float64 `gorm:"type:decimal(5,2);not null" json:"usability"`
	// 稳定性 (0-100)
	Stability float64 `gorm:"type:decimal(5,2);not null" json:"stability"`
	// 综合评分 (0-100)
	OverallScore float64 `gorm:"type:decimal(5,2);not null;index" json:"overall_score"`

	// 详细测试数据
	ResponseTime    int    `gorm:"not null" json:"response_time"`     // 响应时间(ms)
	Throughput      int    `gorm:"not null" json:"throughput"`        // 吞吐量(req/s)
	MemoryUsage     int    `gorm:"not null" json:"memory_usage"`      // 内存使用(MB)
	CPUUsage        float64 `gorm:"type:decimal(5,2);not null" json:"cpu_usage"` // CPU使用率(%)

	// 功能特性
	Features        string `gorm:"type:jsonb" json:"features"`        // JSON 数组
	SupportedSkills int    `gorm:"not null" json:"supported_skills"`  // 支持的 Skills 数量

	// 元数据
	TestDate   time.Time `gorm:"not null" json:"test_date"`
	TestEnv    string    `gorm:"type:varchar(200)" json:"test_env"`
	Status     string    `gorm:"type:varchar(20);not null;default:'active';index" json:"status"` // active, archived

	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// TableName 指定表名
func (ClawBenchmark) TableName() string {
	return "claw_benchmarks"
}

// CalculateOverallScore 计算综合评分
// 性能30% + 功能30% + 易用性20% + 稳定性20%
func (c *ClawBenchmark) CalculateOverallScore() {
	c.OverallScore = c.Performance*0.3 + c.Functionality*0.3 + c.Usability*0.2 + c.Stability*0.2
}
