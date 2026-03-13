package middleware

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/ufatfat/clawbenchmark/pkg/response"
)

func RateLimit(rate string) gin.HandlerFunc {
	rateLimit, err := limiter.NewRateFromFormatted(rate)
	if err != nil {
		panic(err)
	}

	store := memory.NewStore()
	instance := limiter.New(store, rateLimit)

	return func(c *gin.Context) {
		context := limiter.Context{Limit: rateLimit}
		limiterCtx, err := instance.Get(c, c.ClientIP())
		if err != nil {
			response.Error(c, http.StatusInternalServerError, "Rate limiter error")
			c.Abort()
			return
		}

		c.Header("X-RateLimit-Limit", limiterCtx.Limit.String())
		c.Header("X-RateLimit-Remaining", string(rune(limiterCtx.Remaining)))
		c.Header("X-RateLimit-Reset", time.Unix(limiterCtx.Reset, 0).Format(time.RFC3339))

		if limiterCtx.Reached {
			response.Error(c, http.StatusTooManyRequests, "Rate limit exceeded")
			c.Abort()
			return
		}

		context.Limit = rateLimit
		c.Next()
	}
}
