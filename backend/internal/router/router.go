package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ufatfat/clawbenchmark/internal/middleware"
	"github.com/ufatfat/clawbenchmark/pkg/response"
	"go.uber.org/zap"
)

func Setup(log *zap.Logger) *gin.Engine {
	r := gin.New()

	// 中间件
	r.Use(middleware.CORS())
	r.Use(middleware.Logger(log))
	r.Use(middleware.Recovery(log))
	r.Use(middleware.ErrorHandler())
	r.Use(middleware.RateLimit("100-M")) // 100 requests per minute

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		response.Success(c, gin.H{"status": "ok"})
	})

	// API 路由组
	api := r.Group("/api/v1")
	{
		api.GET("/ping", func(c *gin.Context) {
			response.Success(c, gin.H{"message": "pong"})
		})
	}

	return r
}
