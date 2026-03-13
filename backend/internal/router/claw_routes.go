package router

import (
	"github.com/gin-gonic/gin"
	"github.com/ufatfat/clawbenchmark/internal/handler"
)

// RegisterClawRoutes 注册 Claw 测评相关路由
func RegisterClawRoutes(rg *gin.RouterGroup, h *handler.ClawHandler) {
	claws := rg.Group("/claws")
	{
		claws.GET("", h.List)
		claws.POST("", h.Create)
		claws.GET("/top", h.GetTop)
		claws.GET("/:id", h.GetByID)
		claws.PUT("/:id", h.Update)
		claws.DELETE("/:id", h.Delete)
	}
}
