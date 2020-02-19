package tymiddleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// CorsHandler 跨域处理
func CorsHandler(originUrls map[string]bool) gin.HandlerFunc {
	if len(originUrls) > 0 {

	}

	cfg := cors.DefaultConfig()
	cfg.AllowOriginFunc = func(origin string) bool {
		// allowed, ok := originUrls[origin]
		// return ok && allowed
		return true
	}
	cfg.AllowOrigins = []string{"*"}
	cfg.AllowMethods = []string{"POST", "GET"}
	cfg.AllowCredentials = true
	return cors.New(cfg)
}
