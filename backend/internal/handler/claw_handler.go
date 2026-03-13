package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ufatfat/clawbenchmark/internal/service"
	"github.com/ufatfat/clawbenchmark/pkg/response"
	"go.uber.org/zap"
)

type ClawHandler struct {
	svc service.ClawService
	log *zap.Logger
}

func NewClawHandler(svc service.ClawService, log *zap.Logger) *ClawHandler {
	return &ClawHandler{svc: svc, log: log}
}

// Create godoc
// @Summary      创建 Claw 测评记录
// @Description  创建新的 Claw 项目测评数据
// @Tags         claw
// @Accept       json
// @Produce      json
// @Param        body  body      service.CreateBenchmarkRequest  true  "测评数据"
// @Success      201   {object}  response.Response{data=models.ClawBenchmark}
// @Failure      400   {object}  response.Response
// @Failure      500   {object}  response.Response
// @Router       /api/v1/claws [post]
func (h *ClawHandler) Create(c *gin.Context) {
	var req service.CreateBenchmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	benchmark, err := h.svc.CreateBenchmark(c.Request.Context(), &req)
	if err != nil {
		h.log.Error("create benchmark failed", zap.Error(err))
		response.Error(c, http.StatusConflict, err.Error())
		return
	}

	c.JSON(http.StatusCreated, response.Response{Code: 0, Message: "success", Data: benchmark})
}

// GetByID godoc
// @Summary      获取测评记录
// @Description  根据 ID 获取 Claw 测评记录
// @Tags         claw
// @Produce      json
// @Param        id   path      int  true  "测评 ID"
// @Success      200  {object}  response.Response{data=models.ClawBenchmark}
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/claws/{id} [get]
func (h *ClawHandler) GetByID(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	benchmark, err := h.svc.GetBenchmark(c.Request.Context(), id)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, benchmark)
}

// Update godoc
// @Summary      更新测评记录
// @Description  更新 Claw 测评数据（支持部分更新）
// @Tags         claw
// @Accept       json
// @Produce      json
// @Param        id    path      int                             true  "测评 ID"
// @Param        body  body      service.UpdateBenchmarkRequest  true  "更新数据"
// @Success      200   {object}  response.Response{data=models.ClawBenchmark}
// @Failure      400   {object}  response.Response
// @Failure      404   {object}  response.Response
// @Router       /api/v1/claws/{id} [put]
func (h *ClawHandler) Update(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	var req service.UpdateBenchmarkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	benchmark, err := h.svc.UpdateBenchmark(c.Request.Context(), id, &req)
	if err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, benchmark)
}

// Delete godoc
// @Summary      删除测评记录
// @Description  根据 ID 删除 Claw 测评记录
// @Tags         claw
// @Produce      json
// @Param        id   path      int  true  "测评 ID"
// @Success      200  {object}  response.Response
// @Failure      400  {object}  response.Response
// @Failure      404  {object}  response.Response
// @Router       /api/v1/claws/{id} [delete]
func (h *ClawHandler) Delete(c *gin.Context) {
	id, err := parseUintParam(c, "id")
	if err != nil {
		response.BadRequest(c, "invalid id")
		return
	}

	if err := h.svc.DeleteBenchmark(c.Request.Context(), id); err != nil {
		response.NotFound(c, err.Error())
		return
	}

	response.Success(c, nil)
}

// List godoc
// @Summary      获取测评列表
// @Description  分页获取 Claw 测评列表，支持排序
// @Tags         claw
// @Produce      json
// @Param        page       query     int     false  "页码"       default(1)
// @Param        page_size  query     int     false  "每页数量"   default(10)
// @Param        sort_by    query     string  false  "排序字段"   default(overall_score)
// @Param        order      query     string  false  "排序方向"   Enums(asc, desc) default(desc)
// @Success      200        {object}  response.Response{data=service.ListBenchmarksResponse}
// @Failure      400        {object}  response.Response
// @Router       /api/v1/claws [get]
func (h *ClawHandler) List(c *gin.Context) {
	var req service.ListBenchmarksRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		response.BadRequest(c, err.Error())
		return
	}

	result, err := h.svc.ListBenchmarks(c.Request.Context(), &req)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, result)
}

// GetTop godoc
// @Summary      获取 Top N 测评
// @Description  获取综合评分最高的 N 个 Claw 测评记录
// @Tags         claw
// @Produce      json
// @Param        n    query     int  false  "数量"  default(10)
// @Success      200  {object}  response.Response{data=[]models.ClawBenchmark}
// @Failure      500  {object}  response.Response
// @Router       /api/v1/claws/top [get]
func (h *ClawHandler) GetTop(c *gin.Context) {
	n := 10
	if nStr := c.Query("n"); nStr != "" {
		if parsed, err := strconv.Atoi(nStr); err == nil && parsed > 0 {
			n = parsed
		}
	}

	benchmarks, err := h.svc.GetTopBenchmarks(c.Request.Context(), n)
	if err != nil {
		response.InternalError(c, err.Error())
		return
	}

	response.Success(c, benchmarks)
}

func parseUintParam(c *gin.Context, param string) (uint, error) {
	val, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(val), nil
}
