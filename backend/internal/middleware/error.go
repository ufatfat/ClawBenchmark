package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ufatfat/clawbenchmark/pkg/response"
	"go.uber.org/zap"
)

func Recovery(log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.Error("panic recovered",
					zap.Any("error", err),
					zap.String("path", c.Request.URL.Path),
				)
				response.InternalError(c, "Internal server error")
				c.Abort()
			}
		}()
		c.Next()
	}
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			response.Error(c, http.StatusInternalServerError, err.Error())
		}
	}
}
