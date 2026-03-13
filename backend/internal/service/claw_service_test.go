package service

import (
	"context"
	"errors"
	"testing"

	"github.com/ufatfat/clawbenchmark/internal/models"
	"go.uber.org/zap"
)

// mockClawRepository 模拟 Repository
type mockClawRepository struct {
	benchmarks map[uint]*models.ClawBenchmark
	nextID     uint
}

func newMockRepo() *mockClawRepository {
	return &mockClawRepository{
		benchmarks: make(map[uint]*models.ClawBenchmark),
		nextID:     1,
	}
}

func (m *mockClawRepository) Create(_ context.Context, b *models.ClawBenchmark) error {
	for _, existing := range m.benchmarks {
		if existing.Name == b.Name {
			return errors.New("duplicate name")
		}
	}
	b.ID = m.nextID
	m.nextID++
	m.benchmarks[b.ID] = b
	return nil
}

func (m *mockClawRepository) GetByID(_ context.Context, id uint) (*models.ClawBenchmark, error) {
	b, ok := m.benchmarks[id]
	if !ok {
		return nil, errors.New("benchmark not found")
	}
	return b, nil
}

func (m *mockClawRepository) GetByName(_ context.Context, name string) (*models.ClawBenchmark, error) {
	for _, b := range m.benchmarks {
		if b.Name == name {
			return b, nil
		}
	}
	return nil, errors.New("benchmark not found")
}

func (m *mockClawRepository) Update(_ context.Context, b *models.ClawBenchmark) error {
	if _, ok := m.benchmarks[b.ID]; !ok {
		return errors.New("benchmark not found")
	}
	m.benchmarks[b.ID] = b
	return nil
}

func (m *mockClawRepository) Delete(_ context.Context, id uint) error {
	if _, ok := m.benchmarks[id]; !ok {
		return errors.New("benchmark not found")
	}
	delete(m.benchmarks, id)
	return nil
}

func (m *mockClawRepository) List(_ context.Context, page, pageSize int, _, _ string) ([]*models.ClawBenchmark, int64, error) {
	var all []*models.ClawBenchmark
	for _, b := range m.benchmarks {
		all = append(all, b)
	}
	total := int64(len(all))
	start := (page - 1) * pageSize
	if start >= len(all) {
		return []*models.ClawBenchmark{}, total, nil
	}
	end := start + pageSize
	if end > len(all) {
		end = len(all)
	}
	return all[start:end], total, nil
}

func (m *mockClawRepository) GetTopN(_ context.Context, n int) ([]*models.ClawBenchmark, error) {
	var all []*models.ClawBenchmark
	for _, b := range m.benchmarks {
		all = append(all, b)
	}
	if n > len(all) {
		n = len(all)
	}
	return all[:n], nil
}

// --- 测试用例 ---

func newTestService() ClawService {
	return NewClawService(newMockRepo(), zap.NewNop())
}

func validCreateReq() *CreateBenchmarkRequest {
	return &CreateBenchmarkRequest{
		Name:          "TestClaw",
		Version:       "v1.0.0",
		Performance:   80,
		Functionality: 90,
		Usability:     70,
		Stability:     85,
		ResponseTime:  100,
		Throughput:    500,
		MemoryUsage:   256,
		CPUUsage:      30,
	}
}

// TestCalculateOverallScore 验证综合评分算法
func TestCalculateOverallScore(t *testing.T) {
	b := &models.ClawBenchmark{
		Performance:   80,
		Functionality: 90,
		Usability:     70,
		Stability:     85,
	}
	b.CalculateOverallScore()

	// 80*0.3 + 90*0.3 + 70*0.2 + 85*0.2 = 24 + 27 + 14 + 17 = 82
	expected := 82.0
	if b.OverallScore != expected {
		t.Errorf("expected overall_score=%.2f, got=%.2f", expected, b.OverallScore)
	}
}

// TestCreateBenchmark 验证创建测评记录
func TestCreateBenchmark(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	b, err := svc.CreateBenchmark(ctx, validCreateReq())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if b.ID == 0 {
		t.Error("expected non-zero ID")
	}
	if b.OverallScore == 0 {
		t.Error("expected non-zero overall_score")
	}
}

// TestCreateBenchmark_DuplicateName 验证重复名称报错
func TestCreateBenchmark_DuplicateName(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	req := validCreateReq()
	if _, err := svc.CreateBenchmark(ctx, req); err != nil {
		t.Fatalf("first create failed: %v", err)
	}

	_, err := svc.CreateBenchmark(ctx, req)
	if err == nil {
		t.Error("expected error for duplicate name, got nil")
	}
}

// TestGetBenchmark 验证获取测评记录
func TestGetBenchmark(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	created, _ := svc.CreateBenchmark(ctx, validCreateReq())

	b, err := svc.GetBenchmark(ctx, created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if b.Name != created.Name {
		t.Errorf("expected name=%s, got=%s", created.Name, b.Name)
	}
}

// TestGetBenchmark_NotFound 验证不存在的记录
func TestGetBenchmark_NotFound(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	_, err := svc.GetBenchmark(ctx, 9999)
	if err == nil {
		t.Error("expected error for non-existent benchmark")
	}
}

// TestUpdateBenchmark 验证更新测评记录并重新计算评分
func TestUpdateBenchmark(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	created, _ := svc.CreateBenchmark(ctx, validCreateReq())
	originalScore := created.OverallScore

	newPerf := 100.0
	updated, err := svc.UpdateBenchmark(ctx, created.ID, &UpdateBenchmarkRequest{
		Performance: &newPerf,
	})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if updated.OverallScore <= originalScore {
		t.Errorf("expected higher score after performance update, got %.2f <= %.2f", updated.OverallScore, originalScore)
	}
}

// TestDeleteBenchmark 验证删除测评记录
func TestDeleteBenchmark(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	created, _ := svc.CreateBenchmark(ctx, validCreateReq())

	if err := svc.DeleteBenchmark(ctx, created.ID); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	_, err := svc.GetBenchmark(ctx, created.ID)
	if err == nil {
		t.Error("expected error after deletion")
	}
}

// TestListBenchmarks 验证分页查询
func TestListBenchmarks(t *testing.T) {
	svc := newTestService()
	ctx := context.Background()

	for i := 0; i < 5; i++ {
		req := validCreateReq()
		req.Name = "Claw" + string(rune('A'+i))
		svc.CreateBenchmark(ctx, req)
	}

	result, err := svc.ListBenchmarks(ctx, &ListBenchmarksRequest{Page: 1, PageSize: 3})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result.Total != 5 {
		t.Errorf("expected total=5, got=%d", result.Total)
	}
	if len(result.Items) != 3 {
		t.Errorf("expected 3 items, got=%d", len(result.Items))
	}
	if result.TotalPages != 2 {
		t.Errorf("expected total_pages=2, got=%d", result.TotalPages)
	}
}
