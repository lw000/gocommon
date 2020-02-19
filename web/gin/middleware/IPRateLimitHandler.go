package tymiddleware

import (
	"github.com/gin-gonic/gin"
	tylimiter "gocommon/web/limiter"
	"net/http"
)

func IPRateLimitHandler(limiter *tylimiter.IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		pass := limiter.Allow(clientIP)
		if !pass {
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"c": 0, "m": limiter.ErrText(), "d": gin.H{}})
			return
		}
		c.Next()
	}
}
